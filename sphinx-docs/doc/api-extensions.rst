API extensions
==============

The changes below were introduced to the LXD API after the 1.0 API was
finalized.

They are all backward compatible and can be detected by client tools by
looking at the ``api_extensions`` field in ``GET /1.0/``.

storage_zfs_remove_snapshots
----------------------------

A ``storage.zfs_remove_snapshots`` daemon configuration key was
introduced.

It’s a boolean that defaults to false and that when set to true
instructs LXD to remove any needed snapshot when attempting to restore
another.

This is needed as ZFS will only let you restore the latest snapshot.

container_host_shutdown_timeout
-------------------------------

A ``boot.host_shutdown_timeout`` container configuration key was
introduced.

It’s an integer which indicates how long LXD should wait for the
container to stop before killing it.

Its value is only used on clean LXD daemon shutdown. It defaults to 30s.

container_stop_priority
-----------------------

A ``boot.stop.priority`` container configuration key was introduced.

It’s an integer which indicates the priority of a container during
shutdown.

Containers will shutdown starting with the highest priority level.

Containers with the same priority will shutdown in parallel. It defaults
to 0.

container_syscall_filtering
---------------------------

A number of new syscalls related container configuration keys were
introduced.

-  ``security.syscalls.blacklist_default``
-  ``security.syscalls.blacklist_compat``
-  ``security.syscalls.blacklist``
-  ``security.syscalls.whitelist``

See `configuration.md <configuration.md>`__ for how to use them.

auth_pki
--------

This indicates support for PKI authentication mode.

In this mode, the client and server both must use certificates issued by
the same PKI.

See `security.md <security.md>`__ for details.

container_last_used_at
----------------------

A ``last_used_at`` field was added to the ``GET /1.0/containers/<name>``
endpoint.

It is a timestamp of the last time the container was started.

If a container has been created but not started yet, ``last_used_at``
field will be ``1970-01-01T00:00:00Z``

etag
----

Add support for the ETag header on all relevant endpoints.

This adds the following HTTP header on answers to GET:

-  ETag (SHA-256 of user modifiable content)

And adds support for the following HTTP header on PUT requests:

-  If-Match (ETag value retrieved through previous GET)

This makes it possible to GET a LXD object, modify it and PUT it without
risking to hit a race condition where LXD or another client modified the
object in the meantime.

patch
-----

Add support for the HTTP PATCH method.

PATCH allows for partial update of an object in place of PUT.

usb_devices
-----------

Add support for USB hotplug.

https_allowed_credentials
-------------------------

To use LXD API with all Web Browsers (via SPAs) you must send
credentials (certificate) with each XHR (in order for this to happen,
you should set
`“withCredentials=true” <https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/withCredentials>`__
flag to each XHR Request).

Some browsers like Firefox and Safari can’t accept server response
without ``Access-Control-Allow-Credentials: true`` header. To ensure
that the server will return a response with that header, set
``core.https_allowed_credentials=true``.

image_compression_algorithm
---------------------------

This adds support for a ``compression_algorithm`` property when creating
an image (``POST /1.0/images``).

Setting this property overrides the server default value
(``images.compression_algorithm``).

directory_manipulation
----------------------

This allows for creating and listing directories via the LXD API, and
exports the file type via the X-LXD-type header, which can be either
“file” or “directory” right now.

container_cpu_time
------------------

This adds support for retrieving cpu time for a running container.

storage_zfs_use_refquota
------------------------

Introduces a new server property ``storage.zfs_use_refquota`` which
instructs LXD to set the “refquota” property instead of “quota” when
setting a size limit on a container. LXD will also then use
“usedbydataset” in place of “used” when being queried about disk
utilization.

This effectively controls whether disk usage by snapshots should be
considered as part of the container’s disk space usage.

storage_lvm_mount_options
-------------------------

Adds a new ``storage.lvm_mount_options`` daemon configuration option
which defaults to “discard” and allows the user to set addition mount
options for the filesystem used by the LVM LV.

network
-------

Network management API for LXD.

This includes:

-  Addition of the “managed” property on ``/1.0/networks`` entries
-  All the network configuration options (see
   `configuration.md <configuration.md>`__ for details)
-  ``POST /1.0/networks`` (see `RESTful API <rest-api.md>`__ for
   details)
-  ``PUT /1.0/networks/<entry>`` (see `RESTful API <rest-api.md>`__ for
   details)
-  ``PATCH /1.0/networks/<entry>`` (see `RESTful API <rest-api.md>`__
   for details)
-  ``DELETE /1.0/networks/<entry>`` (see `RESTful API <rest-api.md>`__
   for details)
-  ``ipv4.address`` property on “nic” type devices (when nictype is
   “bridged”)
-  ``ipv6.address`` property on “nic” type devices (when nictype is
   “bridged”)
-  ``security.mac_filtering`` property on “nic” type devices (when
   nictype is “bridged”)

profile_usedby
--------------

Adds a new used_by field to profile entries listing the containers that
are using it.

container_push
--------------

When a container is created in push mode, the client serves as a proxy
between the source and target server. This is useful in cases where the
target server is behind a NAT or firewall and cannot directly
communicate with the source server and operate in pull mode.

container_exec_recording
------------------------

Introduces a new boolean “record-output”, parameter to
``/1.0/containers/<name>/exec`` which when set to “true” and combined
with with “wait-for-websocket” set to false, will record stdout and
stderr to disk and make them available through the logs interface.

The URL to the recorded output is included in the operation metadata
once the command is done running.

That output will expire similarly to other log files, typically after 48
hours.

certificate_update
------------------

Adds the following to the REST API:

-  ETag header on GET of a certificate
-  PUT of certificate entries
-  PATCH of certificate entries

container_exec_signal_handling
------------------------------

Adds support ``/1.0/containers/<name>/exec`` for forwarding signals sent
to the client to the processes executing in the container. Currently
SIGTERM and SIGHUP are forwarded. Further signals that can be forwarded
might be added later.

gpu_devices
-----------

Enables adding GPUs to a container.

container_image_properties
--------------------------

Introduces a new ``image`` config key space. Read-only, includes the
properties of the parent image.

migration_progress
------------------

Transfer progress is now exported as part of the operation, on both
sending and receiving ends. This shows up as a “fs_progress” attribute
in the operation metadata.

id_map
------

Enables setting the ``security.idmap.isolated`` and
``security.idmap.isolated``, ``security.idmap.size``, and ``raw.id_map``
fields.

network_firewall_filtering
--------------------------

Add two new keys, ``ipv4.firewall`` and ``ipv6.firewall`` which if set
to false will turn off the generation of iptables FORWARDING rules. NAT
rules will still be added so long as the matching ``ipv4.nat`` or
``ipv6.nat`` key is set to true.

Rules necessary for dnsmasq to work (DHCP/DNS) will always be applied if
dnsmasq is enabled on the bridge.

network_routes
--------------

Introduces ``ipv4.routes`` and ``ipv6.routes`` which allow routing
additional subnets to a LXD bridge.

storage
-------

Storage management API for LXD.

This includes:

-  ``GET /1.0/storage-pools``
-  ``POST /1.0/storage-pools`` (see `RESTful API <rest-api.md>`__ for
   details)

