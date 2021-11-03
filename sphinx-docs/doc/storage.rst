Storage configuration
=====================

LXD supports creating and managing storage pools and storage volumes.
General keys are top-level. Driver specific keys are namespaced by
driver name. Volume keys apply to any volume created in the pool unless
the value is overridden on a per-volume basis.

Storage pool configuration
--------------------------

+-----+-------+------------------+-------------+----------------------+
| Key | Type  | Condition        | Default     | Description          |
+=====+=======+==================+=============+======================+
| siz | strin | appropriate      | 0           | Size of the storage  |
| e   | g     | driver and       |             | pool in bytes        |
|     |       | source           |             | (suffixes            |
|     |       |                  |             | supported).          |
|     |       |                  |             | (Currently valid for |
|     |       |                  |             | loop based pools and |
|     |       |                  |             | zfs.)                |
+-----+-------+------------------+-------------+----------------------+
| sou | strin | -                | -           | Path to block device |
| rce | g     |                  |             | or loop file or      |
|     |       |                  |             | filesystem entry     |
+-----+-------+------------------+-------------+----------------------+
| btr | strin | btrfs driver     | user_subvol | Mount options for    |
| fs. | g     |                  | _rm_allowed | block devices        |
| mou |       |                  |             |                      |
| nt_ |       |                  |             |                      |
| opt |       |                  |             |                      |
| ion |       |                  |             |                      |
| s   |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| cep | strin | ceph driver      | ceph        | Name of the ceph     |
| h.c | g     |                  |             | cluster in which to  |
| lus |       |                  |             | create new storage   |
| ter |       |                  |             | pools.               |
| _na |       |                  |             |                      |
| me  |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| cep | bool  | ceph driver      | false       | Force using an osd   |
| h.o |       |                  |             | storage pool that is |
| sd. |       |                  |             | already in use by    |
| for |       |                  |             | another LXD          |
| ce_ |       |                  |             | instance.            |
| reu |       |                  |             |                      |
| se  |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| cep | strin | ceph driver      | 32          | Number of placement  |
| h.o | g     |                  |             | groups for the osd   |
| sd. |       |                  |             | storage pool.        |
| pg_ |       |                  |             |                      |
| num |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| cep | strin | ceph driver      | name of the | Name of the osd      |
| h.o | g     |                  | pool        | storage pool.        |
| sd. |       |                  |             |                      |
| poo |       |                  |             |                      |
| l_n |       |                  |             |                      |
| ame |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| cep | strin | ceph driver      | -           | Name of the osd data |
| h.o | g     |                  |             | pool.                |
| sd. |       |                  |             |                      |
| dat |       |                  |             |                      |
| a_p |       |                  |             |                      |
| ool |       |                  |             |                      |
| _na |       |                  |             |                      |
| me  |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| cep | strin | ceph driver      | true        | Whether to use RBD   |
| h.r | g     |                  |             | lightweight clones   |
| bd. |       |                  |             | rather than full     |
| clo |       |                  |             | dataset copies.      |
| ne_ |       |                  |             |                      |
| cop |       |                  |             |                      |
| y   |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| cep | strin | ceph driver      | layering    | Comma separate list  |
| h.r | g     |                  |             | of RBD features to   |
| bd. |       |                  |             | enable on the        |
| fea |       |                  |             | volumes.             |
| tur |       |                  |             |                      |
| es  |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| cep | strin | ceph driver      | admin       | The ceph user to use |
| h.u | g     |                  |             | when creating        |
| ser |       |                  |             | storage pools and    |
| .na |       |                  |             | volumes.             |
| me  |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| cep | strin | cephfs driver    | ceph        | Name of the ceph     |
| hfs | g     |                  |             | cluster in which to  |
| .cl |       |                  |             | create new storage   |
| ust |       |                  |             | pools.               |
| er_ |       |                  |             |                      |
| nam |       |                  |             |                      |
| e   |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| cep | strin | cephfs driver    | /           | The base path for    |
| hfs | g     |                  |             | the CEPHFS mount     |
| .pa |       |                  |             |                      |
| th  |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| cep | strin | cephfs driver    | admin       | The ceph user to use |
| hfs | g     |                  |             | when creating        |
| .us |       |                  |             | storage pools and    |
| er. |       |                  |             | volumes.             |
| nam |       |                  |             |                      |
| e   |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| lvm | strin | lvm driver       | LXDThinPool | Thin pool where      |
| .th | g     |                  |             | volumes are created. |
| inp |       |                  |             |                      |
| ool |       |                  |             |                      |
| _na |       |                  |             |                      |
| me  |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| lvm | bool  | lvm driver       | true        | Whether the storage  |
| .us |       |                  |             | pool uses a thinpool |
| e_t |       |                  |             | for logical volumes. |
| hin |       |                  |             |                      |
| poo |       |                  |             |                      |
| l   |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| lvm | strin | lvm driver       | name of the | Name of the volume   |
| .vg | g     |                  | pool        | group to create.     |
| _na |       |                  |             |                      |
| me  |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| lvm | bool  | lvm driver       | false       | Force using an       |
| .vg |       |                  |             | existing non-empty   |
| .fo |       |                  |             | volume group.        |
| rce |       |                  |             |                      |
| _re |       |                  |             |                      |
| use |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| vol | strin | lvm driver       | -           | Number of stripes to |
| ume | g     |                  |             | use for new volumes  |
| .lv |       |                  |             | (or thin pool        |
| m.s |       |                  |             | volume).             |
| tri |       |                  |             |                      |
| pes |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| vol | strin | lvm driver       | -           | Size of stripes to   |
| ume | g     |                  |             | use (at least 4096   |
| .lv |       |                  |             | bytes and multiple   |
| m.s |       |                  |             | of 512bytes).        |
| tri |       |                  |             |                      |
| pes |       |                  |             |                      |
| .si |       |                  |             |                      |
| ze  |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| rsy | strin | -                | 0 (no       | Specifies the upper  |
| nc. | g     |                  | limit)      | limit to be placed   |
| bwl |       |                  |             | on the socket I/O    |
| imi |       |                  |             | whenever rsync has   |
| t   |       |                  |             | to be used to        |
|     |       |                  |             | transfer storage     |
|     |       |                  |             | entities.            |
+-----+-------+------------------+-------------+----------------------+
| rsy | bool  | appropriate      | true        | Whether to use       |
| nc. |       | driver           |             | compression while    |
| com |       |                  |             | migrating storage    |
| pre |       |                  |             | pools.               |
| ssi |       |                  |             |                      |
| on  |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| vol | strin | -                | -           | Records the actual   |
| ati | g     |                  |             | source passed during |
| le. |       |                  |             | creating             |
| ini |       |                  |             | (e.g. /dev/sdb).     |
| tia |       |                  |             |                      |
| l_s |       |                  |             |                      |
| our |       |                  |             |                      |
| ce  |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| vol | strin | -                | true        | Whether the pool has |
| ati | g     |                  |             | been empty on        |
| le. |       |                  |             | creation time.       |
| poo |       |                  |             |                      |
| l.p |       |                  |             |                      |
| ris |       |                  |             |                      |
| tin |       |                  |             |                      |
| e   |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| vol | strin | block based      | ext4        | Filesystem to use    |
| ume | g     | driver (lvm)     |             | for new volumes      |
| .bl |       |                  |             |                      |
| ock |       |                  |             |                      |
| .fi |       |                  |             |                      |
| les |       |                  |             |                      |
| yst |       |                  |             |                      |
| em  |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| vol | strin | block based      | discard     | Mount options for    |
| ume | g     | driver (lvm)     |             | block devices        |
| .bl |       |                  |             |                      |
| ock |       |                  |             |                      |
| .mo |       |                  |             |                      |
| unt |       |                  |             |                      |
| _op |       |                  |             |                      |
| tio |       |                  |             |                      |
| ns  |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| vol | strin | appropriate      | unlimited   | Default volume size  |
| ume | g     | driver           | (10GB for   |                      |
| .si |       |                  | block)      |                      |
| ze  |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| vol | bool  | zfs driver       | false       | Remove snapshots as  |
| ume |       |                  |             | needed               |
| .zf |       |                  |             |                      |
| s.r |       |                  |             |                      |
| emo |       |                  |             |                      |
| ve_ |       |                  |             |                      |
| sna |       |                  |             |                      |
| psh |       |                  |             |                      |
| ots |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| vol | bool  | zfs driver       | false       | Use refquota instead |
| ume |       |                  |             | of quota for space.  |
| .zf |       |                  |             |                      |
| s.u |       |                  |             |                      |
| se_ |       |                  |             |                      |
| ref |       |                  |             |                      |
| quo |       |                  |             |                      |
| ta  |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| zfs | strin | zfs driver       | true        | Whether to use ZFS   |
| .cl | g     |                  |             | lightweight clones   |
| one |       |                  |             | rather than full     |
| _co |       |                  |             | dataset copies       |
| py  |       |                  |             | (boolean) or         |
|     |       |                  |             | “rebase” to copy     |
|     |       |                  |             | based on the initial |
|     |       |                  |             | image.               |
+-----+-------+------------------+-------------+----------------------+
| zfs | strin | zfs driver       | name of the | Name of the zpool    |
| .po | g     |                  | pool        |                      |
| ol_ |       |                  |             |                      |
| nam |       |                  |             |                      |
| e   |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+

