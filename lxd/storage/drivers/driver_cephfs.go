package drivers

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/lxc/lxd/lxd/migration"
	"github.com/lxc/lxd/lxd/operations"
	"github.com/lxc/lxd/shared"
	"github.com/lxc/lxd/shared/api"
)

var cephfsVersion string
var cephfsLoaded bool

type cephfs struct {
	common
}

// load is used to run one-time action per-driver rather than per-pool.
func (d *cephfs) load() error {
	if cephfsLoaded {
		return nil
	}

	// Validate the required binaries.
	for _, tool := range []string{"ceph", "rdb"} {
		_, err := exec.LookPath(tool)
		if err != nil {
			return fmt.Errorf("Required tool '%s' is missing", tool)
		}
	}

	// Detect and record the version.
	if cephfsVersion == "" {
		out, err := shared.RunCommand("rbd", "--version")
		if err != nil {
			return err
		}

		cephfsVersion = strings.TrimSpace(out)
	}

	cephfsLoaded = true
	return nil
}

// Info returns the pool driver information.
func (d *cephfs) Info() Info {
	return Info{
		Name:                  "cephfs",
		Version:               cephfsVersion,
		OptimizedImages:       false,
		PreservesInodes:       false,
		Remote:                true,
		VolumeTypes:           []VolumeType{VolumeTypeCustom},
		BlockBacking:          false,
		RunningQuotaResize:    true,
		RunningSnapshotFreeze: false,
	}
}

// Create is called during pool creation and is effectively using an empty driver struct.
// WARNING: The Create() function cannot rely on any of the struct attributes being set.
func (d *cephfs) Create() error {
	// Config validation.
	if d.config["source"] == "" {
		return fmt.Errorf("Missing required source name/path")
	}

	if d.config["cephfs.path"] != "" && d.config["cephfs.path"] != d.config["source"] {
		return fmt.Errorf("cephfs.path must match the source")
	}

	// Set default properties if missing.
	if d.config["cephfs.cluster_name"] == "" {
		d.config["cephfs.cluster_name"] = "ceph"
	}

	if d.config["cephfs.user.name"] == "" {
		d.config["cephfs.user.name"] = "admin"
	}

	d.config["cephfs.path"] = d.config["source"]

	// Parse the namespace / path.
	fields := strings.SplitN(d.config["cephfs.path"], "/", 2)
	fsName := fields[0]
	fsPath := "/"
	if len(fields) > 1 {
		fsPath = fields[1]
	}

	// Check that the filesystem exists.
	if !d.fsExists(d.config["cephfs.cluster_name"], d.config["cephfs.user.name"], fsName) {
		return fmt.Errorf("The requested '%v' CEPHFS doesn't exist", fsName)
	}

	// Create a temporary mountpoint.
	mountPath, err := ioutil.TempDir("", "lxd_cephfs_")
	if err != nil {
		return err
	}
	defer os.RemoveAll(mountPath)

	err = os.Chmod(mountPath, 0700)
	if err != nil {
		return err
	}

	mountPoint := filepath.Join(mountPath, "mount")

	err = os.Mkdir(mountPoint, 0700)
	if err != nil {
		return err
	}

	// Get the credentials and host.
	monAddresses, userSecret, err := d.getConfig(d.config["cephfs.cluster_name"], d.config["cephfs.user.name"])
	if err != nil {
		return err
	}

	// Mount the pool.
	uri := fmt.Sprintf("%s:6789:/", strings.Join(monAddresses, ","))
	err = tryMount(uri, mountPoint, "ceph", 0, fmt.Sprintf("name=%v,secret=%v,mds_namespace=%v", d.config["cephfs.user.name"], userSecret, fsName))
	if err != nil {
		return err
	}
	defer forceUnmount(mountPoint)

	// Create the path if missing.
	err = os.MkdirAll(filepath.Join(mountPoint, fsPath), 0755)
	if err != nil {
		return err
	}

	// Check that the existing path is empty.
	ok, _ := shared.PathIsEmpty(filepath.Join(mountPoint, fsPath))
	if !ok {
		return fmt.Errorf("Only empty CEPHFS paths can be used as a LXD storage pool")
	}

	return nil
}

