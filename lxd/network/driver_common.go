package network

import (
	"fmt"
	"net"
	"strings"

	"github.com/pkg/errors"

	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/lxd/cluster"
	"github.com/lxc/lxd/lxd/instance"
	"github.com/lxc/lxd/lxd/state"
	"github.com/lxc/lxd/shared"
	"github.com/lxc/lxd/shared/api"
	log "github.com/lxc/lxd/shared/log15"
	"github.com/lxc/lxd/shared/logger"
	"github.com/lxc/lxd/shared/logging"
)

// DHCPRange represents a range of IPs from start to end.
type DHCPRange struct {
	Start net.IP
	End   net.IP
}

// common represents a generic LXD network.
type common struct {
	logger      logger.Logger
	state       *state.State
	id          int64
	name        string
	netType     string
	description string
	config      map[string]string
}

// init initialise internal variables.
func (n *common) init(state *state.State, id int64, name string, netType string, description string, config map[string]string) {
	n.logger = logging.AddContext(logger.Log, log.Ctx{"driver": netType, "network": name})
	n.id = id
	n.name = name
	n.netType = netType
	n.config = config
	n.state = state
	n.description = description
}

// commonRules returns a map of config rules common to all drivers.
func (n *common) commonRules() map[string]func(string) error {
	return map[string]func(string) error{}
}

// validate a network config against common rules and optional driver specific rules.
func (n *common) validate(config map[string]string, driverRules map[string]func(value string) error) error {
	checkedFields := map[string]struct{}{}

	// Get rules common for all drivers.
	rules := n.commonRules()

	// Merge driver specific rules into common rules.
	for field, validator := range driverRules {
		rules[field] = validator
	}

	// Run the validator against each field.
	for k, validator := range rules {
		checkedFields[k] = struct{}{} //Mark field as checked.
		err := validator(config[k])
		if err != nil {
			return errors.Wrapf(err, "Invalid value for network %q option %q", n.name, k)
		}
	}

	// Look for any unchecked fields, as these are unknown fields and validation should fail.
	for k := range config {
		_, checked := checkedFields[k]
		if checked {
			continue
		}

		// User keys are not validated.
		if strings.HasPrefix(k, "user.") {
			continue
		}

		return fmt.Errorf("Invalid option for network %q option %q", n.name, k)
	}

	return nil
}

// Name returns the network name.
func (n *common) Name() string {
	return n.name
}

// Type returns the network type.
func (n *common) Type() string {
	return n.netType
}

// Config returns the network config.
func (n *common) Config() map[string]string {
	return n.config
}

// IsUsed returns whether the network is used by any instances.
func (n *common) IsUsed() bool {
	// Look for instances using the interface
	insts, err := instance.LoadFromAllProjects(n.state)
	if err != nil {
		return true
	}

	for _, inst := range insts {
		if IsInUseByInstance(inst, n.name) {
			return true
		}
	}

	return false
}

// HasDHCPv4 indicates whether the network has DHCPv4 enabled.
func (n *common) HasDHCPv4() bool {
	if n.config["ipv4.dhcp"] == "" || shared.IsTrue(n.config["ipv4.dhcp"]) {
		return true
	}

	return false
}

// HasDHCPv6 indicates whether the network has DHCPv6 enabled (includes stateless SLAAC router advertisement mode).
// Technically speaking stateless SLAAC RA mode isn't DHCPv6, but for consistency with LXD's config paradigm, DHCP
// here means "an ability to automatically allocate IPs and routes", rather than stateful DHCP with leases.
// To check if true stateful DHCPv6 is enabled check the "ipv6.dhcp.stateful" config key.
func (n *common) HasDHCPv6() bool {
	if n.config["ipv6.dhcp"] == "" || shared.IsTrue(n.config["ipv6.dhcp"]) {
		return true
	}

	return false
}