Storage pool configuration keys can be set using the lxc tool with:

.. code:: bash

   lxc storage set [<remote>:]<pool> <key> <value>

Storage volume configuration
----------------------------

+-----+-------+------------------+-------------+----------------------+
| Key | Type  | Condition        | Default     | Description          |
+=====+=======+==================+=============+======================+
| siz | strin | appropriate      | same as     | Size of the storage  |
| e   | g     | driver           | volume.size | volume               |
+-----+-------+------------------+-------------+----------------------+
| blo | strin | block based      | same as     | Filesystem of the    |
| ck. | g     | driver           | volume.bloc | storage volume       |
| fil |       |                  | k.filesyste |                      |
| esy |       |                  | m           |                      |
| ste |       |                  |             |                      |
| m   |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| blo | strin | block based      | same as     | Mount options for    |
| ck. | g     | driver           | volume.bloc | block devices        |
| mou |       |                  | k.mount_opt |                      |
| nt_ |       |                  | ions        |                      |
| opt |       |                  |             |                      |
| ion |       |                  |             |                      |
| s   |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| sec | bool  | custom volume    | false       | Enable id shifting   |
| uri |       |                  |             | overlay (allows      |
| ty. |       |                  |             | attach by multiple   |
| shi |       |                  |             | isolated instances)  |
| fte |       |                  |             |                      |
| d   |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| sec | bool  | custom volume    | false       | Disable id mapping   |
| uri |       |                  |             | for the volume       |
| ty. |       |                  |             |                      |
| unm |       |                  |             |                      |
| app |       |                  |             |                      |
| ed  |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| lvm | strin | lvm driver       | -           | Number of stripes to |
| .st | g     |                  |             | use for new volumes  |
| rip |       |                  |             | (or thin pool        |
| es  |       |                  |             | volume).             |
+-----+-------+------------------+-------------+----------------------+
| lvm | strin | lvm driver       | -           | Size of stripes to   |
| .st | g     |                  |             | use (at least 4096   |
| rip |       |                  |             | bytes and multiple   |
| es. |       |                  |             | of 512bytes).        |
| siz |       |                  |             |                      |
| e   |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| sna | strin | custom volume    | -           | Controls when        |
| psh | g     |                  |             | snapshots are to be  |
| ots |       |                  |             | deleted (expects     |
| .ex |       |                  |             | expression like      |
| pir |       |                  |             | ``1M 2H 3d 4w 5m 6y` |
| y   |       |                  |             | `)                   |
+-----+-------+------------------+-------------+----------------------+
| sna | strin | custom volume    | -           | Cron expression      |
| psh | g     |                  |             | (``<minute> <hour> < |
| ots |       |                  |             | dom> <month> <dow>`` |
| .sc |       |                  |             | ),                   |
| hed |       |                  |             | or a comma separated |
| ule |       |                  |             | list of schedule     |
|     |       |                  |             | aliases              |
|     |       |                  |             | ``<@hourly> <@daily> |
|     |       |                  |             |  <@midnight> <@weekl |
|     |       |                  |             | y> <@monthly> <@annu |
|     |       |                  |             | ally> <@yearly>``    |
+-----+-------+------------------+-------------+----------------------+
| sna | strin | custom volume    | snap%d      | Pongo2 template      |
| psh | g     |                  |             | string which         |
| ots |       |                  |             | represents the       |
| .pa |       |                  |             | snapshot name (used  |
| tte |       |                  |             | for scheduled        |
| rn  |       |                  |             | snapshots and        |
|     |       |                  |             | unnamed snapshots)   |
+-----+-------+------------------+-------------+----------------------+
| zfs | strin | zfs driver       | same as     | Remove snapshots as  |
| .re | g     |                  | volume.zfs. | needed               |
| mov |       |                  | remove_snap |                      |
| e_s |       |                  | shots       |                      |
| nap |       |                  |             |                      |
| sho |       |                  |             |                      |
| ts  |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+
| zfs | strin | zfs driver       | same as     | Use refquota instead |
| .us | g     |                  | volume.zfs. | of quota for space   |
| e_r |       |                  | zfs_requota |                      |
| efq |       |                  |             |                      |
| uot |       |                  |             |                      |
| a   |       |                  |             |                      |
+-----+-------+------------------+-------------+----------------------+