-  ``GET /1.0/storage-pools/<name>`` (see `RESTful API <rest-api.md>`__
   for details)
-  ``POST /1.0/storage-pools/<name>`` (see `RESTful API <rest-api.md>`__
   for details)
-  ``PUT /1.0/storage-pools/<name>`` (see `RESTful API <rest-api.md>`__
   for details)
-  ``PATCH /1.0/storage-pools/<name>`` (see `RESTful
   API <rest-api.md>`__ for details)
-  ``DELETE /1.0/storage-pools/<name>`` (see `RESTful
   API <rest-api.md>`__ for details)

-  ``GET /1.0/storage-pools/<name>/volumes`` (see `RESTful
   API <rest-api.md>`__ for details)

-  ``GET /1.0/storage-pools/<name>/volumes/<volume_type>`` (see `RESTful
   API <rest-api.md>`__ for details)
-  ``POST /1.0/storage-pools/<name>/volumes/<volume_type>`` (see
   `RESTful API <rest-api.md>`__ for details)

-  ``GET /1.0/storage-pools/<pool>/volumes/<volume_type>/<name>`` (see
   `RESTful API <rest-api.md>`__ for details)
-  ``POST /1.0/storage-pools/<pool>/volumes/<volume_type>/<name>`` (see
   `RESTful API <rest-api.md>`__ for details)
-  ``PUT /1.0/storage-pools/<pool>/volumes/<volume_type>/<name>`` (see
   `RESTful API <rest-api.md>`__ for details)
-  ``PATCH /1.0/storage-pools/<pool>/volumes/<volume_type>/<name>`` (see
   `RESTful API <rest-api.md>`__ for details)
-  ``DELETE /1.0/storage-pools/<pool>/volumes/<volume_type>/<name>``
   (see `RESTful API <rest-api.md>`__ for details)

-  All storage configuration options (see
   `configuration.md <configuration.md>`__ for details)

file_delete
-----------

Implements ``DELETE`` in ``/1.0/containers/<name>/files``

file_append
-----------

Implements the ``X-LXD-write`` header which can be one of ``overwrite``
or ``append``.

network_dhcp_expiry
-------------------

Introduces ``ipv4.dhcp.expiry`` and ``ipv6.dhcp.expiry`` allowing to set
the DHCP lease expiry time.

storage_lvm_vg_rename
---------------------

Introduces the ability to rename a volume group by setting
``storage.lvm.vg_name``.

storage_lvm_thinpool_rename
---------------------------

Introduces the ability to rename a thinpool name by setting
``storage.thinpool_name``.

network_vlan
------------

This adds a new ``vlan`` property to ``macvlan`` network devices.

When set, this will instruct LXD to attach to the specified VLAN. LXD
will look for an existing interface for that VLAN on the host. If one
can’t be found it will create one itself and then use that as the
macvlan parent.

image_create_aliases
--------------------

Adds a new ``aliases`` field to ``POST /1.0/images`` allowing for
aliases to be set at image creation/import time.

container_stateless_copy
------------------------

This introduces a new ``live`` attribute in
``POST /1.0/containers/<name>``. Setting it to false tells LXD not to
attempt running state transfer.

container_only_migration
------------------------

Introduces a new boolean ``container_only`` attribute. When set to true
only the container will be copied or moved.

storage_zfs_clone_copy
----------------------

Introduces a new boolean ``storage_zfs_clone_copy`` property for ZFS
storage pools. When set to false copying a container will be done
through zfs send and receive. This will make the target container
independent of its source container thus avoiding the need to keep
dependent snapshots in the ZFS pool around. However, this also entails
less efficient storage usage for the affected pool. The default value
for this property is true, i.e. space-efficient snapshots will be used
unless explicitly set to “false”.

unix_device_rename
------------------

Introduces the ability to rename the unix-block/unix-char device inside
container by setting ``path``, and the ``source`` attribute is added to
specify the device on host. If ``source`` is set without a ``path``, we
should assume that ``path`` will be the same as ``source``. If ``path``
is set without ``source`` and ``major``/``minor`` isn’t set, we should
assume that ``source`` will be the same as ``path``. So at least one of
them must be set.

storage_rsync_bwlimit
---------------------

When rsync has to be invoked to transfer storage entities setting
``rsync.bwlimit`` places an upper limit on the amount of socket I/O
allowed.

network_vxlan_interface
-----------------------

This introduces a new ``tunnel.NAME.interface`` option for networks.

This key control what host network interface is used for a VXLAN tunnel.

storage_btrfs_mount_options
---------------------------

This introduces the ``btrfs.mount_options`` property for btrfs storage
pools.

This key controls what mount options will be used for the btrfs storage
pool.

entity_description
------------------

This adds descriptions to entities like containers, snapshots, networks,
storage pools and volumes.

image_force_refresh
-------------------

This allows forcing a refresh for an existing image.

storage_lvm_lv_resizing
-----------------------

This introduces the ability to resize logical volumes by setting the
``size`` property in the containers root disk device.

id_map_base
-----------

This introduces a new ``security.idmap.base`` allowing the user to skip
the map auto-selection process for isolated containers and specify what
host uid/gid to use as the base.

file_symlinks
-------------

This adds support for transferring symlinks through the file API.
X-LXD-type can now be “symlink” with the request content being the
target path.

container_push_target
---------------------

This adds the ``target`` field to ``POST /1.0/containers/<name>`` which
can be used to have the source LXD host connect to the target during
migration.

network_vlan_physical
---------------------

Allows use of ``vlan`` property with ``physical`` network devices.

When set, this will instruct LXD to attach to the specified VLAN on the
``parent`` interface. LXD will look for an existing interface for that
``parent`` and VLAN on the host. If one can’t be found it will create
one itself. Then, LXD will directly attach this interface to the
container.

storage_images_delete
---------------------

This enabled the storage API to delete storage volumes for images from a
specific storage pool.

container_edit_metadata
-----------------------

This adds support for editing a container metadata.yaml and related
templates via API, by accessing urls under
``/1.0/containers/<name>/metadata``. It can be used to edit a container
before publishing an image from it.

container_snapshot_stateful_migration
-------------------------------------

This enables migrating stateful container snapshots to new containers.

storage_driver_ceph
-------------------

This adds a ceph storage driver.

storage_ceph_user_name
----------------------

This adds the ability to specify the ceph user.

instance_types
--------------

This adds the ``instance_type`` field to the container creation request.
Its value is expanded to LXD resource limits.

storage_volatile_initial_source
-------------------------------

This records the actual source passed to LXD during storage pool
creation.

storage_ceph_force_osd_reuse
----------------------------

This introduces the ``ceph.osd.force_reuse`` property for the ceph
storage driver. When set to ``true`` LXD will reuse a osd storage pool
that is already in use by another LXD instance.

storage_block_filesystem_btrfs
------------------------------

This adds support for btrfs as a storage volume filesystem, in addition
to ext4 and xfs.

resources
---------

This adds support for querying a LXD daemon for the system resources it
has available.

kernel_limits
-------------

This adds support for setting process limits such as maximum number of
open files for the container via ``nofile``. The format is
``limits.kernel.[limit name]``.

storage_api_volume_rename
-------------------------

This adds support for renaming custom storage volumes.