// Delete clears any local and remote data related to this driver instance.
func (d *cephfs) Delete(op *operations.Operation) error {
	// Parse the namespace / path.
	fields := strings.SplitN(d.config["cephfs.path"], "/", 2)
	fsName := fields[0]
	fsPath := "/"
	if len(fields) > 1 {
		fsPath = fields[1]
	}

	// Create a temporary mountpoint.
	mountPath, err := ioutil.TempDir("", "lxd_cephfs_")
	if err != nil {
		return err
	}
	defer os.RemoveAll(mountPath)

	err = os.Chmod(mountPath, 0700)
	if err != nil {
		return err
	}

	mountPoint := filepath.Join(mountPath, "mount")
	err = os.Mkdir(mountPoint, 0700)
	if err != nil {
		return err
	}

	// Get the credentials and host.
	monAddresses, userSecret, err := d.getConfig(d.config["cephfs.cluster_name"], d.config["cephfs.user.name"])
	if err != nil {
		return err
	}

	// Mount the pool.
	uri := fmt.Sprintf("%s:6789:/", strings.Join(monAddresses, ","))
	err = tryMount(uri, mountPoint, "ceph", 0, fmt.Sprintf("name=%v,secret=%v,mds_namespace=%v", d.config["cephfs.user.name"], userSecret, fsName))
	if err != nil {
		return err
	}
	defer forceUnmount(mountPoint)

	if shared.PathExists(filepath.Join(mountPoint, fsPath)) {
		// Delete the usual directories.
		for _, volType := range d.Info().VolumeTypes {
			for _, dir := range BaseDirectories[volType] {
				if shared.PathExists(filepath.Join(mountPoint, fsPath, dir)) {
					err = os.Remove(filepath.Join(mountPoint, fsPath, dir))
					if err != nil {
						return err
					}
				}
			}
		}

		// Confirm that the path is now empty.
		ok, _ := shared.PathIsEmpty(filepath.Join(mountPoint, fsPath))
		if !ok {
			return fmt.Errorf("Only empty CEPHFS paths can be used as a LXD storage pool")
		}

		// Delete the path itself.
		if fsPath != "" && fsPath != "/" {
			err = os.Remove(filepath.Join(mountPoint, fsPath))
			if err != nil {
				return err
			}
		}
	}

	// On delete, wipe everything in the directory.
	err = wipeDirectory(GetPoolMountPath(d.name))
	if err != nil {
		return err
	}

	// Make sure the existing pool is unmounted.
	_, err = d.Unmount()
	if err != nil {
		return err
	}

	return nil
}

// Validate checks that all provide keys are supported and that no conflicting or missing configuration is present.
func (d *cephfs) Validate(config map[string]string) error {
	return nil
}

// Update applies any driver changes required from a configuration change.
func (d *cephfs) Update(changedConfig map[string]string) error {
	return nil
}

// Mount brings up the driver and sets it up to be used.
func (d *cephfs) Mount() (bool, error) {
	// Check if already mounted.
	if shared.IsMountPoint(GetPoolMountPath(d.name)) {
		return false, nil
	}

	// Parse the namespace / path.
	fields := strings.SplitN(d.config["cephfs.path"], "/", 2)
	fsName := fields[0]
	fsPath := ""
	if len(fields) > 1 {
		fsPath = fields[1]
	}

	// Get the credentials and host.
	monAddresses, userSecret, err := d.getConfig(d.config["cephfs.cluster_name"], d.config["cephfs.user.name"])
	if err != nil {
		return false, err
	}

	// Mount the pool.
	uri := fmt.Sprintf("%s:6789:/%s", strings.Join(monAddresses, ","), fsPath)
	err = tryMount(uri, GetPoolMountPath(d.name), "ceph", 0, fmt.Sprintf("name=%v,secret=%v,mds_namespace=%v", d.config["cephfs.user.name"], userSecret, fsName))
	if err != nil {
		return false, err
	}

	return true, nil
}

// Unmount clears any of the runtime state of the driver.
func (d *cephfs) Unmount() (bool, error) {
	return forceUnmount(GetPoolMountPath(d.name))
}

// GetResources returns the pool resource usage information.
func (d *cephfs) GetResources() (*api.ResourcesStoragePool, error) {
	return d.vfsGetResources()
}

// MigrationTypes returns the supported migration types and options supported by the driver.
func (d *cephfs) MigrationTypes(contentType ContentType, refresh bool) []migration.Type {
	if contentType != ContentTypeFS {
		return nil
	}

	// Do not support xattr transfer on cephfs
	return []migration.Type{
		{
			FSType:   migration.MigrationFSType_RSYNC,
			Features: []string{"delete", "compress", "bidirectional"},
		},
	}
}