Storage volume configuration keys can be set using the lxc tool with:

.. code:: bash

   lxc storage volume set [<remote>:]<pool> <volume> <key> <value>

Storage volume content types
----------------------------

Storage volumes can be either ``filesystem`` or ``block`` type.

Containers and container images are always going to be using
``filesystem``. Virtual machines and virtual machine images are always
going to be using ``block``.

Custom storage volumes can be either types with the default being
``filesystem``. Those custom storage volumes of type ``block`` can only
be attached to virtual machines.

Block custom storage volumes can be created with:

.. code:: bash

   lxc storage volume create [<remote>]:<pool> <name> --type=block

Where to store LXD data
-----------------------

Depending on the storage backends used, LXD can either share the
filesystem with its host or keep its data separate.

Sharing with the host
~~~~~~~~~~~~~~~~~~~~~

This is usually the most space efficient way to run LXD and possibly the
easiest to manage. It can be done with:

-  ``dir`` backend on any backing filesystem
-  ``btrfs`` backend if the host is btrfs and you point LXD to a
   dedicated subvolume
-  ``zfs`` backend if the host is zfs and you point LXD to a dedicated
   dataset on your zpool

Dedicated disk/partition
~~~~~~~~~~~~~~~~~~~~~~~~

In this mode, LXD’s storage will be completely independent from the
host. This can be done by having LXD use an empty partition on your main
disk or by having it use a full dedicated disk.