external_authentication
-----------------------

This adds support for external authentication via Macaroons.

network_sriov
-------------

This adds support for SR-IOV enabled network devices.

console
-------

This adds support to interact with the container console device and
console log.

restrict_devlxd
---------------

A new security.devlxd container configuration key was introduced. The
key controls whether the /dev/lxd interface is made available to the
container. If set to false, this effectively prevents the container from
interacting with the LXD daemon.

migration_pre_copy
------------------

This adds support for optimized memory transfer during live migration.

infiniband
----------

This adds support to use infiniband network devices.

maas_network
------------

This adds support for MAAS network integration.

When configured at the daemon level, it’s then possible to attach a
“nic” device to a particular MAAS subnet.

devlxd_events
-------------

This adds a websocket API to the devlxd socket.

When connecting to /1.0/events over the devlxd socket, you will now be
getting a stream of events over websocket.

proxy
-----

This adds a new ``proxy`` device type to containers, allowing forwarding
of connections between the host and container.

network_dhcp_gateway
--------------------

Introduces a new ipv4.dhcp.gateway network config key to set an
alternate gateway.

file_get_symlink
----------------

This makes it possible to retrieve symlinks using the file API.

network_leases
--------------

Adds a new /1.0/networks/NAME/leases API endpoint to query the lease
database on bridges which run a LXD-managed DHCP server.

unix_device_hotplug
-------------------

This adds support for the “required” property for unix devices.

storage_api_local_volume_handling
---------------------------------

This add the ability to copy and move custom storage volumes locally in
the same and between storage pools.

operation_description
---------------------

Adds a “description” field to all operations.

clustering
----------

Clustering API for LXD.

This includes the following new endpoints (see `RESTful
API <rest-api.md>`__ for details):

-  ``GET /1.0/cluster``
-  ``UPDATE /1.0/cluster``

-  ``GET /1.0/cluster/members``

-  ``GET /1.0/cluster/members/<name>``
-  ``POST /1.0/cluster/members/<name>``
-  ``DELETE /1.0/cluster/members/<name>``

The following existing endpoints have been modified:

-  ``POST /1.0/containers`` accepts a new target query parameter
-  ``POST /1.0/storage-pools`` accepts a new target query parameter
-  ``GET /1.0/storage-pool/<name>`` accepts a new target query parameter
-  ``POST /1.0/storage-pool/<pool>/volumes/<type>`` accepts a new target
   query parameter
-  ``GET /1.0/storage-pool/<pool>/volumes/<type>/<name>`` accepts a new
   target query parameter
-  ``POST /1.0/storage-pool/<pool>/volumes/<type>/<name>`` accepts a new
   target query parameter
-  ``PUT /1.0/storage-pool/<pool>/volumes/<type>/<name>`` accepts a new
   target query parameter
-  ``PATCH /1.0/storage-pool/<pool>/volumes/<type>/<name>`` accepts a
   new target query parameter
-  ``DELETE /1.0/storage-pool/<pool>/volumes/<type>/<name>`` accepts a
   new target query parameter
-  ``POST /1.0/networks`` accepts a new target query parameter
-  ``GET /1.0/networks/<name>`` accepts a new target query parameter

event_lifecycle
---------------

This adds a new ``lifecycle`` message type to the events API.

storage_api_remote_volume_handling
----------------------------------

This adds the ability to copy and move custom storage volumes between
remote.

nvidia_runtime
--------------

Adds a ``nvidia_runtime`` config option for containers, setting this to
true will have the NVIDIA runtime and CUDA libraries passed to the
container.

container_mount_propagation
---------------------------

This adds a new “propagation” option to the disk device type, allowing
the configuration of kernel mount propagation.

container_backup
----------------

Add container backup support.

This includes the following new endpoints (see `RESTful
API <rest-api.md>`__ for details):

-  ``GET /1.0/containers/<name>/backups``
-  ``POST /1.0/containers/<name>/backups``

-  ``GET /1.0/containers/<name>/backups/<name>``
-  ``POST /1.0/containers/<name>/backups/<name>``
-  ``DELETE /1.0/containers/<name>/backups/<name>``

-  ``GET /1.0/containers/<name>/backups/<name>/export``

The following existing endpoint has been modified:

-  ``POST /1.0/containers`` accepts the new source type ``backup``

devlxd_images
-------------

Adds a ``security.devlxd.images`` config option for containers which
controls the availability of a ``/1.0/images/FINGERPRINT/export`` API
over devlxd. This can be used by a container running nested LXD to
access raw images from the host.

container_local_cross_pool_handling
-----------------------------------

This enables copying or moving containers between storage pools on the
same LXD instance.

proxy_unix
----------

Add support for both unix sockets and abstract unix sockets in proxy
devices. They can be used by specifying the address as
``unix:/path/to/unix.sock`` (normal socket) or ``unix:@/tmp/unix.sock``
(abstract socket).

Supported connections are now:

-  ``TCP <-> TCP``
-  ``UNIX <-> UNIX``
-  ``TCP <-> UNIX``
-  ``UNIX <-> TCP``

proxy_udp
---------

Add support for udp in proxy devices.

Supported connections are now:

-  ``TCP <-> TCP``
-  ``UNIX <-> UNIX``
-  ``TCP <-> UNIX``
-  ``UNIX <-> TCP``
-  ``UDP <-> UDP``
-  ``TCP <-> UDP``
-  ``UNIX <-> UDP``

clustering_join
---------------

This makes GET /1.0/cluster return information about which storage pools
and networks are required to be created by joining nodes and which
node-specific configuration keys they are required to use when creating
them. Likewise the PUT /1.0/cluster endpoint now accepts the same format
to pass information about storage pools and networks to be automatically
created before attempting to join a cluster.

proxy_tcp_udp_multi_port_handling
---------------------------------

Adds support for forwarding traffic for multiple ports. Forwarding is
allowed between a range of ports if the port range is equal for source
and target (for example ``1.2.3.4 0-1000 -> 5.6.7.8 1000-2000``) and
between a range of source ports and a single target port (for example
``1.2.3.4 0-1000 -> 5.6.7.8 1000``).

network_state
-------------

Adds support for retrieving a network’s state.

This adds the following new endpoint (see `RESTful API <rest-api.md>`__
for details):

-  ``GET /1.0/networks/<name>/state``

proxy_unix_dac_properties
-------------------------

This adds support for gid, uid, and mode properties for non-abstract
unix sockets.

container_protection_delete
---------------------------

Enables setting the ``security.protection.delete`` field which prevents
containers from being deleted if set to true. Snapshots are not affected
by this setting.

proxy_priv_drop
---------------

Adds security.uid and security.gid for the proxy devices, allowing
privilege dropping and effectively changing the uid/gid used for
connections to Unix sockets too.

pprof_http
----------

This adds a new core.debug_address config option to start a debugging
HTTP server.

That server currently includes a pprof API and replaces the old
cpu-profile, memory-profile and print-goroutines debug options.

proxy_haproxy_protocol
----------------------

Adds a proxy_protocol key to the proxy device which controls the use of
the HAProxy PROXY protocol header.

network_hwaddr
--------------

Adds a bridge.hwaddr key to control the MAC address of the bridge.

