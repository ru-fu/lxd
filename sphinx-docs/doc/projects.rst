Project configuration
=====================

LXD supports projects as a way to split your LXD server. Each project
holds its own set of instances and may also have its own images and
profiles.

What a project contains is defined through the ``features``
configuration keys. When a feature is disabled, the project inherits
from the ``default`` project.

By default all new projects get the entire feature set, on upgrade,
existing projects do not get new features enabled.

The key/value configuration is namespaced with the following namespaces
currently supported:

-  ``features`` (What part of the project featureset is in use)
-  ``limits`` (Resource limits applied on containers and VMs belonging
   to the project)
-  ``user`` (free form key/value for user metadata)

+-------------+-------------+-------------+-------------+-------------+
| Key         | Type        | Condition   | Default     | Description |
+=============+=============+=============+=============+=============+
| backups.com | string      | -           | -           | Compression |
| pression_al |             |             |             | algorithm   |
| gorithm     |             |             |             | to use for  |
|             |             |             |             | backups     |
|             |             |             |             | (bzip2,     |
|             |             |             |             | gzip, lzma, |
|             |             |             |             | xz or none) |
|             |             |             |             | in the      |
|             |             |             |             | project     |
+-------------+-------------+-------------+-------------+-------------+
| features.im | boolean     | -           | true        | Separate    |
| ages        |             |             |             | set of      |
|             |             |             |             | images and  |
|             |             |             |             | image       |
|             |             |             |             | aliases for |
|             |             |             |             | the project |
+-------------+-------------+-------------+-------------+-------------+
| features.ne | boolean     | -           | false       | Separate    |
| tworks      |             |             |             | set of      |
|             |             |             |             | networks    |
|             |             |             |             | for the     |
|             |             |             |             | project     |
+-------------+-------------+-------------+-------------+-------------+
| features.pr | boolean     | -           | true        | Separate    |
| ofiles      |             |             |             | set of      |
|             |             |             |             | profiles    |
|             |             |             |             | for the     |
|             |             |             |             | project     |
+-------------+-------------+-------------+-------------+-------------+
| features.st | boolean     | -           | true        | Separate    |
| orage.volum |             |             |             | set of      |
| es          |             |             |             | storage     |
|             |             |             |             | volumes for |
|             |             |             |             | the project |
+-------------+-------------+-------------+-------------+-------------+
| images.auto | boolean     | -           | -           | Whether to  |
| _update_cac |             |             |             | automatical |
| hed         |             |             |             | ly          |
|             |             |             |             | update any  |
|             |             |             |             | image that  |
|             |             |             |             | LXD caches  |
+-------------+-------------+-------------+-------------+-------------+
| images.auto | integer     | -           | -           | Interval in |
| _update_int |             |             |             | hours at    |
| erval       |             |             |             | which to    |
|             |             |             |             | look for    |
|             |             |             |             | update to   |
|             |             |             |             | cached      |
|             |             |             |             | images (0   |
|             |             |             |             | disables    |
|             |             |             |             | it)         |
+-------------+-------------+-------------+-------------+-------------+
| images.comp | string      | -           | -           | Compression |
| ression_alg |             |             |             | algorithm   |
| orithm      |             |             |             | to use for  |
|             |             |             |             | images      |
|             |             |             |             | (bzip2,     |
|             |             |             |             | gzip, lzma, |
|             |             |             |             | xz or none) |
|             |             |             |             | in the      |
|             |             |             |             | project     |
+-------------+-------------+-------------+-------------+-------------+
| images.defa | string      | -           | -           | Default     |
| ult_archite |             |             |             | architectur |
| cture       |             |             |             | e           |
|             |             |             |             | which       |
|             |             |             |             | should be   |
|             |             |             |             | used in     |
|             |             |             |             | mixed       |
|             |             |             |             | architectur |
|             |             |             |             | e           |
|             |             |             |             | cluster     |
+-------------+-------------+-------------+-------------+-------------+
| images.remo | integer     | -           | -           | Number of   |
| te_cache_ex |             |             |             | days after  |
| piry        |             |             |             | which an    |
|             |             |             |             | unused      |
|             |             |             |             | cached      |
|             |             |             |             | remote      |
|             |             |             |             | image will  |
|             |             |             |             | be flushed  |
|             |             |             |             | in the      |
|             |             |             |             | project     |
+-------------+-------------+-------------+-------------+-------------+
| limits.cont | integer     | -           | -           | Maximum     |
| ainers      |             |             |             | number of   |
|             |             |             |             | containers  |
|             |             |             |             | that can be |
|             |             |             |             | created in  |
|             |             |             |             | the project |
+-------------+-------------+-------------+-------------+-------------+
| limits.cpu  | integer     | -           | -           | Maximum     |
|             |             |             |             | value for   |
|             |             |             |             | the sum of  |
|             |             |             |             | individual  |
|             |             |             |             | “limits.cpu |
|             |             |             |             | ”           |
|             |             |             |             | configs set |
|             |             |             |             | on the      |
|             |             |             |             | instances   |
|             |             |             |             | of the      |
|             |             |             |             | project     |
+-------------+-------------+-------------+-------------+-------------+
| limits.disk | string      | -           | -           | Maximum     |
|             |             |             |             | value of    |
|             |             |             |             | aggregate   |
|             |             |             |             | disk space  |
|             |             |             |             | used by all |
|             |             |             |             | instances   |
|             |             |             |             | volumes,    |
|             |             |             |             | custom      |
|             |             |             |             | volumes and |
|             |             |             |             | images of   |
|             |             |             |             | the project |
+-------------+-------------+-------------+-------------+-------------+
| limits.inst | integer     | -           | -           | Maximum     |
| ances       |             |             |             | number of   |
|             |             |             |             | total       |
|             |             |             |             | instances   |
|             |             |             |             | that can be |
|             |             |             |             | created in  |
|             |             |             |             | the project |
+-------------+-------------+-------------+-------------+-------------+
| limits.memo | string      | -           | -           | Maximum     |
| ry          |             |             |             | value for   |
|             |             |             |             | the sum of  |
|             |             |             |             | individual  |
|             |             |             |             | “limits.mem |
|             |             |             |             | ory”        |
|             |             |             |             | configs set |
|             |             |             |             | on the      |
|             |             |             |             | instances   |
|             |             |             |             | of the      |
|             |             |             |             | project     |
+-------------+-------------+-------------+-------------+-------------+
| limits.netw | integer     | -           | -           | Maximum     |
| orks        |             |             |             | value for   |
|             |             |             |             | the number  |
|             |             |             |             | of networks |
|             |             |             |             | this        |
|             |             |             |             | project can |
|             |             |             |             | have        |
+-------------+-------------+-------------+-------------+-------------+
| limits.proc | integer     | -           | -           | Maximum     |
| esses       |             |             |             | value for   |
|             |             |             |             | the sum of  |
|             |             |             |             | individual  |
|             |             |             |             | “limits.pro |
|             |             |             |             | cesses”     |
|             |             |             |             | configs set |
|             |             |             |             | on the      |
|             |             |             |             | instances   |
|             |             |             |             | of the      |
|             |             |             |             | project     |
+-------------+-------------+-------------+-------------+-------------+
| limits.virt | integer     | -           | -           | Maximum     |
| ual-machine |             |             |             | number of   |
| s           |             |             |             | VMs that    |
|             |             |             |             | can be      |
|             |             |             |             | created in  |
|             |             |             |             | the project |
+-------------+-------------+-------------+-------------+-------------+
| restricted  | boolean     | -           | false       | Block       |
|             |             |             |             | access to   |
|             |             |             |             | security-se |
|             |             |             |             | nsitive     |
|             |             |             |             | features    |
+-------------+-------------+-------------+-------------+-------------+
| restricted. | string      | -           | block       | Prevents    |
| backups     |             |             |             | the         |
|             |             |             |             | creation of |
|             |             |             |             | any         |
|             |             |             |             | instance or |
|             |             |             |             | volume      |
|             |             |             |             | backups.    |
+-------------+-------------+-------------+-------------+-------------+
| restricted. | string      | -           | block       | Prevents    |
| cluster.tar |             |             |             | direct      |
| get         |             |             |             | targeting   |
|             |             |             |             | of cluster  |
|             |             |             |             | members     |
|             |             |             |             | when        |
|             |             |             |             | creating or |
|             |             |             |             | moving      |
|             |             |             |             | instances.  |
+-------------+-------------+-------------+-------------+-------------+
| restricted. | string      | -           | block       | Prevents    |
| containers. |             |             |             | use of      |
| lowlevel    |             |             |             | low-level   |
|             |             |             |             | container   |
|             |             |             |             | options     |
|             |             |             |             | like        |
|             |             |             |             | raw.lxc,    |
|             |             |             |             | raw.idmap,  |
|             |             |             |             | volatile,   |
|             |             |             |             | etc.        |
+-------------+-------------+-------------+-------------+-------------+
| restricted. | string      | -           | block       | Prevents    |
| containers. |             |             |             | setting     |
| nesting     |             |             |             | security.ne |
|             |             |             |             | sting=true. |
+-------------+-------------+-------------+-------------+-------------+
| restricted. | string      | -           | unprivilige | If          |
| containers. |             |             | d           | “unprivilig |
| privilege   |             |             |             | ed”,        |
|             |             |             |             | prevents    |
|             |             |             |             | setting     |
|             |             |             |             | security.pr |
|             |             |             |             | ivileged=tr |
|             |             |             |             | ue.         |
|             |             |             |             | If          |
|             |             |             |             | “isolated”, |
|             |             |             |             | prevents    |
|             |             |             |             | setting     |
|             |             |             |             | security.pr |
|             |             |             |             | ivileged=tr |
|             |             |             |             | ue          |
|             |             |             |             | and also    |
|             |             |             |             | security.id |
|             |             |             |             | map.isolate |
|             |             |             |             | d=true.     |
|             |             |             |             | If “allow”, |
|             |             |             |             | no          |
|             |             |             |             | restriction |
|             |             |             |             | apply.      |
+-------------+-------------+-------------+-------------+-------------+
| restricted. | string      | -           | managed     | If “block”  |
| devices.dis |             |             |             | prevent use |
| k           |             |             |             | of disk     |
|             |             |             |             | devices     |
|             |             |             |             | except the  |
|             |             |             |             | root one.   |
|             |             |             |             | If          |
|             |             |             |             | “managed”   |
|             |             |             |             | allow use   |
|             |             |             |             | of disk     |
|             |             |             |             | devices     |
|             |             |             |             | only if     |
|             |             |             |             | “pool=” is  |
|             |             |             |             | set. If     |
|             |             |             |             | “allow”, no |
|             |             |             |             | restriction |
|             |             |             |             | s           |
|             |             |             |             | apply.      |
+-------------+-------------+-------------+-------------+-------------+
| restricted. | string      | -           | block       | Prevents    |
| devices.gpu |             |             |             | use of      |
|             |             |             |             | devices of  |
|             |             |             |             | type “gpu”  |
+-------------+-------------+-------------+-------------+-------------+
| restricted. | string      | -           | block       | Prevents    |
| devices.inf |             |             |             | use of      |
| iniband     |             |             |             | devices of  |
|             |             |             |             | type        |
|             |             |             |             | “infiniband |
|             |             |             |             | ”           |
+-------------+-------------+-------------+-------------+-------------+
| restricted. | string      | -           | managed     | If “block”  |
| devices.nic |             |             |             | prevent use |
|             |             |             |             | of all      |
|             |             |             |             | network     |
|             |             |             |             | devices. If |
|             |             |             |             | “managed”   |
|             |             |             |             | allow use   |
|             |             |             |             | of network  |
|             |             |             |             | devices     |
|             |             |             |             | only if     |
|             |             |             |             | “network=”  |
|             |             |             |             | is set. If  |
|             |             |             |             | “allow”, no |
|             |             |             |             | restriction |
|             |             |             |             | s           |
|             |             |             |             | apply.      |
+-------------+-------------+-------------+-------------+-------------+
| restricted. | string      | -           | block       | Prevents    |
| devices.pci |             |             |             | use of      |
|             |             |             |             | devices of  |
|             |             |             |             | type “pci”  |
+-------------+-------------+-------------+-------------+-------------+
| restricted. | string      | -           | block       | Prevents    |
| devices.pro |             |             |             | use of      |
| xy          |             |             |             | devices of  |
|             |             |             |             | type        |
|             |             |             |             | “proxy”     |
+-------------+-------------+-------------+-------------+-------------+
| restricted. | string      | -           | block       | Prevents    |
| devices.uni |             |             |             | use of      |
| x-block     |             |             |             | devices of  |
|             |             |             |             | type        |
|             |             |             |             | “unix-block |
|             |             |             |             | ”           |
+-------------+-------------+-------------+-------------+-------------+
| restricted. | string      | -           | block       | Prevents    |
| devices.uni |             |             |             | use of      |
| x-char      |             |             |             | devices of  |
|             |             |             |             | type        |
|             |             |             |             | “unix-char” |
+-------------+-------------+-------------+-------------+-------------+
| restricted. | string      | -           | block       | Prevents    |
| devices.uni |             |             |             | use of      |
| x-hotplug   |             |             |             | devices of  |
|             |             |             |             | type        |
|             |             |             |             | “unix-hotpl |
|             |             |             |             | ug”         |
+-------------+-------------+-------------+-------------+-------------+
| restricted. | string      | -           | block       | Prevents    |
| devices.usb |             |             |             | use of      |
|             |             |             |             | devices of  |
|             |             |             |             | type “usb”  |
+-------------+-------------+-------------+-------------+-------------+
| restricted. | string      | -           | block       | Comma       |
| networks.su |             |             |             | delimited   |
| bnets       |             |             |             | list of     |
|             |             |             |             | network     |
|             |             |             |             | subnets     |
|             |             |             |             | from the    |
|             |             |             |             | uplink      |
|             |             |             |             | networks    |
|             |             |             |             | (in the     |
|             |             |             |             | form        |
|             |             |             |             | ``<uplink>: |
|             |             |             |             | <subnet>``) |
|             |             |             |             | that are    |
|             |             |             |             | allocated   |
|             |             |             |             | for use in  |
|             |             |             |             | this        |
|             |             |             |             | project     |
+-------------+-------------+-------------+-------------+-------------+
| restricted. | string      | -           | block       | Comma       |
| networks.up |             |             |             | delimited   |
| links       |             |             |             | list of     |
|             |             |             |             | network     |
|             |             |             |             | names that  |
|             |             |             |             | can be used |
|             |             |             |             | as uplinks  |
|             |             |             |             | for         |
|             |             |             |             | networks in |
|             |             |             |             | this        |
|             |             |             |             | project     |
+-------------+-------------+-------------+-------------+-------------+
| restricted. | string      | -           | block       | Prevents    |
| snapshots   |             |             |             | the         |
|             |             |             |             | creation of |
|             |             |             |             | any         |
|             |             |             |             | instance or |
|             |             |             |             | volume      |
|             |             |             |             | snapshots.  |
+-------------+-------------+-------------+-------------+-------------+
| restricted. | string      | -           | block       | Prevents    |
| virtual-mac |             |             |             | use of      |
| hines.lowle |             |             |             | low-level   |
| vel         |             |             |             | virtual-mac |
|             |             |             |             | hine        |
|             |             |             |             | options     |
|             |             |             |             | like        |
|             |             |             |             | raw.qemu,   |
|             |             |             |             | volatile,   |
|             |             |             |             | etc.        |
+-------------+-------------+-------------+-------------+-------------+