This is supported by all storage drivers except ``dir``, ``ceph`` and
``cephfs``.

Loop disk
~~~~~~~~~

If neither of the options above are possible for you, LXD can create a
loop file on your main drive and then have the selected storage driver
use that.

This is functionally similar to using a disk/partition but uses a large
file on your main drive instead. This comes at a performance penalty as
every writes need to go through the storage driver and then your main
drive’s filesystem. The loop files also usually cannot be shrunk. They
will grow up to the limit you select but deleting instances or images
will not cause the file to shrink.

Storage Backends and supported functions
----------------------------------------

Feature comparison
~~~~~~~~~~~~~~~~~~

LXD supports using ZFS, btrfs, LVM or just plain directories for storage
of images, instances and custom volumes. Where possible, LXD tries to
use the advanced features of each system to optimize operations.

========================================= ========= ===== === === ====
Feature                                   Directory Btrfs LVM ZFS CEPH
========================================= ========= ===== === === ====
Optimized image storage                   no        yes   yes yes yes
Optimized instance creation               no        yes   yes yes yes
Optimized snapshot creation               no        yes   yes yes yes
Optimized image transfer                  no        yes   no  yes yes
Optimized instance transfer               no        yes   no  yes yes
Copy on write                             no        yes   yes yes yes
Block based                               no        no    yes no  yes
Instant cloning                           no        yes   yes yes yes
Storage driver usable inside a container  yes       yes   no  no  no
Restore from older snapshots (not latest) yes       yes   yes no  yes
Storage quotas                            yes(*)    yes   yes yes no
========================================= ========= ===== === === ====