proxy_nat
---------

This adds optimized UDP/TCP proxying. If the configuration allows,
proxying will be done via iptables instead of proxy devices.

network_nat_order
-----------------

This introduces the ``ipv4.nat.order`` and ``ipv6.nat.order``
configuration keys for LXD bridges. Those keys control whether to put
the LXD rules before or after any pre-existing rules in the chain.

container_full
--------------

This introduces a new recursion=2 mode for ``GET /1.0/containers`` which
allows for the retrieval of all container structs, including the state,
snapshots and backup structs.

This effectively allows for “lxc list” to get all it needs in one query.

candid_authentication
---------------------

This introduces the new candid.api.url config option and removes
core.macaroon.endpoint.

backup_compression
------------------

This introduces a new ``backups.compression_algorithm`` config key which
allows configuration of backup compression.

candid_config
-------------

This introduces the config keys ``candid.domains`` and
``candid.expiry``. The former allows specifying allowed/valid Candid
domains, the latter makes the macaroon’s expiry configurable. The
``lxc remote add`` command now has a ``--domain`` flag which allows
specifying a Candid domain.

nvidia_runtime_config
---------------------

This introduces a few extra config keys when using nvidia.runtime and
the libnvidia-container library. Those keys translate pretty much
directly to the matching nvidia-container environment variables:

-  nvidia.driver.capabilities => NVIDIA_DRIVER_CAPABILITIES
-  nvidia.require.cuda => NVIDIA_REQUIRE_CUDA
-  nvidia.require.driver => NVIDIA_REQUIRE_DRIVER

storage_api_volume_snapshots
----------------------------

Add support for storage volume snapshots. They work like container
snapshots, only for volumes.

This adds the following new endpoint (see `RESTful API <rest-api.md>`__
for details):

-  ``GET /1.0/storage-pools/<pool>/volumes/<type>/<name>/snapshots``
-  ``POST /1.0/storage-pools/<pool>/volumes/<type>/<name>/snapshots``

-  ``GET /1.0/storage-pools/<pool>/volumes/<type>/<volume>/snapshots/<name>``
-  ``PUT /1.0/storage-pools/<pool>/volumes/<type>/<volume>/snapshots/<name>``
-  ``POST /1.0/storage-pools/<pool>/volumes/<type>/<volume>/snapshots/<name>``
-  ``DELETE /1.0/storage-pools/<pool>/volumes/<type>/<volume>/snapshots/<name>``

storage_unmapped
----------------

Introduces a new ``security.unmapped`` boolean on storage volumes.

Setting it to true will flush the current map on the volume and prevent
any further idmap tracking and remapping on the volume.

This can be used to share data between isolated containers after
attaching it to the container which requires write access.

projects
--------

Add a new project API, supporting creation, update and deletion of
projects.

Projects can hold containers, profiles or images at this point and let
you get a separate view of your LXD resources by switching to it.

candid_config_key
-----------------

This introduces a new ``candid.api.key`` option which allows for setting
the expected public key for the endpoint, allowing for safe use of a
HTTP-only candid server.

network_vxlan_ttl
-----------------

This adds a new ``tunnel.NAME.ttl`` network configuration option which
makes it possible to raise the ttl on VXLAN tunnels.

container_incremental_copy
--------------------------

This adds support for incremental container copy. When copying a
container using the ``--refresh`` flag, only the missing or outdated
files will be copied over. Should the target container not exist yet, a
normal copy operation is performed.

usb_optional_vendorid
---------------------

As the name implies, the ``vendorid`` field on USB devices attached to
containers has now been made optional, allowing for all USB devices to
be passed to a container (similar to what’s done for GPUs).

snapshot_scheduling
-------------------

This adds support for snapshot scheduling. It introduces three new
configuration keys: ``snapshots.schedule``,
``snapshots.schedule.stopped``, and ``snapshots.pattern``. Snapshots can
be created automatically up to every minute.

snapshots_schedule_aliases
--------------------------

Snapshot schedule can be configured by a comma separated list of
schedule aliases. Available aliases are
``<@hourly> <@daily> <@midnight> <@weekly> <@monthly> <@annually> <@yearly> <@startup>``
for instances, and
``<@hourly> <@daily> <@midnight> <@weekly> <@monthly> <@annually> <@yearly>``
for storage volumes.

container_copy_project
----------------------

Introduces a ``project`` field to the container source dict, allowing
for copy/move of containers between projects.

clustering_server_address
-------------------------

This adds support for configuring a server network address which differs
from the REST API client network address. When bootstrapping a new
cluster, clients can set the new ``cluster.https_address`` config key to
specify the address of the initial server. When joining a new server,
clients can set the ``core.https_address`` config key of the joining
server to the REST API address the joining server should listen at, and
set the ``server_address`` key in the ``PUT /1.0/cluster`` API to the
address the joining server should use for clustering traffic (the value
of ``server_address`` will be automatically copied to the
``cluster.https_address`` config key of the joining server).

clustering_image_replication
----------------------------

Enable image replication across the nodes in the cluster. A new
cluster.images_minimal_replica configuration key was introduced can be
used to specify to the minimal numbers of nodes for image replication.

container_protection_shift
--------------------------

Enables setting the ``security.protection.shift`` option which prevents
containers from having their filesystem shifted.

snapshot_expiry
---------------

This adds support for snapshot expiration. The task is run minutely. The
config option ``snapshots.expiry`` takes an expression in the form of
``1M 2H 3d 4w 5m 6y`` (1 minute, 2 hours, 3 days, 4 weeks, 5 months, 6
years), however not all parts have to be used.

Snapshots which are then created will be given an expiry date based on
the expression. This expiry date, defined by ``expires_at``, can be
manually edited using the API or ``lxc config edit``. Snapshots with a
valid expiry date will be removed when the task in run. Expiry can be
disabled by setting ``expires_at`` to an empty string or
``0001-01-01T00:00:00Z`` (zero time). This is the default if
``snapshots.expiry`` is not set.

This adds the following new endpoint (see `RESTful API <rest-api.md>`__
for details):

-  ``PUT /1.0/containers/<name>/snapshots/<name>``

snapshot_expiry_creation
------------------------

Adds ``expires\_at`` to container creation, allowing for override of a
snapshot’s expiry at creation time.

network_leases_location
-----------------------

Introductes a “Location” field in the leases list. This is used when
querying a cluster to show what node a particular lease was found on.

resources_cpu_socket
--------------------

Add Socket field to CPU resources in case we get out of order socket
information.

resources_gpu
-------------

Add a new GPU struct to the server resources, listing all usable GPUs on
the system.

resources_numa
--------------

Shows the NUMA node for all CPUs and GPUs.

kernel_features
---------------

Exposes the state of optional kernel features through the server
environment.

id_map_current
--------------

This introduces a new internal ``volatile.idmap.current`` key which is
used to track the current mapping for the container.

This effectively gives us:

-  ``volatile.last_state.idmap`` => On-disk idmap
-  ``volatile.idmap.current`` => Current kernel map
-  ``volatile.idmap.next`` => Next on-disk idmap

This is required to implement environments where the on-disk map isn’t
changed but the kernel map is (e.g. shiftfs).

event_location
--------------

Expose the location of the generation of API events.