// DHCPv4Ranges returns a parsed set of DHCPv4 ranges for this network.
func (n *common) DHCPv4Ranges() []DHCPRange {
	dhcpRanges := make([]DHCPRange, 0)
	if n.config["ipv4.dhcp.ranges"] != "" {
		for _, r := range strings.Split(n.config["ipv4.dhcp.ranges"], ",") {
			parts := strings.SplitN(strings.TrimSpace(r), "-", 2)
			if len(parts) == 2 {
				startIP := net.ParseIP(parts[0])
				endIP := net.ParseIP(parts[1])
				dhcpRanges = append(dhcpRanges, DHCPRange{
					Start: startIP.To4(),
					End:   endIP.To4(),
				})
			}
		}
	}

	return dhcpRanges
}

// DHCPv6Ranges returns a parsed set of DHCPv6 ranges for this network.
func (n *common) DHCPv6Ranges() []DHCPRange {
	dhcpRanges := make([]DHCPRange, 0)
	if n.config["ipv6.dhcp.ranges"] != "" {
		for _, r := range strings.Split(n.config["ipv6.dhcp.ranges"], ",") {
			parts := strings.SplitN(strings.TrimSpace(r), "-", 2)
			if len(parts) == 2 {
				startIP := net.ParseIP(parts[0])
				endIP := net.ParseIP(parts[1])
				dhcpRanges = append(dhcpRanges, DHCPRange{
					Start: startIP.To16(),
					End:   endIP.To16(),
				})
			}
		}
	}

	return dhcpRanges
}

// update the internal config variables, and if not cluster notification, notifies all nodes and updates database.
func (n *common) update(applyNetwork api.NetworkPut, clusterNotification bool) error {
	// Update internal config before database has been updated (so that if update is a notification we apply
	// the config being supplied and not that in the database).
	n.description = applyNetwork.Description
	n.config = applyNetwork.Config

	// If this update isn't coming via a cluster notification itself, then notify all nodes of change and then
	// update the database.
	if !clusterNotification {
		// Notify all other nodes to update the network.
		notifier, err := cluster.NewNotifier(n.state, n.state.Endpoints.NetworkCert(), cluster.NotifyAll)
		if err != nil {
			return err
		}

		err = notifier(func(client lxd.InstanceServer) error {
			return client.UpdateNetwork(n.name, applyNetwork, "")
		})
		if err != nil {
			return err
		}

		// Update the database.
		err = n.state.Cluster.UpdateNetwork(n.name, applyNetwork.Description, applyNetwork.Config)
		if err != nil {
			return err
		}
	}

	return nil
}

// configChanged compares supplied new config with existing config. Returns a boolean indicating if differences in
// the config or description were found (and the database record needs updating), and a list of non-user config
// keys that have changed, and a copy of the current internal network config that can be used to revert if needed.
func (n *common) configChanged(newNetwork api.NetworkPut) (bool, []string, api.NetworkPut, error) {
	// Backup the current state.
	oldNetwork := api.NetworkPut{
		Description: n.description,
		Config:      map[string]string{},
	}

	err := shared.DeepCopy(&n.config, &oldNetwork.Config)
	if err != nil {
		return false, nil, oldNetwork, err
	}

	// Diff the configurations.
	changedKeys := []string{}
	dbUpdateNeeded := false

	if newNetwork.Description != n.description {
		dbUpdateNeeded = true
	}

	for k, v := range oldNetwork.Config {
		if v != newNetwork.Config[k] {
			dbUpdateNeeded = true

			// Add non-user changed key to list of changed keys.
			if !strings.HasPrefix(k, "user.") && !shared.StringInSlice(k, changedKeys) {
				changedKeys = append(changedKeys, k)
			}
		}
	}

	for k, v := range newNetwork.Config {
		if v != oldNetwork.Config[k] {
			dbUpdateNeeded = true

			// Add non-user changed key to list of changed keys.
			if !strings.HasPrefix(k, "user.") && !shared.StringInSlice(k, changedKeys) {
				changedKeys = append(changedKeys, k)
			}
		}
	}

	return dbUpdateNeeded, changedKeys, oldNetwork, nil
}