Recommended setup
~~~~~~~~~~~~~~~~~

The two best options for use with LXD are ZFS and btrfs. They have about
similar functionalities but ZFS is more reliable if available on your
particular platform.

Whenever possible, you should dedicate a full disk or partition to your
LXD storage pool. While LXD will let you create loop based storage, this
isn’t recommended for production use.

Similarly, the directory backend is to be considered as a last resort
option. It does support all main LXD features, but is terribly slow and
inefficient as it can’t perform instant copies or snapshots and so needs
to copy the entirety of the instance’s storage every time.

Security Considerations
~~~~~~~~~~~~~~~~~~~~~~~

Currently, the Linux Kernel may not apply mount options and silently
ignore them when a block-based filesystem (e.g. ``ext4``) is already
mounted with different options. This means when dedicated disk devices
are shared between different storage pools with different mount options
set, the second mount may not have the expected mount options. This
becomes security relevant, when e.g. one storage pool is supposed to
provide ``acl`` support and the second one is supposed to not provide
``acl`` support. For this reason it is currently recommended to either
have dedicated disk devices per storage pool or ensure that all storage
pools that share the same dedicated disk device use the same mount
options.

Optimized image storage
~~~~~~~~~~~~~~~~~~~~~~~

All backends but the directory backend have some kind of optimized image
storage format. This is used by LXD to make instance creation near
instantaneous by simply cloning a pre-made image volume rather than
unpack the image tarball from scratch.

As it would be wasteful to prepare such a volume on a storage pool that
may never be used with that image, the volume is generated on demand,
causing the first instance to take longer to create than subsequent
ones.

Optimized instance transfer
~~~~~~~~~~~~~~~~~~~~~~~~~~~

ZFS, btrfs and CEPH RBD have an internal send/receive mechanisms which
allow for optimized volume transfer. LXD uses those features to transfer
instances and snapshots between servers.

When such capabilities aren’t available, either because the storage
driver doesn’t support it or because the storage backend of the source
and target servers differ, LXD will fallback to using rsync to transfer
the individual files instead.

When rsync has to be used LXD allows to specify an upper limit on the
amount of socket I/O by setting the ``rsync.bwlimit`` storage pool
property to a non-zero value.

Default storage pool
--------------------

There is no concept of a default storage pool in LXD. Instead, the pool
to use for the instance’s root is treated as just another “disk” device
in LXD.

The device entry looks like:

.. code:: yaml

     root:
       type: disk
       path: /
       pool: default

And it can be directly set on an instance (“-s” option to “lxc launch”
and “lxc init”) or it can be set through LXD profiles.