storage_api_remote_volume_snapshots
-----------------------------------

This allows migrating storage volumes including their snapshots.

network_nat_address
-------------------

This introduces the ``ipv4.nat.address`` and ``ipv6.nat.address``
configuration keys for LXD bridges. Those keys control the source
address used for outbound traffic from the bridge.

container_nic_routes
--------------------

This introduces the ``ipv4.routes`` and ``ipv6.routes`` properties on
“nic” type devices. This allows adding static routes on host to
container’s nic.

rbac
----

Adds support for RBAC (role based access control). This introduces new
config keys:

-  rbac.api.url
-  rbac.api.key
-  rbac.api.expiry
-  rbac.agent.url
-  rbac.agent.username
-  rbac.agent.private_key
-  rbac.agent.public_key

cluster_internal_copy
---------------------

This makes it possible to do a normal “POST /1.0/containers” to copy a
container between cluster nodes with LXD internally detecting whether a
migration is required.

seccomp_notify
--------------

If the kernel supports seccomp-based syscall interception LXD can be
notified by a container that a registered syscall has been performed.
LXD can then decide to trigger various actions.

lxc_features
------------

This introduces the ``lxc_features`` section output from the
``lxc info`` command via the ``GET /1.0/`` route. It outputs the result
of checks for key features being present in the underlying LXC library.

container_nic_ipvlan
--------------------

This introduces the ``ipvlan`` “nic” device type.

network_vlan_sriov
------------------

This introduces VLAN (``vlan``) and MAC filtering
(``security.mac_filtering``) support for SR-IOV devices.

storage_cephfs
--------------

Add support for CEPHFS as a storage pool driver. This can only be used
for custom volumes, images and containers should be on CEPH (RBD)
instead.

container_nic_ipfilter
----------------------

This introduces container IP filtering (``security.ipv4_filtering`` and
``security.ipv6_filtering``) support for ``bridged`` nic devices.

resources_v2
------------

Rework the resources API at /1.0/resources, especially:

-  CPU

   -  Fix reporting to track sockets, cores and threads
   -  Track NUMA node per core
   -  Track base and turbo frequency per socket
   -  Track current frequency per core
   -  Add CPU cache information
   -  Export the CPU architecture
   -  Show online/offline status of threads

-  Memory

   -  Add hugepages tracking
   -  Track memory consumption per NUMA node too

-  GPU

   -  Split DRM information to separate struct
   -  Export device names and nodes in DRM struct
   -  Export device name and node in NVIDIA struct
   -  Add SR-IOV VF tracking

container_exec_user_group_cwd
-----------------------------

Adds support for specifying User, Group and Cwd during
``POST /1.0/containers/NAME/exec``.

container_syscall_intercept
---------------------------

Adds the ``security.syscalls.intercept.\*`` configuration keys to
control what system calls will be interecepted by LXD and processed with
elevated permissions.

container_disk_shift
--------------------

Adds the ``shift`` property on ``disk`` devices which controls the use
of the shiftfs overlay.

storage_shifted
---------------

Introduces a new ``security.shifted`` boolean on storage volumes.

Setting it to true will allow multiple isolated containers to attach the
same storage volume while keeping the filesystem writable from all of
them.

This makes use of shiftfs as an overlay filesystem.

resources_infiniband
--------------------

Export infiniband character device information (issm, umad, uverb) as
part of the resources API.

daemon_storage
--------------

This introduces two new configuration keys ``storage.images_volume`` and
``storage.backups_volume`` to allow for a storage volume on an existing
pool be used for storing the daemon-wide images and backups artifacts.

instances
---------

This introduces the concept of instances, of which currently the only
type is “container”.

image_types
-----------

This introduces support for a new Type field on images, indicating what
type of images they are.

resources_disk_sata
-------------------

Extends the disk resource API struct to include:

-  Proper detection of sata devices (type)
-  Device path
-  Drive RPM
-  Block size
-  Firmware version
-  Serial number

clustering_roles
----------------

This adds a new ``roles`` attribute to cluster entries, exposing a list
of roles that the member serves in the cluster.

images_expiry
-------------

This allows for editing of the expiry date on images.

resources_network_firmware
--------------------------

Adds a FirmwareVersion field to network card entries.

backup_compression_algorithm
----------------------------

This adds support for a ``compression_algorithm`` property when creating
a backup (``POST /1.0/containers/<name>/backups``).

Setting this property overrides the server default value
(``backups.compression_algorithm``).

ceph_data_pool_name
-------------------

This adds support for an optional argument (``ceph.osd.data_pool_name``)
when creating storage pools using Ceph RBD, when this argument is used
the pool will store it’s actual data in the pool specified with
``data_pool_name`` while keeping the metadata in the pool specified by
``pool_name``.

container_syscall_intercept_mount
---------------------------------

Adds the ``security.syscalls.intercept.mount``,
``security.syscalls.intercept.mount.allowed``, and
``security.syscalls.intercept.mount.shift`` configuration keys to
control whether and how the mount system call will be interecepted by
LXD and processed with elevated permissions.

compression_squashfs
--------------------

Adds support for importing/exporting of images/backups using SquashFS
file system format.

container_raw_mount
-------------------

This adds support for passing in raw mount options for disk devices.

container_nic_routed
--------------------

This introduces the ``routed`` “nic” device type.

container_syscall_intercept_mount_fuse
--------------------------------------

Adds the ``security.syscalls.intercept.mount.fuse`` key. It can be used
to redirect filesystem mounts to their fuse implementation. To this end,
set e.g. ``security.syscalls.intercept.mount.fuse=ext4=fuse2fs``.

container_disk_ceph
-------------------

This allows for existing a CEPH RDB or FS to be directly connected to a
LXD container.

virtual_machines
----------------

Add virtual machine support.

image_profiles
--------------

Allows a list of profiles to be applied to an image when launching a new
container.

clustering_architecture
-----------------------

This adds a new ``architecture`` attribute to cluster members which
indicates a cluster member’s architecture.

resources_disk_id
-----------------

Add a new device_id field in the disk entries on the resources API.

storage_lvm_stripes
-------------------

This adds the ability to use LVM stripes on normal volumes and thin pool
volumes.

vm_boot_priority
----------------

Adds a ``boot.priority`` property on nic and disk devices to control the
boot order.

unix_hotplug_devices
--------------------

Adds support for unix char and block device hotplugging.

api_filtering
-------------

Adds support for filtering the result of a GET request for instances and
images.

instance_nic_network
--------------------

Adds support for the ``network`` property on a NIC device to allow a NIC
to be linked to a managed network. This allows it to inherit some of the
network’s settings and allows better validation of IP settings.

clustering_sizing
-----------------

Support specifying a custom values for database voters and standbys. The
new ``cluster.max_voters`` and ``cluster.max_standby`` configuration
keys were introduced to specify to the ideal number of database voter
and standbys.

firewall_driver
---------------

Adds the ``Firewall`` property to the ServerEnvironment struct
indicating the firewall driver being used.

storage_lvm_vg_force_reuse
--------------------------

Introduces the ability to create a storage pool from an existing
non-empty volume group. This option should be used with care, as LXD can
then not guarantee that volume name conflicts won’t occur with non-LXD
created volumes in the same volume group. This could also potentially
lead to LXD deleting a non-LXD volume should name conflicts occur.