Those keys can be set using the lxc tool with:

.. code:: bash

   lxc project set <project> <key> <value>

Project limits
--------------

Note that to be able to set one of the ``limits.*`` config keys, **all**
instances in the project **must** have that same config key defined,
either directly or via a profile.

In addition to that:

-  The ``limits.cpu`` config key also requires that CPU pinning is
   **not** used.
-  The ``limits.memory`` config key must be set to an absolute value,
   **not** a percentage.

The ``limits.*`` config keys defined on a project act as a hard upper
bound for the **aggregate** value of the individual ``limits.*`` config
keys defined on the project’s instances, either directly or via
profiles.

For example, setting the project’s ``limits.memory`` config key to
``50GB`` means that the sum of the individual values of all
``limits.memory`` config keys defined on the project’s instances will be
kept under ``50GB``. Trying to create or modify an instance assigning it
a ``limits.memory`` value that would make the total sum exceed ``50GB``,
will result in an error.

Similarly, setting the project’s ``limits.cpu`` config key to ``100``,
means that the **sum** of individual ``limits.cpu`` values will be kept
below ``100``.

Project restrictions
--------------------

If the ``restricted`` config key is set to ``true``, then the instances
of the project won’t be able to access security-sensitive features, such
as container nesting, raw LXC configuration, etc.

The exact set of features that the ``restricted`` config key blocks may
grow across LXD releases, as more features are added that are considered
security-sensitive.

Using the various ``restricted.*`` sub-keys, it’s possible to pick
individual features which would be normally blocked by ``restricted``
and allow them, so they can be used by instances of the project.

For example:

.. code:: bash

   lxc project set <project> restricted=true
   lxc project set <project> restricted.containers.nesting=allow

will block all security-sensitive features **except** container nesting.

Each security-sensitive feature has an associated ``restricted.*``
project config sub-key whose default value needs to be explicitly
changed if you want for that feature to be allowed it in the project.

Note that changing the value of a specific ``restricted.*`` config key
has an effect only if the top-level ``restricted`` key itself is
currently set to ``true``. If ``restricted`` is set to ``false``,
changing a ``restricted.*`` sub-key is effectively a no-op.

Most ``'restricted.*`` config keys are binary switches that can be set
to either ``block`` (the default) or ``allow``. However some of them
support other values for more fine-grained control.

Setting all ``restricted.*`` keys to ``allow`` is effectively equivalent
to setting ``restricted`` itself to ``false``.