That latter option is what the default LXD setup (through “lxd init”)
will do for you. The same can be done manually against any profile using
(for the “default” profile):

.. code:: bash

   lxc profile device add default root disk path=/ pool=default

I/O limits
----------

I/O limits in IOp/s or MB/s can be set on storage devices when attached
to an instance (see `Instances <instances.md>`__).

Those are applied through the Linux ``blkio`` cgroup controller which
makes it possible to restrict I/O at the disk level (but nothing finer
grained than that).

Because those apply to a whole physical disk rather than a partition or
path, the following restrictions apply:

-  Limits will not apply to filesystems that are backed by virtual
   devices (e.g. device mapper).
-  If a filesystem is backed by multiple block devices, each device will
   get the same limit.
-  If the instance is passed two disk devices that are each backed by
   the same disk, the limits of the two devices will be averaged.

It’s also worth noting that all I/O limits only apply to actual block
device access, so you will need to consider the filesystem’s own
overhead when setting limits. This also means that access to cached data
will not be affected by the limit.

Notes and examples
------------------

Directory
~~~~~~~~~

-  While this backend is fully functional, it’s also much slower than
   all the others due to it having to unpack images or do instant copies
   of instances, snapshots and images.
-  Quotas are supported with the directory backend when running on
   either ext4 or XFS with project quotas enabled at the filesystem
   level.

The following commands can be used to create directory storage pools
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

-  Create a new directory pool called “pool1”.

.. code:: bash

   lxc storage create pool1 dir

-  Use an existing directory for “pool2”.

.. code:: bash

   lxc storage create pool2 dir source=/data/lxd

CEPH
~~~~

-  Uses RBD images for images, then snapshots and clones to create
   instances and snapshots.
-  Due to the way copy-on-write works in RBD, parent filesystems can’t
   be removed until all children are gone. As a result, LXD will
   automatically prefix any removed but still referenced object with
   “zombie\_” and keep it until such time the references are gone and it
   can safely be removed.
-  Note that LXD will assume it has full control over the osd storage
   pool. It is recommended to not maintain any non-LXD owned filesystem
   entities in a LXD OSD storage pool since LXD might delete them.
-  Note that sharing the same osd storage pool between multiple LXD
   instances is not supported. LXD only allows sharing of an OSD storage
   pool between multiple LXD instances only for backup purposes of
   existing instances via ``lxd import``. In line with this, LXD
   requires the “ceph.osd.force_reuse” property to be set to true. If
   not set, LXD will refuse to reuse an osd storage pool it detected as
   being in use by another LXD instance.
-  When setting up a ceph cluster that LXD is going to use we recommend
   using ``xfs`` as the underlying filesystem for the storage entities
   that are used to hold OSD storage pools. Using ``ext4`` as the
   underlying filesystem for the storage entities is not recommended by
   Ceph upstream. You may see unexpected and erratic failures which are
   unrelated to LXD itself.

The following commands can be used to create Ceph storage pools
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

-  Create a osd storage pool named “pool1” in the CEPH cluster “ceph”.

.. code:: bash

   lxc storage create pool1 ceph

-  Create a osd storage pool named “pool1” in the CEPH cluster
   “my-cluster”.

.. code:: bash

   lxc storage create pool1 ceph ceph.cluster_name=my-cluster

-  Create a osd storage pool named “pool1” with the on-disk name
   “my-osd”.

.. code:: bash

   lxc storage create pool1 ceph ceph.osd.pool_name=my-osd

-  Use the existing osd storage pool “my-already-existing-osd”.

.. code:: bash

   lxc storage create pool1 ceph source=my-already-existing-osd

CEPHFS
~~~~~~

-  Can only be used for custom storage volumes
-  Supports snapshots if enabled on the server side

Btrfs
~~~~~

-  Uses a subvolume per instance, image and snapshot, creating btrfs
   snapshots when creating a new object.