container_syscall_intercept_hugetlbfs
-------------------------------------

When mount syscall interception is enabled and hugetlbfs is specified as
an allowed filesystem type LXD will mount a separate hugetlbfs instance
for the container with the uid and gid mount options set to the
container’s root uid and gid. This ensures that processes in the
container can use hugepages.

limits_hugepages
----------------

This allows to limit the number of hugepages a container can use through
the hugetlb cgroup. This means the hugetlb cgroup needs to be available.
Note, that limiting hugepages is recommended when intercepting the mount
syscall for the hugetlbfs filesystem to avoid allowing the container to
exhaust the host’s hugepages resources.

container_nic_routed_gateway
----------------------------

This introduces the ``ipv4.gateway`` and ``ipv6.gateway`` NIC config
keys that can take a value of either “auto” or “none”. The default value
for the key if unspecified is “auto”. This will cause the current
behaviour of a default gateway being added inside the container and the
same gateway address being added to the host-side interface. If the
value is set to “none” then no default gateway nor will the address be
added to the host-side interface. This allows multiple routed NIC
devices to be added to a container.

projects_restrictions
---------------------

This introduces support for the ``restricted`` configuration key on
project, which can prevent the use of security-sensitive features in a
project.

custom_volume_snapshot_expiry
-----------------------------

This allows custom volume snapshots to expiry. Expiry dates can be set
individually, or by setting the ``snapshots.expiry`` config key on the
parent custom volume which then automatically applies to all created
snapshots.

volume_snapshot_scheduling
--------------------------

This adds support for custom volume snapshot scheduling. It introduces
two new configuration keys: ``snapshots.schedule`` and
``snapshots.pattern``. Snapshots can be created automatically up to
every minute.

trust_ca_certificates
---------------------

This allows for checking client certificates trusted by the provided CA
(``server.ca``). It can be enabled by setting
``core.trust_ca_certificates`` to true. If enabled, it will perform the
check, and bypass the trusted password if true. An exception will be
made if the connecting client certificate is in the provided CRL
(``ca.crl``). In this case, it will ask for the password.

snapshot_disk_usage
-------------------

This adds a new ``size`` field to the output of
``/1.0/instances/<name>/snapshots/<snapshot>`` which represents the disk
usage of the snapshot.

clustering_edit_roles
---------------------

This adds a writable endpoint for cluster members, allowing the editing
of their roles.

container_nic_routed_host_address
---------------------------------

This introduces the ``ipv4.host_address`` and ``ipv6.host_address`` NIC
config keys that can be used to control the host-side veth interface’s
IP addresses. This can be useful when using multiple routed NICs at the
same time and needing a predictable next-hop address to use.

This also alters the behaviour of ``ipv4.gateway`` and ``ipv6.gateway``
NIC config keys. When they are set to “auto” the container will have its
default gateway set to the value of ``ipv4.host_address`` or
``ipv6.host_address`` respectively.

The default values are:

``ipv4.host_address``: 169.254.0.1 ``ipv6.host_address``: fe80::1

This is backward compatible with the previous default behaviour.

container_nic_ipvlan_gateway
----------------------------

This introduces the ``ipv4.gateway`` and ``ipv6.gateway`` NIC config
keys that can take a value of either “auto” or “none”. The default value
for the key if unspecified is “auto”. This will cause the current
behaviour of a default gateway being added inside the container and the
same gateway address being added to the host-side interface. If the
value is set to “none” then no default gateway nor will the address be
added to the host-side interface. This allows multiple ipvlan NIC
devices to be added to a container.

resources_usb_pci
-----------------

This adds USB and PCI devices to the output of ``/1.0/resources``.

resources_cpu_threads_numa
--------------------------

This indicates that the numa_node field is now recorded per-thread
rather than per core as some hardware apparently puts threads in
different NUMA domains.

resources_cpu_core_die
----------------------

Exposes the ``die_id`` information on each core.

api_os
------

This introduces two new fields in ``/1.0``, ``os`` and ``os_version``.

Those are taken from the os-release data on the system.

container_nic_routed_host_table
-------------------------------

This introduces the ``ipv4.host_table`` and ``ipv6.host_table`` NIC
config keys that can be used to add static routes for the instance’s IPs
to a custom policy routing table by ID.

container_nic_ipvlan_host_table
-------------------------------

This introduces the ``ipv4.host_table`` and ``ipv6.host_table`` NIC
config keys that can be used to add static routes for the instance’s IPs
to a custom policy routing table by ID.

container_nic_ipvlan_mode
-------------------------

This introduces the ``mode`` NIC config key that can be used to switch
the ``ipvlan`` mode into either ``l2`` or ``l3s``. If not specified, the
default value is ``l3s`` (which is the old behavior).

In ``l2`` mode the ``ipv4.address`` and ``ipv6.address`` keys will
accept addresses in either CIDR or singular formats. If singular format
is used, the default subnet size is taken to be /24 and /64 for IPv4 and
IPv6 respectively.

In ``l2`` mode the ``ipv4.gateway`` and ``ipv6.gateway`` keys accept
only a singular IP address.

resources_system
----------------

This adds system information to the output of ``/1.0/resources``.

images_push_relay
-----------------

This adds the push and relay modes to image copy. It also introduces the
following new endpoint: - ``POST 1.0/images/<fingerprint>/export``

network_dns_search
------------------

This introduces the ``dns.search`` config option on networks.

container_nic_routed_limits
---------------------------

This introduces ``limits.ingress``, ``limits.egress`` and ``limits.max``
for routed NICs.

instance_nic_bridged_vlan
-------------------------

This introduces the ``vlan`` and ``vlan.tagged`` settings for
``bridged`` NICs.

``vlan`` specifies the untagged VLAN to join, and ``vlan.tagged`` is a
comma delimited list of tagged VLANs to join.

network_state_bond_bridge
-------------------------

This adds a “bridge” and “bond” section to the /1.0/networks/NAME/state
API.

Those contain additional state information relevant to those particular
types.

Bond:

-  Mode
-  Transmit hash
-  Up delay
-  Down delay
-  MII frequency
-  MII state
-  Lower devices

Bridge:

-  ID
-  Forward delay
-  STP mode
-  Default VLAN
-  VLAN filtering
-  Upper devices

resources_cpu_isolated
----------------------

Add an ``Isolated`` property on CPU threads to indicate if the thread is
physically ``Online`` but is configured not to accept tasks.

usedby_consistency
------------------

This extension indicates that UsedBy should now be consistent with
suitable ?project= and ?target= when appropriate.

The 5 entities that have UsedBy are: - Profiles - Projects - Networks -
Storage pools - Storage volumes

custom_block_volumes
--------------------

This adds support for creating and attaching custom block volumes to
instances. It introduces the new ``--type`` flag when creating custom
storage volumes, and accepts the values ``fs`` and ``block``.

clustering_failure_domains
--------------------------

This extension adds a new ``failure\_domain`` field to the
``PUT /1.0/cluster/<node>`` API, which can be used to set the failure
domain of a node.

container_syscall_filtering_allow_deny_syntax
---------------------------------------------

A number of new syscalls related container configuration keys were
updated.

-  ``security.syscalls.deny_default``
-  ``security.syscalls.deny_compat``
-  ``security.syscalls.deny``
-  ``security.syscalls.allow``

resources_gpu_mdev
------------------

Expose available mediated device profiles and devices in /1.0/resources.

console_vga_type
----------------

This extends the ``/1.0/console`` endpoint to take a ``?type=``
argument, which can be set to ``console`` (default) or ``vga`` (the new
type added by this extension).

When POST’ing to ``/1.0/<instance name>/console?type=vga`` the data
websocket returned by the operation in the metadata field will be a
bidirectional proxy attached to a SPICE unix socket of the target
virtual machine.

projects_limits_disk
--------------------

Add ``limits.disk`` to the available project configuration keys. If set,
it limits the total amount of disk space that instances volumes, custom
volumes and images volumes can use in the project.

network_type_macvlan
--------------------

Adds support for additional network type ``macvlan`` and adds ``parent``
configuration key for this network type to specify which parent
interface should be used for creating NIC device interfaces on top of.

Also adds ``network`` configuration key support for ``macvlan`` NICs to
allow them to specify the associated network of the same type that they
should use as the basis for the NIC device.

network_type_sriov
------------------

Adds support for additional network type ``sriov`` and adds ``parent``
configuration key for this network type to specify which parent
interface should be used for creating NIC device interfaces on top of.

Also adds ``network`` configuration key support for ``sriov`` NICs to
allow them to specify the associated network of the same type that they
should use as the basis for the NIC device.

container_syscall_intercept_bpf_devices
---------------------------------------

This adds support to intercept the bpf syscall in containers.
Specifically, it allows to manage device cgroup bpf programs.

network_type_ovn
----------------

Adds support for additional network type ``ovn`` with the ability to
specify a ``bridge`` type network as the ``parent``.

Introduces a new NIC device type of ``ovn`` which allows the ``network``
configuration key to specify which ``ovn`` type network they should
connect to.

Also introduces two new global config keys that apply to all ``ovn``
networks and NIC devices:

-  network.ovn.integration_bridge - the OVS integration bridge to use.
-  network.ovn.northbound_connection - the OVN northbound database
   connection string.

projects_networks
-----------------

Adds the ``features.networks`` config key to projects and the ability
for a project to hold networks.

projects_networks_restricted_uplinks
------------------------------------

Adds the ``restricted.networks.uplinks`` project config key to indicate
(as a comma delimited list) which networks the networks created inside
the project can use as their uplink network.

custom_volume_backup
--------------------

Add custom volume backup support.

This includes the following new endpoints (see `RESTful
API <rest-api.md>`__ for details):

-  ``GET /1.0/storage-pools/<pool>/<type>/<volume>/backups``
-  ``POST /1.0/storage-pools/<pool>/<type>/<volume>/backups``

-  ``GET /1.0/storage-pools/<pool>/<type>/<volume>/backups/<name>``
-  ``POST /1.0/storage-pools/<pool>/<type>/<volume>/backups/<name>``
-  ``DELETE /1.0/storage-pools/<pool>/<type>/<volume>/backups/<name>``

-  ``GET /1.0/storage-pools/<pool>/<type>/<volume>/backups/<name>/export``

The following existing endpoint has been modified:

-  ``POST /1.0/storage-pools/<pool>/<type>/<volume>`` accepts the new
   source type ``backup``

backup_override_name
--------------------

Adds ``Name`` field to ``InstanceBackupArgs`` to allow specifying a
different instance name when restoring a backup.

Adds ``Name`` and ``PoolName`` fields to ``StoragePoolVolumeBackupArgs``
to allow specifying a different volume name when restoring a custom
volume backup.

storage_rsync_compression
-------------------------

Adds ``rsync.compression`` config key to storage pools. This key can be
used to disable compression in rsync while migrating storage pools.

network_type_physical
---------------------

Adds support for additional network type ``physical`` that can be used
as an uplink for ``ovn`` networks.

The interface specified by ``parent`` on the ``physical`` network will
be connected to the ``ovn`` network’s gateway.

network_ovn_external_subnets
----------------------------

Adds support for ``ovn`` networks to use external subnets from uplink
networks.

Introduces the ``ipv4.routes`` and ``ipv6.routes`` setting on
``physical`` networks that defines the external routes allowed to be
used in child OVN networks in their ``ipv4.routes.external`` and
``ipv6.routes.external`` settings.

Introduces the ``restricted.networks.subnets`` project setting that
specifies which external subnets are allowed to be used by OVN networks
inside the project (if not set then all routes defined on the uplink
network are allowed).

network_ovn_nat
---------------

Adds support for ``ipv4.nat`` and ``ipv6.nat`` settings on ``ovn``
networks.

When creating the network if these settings are unspecified, and an
equivalent IP address is being generated for the subnet, then the
appropriate NAT setting will added set to ``true``.

If the setting is missing then the value is taken as ``false``.

network_ovn_external_routes_remove
----------------------------------

Removes the settings ``ipv4.routes.external`` and
``ipv6.routes.external`` from ``ovn`` networks.

The equivalent settings on the ``ovn`` NIC type can be used instead for
this, rather than having to specify them both at the network and NIC
level.

tpm_device_type
---------------

This introduces the ``tpm`` device type.

storage_zfs_clone_copy_rebase
-----------------------------

This introduces ``rebase`` as a value for zfs.clone_copy causing LXD to
track down any “image” dataset in the ancestry line and then perform
send/receive on top of that.

gpu_mdev
--------

This adds support for virtual GPUs. It introduces the ``mdev`` config
key for GPU devices which takes a supported mdev type,
e.g. i915-GVTg_V5_4.

resources_pci_iommu
-------------------

This adds the IOMMUGroup field for PCI entries in the resources API.

resources_network_usb
---------------------

Adds the usb_address field to the network card entries in the resources
API.

resources_disk_address
----------------------

Adds the usb_address and pci_address fields to the disk entries in the
resources API.

network_physical_ovn_ingress_mode
---------------------------------

Adds ``ovn.ingress_mode`` setting for ``physical`` networks.

Sets the method that OVN NIC external IPs will be advertised on uplink
network.

Either ``l2proxy`` (proxy ARP/NDP) or ``routed``.

network_ovn_dhcp
----------------

Adds ``ipv4.dhcp`` and ``ipv6.dhcp`` settings for ``ovn`` networks.

Allows DHCP (and RA for IPv6) to be disabled. Defaults to on.

network_physical_routes_anycast
-------------------------------

Adds ``ipv4.routes.anycast`` and ``ipv6.routes.anycast`` boolean
settings for ``physical`` networks. Defaults to false.

Allows OVN networks using physical network as uplink to relax external
subnet/route overlap detection when used with
``ovn.ingress_mode=routed``.

projects_limits_instances
-------------------------

Adds ``limits.instances`` to the available project configuration keys.
If set, it limits the total number of instances (VMs and containers)
that can be used in the project.

network_state_vlan
------------------

This adds a “vlan” section to the /1.0/networks/NAME/state API.

Those contain additional state information relevant to VLAN interfaces:
- lower_device - vid

instance_nic_bridged_port_isolation
-----------------------------------