-  btrfs can be used as a storage backend inside a container (nesting),
   so long as the parent container is itself on btrfs. (But see notes
   about btrfs quota via qgroups.)
-  btrfs supports storage quotas via qgroups. While btrfs qgroups are
   hierarchical, new subvolumes will not automatically be added to the
   qgroups of their parent subvolumes. This means that users can
   trivially escape any quotas that are set. If adherence to strict
   quotas is a necessity users should be mindful of this and maybe
   consider using a zfs storage pool with refquotas.

The following commands can be used to create BTRFS storage pools
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

-  Create loop-backed pool named “pool1”.

.. code:: bash

   lxc storage create pool1 btrfs

-  Create a new pool called “pool1” using an existing btrfs filesystem
   at ``/some/path``.

.. code:: bash

   lxc storage create pool1 btrfs source=/some/path

-  Create a new pool called “pool1” on ``/dev/sdX``.

.. code:: bash

   lxc storage create pool1 btrfs source=/dev/sdX

Growing a loop backed btrfs pool
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

LXD doesn’t let you directly grow a loop backed btrfs pool, but you can
do so with:

.. code:: bash

   sudo truncate -s +5G /var/lib/lxd/disks/<POOL>.img
   sudo losetup -c <LOOPDEV>
   sudo btrfs filesystem resize max /var/lib/lxd/storage-pools/<POOL>/

(NOTE: For users of the snap, use ``/var/snap/lxd/common/lxd/`` instead
of ``/var/lib/lxd/``)

LVM
~~~

-  Uses LVs for images, then LV snapshots for instances and instance
   snapshots.
-  The filesystem used for the LVs is ext4 (can be configured to use xfs
   instead).
-  By default, all LVM storage pools use an LVM thinpool in which
   logical volumes for all LXD storage entities (images, instances,
   etc.) are created. This behavior can be changed by setting
   “lvm.use_thinpool” to “false”. In this case, LXD will use normal
   logical volumes for all non-instance snapshot storage entities
   (images, instances, etc.). This means most storage operations will
   need to fallback to rsyncing since non-thinpool logical volumes do
   not support snapshots of snapshots. Note that this entails serious
   performance impacts for the LVM driver causing it to be close to the
   fallback DIR driver both in speed and storage usage. This option
   should only be chosen if the use-case renders it necessary.
-  For environments with high instance turn over (e.g continuous
   integration) it may be important to tweak the archival ``retain_min``
   and ``retain_days`` settings in ``/etc/lvm/lvm.conf`` to avoid
   slowdowns when interacting with LXD.

The following commands can be used to create LVM storage pools
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

-  Create a loop-backed pool named “pool1”. The LVM Volume Group will
   also be called “pool1”.

.. code:: bash

   lxc storage create pool1 lvm

-  Use the existing LVM Volume Group called “my-pool”

.. code:: bash

   lxc storage create pool1 lvm source=my-pool

-  Use the existing LVM Thinpool called “my-pool” in Volume Group
   “my-vg”.

.. code:: bash

   lxc storage create pool1 lvm source=my-vg lvm.thinpool_name=my-pool

-  Create a new pool named “pool1” on ``/dev/sdX``. The LVM Volume Group
   will also be called “pool1”.

.. code:: bash

   lxc storage create pool1 lvm source=/dev/sdX

-  Create a new pool called “pool1” using ``/dev/sdX`` with the LVM
   Volume Group called “my-pool”.

.. code:: bash

   lxc storage create pool1 lvm source=/dev/sdX lvm.vg_name=my-pool

ZFS
~~~

-  When LXD creates a ZFS pool, compression is enabled by default.
-  Uses ZFS filesystems for images, then snapshots and clones to create
   instances and snapshots.
-  Due to the way copy-on-write works in ZFS, parent filesystems can’t
   be removed until all children are gone. As a result, LXD will
   automatically rename any removed but still referenced object to a
   random deleted/ path and keep it until such time the references are
   gone and it can safely be removed.