This adds the ``security.port_isolation`` field for bridged NIC
instances.

instance_bulk_state_change
--------------------------

Adds the following endpoint for bulk state change (see `RESTful
API <rest-api.md>`__ for details):

-  ``PUT /1.0/instances``

network_gvrp
------------

This adds an optional ``gvrp`` property to ``macvlan`` and ``physical``
networks, and to ``ipvlan``, ``macvlan``, ``routed`` and ``physical``
NIC devices.

When set, this specifies whether the VLAN should be registered using
GARP VLAN Registration Protocol. Defaults to false.

instance_pool_move
------------------

This adds a ``pool`` field to the ``POST /1.0/instances/NAME`` API,
allowing for easy move of an instance root disk between pools.

gpu_sriov
---------

This adds support for SR-IOV enabled GPUs. It introduces the ``sriov``
gpu type property.

pci_device_type
---------------

This introduces the ``pci`` device type.

storage_volume_state
--------------------

Add new ``/1.0/storage-pools/POOL/volumes/VOLUME/state`` API endpoint to
get usage data on a volume.

network_acl
-----------

This adds the concept of network ACLs to API under the API endpoint
prefix ``/1.0/network-acls``.

migration_stateful
------------------

Add a new ``migration.stateful`` config key.

disk_state_quota
----------------

This introduces the ``size.state`` device config key on ``disk``
devices.

storage_ceph_features
---------------------

Adds a new ``ceph.rbd.features`` config key on storage pools to control
the RBD features used for new volumes.

projects_compression
--------------------

Adds new ``backups.compression_algorithm`` and
``images.compression_algorithm`` config keys which allows configuration
of backup and image compression per-project.

projects_images_remote_cache_expiry
-----------------------------------

Add new ``images.remote_cache_expiry`` config key to projects, allowing
for set number of days after which an unused cached remote image will be
flushed.

certificate_project
-------------------

Adds a new ``restricted`` property to certificates in the API as well as
``projects`` holding a list of project names that the certificate has
access to.

network_ovn_acl
---------------

Adds a new ``security.acls`` property to OVN networks and OVN NICs,
allowing Network ACLs to be applied.

projects_images_auto_update
---------------------------

Adds new ``images.auto_update_cached`` and
``images.auto_update_interval`` config keys which allows configuration
of images auto update in projects

projects_restricted_cluster_target
----------------------------------

Adds new ``restricted.cluster.target`` config key to project which
prevent the user from using –target to specify what cluster member to
place a workload on or the ability to move a workload between members.

images_default_architecture
---------------------------

Adds new ``images.default_architecture`` global config key and matching
per-project key which lets user tell LXD what architecture to go with
when no specific one is specified as part of the image request.

network_ovn_acl_defaults
------------------------

Adds new ``security.acls.default.{in,e}gress.action`` and
``security.acls.default.{in,e}gress.logged`` config keys for OVN
networks and NICs. This replaces the removed ACL ``default.action`` and
``default.logged`` keys.

gpu_mig
-------

This adds support for NVIDIA MIG. It introduces the ``mig`` gputype and
associaetd config keys.

project_usage
-------------

Adds an API endpoint to get current resource allocations in a project.
Accessible at API ``GET /1.0/projects/<name>/state``.

network_bridge_acl
------------------

Adds a new ``security.acls`` config key to ``bridge`` networks, allowing
Network ACLs to be applied.

Also adds ``security.acls.default.{in,e}gress.action`` and
``security.acls.default.{in,e}gress.logged`` config keys for specifying
the default behaviour for unmatched traffic.

warnings
--------

Warning API for LXD.

This includes the following endpoints (see `Restful API <rest-api.md>`__
for details):

-  ``GET /1.0/warnings``

-  ``GET /1.0/warnings/<uuid>``
-  ``PUT /1.0/warnings/<uuid>``
-  ``DELETE /1.0/warnings/<uuid>``

projects_restricted_backups_and_snapshots
-----------------------------------------

Adds new ``restricted.backups`` and ``restricted.snapshots`` config keys
to project which prevents the user from creation of backups and
snapshots.

clustering_join_token
---------------------

Adds ``POST /1.0/cluster/members`` API endpoint for requesting a join
token used when adding new cluster members without using the trust
password.

clustering_description
----------------------

Adds an editable description to the cluster members.

server_trusted_proxy
--------------------

This introduces support for ``core.https_trusted_proxy`` which has LXD
parse a HAProxy style connection header on such connections and if
present, will rewrite the request’s source address to that provided by
the proxy server.

clustering_update_cert
----------------------

Adds ``PUT /1.0/cluster/certificate`` endpoint for updating the cluster
certificate across the whole cluster

storage_api_project
-------------------

This adds support for copy/move custom storage volumes between projects.

server_instance_driver_operational
----------------------------------

This modifies the ``driver`` output for the ``/1.0`` endpoint to only
include drivers which are actually supported and operational on the
server (as opposed to being included in LXD but not operational on the
server).

server_supported_storage_drivers
--------------------------------

This adds supported storage driver info to server environment info.

event_lifecycle_requestor_address
---------------------------------

Adds a new address field to lifecycle requestor.

resources_gpu_usb
-----------------

Add a new USBAddress (usb_address) field to ResourcesGPUCard (GPU
entries) in the resources API.

clustering_evacuation
---------------------

Adds ``POST /1.0/cluster/members/<name>/state`` endpoint for evacuating
and restoring cluster members. It also adds the config keys
``cluster.evacuate`` and ``volatile.evacuate.origin`` for setting the
evacuation method (``auto``, ``stop`` or ``migrate``) and the origin of
any migrated instance respectively.

network_ovn_nat_address
-----------------------

This introduces the ``ipv4.nat.address`` and ``ipv6.nat.address``
configuration keys for LXD ``ovn`` networks. Those keys control the
source address used for outbound traffic from the OVN virtual network.
These keys can only be specified when the OVN network’s uplink network
has ``ovn.ingress_mode=routed``.

network_bgp
-----------

This introduces support for LXD acting as a BGP router to advertise
routes to ``bridge`` and ``ovn`` networks.

This comes with the addition to global config of:

-  ``core.bgp_address``
-  ``core.bgp_asn``
-  ``core.bgp_routerid``

The following network configurations keys (``bridge`` and ``physical``):

-  ``bgp.peers.<name>.address``
-  ``bgp.peers.<name>.asn``
-  ``bgp.peers.<name>.password``
-  ``bgp.ipv4.nexthop``
-  ``bgp.ipv6.nexthop``

And the following NIC-specific configuration keys (``bridged`` nictype):

-  ``ipv4.routes.external``
-  ``ipv6.routes.external``

network_forward
---------------

This introduces the networking address forward functionality. Allowing
for ``bridge`` and ``ovn`` networks to define external IP addresses that
can be forwarded to internal IP(s) inside their respective networks.

custom_volume_refresh
---------------------

Adds support for refresh during volume migration.

network_counters_errors_dropped
-------------------------------

This adds the received and sent errors as well as inbound and outbound
dropped packets to the network counters.

metrics
-------

This adds metrics to LXD. It returns metrics of running instances using
the OpenMetrics format.

This includes the following endpoints:

-  ``GET /1.0/metrics``