-  ZFS as it is today doesn’t support delegating part of a pool to a
   container user. Upstream is actively working on this.
-  ZFS doesn’t support restoring from snapshots other than the latest
   one. You can however create new instances from older snapshots which
   makes it possible to confirm the snapshots is indeed what you want to
   restore before you remove the newer snapshots.

   LXD can be configured to automatically discard the newer snapshots
   during restore. This can be configured through the
   ``volume.zfs.remove_snapshots`` pool option.

   However note that instance copies use ZFS snapshots too, so you also
   cannot restore an instance to a snapshot taken before the last copy
   without having to also delete all its descendants.

   Copying the wanted snapshot into a new instance and then deleting the
   old instance does however work, at the cost of losing any other
   snapshot the instance may have had.

-  Note that LXD will assume it has full control over the ZFS pool or
   dataset. It is recommended to not maintain any non-LXD owned
   filesystem entities in a LXD zfs pool or dataset since LXD might
   delete them.
-  When quotas are used on a ZFS dataset LXD will set the ZFS “quota”
   property. In order to have LXD set the ZFS “refquota” property,
   either set “zfs.use_refquota” to “true” for the given dataset or set
   “volume.zfs.use_refquota” to true on the storage pool. The former
   option will make LXD use refquota only for the given storage volume
   the latter will make LXD use refquota for all storage volumes in the
   storage pool.
-  I/O quotas (IOps/MBs) are unlikely to affect ZFS filesystems very
   much. That’s because of ZFS being a port of a Solaris module (using
   SPL) and not a native Linux filesystem using the Linux VFS API which
   is where I/O limits are applied.

The following commands can be used to create ZFS storage pools
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

-  Create a loop-backed pool named “pool1”. The ZFS Zpool will also be
   called “pool1”.

.. code:: bash

   lxc storage create pool1 zfs

-  Create a loop-backed pool named “pool1” with the ZFS Zpool called
   “my-tank”.

.. code:: bash

   lxc storage create pool1 zfs zfs.pool_name=my-tank

-  Use the existing ZFS Zpool “my-tank”.

.. code:: bash

   lxc storage create pool1 zfs source=my-tank

-  Use the existing ZFS dataset “my-tank/slice”.

.. code:: bash

   lxc storage create pool1 zfs source=my-tank/slice

-  Create a new pool called “pool1” on ``/dev/sdX``. The ZFS Zpool will
   also be called “pool1”.

.. code:: bash

   lxc storage create pool1 zfs source=/dev/sdX

-  Create a new pool on ``/dev/sdX`` with the ZFS Zpool called
   “my-tank”.

.. code:: bash

   lxc storage create pool1 zfs source=/dev/sdX zfs.pool_name=my-tank

Growing a loop backed ZFS pool
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

LXD doesn’t let you directly grow a loop backed ZFS pool, but you can do
so with:

.. code:: bash

   sudo truncate -s +5G /var/lib/lxd/disks/<POOL>.img
   sudo zpool set autoexpand=on lxd
   sudo zpool online -e lxd /var/lib/lxd/disks/<POOL>.img
   sudo zpool set autoexpand=off lxd

(NOTE: For users of the snap, use ``/var/snap/lxd/common/lxd/`` instead
of ``/var/lib/lxd/``)

Enabling TRIM on existing pools
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

LXD will automatically enable trimming support on all newly created
pools on ZFS 0.8 or later.

This helps with the lifetime of SSDs by allowing better block re-use by
the controller. This also will allow freeing space on the root
filesystem when using a loop backed ZFS pool.

For systems which were upgraded from pre-0.8 to 0.8, this can be enabled
with a one time action of:

-  zpool upgrade ZPOOL-NAME
-  zpool set autotrim=on ZPOOL-NAME
-  zpool trim ZPOOL-NAME

This will make sure that TRIM is automatically issued in the future as
well as cause TRIM on all currently unused space.
