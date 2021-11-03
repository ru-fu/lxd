.. role:: raw-latex(raw)
   :format: latex
..

Instance configuration
======================

Properties
----------

The following are direct instance properties and can’t be part of a
profile:

-  ``name``
-  ``architecture``

Name is the instance name and can only be changed by renaming the
instance.

Valid instance names must:

-  Be between 1 and 63 characters long
-  Be made up exclusively of letters, numbers and dashes from the ASCII
   table
-  Not start with a digit or a dash
-  Not end with a dash

This requirement is so that the instance name may properly be used in
DNS records, on the filesystem, in various security profiles as well as
the hostname of the instance itself.

Key/value configuration
-----------------------

The key/value configuration is namespaced with the following namespaces
currently supported:

-  ``boot`` (boot related options, timing, dependencies, …)
-  ``environment`` (environment variables)
-  ``image`` (copy of the image properties at time of creation)
-  ``limits`` (resource limits)
-  ``nvidia`` (NVIDIA and CUDA configuration)
-  ``raw`` (raw instance configuration overrides)
-  ``security`` (security policies)
-  ``user`` (storage for user properties, searchable)
-  ``volatile`` (used internally by LXD to store internal data specific
   to an instance)

The currently supported keys are:

+---+-----+---------+---------------+---------------+---------------+
| K | Typ | Default | Live update   | Condition     | Description   |
| e | e   |         |               |               |               |
| y |     |         |               |               |               |
+===+=====+=========+===============+===============+===============+
| b | boo | -       | n/a           | -             | Always start  |
| o | lea |         |               |               | the instance  |
| o | n   |         |               |               | when LXD      |
| t |     |         |               |               | starts (if    |
| . |     |         |               |               | not set,      |
| a |     |         |               |               | restore last  |
| u |     |         |               |               | state)        |
| t |     |         |               |               |               |
| o |     |         |               |               |               |
| s |     |         |               |               |               |
| t |     |         |               |               |               |
| a |     |         |               |               |               |
| r |     |         |               |               |               |
| t |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| b | int | 0       | n/a           | -             | Number of     |
| o | ege |         |               |               | seconds to    |
| o | r   |         |               |               | wait after    |
| t |     |         |               |               | the instance  |
| . |     |         |               |               | started       |
| a |     |         |               |               | before        |
| u |     |         |               |               | starting the  |
| t |     |         |               |               | next one      |
| o |     |         |               |               |               |
| s |     |         |               |               |               |
| t |     |         |               |               |               |
| a |     |         |               |               |               |
| r |     |         |               |               |               |
| t |     |         |               |               |               |
| . |     |         |               |               |               |
| d |     |         |               |               |               |
| e |     |         |               |               |               |
| l |     |         |               |               |               |
| a |     |         |               |               |               |
| y |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| b | int | 0       | n/a           | -             | What order to |
| o | ege |         |               |               | start the     |
| o | r   |         |               |               | instances in  |
| t |     |         |               |               | (starting     |
| . |     |         |               |               | with highest) |
| a |     |         |               |               |               |
| u |     |         |               |               |               |
| t |     |         |               |               |               |
| o |     |         |               |               |               |
| s |     |         |               |               |               |
| t |     |         |               |               |               |
| a |     |         |               |               |               |
| r |     |         |               |               |               |
| t |     |         |               |               |               |
| . |     |         |               |               |               |
| p |     |         |               |               |               |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| o |     |         |               |               |               |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| t |     |         |               |               |               |
| y |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| b | int | 30      | yes           | -             | Seconds to    |
| o | ege |         |               |               | wait for      |
| o | r   |         |               |               | instance to   |
| t |     |         |               |               | shutdown      |
| . |     |         |               |               | before it is  |
| h |     |         |               |               | force stopped |
| o |     |         |               |               |               |
| s |     |         |               |               |               |
| t |     |         |               |               |               |
| _ |     |         |               |               |               |
| s |     |         |               |               |               |
| h |     |         |               |               |               |
| u |     |         |               |               |               |
| t |     |         |               |               |               |
| d |     |         |               |               |               |
| o |     |         |               |               |               |
| w |     |         |               |               |               |
| n |     |         |               |               |               |
| _ |     |         |               |               |               |
| t |     |         |               |               |               |
| i |     |         |               |               |               |
| m |     |         |               |               |               |
| e |     |         |               |               |               |
| o |     |         |               |               |               |
| u |     |         |               |               |               |
| t |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| b | int | 0       | n/a           | -             | What order to |
| o | ege |         |               |               | shutdown the  |
| o | r   |         |               |               | instances     |
| t |     |         |               |               | (starting     |
| . |     |         |               |               | with highest) |
| s |     |         |               |               |               |
| t |     |         |               |               |               |
| o |     |         |               |               |               |
| p |     |         |               |               |               |
| . |     |         |               |               |               |
| p |     |         |               |               |               |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| o |     |         |               |               |               |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| t |     |         |               |               |               |
| y |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| c | str | auto    | n/a           | -             | What to do    |
| l | ing |         |               |               | when          |
| u |     |         |               |               | evacuating    |
| s |     |         |               |               | the instance  |
| t |     |         |               |               | (auto,        |
| e |     |         |               |               | migrate, or   |
| r |     |         |               |               | stop)         |
| . |     |         |               |               |               |
| e |     |         |               |               |               |
| v |     |         |               |               |               |
| a |     |         |               |               |               |
| c |     |         |               |               |               |
| u |     |         |               |               |               |
| a |     |         |               |               |               |
| t |     |         |               |               |               |
| e |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| e | str | -       | yes (exec)    | -             | key/value     |
| n | ing |         |               |               | environment   |
| v |     |         |               |               | variables to  |
| i |     |         |               |               | export to the |
| r |     |         |               |               | instance and  |
| o |     |         |               |               | set on exec   |
| n |     |         |               |               |               |
| m |     |         |               |               |               |
| e |     |         |               |               |               |
| n |     |         |               |               |               |
| t |     |         |               |               |               |
| . |     |         |               |               |               |
| \ |     |         |               |               |               |
| * |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| l | str | -       | yes           | -             | Number or     |
| i | ing |         |               |               | range of CPUs |
| m |     |         |               |               | to expose to  |
| i |     |         |               |               | the instance  |
| t |     |         |               |               | (defaults to  |
| s |     |         |               |               | 1 CPU for     |
| . |     |         |               |               | VMs)          |
| c |     |         |               |               |               |
| p |     |         |               |               |               |
| u |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| l | str | 100%    | yes           | container     | How much of   |
| i | ing |         |               |               | the CPU can   |
| m |     |         |               |               | be used. Can  |
| i |     |         |               |               | be a          |
| t |     |         |               |               | percentage    |
| s |     |         |               |               | (e.g. 50%)    |
| . |     |         |               |               | for a soft    |
| c |     |         |               |               | limit or hard |
| p |     |         |               |               | a chunk of    |
| u |     |         |               |               | time          |
| . |     |         |               |               | (25ms/100ms)  |
| a |     |         |               |               |               |
| l |     |         |               |               |               |
| l |     |         |               |               |               |
| o |     |         |               |               |               |
| w |     |         |               |               |               |
| a |     |         |               |               |               |
| n |     |         |               |               |               |
| c |     |         |               |               |               |
| e |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| l | int | 10      | yes           | container     | CPU           |
| i | ege | (maximu |               |               | scheduling    |
| m | r   | m)      |               |               | priority      |
| i |     |         |               |               | compared to   |
| t |     |         |               |               | other         |
| s |     |         |               |               | instances     |
| . |     |         |               |               | sharing the   |
| c |     |         |               |               | same CPUs     |
| p |     |         |               |               | (overcommit)  |
| u |     |         |               |               | (integer      |
| . |     |         |               |               | between 0 and |
| p |     |         |               |               | 10)           |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| o |     |         |               |               |               |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| t |     |         |               |               |               |
| y |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| l | int | 5       | yes           | -             | When under    |
| i | ege | (medium |               |               | load, how     |
| m | r   | )       |               |               | much priority |
| i |     |         |               |               | to give to    |
| t |     |         |               |               | the           |
| s |     |         |               |               | instance’s    |
| . |     |         |               |               | I/O requests  |
| d |     |         |               |               | (integer      |
| i |     |         |               |               | between 0 and |
| s |     |         |               |               | 10)           |
| k |     |         |               |               |               |
| . |     |         |               |               |               |
| p |     |         |               |               |               |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| o |     |         |               |               |               |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| t |     |         |               |               |               |
| y |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| l | str | -       | yes           | container     | Fixed value   |
| i | ing |         |               |               | in bytes      |
| m |     |         |               |               | (various      |
| i |     |         |               |               | suffixes      |
| t |     |         |               |               | supported,    |
| s |     |         |               |               | see below) to |
| . |     |         |               |               | limit number  |
| h |     |         |               |               | of 64 KB      |
| u |     |         |               |               | hugepages     |
| g |     |         |               |               | (Available    |
| e |     |         |               |               | hugepage      |
| p |     |         |               |               | sizes are     |
| a |     |         |               |               | architecture  |
| g |     |         |               |               | dependent.)   |
| e |     |         |               |               |               |
| s |     |         |               |               |               |
| . |     |         |               |               |               |
| 6 |     |         |               |               |               |
| 4 |     |         |               |               |               |
| K |     |         |               |               |               |
| B |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| l | str | -       | yes           | container     | Fixed value   |
| i | ing |         |               |               | in bytes      |
| m |     |         |               |               | (various      |
| i |     |         |               |               | suffixes      |
| t |     |         |               |               | supported,    |
| s |     |         |               |               | see below) to |
| . |     |         |               |               | limit number  |
| h |     |         |               |               | of 1 MB       |
| u |     |         |               |               | hugepages     |
| g |     |         |               |               | (Available    |
| e |     |         |               |               | hugepage      |
| p |     |         |               |               | sizes are     |
| a |     |         |               |               | architecture  |
| g |     |         |               |               | dependent.)   |
| e |     |         |               |               |               |
| s |     |         |               |               |               |
| . |     |         |               |               |               |
| 1 |     |         |               |               |               |
| M |     |         |               |               |               |
| B |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| l | str | -       | yes           | container     | Fixed value   |
| i | ing |         |               |               | in bytes      |
| m |     |         |               |               | (various      |
| i |     |         |               |               | suffixes      |
| t |     |         |               |               | supported,    |
| s |     |         |               |               | see below) to |
| . |     |         |               |               | limit number  |
| h |     |         |               |               | of 2 MB       |
| u |     |         |               |               | hugepages     |
| g |     |         |               |               | (Available    |
| e |     |         |               |               | hugepage      |
| p |     |         |               |               | sizes are     |
| a |     |         |               |               | architecture  |
| g |     |         |               |               | dependent.)   |
| e |     |         |               |               |               |
| s |     |         |               |               |               |
| . |     |         |               |               |               |
| 2 |     |         |               |               |               |
| M |     |         |               |               |               |
| B |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| l | str | -       | yes           | container     | Fixed value   |
| i | ing |         |               |               | in bytes      |
| m |     |         |               |               | (various      |
| i |     |         |               |               | suffixes      |
| t |     |         |               |               | supported,    |
| s |     |         |               |               | see below) to |
| . |     |         |               |               | limit number  |
| h |     |         |               |               | of 1 GB       |
| u |     |         |               |               | hugepages     |
| g |     |         |               |               | (Available    |
| e |     |         |               |               | hugepage      |
| p |     |         |               |               | sizes are     |
| a |     |         |               |               | architecture  |
| g |     |         |               |               | dependent.)   |
| e |     |         |               |               |               |
| s |     |         |               |               |               |
| . |     |         |               |               |               |
| 1 |     |         |               |               |               |
| G |     |         |               |               |               |
| B |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| l | str | -       | no            | container     | This limits   |
| i | ing |         |               |               | kernel        |
| m |     |         |               |               | resources per |
| i |     |         |               |               | instance      |
| t |     |         |               |               | (e.g. number  |
| s |     |         |               |               | of open       |
| . |     |         |               |               | files)        |
| k |     |         |               |               |               |
| e |     |         |               |               |               |
| r |     |         |               |               |               |
| n |     |         |               |               |               |
| e |     |         |               |               |               |
| l |     |         |               |               |               |
| . |     |         |               |               |               |
| \ |     |         |               |               |               |
| * |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| l | str | -       | yes           | -             | Percentage of |
| i | ing |         |               |               | the host’s    |
| m |     |         |               |               | memory or     |
| i |     |         |               |               | fixed value   |
| t |     |         |               |               | in bytes      |
| s |     |         |               |               | (various      |
| . |     |         |               |               | suffixes      |
| m |     |         |               |               | supported,    |
| e |     |         |               |               | see below)    |
| m |     |         |               |               | (defaults to  |
| o |     |         |               |               | 1GiB for VMs) |
| r |     |         |               |               |               |
| y |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| l | str | hard    | yes           | container     | If hard,      |
| i | ing |         |               |               | instance      |
| m |     |         |               |               | can’t exceed  |
| i |     |         |               |               | its memory    |
| t |     |         |               |               | limit. If     |
| s |     |         |               |               | soft, the     |
| . |     |         |               |               | instance can  |
| m |     |         |               |               | exceed its    |
| e |     |         |               |               | memory limit  |
| m |     |         |               |               | when extra    |
| o |     |         |               |               | host memory   |
| r |     |         |               |               | is available  |
| y |     |         |               |               |               |
| . |     |         |               |               |               |
| e |     |         |               |               |               |
| n |     |         |               |               |               |
| f |     |         |               |               |               |
| o |     |         |               |               |               |
| r |     |         |               |               |               |
| c |     |         |               |               |               |
| e |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| l | boo | false   | no            | virtual-machi | Controls      |
| i | lea |         |               | ne            | whether to    |
| m | n   |         |               |               | back the      |
| i |     |         |               |               | instance      |
| t |     |         |               |               | using         |
| s |     |         |               |               | hugepages     |
| . |     |         |               |               | rather than   |
| m |     |         |               |               | regular       |
| e |     |         |               |               | system memory |
| m |     |         |               |               |               |
| o |     |         |               |               |               |
| r |     |         |               |               |               |
| y |     |         |               |               |               |
| . |     |         |               |               |               |
| h |     |         |               |               |               |
| u |     |         |               |               |               |
| g |     |         |               |               |               |
| e |     |         |               |               |               |
| p |     |         |               |               |               |
| a |     |         |               |               |               |
| g |     |         |               |               |               |
| e |     |         |               |               |               |
| s |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| l | boo | true    | yes           | container     | Controls      |
| i | lea |         |               |               | whether to    |
| m | n   |         |               |               | encourage/dis |
| i |     |         |               |               | courage       |
| t |     |         |               |               | swapping less |
| s |     |         |               |               | used pages    |
| . |     |         |               |               | for this      |
| m |     |         |               |               | instance      |
| e |     |         |               |               |               |
| m |     |         |               |               |               |
| o |     |         |               |               |               |
| r |     |         |               |               |               |
| y |     |         |               |               |               |
| . |     |         |               |               |               |
| s |     |         |               |               |               |
| w |     |         |               |               |               |
| a |     |         |               |               |               |
| p |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| l | int | 10      | yes           | container     | The higher    |
| i | ege | (maximu |               |               | this is set,  |
| m | r   | m)      |               |               | the least     |
| i |     |         |               |               | likely the    |
| t |     |         |               |               | instance is   |
| s |     |         |               |               | to be swapped |
| . |     |         |               |               | to disk       |
| m |     |         |               |               | (integer      |
| e |     |         |               |               | between 0 and |
| m |     |         |               |               | 10)           |
| o |     |         |               |               |               |
| r |     |         |               |               |               |
| y |     |         |               |               |               |
| . |     |         |               |               |               |
| s |     |         |               |               |               |
| w |     |         |               |               |               |
| a |     |         |               |               |               |
| p |     |         |               |               |               |
| . |     |         |               |               |               |
| p |     |         |               |               |               |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| o |     |         |               |               |               |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| t |     |         |               |               |               |
| y |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| l | int | 0       | yes           | -             | When under    |
| i | ege | (minimu |               |               | load, how     |
| m | r   | m)      |               |               | much priority |
| i |     |         |               |               | to give to    |
| t |     |         |               |               | the           |
| s |     |         |               |               | instance’s    |
| . |     |         |               |               | network       |
| n |     |         |               |               | requests      |
| e |     |         |               |               | (integer      |
| t |     |         |               |               | between 0 and |
| w |     |         |               |               | 10)           |
| o |     |         |               |               |               |
| r |     |         |               |               |               |
| k |     |         |               |               |               |
| . |     |         |               |               |               |
| p |     |         |               |               |               |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| o |     |         |               |               |               |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| t |     |         |               |               |               |
| y |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| l | int | - (max) | yes           | container     | Maximum       |
| i | ege |         |               |               | number of     |
| m | r   |         |               |               | processes     |
| i |     |         |               |               | that can run  |
| t |     |         |               |               | in the        |
| s |     |         |               |               | instance      |
| . |     |         |               |               |               |
| p |     |         |               |               |               |
| r |     |         |               |               |               |
| o |     |         |               |               |               |
| c |     |         |               |               |               |
| e |     |         |               |               |               |
| s |     |         |               |               |               |
| s |     |         |               |               |               |
| e |     |         |               |               |               |
| s |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| l | str | -       | yes           | container     | Comma         |
| i | ing |         |               |               | separated     |
| n |     |         |               |               | list of       |
| u |     |         |               |               | kernel        |
| x |     |         |               |               | modules to    |
| . |     |         |               |               | load before   |
| k |     |         |               |               | starting the  |
| e |     |         |               |               | instance      |
| r |     |         |               |               |               |
| n |     |         |               |               |               |
| e |     |         |               |               |               |
| l |     |         |               |               |               |
| _ |     |         |               |               |               |
| m |     |         |               |               |               |
| o |     |         |               |               |               |
| d |     |         |               |               |               |
| u |     |         |               |               |               |
| l |     |         |               |               |               |
| e |     |         |               |               |               |
| s |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| m | boo | false   | yes           | container     | Incremental   |
| i | lea |         |               |               | memory        |
| g | n   |         |               |               | transfer of   |
| r |     |         |               |               | the           |
| a |     |         |               |               | instance’s    |
| t |     |         |               |               | memory to     |
| i |     |         |               |               | reduce        |
| o |     |         |               |               | downtime      |
| n |     |         |               |               |               |
| . |     |         |               |               |               |
| i |     |         |               |               |               |
| n |     |         |               |               |               |
| c |     |         |               |               |               |
| r |     |         |               |               |               |
| e |     |         |               |               |               |
| m |     |         |               |               |               |
| e |     |         |               |               |               |
| n |     |         |               |               |               |
| t |     |         |               |               |               |
| a |     |         |               |               |               |
| l |     |         |               |               |               |
| . |     |         |               |               |               |
| m |     |         |               |               |               |
| e |     |         |               |               |               |
| m |     |         |               |               |               |
| o |     |         |               |               |               |
| r |     |         |               |               |               |
| y |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| m | int | 70      | yes           | container     | Percentage of |
| i | ege |         |               |               | memory to     |
| g | r   |         |               |               | have in sync  |
| r |     |         |               |               | before        |
| a |     |         |               |               | stopping the  |
| t |     |         |               |               | instance      |
| i |     |         |               |               |               |
| o |     |         |               |               |               |
| n |     |         |               |               |               |
| . |     |         |               |               |               |
| i |     |         |               |               |               |
| n |     |         |               |               |               |
| c |     |         |               |               |               |
| r |     |         |               |               |               |
| e |     |         |               |               |               |
| m |     |         |               |               |               |
| e |     |         |               |               |               |
| n |     |         |               |               |               |
| t |     |         |               |               |               |
| a |     |         |               |               |               |
| l |     |         |               |               |               |
| . |     |         |               |               |               |
| m |     |         |               |               |               |
| e |     |         |               |               |               |
| m |     |         |               |               |               |
| o |     |         |               |               |               |
| r |     |         |               |               |               |
| y |     |         |               |               |               |
| . |     |         |               |               |               |
| g |     |         |               |               |               |
| o |     |         |               |               |               |
| a |     |         |               |               |               |
| l |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| m | int | 10      | yes           | container     | Maximum       |
| i | ege |         |               |               | number of     |
| g | r   |         |               |               | transfer      |
| r |     |         |               |               | operations to |
| a |     |         |               |               | go through    |
| t |     |         |               |               | before        |
| i |     |         |               |               | stopping the  |
| o |     |         |               |               | instance      |
| n |     |         |               |               |               |
| . |     |         |               |               |               |
| i |     |         |               |               |               |
| n |     |         |               |               |               |
| c |     |         |               |               |               |
| r |     |         |               |               |               |
| e |     |         |               |               |               |
| m |     |         |               |               |               |
| e |     |         |               |               |               |
| n |     |         |               |               |               |
| t |     |         |               |               |               |
| a |     |         |               |               |               |
| l |     |         |               |               |               |
| . |     |         |               |               |               |
| m |     |         |               |               |               |
| e |     |         |               |               |               |
| m |     |         |               |               |               |
| o |     |         |               |               |               |
| r |     |         |               |               |               |
| y |     |         |               |               |               |
| . |     |         |               |               |               |
| i |     |         |               |               |               |
| t |     |         |               |               |               |
| e |     |         |               |               |               |
| r |     |         |               |               |               |
| a |     |         |               |               |               |
| t |     |         |               |               |               |
| i |     |         |               |               |               |
| o |     |         |               |               |               |
| n |     |         |               |               |               |
| s |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| m | boo | false   | no            | virtual-machi | Allow for     |
| i | lea |         |               | ne            | stateful      |
| g | n   |         |               |               | stop/start    |
| r |     |         |               |               | and           |
| a |     |         |               |               | snapshots.    |
| t |     |         |               |               | This will     |
| i |     |         |               |               | prevent the   |
| o |     |         |               |               | use of some   |
| n |     |         |               |               | features that |
| . |     |         |               |               | are           |
| s |     |         |               |               | incompatible  |
| t |     |         |               |               | with it       |
| a |     |         |               |               |               |
| t |     |         |               |               |               |
| e |     |         |               |               |               |
| f |     |         |               |               |               |
| u |     |         |               |               |               |
| l |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| n | str | compute | no            | container     | What driver   |
| v | ing | ,utilit |               |               | capabilities  |
| i |     | y       |               |               | the instance  |
| d |     |         |               |               | needs (sets   |
| i |     |         |               |               | libnvidia-con |
| a |     |         |               |               | tainer        |
| . |     |         |               |               | NVIDIA_DRIVER |
| d |     |         |               |               | _CAPABILITIES |
| r |     |         |               |               | )             |
| i |     |         |               |               |               |
| v |     |         |               |               |               |
| e |     |         |               |               |               |
| r |     |         |               |               |               |
| . |     |         |               |               |               |
| c |     |         |               |               |               |
| a |     |         |               |               |               |
| p |     |         |               |               |               |
| a |     |         |               |               |               |
| b |     |         |               |               |               |
| i |     |         |               |               |               |
| l |     |         |               |               |               |
| i |     |         |               |               |               |
| t |     |         |               |               |               |
| i |     |         |               |               |               |
| e |     |         |               |               |               |
| s |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| n | boo | false   | no            | container     | Pass the host |
| v | lea |         |               |               | NVIDIA and    |
| i | n   |         |               |               | CUDA runtime  |
| d |     |         |               |               | libraries     |
| i |     |         |               |               | into the      |
| a |     |         |               |               | instance      |
| . |     |         |               |               |               |
| r |     |         |               |               |               |
| u |     |         |               |               |               |
| n |     |         |               |               |               |
| t |     |         |               |               |               |
| i |     |         |               |               |               |
| m |     |         |               |               |               |
| e |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| n | str | -       | no            | container     | Version       |
| v | ing |         |               |               | expression    |
| i |     |         |               |               | for the       |
| d |     |         |               |               | required CUDA |
| i |     |         |               |               | version (sets |
| a |     |         |               |               | libnvidia-con |
| . |     |         |               |               | tainer        |
| r |     |         |               |               | NVIDIA_REQUIR |
| e |     |         |               |               | E_CUDA)       |
| q |     |         |               |               |               |
| u |     |         |               |               |               |
| i |     |         |               |               |               |
| r |     |         |               |               |               |
| e |     |         |               |               |               |
| . |     |         |               |               |               |
| c |     |         |               |               |               |
| u |     |         |               |               |               |
| d |     |         |               |               |               |
| a |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| n | str | -       | no            | container     | Version       |
| v | ing |         |               |               | expression    |
| i |     |         |               |               | for the       |
| d |     |         |               |               | required      |
| i |     |         |               |               | driver        |
| a |     |         |               |               | version (sets |
| . |     |         |               |               | libnvidia-con |
| r |     |         |               |               | tainer        |
| e |     |         |               |               | NVIDIA_REQUIR |
| q |     |         |               |               | E_DRIVER)     |
| u |     |         |               |               |               |
| i |     |         |               |               |               |
| r |     |         |               |               |               |
| e |     |         |               |               |               |
| . |     |         |               |               |               |
| d |     |         |               |               |               |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| v |     |         |               |               |               |
| e |     |         |               |               |               |
| r |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| r | blo | -       | yes           | -             | Apparmor      |
| a | b   |         |               |               | profile       |
| w |     |         |               |               | entries to be |
| . |     |         |               |               | appended to   |
| a |     |         |               |               | the generated |
| p |     |         |               |               | profile       |
| p |     |         |               |               |               |
| a |     |         |               |               |               |
| r |     |         |               |               |               |
| m |     |         |               |               |               |
| o |     |         |               |               |               |
| r |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| r | blo | -       | no            | unprivileged  | Raw idmap     |
| a | b   |         |               | container     | configuration |
| w |     |         |               |               | (e.g. “both   |
| . |     |         |               |               | 1000 1000”)   |
| i |     |         |               |               |               |
| d |     |         |               |               |               |
| m |     |         |               |               |               |
| a |     |         |               |               |               |
| p |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| r | blo | -       | no            | container     | Raw LXC       |
| a | b   |         |               |               | configuration |
| w |     |         |               |               | to be         |
| . |     |         |               |               | appended to   |
| l |     |         |               |               | the generated |
| x |     |         |               |               | one           |
| c |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| r | blo | -       | no            | virtual-machi | Raw Qemu      |
| a | b   |         |               | ne            | configuration |
| w |     |         |               |               | to be         |
| . |     |         |               |               | appended to   |
| q |     |         |               |               | the generated |
| e |     |         |               |               | command line  |
| m |     |         |               |               |               |
| u |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| r | blo | -       | no            | container     | Raw Seccomp   |
| a | b   |         |               |               | configuration |
| w |     |         |               |               |               |
| . |     |         |               |               |               |
| s |     |         |               |               |               |
| e |     |         |               |               |               |
| c |     |         |               |               |               |
| c |     |         |               |               |               |
| o |     |         |               |               |               |
| m |     |         |               |               |               |
| p |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | boo | true    | no            | -             | Controls the  |
| e | lea |         |               |               | presence of   |
| c | n   |         |               |               | /dev/lxd in   |
| u |     |         |               |               | the instance  |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| t |     |         |               |               |               |
| y |     |         |               |               |               |
| . |     |         |               |               |               |
| d |     |         |               |               |               |
| e |     |         |               |               |               |
| v |     |         |               |               |               |
| l |     |         |               |               |               |
| x |     |         |               |               |               |
| d |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | boo | false   | no            | container     | Controls the  |
| e | lea |         |               |               | availability  |
| c | n   |         |               |               | of the        |
| u |     |         |               |               | /1.0/images   |
| r |     |         |               |               | API over      |
| i |     |         |               |               | devlxd        |
| t |     |         |               |               |               |
| y |     |         |               |               |               |
| . |     |         |               |               |               |
| d |     |         |               |               |               |
| e |     |         |               |               |               |
| v |     |         |               |               |               |
| l |     |         |               |               |               |
| x |     |         |               |               |               |
| d |     |         |               |               |               |
| . |     |         |               |               |               |
| i |     |         |               |               |               |
| m |     |         |               |               |               |
| a |     |         |               |               |               |
| g |     |         |               |               |               |
| e |     |         |               |               |               |
| s |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | int | -       | no            | unprivileged  | The base host |
| e | ege |         |               | container     | ID to use for |
| c | r   |         |               |               | the           |
| u |     |         |               |               | allocation    |
| r |     |         |               |               | (overrides    |
| i |     |         |               |               | auto-detectio |
| t |     |         |               |               | n)            |
| y |     |         |               |               |               |
| . |     |         |               |               |               |
| i |     |         |               |               |               |
| d |     |         |               |               |               |
| m |     |         |               |               |               |
| a |     |         |               |               |               |
| p |     |         |               |               |               |
| . |     |         |               |               |               |
| b |     |         |               |               |               |
| a |     |         |               |               |               |
| s |     |         |               |               |               |
| e |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | boo | false   | no            | unprivileged  | Use an idmap  |
| e | lea |         |               | container     | for this      |
| c | n   |         |               |               | instance that |
| u |     |         |               |               | is unique     |
| r |     |         |               |               | among         |
| i |     |         |               |               | instances     |
| t |     |         |               |               | with isolated |
| y |     |         |               |               | set           |
| . |     |         |               |               |               |
| i |     |         |               |               |               |
| d |     |         |               |               |               |
| m |     |         |               |               |               |
| a |     |         |               |               |               |
| p |     |         |               |               |               |
| . |     |         |               |               |               |
| i |     |         |               |               |               |
| s |     |         |               |               |               |
| o |     |         |               |               |               |
| l |     |         |               |               |               |
| a |     |         |               |               |               |
| t |     |         |               |               |               |
| e |     |         |               |               |               |
| d |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | int | -       | no            | unprivileged  | The size of   |
| e | ege |         |               | container     | the idmap to  |
| c | r   |         |               |               | use           |
| u |     |         |               |               |               |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| t |     |         |               |               |               |
| y |     |         |               |               |               |
| . |     |         |               |               |               |
| i |     |         |               |               |               |
| d |     |         |               |               |               |
| m |     |         |               |               |               |
| a |     |         |               |               |               |
| p |     |         |               |               |               |
| . |     |         |               |               |               |
| s |     |         |               |               |               |
| i |     |         |               |               |               |
| z |     |         |               |               |               |
| e |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | boo | false   | yes           | container     | Support       |
| e | lea |         |               |               | running lxd   |
| c | n   |         |               |               | (nested)      |
| u |     |         |               |               | inside the    |
| r |     |         |               |               | instance      |
| i |     |         |               |               |               |
| t |     |         |               |               |               |
| y |     |         |               |               |               |
| . |     |         |               |               |               |
| n |     |         |               |               |               |
| e |     |         |               |               |               |
| s |     |         |               |               |               |
| t |     |         |               |               |               |
| i |     |         |               |               |               |
| n |     |         |               |               |               |
| g |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | boo | false   | no            | container     | Runs the      |
| e | lea |         |               |               | instance in   |
| c | n   |         |               |               | privileged    |
| u |     |         |               |               | mode          |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| t |     |         |               |               |               |
| y |     |         |               |               |               |
| . |     |         |               |               |               |
| p |     |         |               |               |               |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| v |     |         |               |               |               |
| i |     |         |               |               |               |
| l |     |         |               |               |               |
| e |     |         |               |               |               |
| g |     |         |               |               |               |
| e |     |         |               |               |               |
| d |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | boo | false   | yes           | -             | Prevents the  |
| e | lea |         |               |               | instance from |
| c | n   |         |               |               | being deleted |
| u |     |         |               |               |               |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| t |     |         |               |               |               |
| y |     |         |               |               |               |
| . |     |         |               |               |               |
| p |     |         |               |               |               |
| r |     |         |               |               |               |
| o |     |         |               |               |               |
| t |     |         |               |               |               |
| e |     |         |               |               |               |
| c |     |         |               |               |               |
| t |     |         |               |               |               |
| i |     |         |               |               |               |
| o |     |         |               |               |               |
| n |     |         |               |               |               |
| . |     |         |               |               |               |
| d |     |         |               |               |               |
| e |     |         |               |               |               |
| l |     |         |               |               |               |
| e |     |         |               |               |               |
| t |     |         |               |               |               |
| e |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | boo | false   | yes           | container     | Prevents the  |
| e | lea |         |               |               | instance’s    |
| c | n   |         |               |               | filesystem    |
| u |     |         |               |               | from being    |
| r |     |         |               |               | uid/gid       |
| i |     |         |               |               | shifted on    |
| t |     |         |               |               | startup       |
| y |     |         |               |               |               |
| . |     |         |               |               |               |
| p |     |         |               |               |               |
| r |     |         |               |               |               |
| o |     |         |               |               |               |
| t |     |         |               |               |               |
| e |     |         |               |               |               |
| c |     |         |               |               |               |
| t |     |         |               |               |               |
| i |     |         |               |               |               |
| o |     |         |               |               |               |
| n |     |         |               |               |               |
| . |     |         |               |               |               |
| s |     |         |               |               |               |
| h |     |         |               |               |               |
| i |     |         |               |               |               |
| f |     |         |               |               |               |
| t |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | boo | true    | no            | virtual-machi | Controls      |
| e | lea |         |               | ne            | whether UEFI  |
| c | n   |         |               |               | secure boot   |
| u |     |         |               |               | is enabled    |
| r |     |         |               |               | with the      |
| i |     |         |               |               | default       |
| t |     |         |               |               | Microsoft     |
| y |     |         |               |               | keys          |
| . |     |         |               |               |               |
| s |     |         |               |               |               |
| e |     |         |               |               |               |
| c |     |         |               |               |               |
| u |     |         |               |               |               |
| r |     |         |               |               |               |
| e |     |         |               |               |               |
| b |     |         |               |               |               |
| o |     |         |               |               |               |
| o |     |         |               |               |               |
| t |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | str | -       | no            | container     | A             |
| e | ing |         |               |               | ‘:raw-latex:` |
| c |     |         |               |               | \n`’          |
| u |     |         |               |               | separated     |
| r |     |         |               |               | list of       |
| i |     |         |               |               | syscalls to   |
| t |     |         |               |               | allow         |
| y |     |         |               |               | (mutually     |
| . |     |         |               |               | exclusive     |
| s |     |         |               |               | with          |
| y |     |         |               |               | security.sysc |
| s |     |         |               |               | alls.deny*)   |
| c |     |         |               |               |               |
| a |     |         |               |               |               |
| l |     |         |               |               |               |
| l |     |         |               |               |               |
| s |     |         |               |               |               |
| . |     |         |               |               |               |
| a |     |         |               |               |               |
| l |     |         |               |               |               |
| l |     |         |               |               |               |
| o |     |         |               |               |               |
| w |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | str | -       | no            | container     | A             |
| e | ing |         |               |               | ‘:raw-latex:` |
| c |     |         |               |               | \n`’          |
| u |     |         |               |               | separated     |
| r |     |         |               |               | list of       |
| i |     |         |               |               | syscalls to   |
| t |     |         |               |               | deny          |
| y |     |         |               |               |               |
| . |     |         |               |               |               |
| s |     |         |               |               |               |
| y |     |         |               |               |               |
| s |     |         |               |               |               |
| c |     |         |               |               |               |
| a |     |         |               |               |               |
| l |     |         |               |               |               |
| l |     |         |               |               |               |
| s |     |         |               |               |               |
| . |     |         |               |               |               |
| d |     |         |               |               |               |
| e |     |         |               |               |               |
| n |     |         |               |               |               |
| y |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | boo | false   | no            | container     | On x86_64     |
| e | lea |         |               |               | this enables  |
| c | n   |         |               |               | blocking of   |
| u |     |         |               |               | compat_\*     |
| r |     |         |               |               | syscalls, it  |
| i |     |         |               |               | is a no-op on |
| t |     |         |               |               | other arches  |
| y |     |         |               |               |               |
| . |     |         |               |               |               |
| s |     |         |               |               |               |
| y |     |         |               |               |               |
| s |     |         |               |               |               |
| c |     |         |               |               |               |
| a |     |         |               |               |               |
| l |     |         |               |               |               |
| l |     |         |               |               |               |
| s |     |         |               |               |               |
| . |     |         |               |               |               |
| d |     |         |               |               |               |
| e |     |         |               |               |               |
| n |     |         |               |               |               |
| y |     |         |               |               |               |
| _ |     |         |               |               |               |
| c |     |         |               |               |               |
| o |     |         |               |               |               |
| m |     |         |               |               |               |
| p |     |         |               |               |               |
| a |     |         |               |               |               |
| t |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | boo | true    | no            | container     | Enables the   |
| e | lea |         |               |               | default       |
| c | n   |         |               |               | syscall deny  |
| u |     |         |               |               |               |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| t |     |         |               |               |               |
| y |     |         |               |               |               |
| . |     |         |               |               |               |
| s |     |         |               |               |               |
| y |     |         |               |               |               |
| s |     |         |               |               |               |
| c |     |         |               |               |               |
| a |     |         |               |               |               |
| l |     |         |               |               |               |
| l |     |         |               |               |               |
| s |     |         |               |               |               |
| . |     |         |               |               |               |
| d |     |         |               |               |               |
| e |     |         |               |               |               |
| n |     |         |               |               |               |
| y |     |         |               |               |               |
| _ |     |         |               |               |               |
| d |     |         |               |               |               |
| e |     |         |               |               |               |
| f |     |         |               |               |               |
| a |     |         |               |               |               |
| u |     |         |               |               |               |
| l |     |         |               |               |               |
| t |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | boo | false   | no            | container     | Handles the   |
| e | lea |         |               |               | ``bpf``       |
| c | n   |         |               |               | system call   |
| u |     |         |               |               |               |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| t |     |         |               |               |               |
| y |     |         |               |               |               |
| . |     |         |               |               |               |
| s |     |         |               |               |               |
| y |     |         |               |               |               |
| s |     |         |               |               |               |
| c |     |         |               |               |               |
| a |     |         |               |               |               |
| l |     |         |               |               |               |
| l |     |         |               |               |               |
| s |     |         |               |               |               |
| . |     |         |               |               |               |
| i |     |         |               |               |               |
| n |     |         |               |               |               |
| t |     |         |               |               |               |
| e |     |         |               |               |               |
| r |     |         |               |               |               |
| c |     |         |               |               |               |
| e |     |         |               |               |               |
| p |     |         |               |               |               |
| t |     |         |               |               |               |
| . |     |         |               |               |               |
| b |     |         |               |               |               |
| p |     |         |               |               |               |
| f |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | boo | false   | no            | container     | Allows        |
| e | lea |         |               |               | ``bpf``       |
| c | n   |         |               |               | programs for  |
| u |     |         |               |               | the devices   |
| r |     |         |               |               | cgroup in the |
| i |     |         |               |               | unified       |
| t |     |         |               |               | hierarchy to  |
| y |     |         |               |               | be loaded.    |
| . |     |         |               |               |               |
| s |     |         |               |               |               |
| y |     |         |               |               |               |
| s |     |         |               |               |               |
| c |     |         |               |               |               |
| a |     |         |               |               |               |
| l |     |         |               |               |               |
| l |     |         |               |               |               |
| s |     |         |               |               |               |
| . |     |         |               |               |               |
| i |     |         |               |               |               |
| n |     |         |               |               |               |
| t |     |         |               |               |               |
| e |     |         |               |               |               |
| r |     |         |               |               |               |
| c |     |         |               |               |               |
| e |     |         |               |               |               |
| p |     |         |               |               |               |
| t |     |         |               |               |               |
| . |     |         |               |               |               |
| b |     |         |               |               |               |
| p |     |         |               |               |               |
| f |     |         |               |               |               |
| . |     |         |               |               |               |
| d |     |         |               |               |               |
| e |     |         |               |               |               |
| v |     |         |               |               |               |
| i |     |         |               |               |               |
| c |     |         |               |               |               |
| e |     |         |               |               |               |
| s |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | boo | false   | no            | container     | Handles the   |
| e | lea |         |               |               | ``mknod`` and |
| c | n   |         |               |               | ``mknodat``   |
| u |     |         |               |               | system calls  |
| r |     |         |               |               | (allows       |
| i |     |         |               |               | creation of a |
| t |     |         |               |               | limited       |
| y |     |         |               |               | subset of     |
| . |     |         |               |               | char/block    |
| s |     |         |               |               | devices)      |
| y |     |         |               |               |               |
| s |     |         |               |               |               |
| c |     |         |               |               |               |
| a |     |         |               |               |               |
| l |     |         |               |               |               |
| l |     |         |               |               |               |
| s |     |         |               |               |               |
| . |     |         |               |               |               |
| i |     |         |               |               |               |
| n |     |         |               |               |               |
| t |     |         |               |               |               |
| e |     |         |               |               |               |
| r |     |         |               |               |               |
| c |     |         |               |               |               |
| e |     |         |               |               |               |
| p |     |         |               |               |               |
| t |     |         |               |               |               |
| . |     |         |               |               |               |
| m |     |         |               |               |               |
| k |     |         |               |               |               |
| n |     |         |               |               |               |
| o |     |         |               |               |               |
| d |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | boo | false   | no            | container     | Handles the   |
| e | lea |         |               |               | ``mount``     |
| c | n   |         |               |               | system call   |
| u |     |         |               |               |               |
| r |     |         |               |               |               |
| i |     |         |               |               |               |
| t |     |         |               |               |               |
| y |     |         |               |               |               |
| . |     |         |               |               |               |
| s |     |         |               |               |               |
| y |     |         |               |               |               |
| s |     |         |               |               |               |
| c |     |         |               |               |               |
| a |     |         |               |               |               |
| l |     |         |               |               |               |
| l |     |         |               |               |               |
| s |     |         |               |               |               |
| . |     |         |               |               |               |
| i |     |         |               |               |               |
| n |     |         |               |               |               |
| t |     |         |               |               |               |
| e |     |         |               |               |               |
| r |     |         |               |               |               |
| c |     |         |               |               |               |
| e |     |         |               |               |               |
| p |     |         |               |               |               |
| t |     |         |               |               |               |
| . |     |         |               |               |               |
| m |     |         |               |               |               |
| o |     |         |               |               |               |
| u |     |         |               |               |               |
| n |     |         |               |               |               |
| t |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | str | -       | yes           | container     | Specify a     |
| e | ing |         |               |               | comma-separat |
| c |     |         |               |               | ed            |
| u |     |         |               |               | list of       |
| r |     |         |               |               | filesystems   |
| i |     |         |               |               | that are safe |
| t |     |         |               |               | to mount for  |
| y |     |         |               |               | processes     |
| . |     |         |               |               | inside the    |
| s |     |         |               |               | instance      |
| y |     |         |               |               |               |
| s |     |         |               |               |               |
| c |     |         |               |               |               |
| a |     |         |               |               |               |
| l |     |         |               |               |               |
| l |     |         |               |               |               |
| s |     |         |               |               |               |
| . |     |         |               |               |               |
| i |     |         |               |               |               |
| n |     |         |               |               |               |
| t |     |         |               |               |               |
| e |     |         |               |               |               |
| r |     |         |               |               |               |
| c |     |         |               |               |               |
| e |     |         |               |               |               |
| p |     |         |               |               |               |
| t |     |         |               |               |               |
| . |     |         |               |               |               |
| m |     |         |               |               |               |
| o |     |         |               |               |               |
| u |     |         |               |               |               |
| n |     |         |               |               |               |
| t |     |         |               |               |               |
| . |     |         |               |               |               |
| a |     |         |               |               |               |
| l |     |         |               |               |               |
| l |     |         |               |               |               |
| o |     |         |               |               |               |
| w |     |         |               |               |               |
| e |     |         |               |               |               |
| d |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | str | -       | yes           | container     | Whether to    |
| e | ing |         |               |               | redirect      |
| c |     |         |               |               | mounts of a   |
| u |     |         |               |               | given         |
| r |     |         |               |               | filesystem to |
| i |     |         |               |               | their fuse    |
| t |     |         |               |               | implemenation |
| y |     |         |               |               | (e.g. ext4=fu |
| . |     |         |               |               | se2fs)        |
| s |     |         |               |               |               |
| y |     |         |               |               |               |
| s |     |         |               |               |               |
| c |     |         |               |               |               |
| a |     |         |               |               |               |
| l |     |         |               |               |               |
| l |     |         |               |               |               |
| s |     |         |               |               |               |
| . |     |         |               |               |               |
| i |     |         |               |               |               |
| n |     |         |               |               |               |
| t |     |         |               |               |               |
| e |     |         |               |               |               |
| r |     |         |               |               |               |
| c |     |         |               |               |               |
| e |     |         |               |               |               |
| p |     |         |               |               |               |
| t |     |         |               |               |               |
| . |     |         |               |               |               |
| m |     |         |               |               |               |
| o |     |         |               |               |               |
| u |     |         |               |               |               |
| n |     |         |               |               |               |
| t |     |         |               |               |               |
| . |     |         |               |               |               |
| f |     |         |               |               |               |
| u |     |         |               |               |               |
| s |     |         |               |               |               |
| e |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | boo | false   | yes           | container     | Whether to    |
| e | lea |         |               |               | mount shiftfs |
| c | n   |         |               |               | on top of     |
| u |     |         |               |               | filesystems   |
| r |     |         |               |               | handled       |
| i |     |         |               |               | through mount |
| t |     |         |               |               | syscall       |
| y |     |         |               |               | interception  |
| . |     |         |               |               |               |
| s |     |         |               |               |               |
| y |     |         |               |               |               |
| s |     |         |               |               |               |
| c |     |         |               |               |               |
| a |     |         |               |               |               |
| l |     |         |               |               |               |
| l |     |         |               |               |               |
| s |     |         |               |               |               |
| . |     |         |               |               |               |
| i |     |         |               |               |               |
| n |     |         |               |               |               |
| t |     |         |               |               |               |
| e |     |         |               |               |               |
| r |     |         |               |               |               |
| c |     |         |               |               |               |
| e |     |         |               |               |               |
| p |     |         |               |               |               |
| t |     |         |               |               |               |
| . |     |         |               |               |               |
| m |     |         |               |               |               |
| o |     |         |               |               |               |
| u |     |         |               |               |               |
| n |     |         |               |               |               |
| t |     |         |               |               |               |
| . |     |         |               |               |               |
| s |     |         |               |               |               |
| h |     |         |               |               |               |
| i |     |         |               |               |               |
| f |     |         |               |               |               |
| t |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | boo | false   | no            | container     | Handles the   |
| e | lea |         |               |               | ``setxattr``  |
| c | n   |         |               |               | system call   |
| u |     |         |               |               | (allows       |
| r |     |         |               |               | setting a     |
| i |     |         |               |               | limited       |
| t |     |         |               |               | subset of     |
| y |     |         |               |               | restricted    |
| . |     |         |               |               | extended      |
| s |     |         |               |               | attributes)   |
| y |     |         |               |               |               |
| s |     |         |               |               |               |
| c |     |         |               |               |               |
| a |     |         |               |               |               |
| l |     |         |               |               |               |
| l |     |         |               |               |               |
| s |     |         |               |               |               |
| . |     |         |               |               |               |
| i |     |         |               |               |               |
| n |     |         |               |               |               |
| t |     |         |               |               |               |
| e |     |         |               |               |               |
| r |     |         |               |               |               |
| c |     |         |               |               |               |
| e |     |         |               |               |               |
| p |     |         |               |               |               |
| t |     |         |               |               |               |
| . |     |         |               |               |               |
| s |     |         |               |               |               |
| e |     |         |               |               |               |
| t |     |         |               |               |               |
| x |     |         |               |               |               |
| a |     |         |               |               |               |
| t |     |         |               |               |               |
| t |     |         |               |               |               |
| r |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | str | -       | no            | -             | Cron          |
| n | ing |         |               |               | expression    |
| a |     |         |               |               | (``<minute> < |
| p |     |         |               |               | hour> <dom> < |
| s |     |         |               |               | month> <dow>` |
| h |     |         |               |               | `),           |
| o |     |         |               |               | or a comma    |
| t |     |         |               |               | separated     |
| s |     |         |               |               | list of       |
| . |     |         |               |               | schedule      |
| s |     |         |               |               | aliases       |
| c |     |         |               |               | ``<@hourly> < |
| h |     |         |               |               | @daily> <@mid |
| e |     |         |               |               | night> <@week |
| d |     |         |               |               | ly> <@monthly |
| u |     |         |               |               | > <@annually> |
| l |     |         |               |               |  <@yearly> <@ |
| e |     |         |               |               | startup>``    |
+---+-----+---------+---------------+---------------+---------------+
| s | boo | false   | no            | -             | Controls      |
| n | l   |         |               |               | whether or    |
| a |     |         |               |               | not stopped   |
| p |     |         |               |               | instances are |
| s |     |         |               |               | to be         |
| h |     |         |               |               | snapshoted    |
| o |     |         |               |               | automatically |
| t |     |         |               |               |               |
| s |     |         |               |               |               |
| . |     |         |               |               |               |
| s |     |         |               |               |               |
| c |     |         |               |               |               |
| h |     |         |               |               |               |
| e |     |         |               |               |               |
| d |     |         |               |               |               |
| u |     |         |               |               |               |
| l |     |         |               |               |               |
| e |     |         |               |               |               |
| . |     |         |               |               |               |
| s |     |         |               |               |               |
| t |     |         |               |               |               |
| o |     |         |               |               |               |
| p |     |         |               |               |               |
| p |     |         |               |               |               |
| e |     |         |               |               |               |
| d |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | str | snap%d  | no            | -             | Pongo2        |
| n | ing |         |               |               | template      |
| a |     |         |               |               | string which  |
| p |     |         |               |               | represents    |
| s |     |         |               |               | the snapshot  |
| h |     |         |               |               | name (used    |
| o |     |         |               |               | for scheduled |
| t |     |         |               |               | snapshots and |
| s |     |         |               |               | unnamed       |
| . |     |         |               |               | snapshots)    |
| p |     |         |               |               |               |
| a |     |         |               |               |               |
| t |     |         |               |               |               |
| t |     |         |               |               |               |
| e |     |         |               |               |               |
| r |     |         |               |               |               |
| n |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| s | str | -       | no            | -             | Controls when |
| n | ing |         |               |               | snapshots are |
| a |     |         |               |               | to be deleted |
| p |     |         |               |               | (expects      |
| s |     |         |               |               | expression    |
| h |     |         |               |               | like          |
| o |     |         |               |               | ``1M 2H 3d 4w |
| t |     |         |               |               |  5m 6y``)     |
| s |     |         |               |               |               |
| . |     |         |               |               |               |
| e |     |         |               |               |               |
| x |     |         |               |               |               |
| p |     |         |               |               |               |
| i |     |         |               |               |               |
| r |     |         |               |               |               |
| y |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+
| u | str | -       | n/a           | -             | Free form     |
| s | ing |         |               |               | user          |
| e |     |         |               |               | key/value     |
| r |     |         |               |               | storage (can  |
| . |     |         |               |               | be used in    |
| \ |     |         |               |               | search)       |
| * |     |         |               |               |               |
+---+-----+---------+---------------+---------------+---------------+

The following volatile keys are currently internally used by LXD:

+-------+----------+-------------------+------------------------------+
| Key   | Type     | Default           | Description                  |
+=======+==========+===================+==============================+
| volat | string   | -                 | The name of a template hook  |
| ile.a |          |                   | which should be triggered    |
| pply_ |          |                   | upon next startup            |
| templ |          |                   |                              |
| ate   |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| volat | string   | -                 | The hash of the image the    |
| ile.b |          |                   | instance was created from,   |
| ase_i |          |                   | if any                       |
| mage  |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| volat | string   | -                 | The origin (cluster member)  |
| ile.e |          |                   | of the evacuated instance    |
| vacua |          |                   |                              |
| te.or |          |                   |                              |
| igin  |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| volat | integer  | -                 | The first id in the          |
| ile.i |          |                   | instance’s primary idmap     |
| dmap. |          |                   | range                        |
| base  |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| volat | string   | -                 | The idmap currently in use   |
| ile.i |          |                   | by the instance              |
| dmap. |          |                   |                              |
| curre |          |                   |                              |
| nt    |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| volat | string   | -                 | The idmap to use next time   |
| ile.i |          |                   | the instance starts          |
| dmap. |          |                   |                              |
| next  |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| volat | string   | -                 | Serialized instance uid/gid  |
| ile.l |          |                   | map                          |
| ast_s |          |                   |                              |
| tate. |          |                   |                              |
| idmap |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| volat | string   | -                 | Instance state as of last    |
| ile.l |          |                   | host shutdown                |
| ast_s |          |                   |                              |
| tate. |          |                   |                              |
| power |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| volat | string   | -                 | Instance vsock ID used as of |
| ile.v |          |                   | last start                   |
| sock_ |          |                   |                              |
| id    |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| volat | string   | -                 | Instance UUID (globally      |
| ile.u |          |                   | unique across all servers    |
| uid   |          |                   | and projects)                |
+-------+----------+-------------------+------------------------------+
| volat | string   | -                 | Disk quota to be applied on  |
| ile.< |          |                   | next instance start          |
| name> |          |                   |                              |
| .appl |          |                   |                              |
| y_quo |          |                   |                              |
| ta    |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| volat | string   | -                 | RBD device path for Ceph     |
| ile.< |          |                   | disk devices                 |
| name> |          |                   |                              |
| .ceph |          |                   |                              |
| _rbd  |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| volat | string   | -                 | Network device name on the   |
| ile.< |          |                   | host                         |
| name> |          |                   |                              |
| .host |          |                   |                              |
| _name |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| volat | string   | -                 | Network device MAC address   |
| ile.< |          |                   | (when no hwaddr property is  |
| name> |          |                   | set on the device itself)    |
| .hwad |          |                   |                              |
| dr    |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| volat | string   | -                 | Whether or not the network   |
| ile.< |          |                   | device physical device was   |
| name> |          |                   | created (“true” or “false”)  |
| .last |          |                   |                              |
| _stat |          |                   |                              |
| e.cre |          |                   |                              |
| ated  |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| volat | string   | -                 | Network device original MTU  |
| ile.< |          |                   | used when moving a physical  |
| name> |          |                   | device into an instance      |
| .last |          |                   |                              |
| _stat |          |                   |                              |
| e.mtu |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| volat | string   | -                 | Network device original MAC  |
| ile.< |          |                   | used when moving a physical  |
| name> |          |                   | device into an instance      |
| .last |          |                   |                              |
| _stat |          |                   |                              |
| e.hwa |          |                   |                              |
| ddr   |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| volat | string   | -                 | SR-IOV Virtual function ID   |
| ile.< |          |                   | used when moving a VF into   |
| name> |          |                   | an instance                  |
| .last |          |                   |                              |
| _stat |          |                   |                              |
| e.vf. |          |                   |                              |
| id    |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| volat | string   | -                 | SR-IOV Virtual function      |
| ile.< |          |                   | original MAC used when       |
| name> |          |                   | moving a VF into an instance |
| .last |          |                   |                              |
| _stat |          |                   |                              |
| e.vf. |          |                   |                              |
| hwadd |          |                   |                              |
| r     |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| volat | string   | -                 | SR-IOV Virtual function      |
| ile.< |          |                   | original VLAN used when      |
| name> |          |                   | moving a VF into an instance |
| .last |          |                   |                              |
| _stat |          |                   |                              |
| e.vf. |          |                   |                              |
| vlan  |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| volat | string   | -                 | SR-IOV Virtual function      |
| ile.< |          |                   | original spoof check setting |
| name> |          |                   | used when moving a VF into   |
| .last |          |                   | an instance                  |
| _stat |          |                   |                              |
| e.vf. |          |                   |                              |
| spoof |          |                   |                              |
| check |          |                   |                              |
+-------+----------+-------------------+------------------------------+

Additionally, those user keys have become common with images (support
isn’t guaranteed):

+-------+----------+-------------------+------------------------------+
| Key   | Type     | Default           | Description                  |
+=======+==========+===================+==============================+
| user. | string   | -                 | Cloud-init meta-data,        |
| meta- |          |                   | content is appended to seed  |
| data  |          |                   | value                        |
+-------+----------+-------------------+------------------------------+
| user. | string   | DHCP on eth0      | Cloud-init network-config,   |
| netwo |          |                   | content is used as seed      |
| rk-co |          |                   | value                        |
| nfig  |          |                   |                              |
+-------+----------+-------------------+------------------------------+
| user. | string   | dhcp              | One of “dhcp” or             |
| netwo |          |                   | “link-local”. Used to        |
| rk_mo |          |                   | configure network in         |
| de    |          |                   | supported images             |
+-------+----------+-------------------+------------------------------+
| user. | string   | #!cloud-config    | Cloud-init user-data,        |
| user- |          |                   | content is used as seed      |
| data  |          |                   | value                        |
+-------+----------+-------------------+------------------------------+
| user. | string   | #!cloud-config    | Cloud-init vendor-data,      |
| vendo |          |                   | content is used as seed      |
| r-dat |          |                   | value                        |
| a     |          |                   |                              |
+-------+----------+-------------------+------------------------------+

Note that while a type is defined above as a convenience, all values are
stored as strings and should be exported over the REST API as strings
(which makes it possible to support any extra values without breaking
backward compatibility).

Those keys can be set using the lxc tool with:

.. code:: bash

   lxc config set <instance> <key> <value>

Volatile keys can’t be set by the user and can only be set directly
against an instance.

The raw keys allow direct interaction with the backend features that LXD
itself uses, setting those may very well break LXD in non-obvious ways
and should whenever possible be avoided.

CPU limits
~~~~~~~~~~

The CPU limits are implemented through a mix of the ``cpuset`` and
``cpu`` CGroup controllers.

``limits.cpu`` results in CPU pinning through the ``cpuset`` controller.
A set of CPUs (e.g. ``1,2,3``) or a CPU range (e.g. ``0-3``) can be
specified.

When a number of CPUs is specified instead (e.g. ``4``), LXD will do
dynamic load-balancing of all instances that aren’t pinned to specific
CPUs, trying to spread the load on the machine. Instances will then be
re-balanced every time an instance starts or stops as well as whenever a
CPU is added to the system.

To pin to a single CPU, you have to use the range syntax (e.g. ``1-1``)
to differentiate it from a number of CPUs.

``limits.cpu.allowance`` drives either the CFS scheduler quotas when
passed a time constraint, or the generic CPU shares mechanism when
passed a percentage value.

The time constraint (e.g. ``20ms/50ms``) is relative to one CPU worth of
time, so to restrict to two CPUs worth of time, something like
100ms/50ms should be used.

When using a percentage value, the limit will only be applied when under
load and will be used to calculate the scheduler priority for the
instance, relative to any other instance which is using the same CPU(s).

``limits.cpu.priority`` is another knob which is used to compute that
scheduler priority score when a number of instances sharing a set of
CPUs have the same percentage of CPU assigned to them.

Devices configuration
---------------------

LXD will always provide the instance with the basic devices which are
required for a standard POSIX system to work. These aren’t visible in
instance or profile configuration and may not be overridden.

Those include:

-  ``/dev/null`` (character device)
-  ``/dev/zero`` (character device)
-  ``/dev/full`` (character device)
-  ``/dev/console`` (character device)
-  ``/dev/tty`` (character device)
-  ``/dev/random`` (character device)
-  ``/dev/urandom`` (character device)
-  ``/dev/net/tun`` (character device)
-  ``/dev/fuse`` (character device)
-  ``lo`` (network interface)

Anything else has to be defined in the instance configuration or in one
of its profiles. The default profile will typically contain a network
interface to become ``eth0`` in the instance.

To add extra devices to an instance, device entries can be added
directly to an instance, or to a profile.

Devices may be added or removed while the instance is running.

Every device entry is identified by a unique name. If the same name is
used in a subsequent profile or in the instance’s own configuration, the
whole entry is overridden by the new definition.

Device entries are added to an instance through:

.. code:: bash

   lxc config device add <instance> <name> <type> [key=value]...

or to a profile with:

.. code:: bash

   lxc profile device add <profile> <name> <type> [key=value]...

Device types
------------

LXD supports the following device types:

============= ===================================== =========
==============================
ID (database) Name                                  Condition Description
============= ===================================== =========
==============================
0             `none <#type-none>`__                 -         Inheritance blocker
1             `nic <#type-nic>`__                   -         Network interface
2             `disk <#type-disk>`__                 -         Mountpoint inside the instance
3             `unix-char <#type-unix-char>`__       container Unix character device
4             `unix-block <#type-unix-block>`__     container Unix block device
5             `usb <#type-usb>`__                   -         USB device
6             `gpu <#type-gpu>`__                   -         GPU device
7             `infiniband <#type-infiniband>`__     container Infiniband device
8             `proxy <#type-proxy>`__               container Proxy device
9             `unix-hotplug <#type-unix-hotplug>`__ container Unix hotplug device
10            `tpm <#type-tpm>`__                   -         TPM device
11            `pci <#type-pci>`__                   VM        PCI device
============= ===================================== =========
==============================

Type: none
~~~~~~~~~~

Supported instance types: container, VM

A none type device doesn’t have any property and doesn’t create anything
inside the instance.

It’s only purpose it to stop inheritance of devices coming from
profiles.

To do so, just add a none type device with the same name of the one you
wish to skip inheriting. It can be added in a profile being applied
after the profile it originated from or directly on the instance.

Type: nic
~~~~~~~~~

LXD supports several different kinds of network devices (referred to as
Network Interface Controller or NIC).

When adding a network device to an instance, there are two ways to
specify the type of device you want to add; either by specifying the
``nictype`` property or using the ``network`` property.

Specifying a NIC using the ``network`` property
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

When specifying the ``network`` property, the NIC is linked to an
existing managed network and the ``nictype`` is automatically detected
based on the network’s type.

Some of the NICs properties are inherited from the network rather than
being customisable for each NIC.

These are detailed in the “Managed” column in the NIC specific sections
below.

NICs Available:
^^^^^^^^^^^^^^^

See the NIC’s settings below for details about which properties are
available.

The following NICs can be specified using the ``nictype`` or ``network``
properties:

-  `bridged <#nic-bridged>`__: Uses an existing bridge on the host and
   creates a virtual device pair to connect the host bridge to the
   instance.
-  `macvlan <#nic-macvlan>`__: Sets up a new network device based on an
   existing one but using a different MAC address.
-  `sriov <#nic-sriov>`__: Passes a virtual function of an SR-IOV
   enabled physical network device into the instance.

The following NICs can be specified using only the ``network`` property:

-  `ovn <#nic-ovn>`__: Uses an existing OVN network and creates a
   virtual device pair to connect the instance to it.

The following NICs can be specified using only the ``nictype`` property:

-  `physical <#nic-physical>`__: Straight physical device passthrough
   from the host. The targeted device will vanish from the host and
   appear in the instance.
-  `ipvlan <#nic-ipvlan>`__: Sets up a new network device based on an
   existing one using the same MAC address but a different IP.
-  `p2p <#nic-p2p>`__: Creates a virtual device pair, putting one side
   in the instance and leaving the other side on the host.
-  `routed <#nic-routed>`__: Creates a virtual device pair to connect
   the host to the instance and sets up static routes and proxy ARP/NDP
   entries to allow the instance to join the network of a designated
   parent interface.

nic: bridged
^^^^^^^^^^^^

Supported instance types: container, VM

Selected using: ``nictype``, ``network``

Uses an existing bridge on the host and creates a virtual device pair to
connect the host bridge to the instance.

Device configuration properties:

+-----------+-----------+-----------+-----------+-----------+-----------+
| Key       | Type      | Default   | Required  | Managed   | Descripti |
|           |           |           |           |           | on        |
+===========+===========+===========+===========+===========+===========+
| parent    | string    | -         | yes       | yes       | The name  |
|           |           |           |           |           | of the    |
|           |           |           |           |           | host      |
|           |           |           |           |           | device    |
+-----------+-----------+-----------+-----------+-----------+-----------+
| network   | string    | -         | yes       | no        | The LXD   |
|           |           |           |           |           | network   |
|           |           |           |           |           | to link   |
|           |           |           |           |           | device to |
|           |           |           |           |           | (instead  |
|           |           |           |           |           | of        |
|           |           |           |           |           | parent)   |
+-----------+-----------+-----------+-----------+-----------+-----------+
| name      | string    | kernel    | no        | no        | The name  |
|           |           | assigned  |           |           | of the    |
|           |           |           |           |           | interface |
|           |           |           |           |           | inside    |
|           |           |           |           |           | the       |
|           |           |           |           |           | instance  |
+-----------+-----------+-----------+-----------+-----------+-----------+
| mtu       | integer   | parent    | no        | yes       | The MTU   |
|           |           | MTU       |           |           | of the    |
|           |           |           |           |           | new       |
|           |           |           |           |           | interface |
+-----------+-----------+-----------+-----------+-----------+-----------+
| hwaddr    | string    | randomly  | no        | no        | The MAC   |
|           |           | assigned  |           |           | address   |
|           |           |           |           |           | of the    |
|           |           |           |           |           | new       |
|           |           |           |           |           | interface |
+-----------+-----------+-----------+-----------+-----------+-----------+
| host_name | string    | randomly  | no        | no        | The name  |
|           |           | assigned  |           |           | of the    |
|           |           |           |           |           | interface |
|           |           |           |           |           | inside    |
|           |           |           |           |           | the host  |
+-----------+-----------+-----------+-----------+-----------+-----------+
| limits.in | string    | -         | no        | no        | I/O limit |
| gress     |           |           |           |           | in bit/s  |
|           |           |           |           |           | for       |
|           |           |           |           |           | incoming  |
|           |           |           |           |           | traffic   |
|           |           |           |           |           | (various  |
|           |           |           |           |           | suffixes  |
|           |           |           |           |           | supported |
|           |           |           |           |           | ,         |
|           |           |           |           |           | see       |
|           |           |           |           |           | below)    |
+-----------+-----------+-----------+-----------+-----------+-----------+
| limits.eg | string    | -         | no        | no        | I/O limit |
| ress      |           |           |           |           | in bit/s  |
|           |           |           |           |           | for       |
|           |           |           |           |           | outgoing  |
|           |           |           |           |           | traffic   |
|           |           |           |           |           | (various  |
|           |           |           |           |           | suffixes  |
|           |           |           |           |           | supported |
|           |           |           |           |           | ,         |
|           |           |           |           |           | see       |
|           |           |           |           |           | below)    |
+-----------+-----------+-----------+-----------+-----------+-----------+
| limits.ma | string    | -         | no        | no        | Same as   |
| x         |           |           |           |           | modifying |
|           |           |           |           |           | both      |
|           |           |           |           |           | limits.in |
|           |           |           |           |           | gress     |
|           |           |           |           |           | and       |
|           |           |           |           |           | limits.eg |
|           |           |           |           |           | ress      |
+-----------+-----------+-----------+-----------+-----------+-----------+
| ipv4.addr | string    | -         | no        | no        | An IPv4   |
| ess       |           |           |           |           | address   |
|           |           |           |           |           | to assign |
|           |           |           |           |           | to the    |
|           |           |           |           |           | instance  |
|           |           |           |           |           | through   |
|           |           |           |           |           | DHCP      |
+-----------+-----------+-----------+-----------+-----------+-----------+
| ipv6.addr | string    | -         | no        | no        | An IPv6   |
| ess       |           |           |           |           | address   |
|           |           |           |           |           | to assign |
|           |           |           |           |           | to the    |
|           |           |           |           |           | instance  |
|           |           |           |           |           | through   |
|           |           |           |           |           | DHCP      |
+-----------+-----------+-----------+-----------+-----------+-----------+
| ipv4.rout | string    | -         | no        | no        | Comma     |
| es        |           |           |           |           | delimited |
|           |           |           |           |           | list of   |
|           |           |           |           |           | IPv4      |
|           |           |           |           |           | static    |
|           |           |           |           |           | routes to |
|           |           |           |           |           | add on    |
|           |           |           |           |           | host to   |
|           |           |           |           |           | NIC       |
+-----------+-----------+-----------+-----------+-----------+-----------+
| ipv6.rout | string    | -         | no        | no        | Comma     |
| es        |           |           |           |           | delimited |
|           |           |           |           |           | list of   |
|           |           |           |           |           | IPv6      |
|           |           |           |           |           | static    |
|           |           |           |           |           | routes to |
|           |           |           |           |           | add on    |
|           |           |           |           |           | host to   |
|           |           |           |           |           | NIC       |
+-----------+-----------+-----------+-----------+-----------+-----------+
| ipv4.rout | string    | -         | no        | no        | Comma     |
| es.extern |           |           |           |           | delimited |
| al        |           |           |           |           | list of   |
|           |           |           |           |           | IPv4      |
|           |           |           |           |           | static    |
|           |           |           |           |           | routes to |
|           |           |           |           |           | route to  |
|           |           |           |           |           | the NIC   |
|           |           |           |           |           | and       |
|           |           |           |           |           | publish   |
|           |           |           |           |           | on uplink |
|           |           |           |           |           | network   |
|           |           |           |           |           | (BGP)     |
+-----------+-----------+-----------+-----------+-----------+-----------+
| ipv6.rout | string    | -         | no        | no        | Comma     |
| es.extern |           |           |           |           | delimited |
| al        |           |           |           |           | list of   |
|           |           |           |           |           | IPv6      |
|           |           |           |           |           | static    |
|           |           |           |           |           | routes to |
|           |           |           |           |           | route to  |
|           |           |           |           |           | the NIC   |
|           |           |           |           |           | and       |
|           |           |           |           |           | publish   |
|           |           |           |           |           | on uplink |
|           |           |           |           |           | network   |
|           |           |           |           |           | (BGP)     |
+-----------+-----------+-----------+-----------+-----------+-----------+
| security. | boolean   | false     | no        | no        | Prevent   |
| mac_filte |           |           |           |           | the       |
| ring      |           |           |           |           | instance  |
|           |           |           |           |           | from      |
|           |           |           |           |           | spoofing  |
|           |           |           |           |           | another’s |
|           |           |           |           |           | MAC       |
|           |           |           |           |           | address   |
+-----------+-----------+-----------+-----------+-----------+-----------+
| security. | boolean   | false     | no        | no        | Prevent   |
| ipv4_filt |           |           |           |           | the       |
| ering     |           |           |           |           | instance  |
|           |           |           |           |           | from      |
|           |           |           |           |           | spoofing  |
|           |           |           |           |           | another’s |
|           |           |           |           |           | IPv4      |
|           |           |           |           |           | address   |
|           |           |           |           |           | (enables  |
|           |           |           |           |           | mac_filte |
|           |           |           |           |           | ring)     |
+-----------+-----------+-----------+-----------+-----------+-----------+
| security. | boolean   | false     | no        | no        | Prevent   |
| ipv6_filt |           |           |           |           | the       |
| ering     |           |           |           |           | instance  |
|           |           |           |           |           | from      |
|           |           |           |           |           | spoofing  |
|           |           |           |           |           | another’s |
|           |           |           |           |           | IPv6      |
|           |           |           |           |           | address   |
|           |           |           |           |           | (enables  |
|           |           |           |           |           | mac_filte |
|           |           |           |           |           | ring)     |
+-----------+-----------+-----------+-----------+-----------+-----------+
| maas.subn | string    | -         | no        | yes       | MAAS IPv4 |
| et.ipv4   |           |           |           |           | subnet to |
|           |           |           |           |           | register  |
|           |           |           |           |           | the       |
|           |           |           |           |           | instance  |
|           |           |           |           |           | in        |
+-----------+-----------+-----------+-----------+-----------+-----------+
| maas.subn | string    | -         | no        | yes       | MAAS IPv6 |
| et.ipv6   |           |           |           |           | subnet to |
|           |           |           |           |           | register  |
|           |           |           |           |           | the       |
|           |           |           |           |           | instance  |
|           |           |           |           |           | in        |
+-----------+-----------+-----------+-----------+-----------+-----------+
| boot.prio | integer   | -         | no        | no        | Boot      |
| rity      |           |           |           |           | priority  |
|           |           |           |           |           | for VMs   |
|           |           |           |           |           | (higher   |
|           |           |           |           |           | boots     |
|           |           |           |           |           | first)    |
+-----------+-----------+-----------+-----------+-----------+-----------+
| vlan      | integer   | -         | no        | no        | The VLAN  |
|           |           |           |           |           | ID to use |
|           |           |           |           |           | for       |
|           |           |           |           |           | untagged  |
|           |           |           |           |           | traffic   |
|           |           |           |           |           | (Can be   |
|           |           |           |           |           | ``none``  |
|           |           |           |           |           | to remove |
|           |           |           |           |           | port from |
|           |           |           |           |           | default   |
|           |           |           |           |           | VLAN)     |
+-----------+-----------+-----------+-----------+-----------+-----------+
| vlan.tagg | integer   | -         | no        | no        | Comma     |
| ed        |           |           |           |           | delimited |
|           |           |           |           |           | list of   |
|           |           |           |           |           | VLAN IDs  |
|           |           |           |           |           | to join   |
|           |           |           |           |           | for       |
|           |           |           |           |           | tagged    |
|           |           |           |           |           | traffic   |
+-----------+-----------+-----------+-----------+-----------+-----------+
| security. | boolean   | false     | no        | no        | Prevent   |
| port_isol |           |           |           |           | the NIC   |
| ation     |           |           |           |           | from      |
|           |           |           |           |           | communica |
|           |           |           |           |           | ting      |
|           |           |           |           |           | with      |
|           |           |           |           |           | other     |
|           |           |           |           |           | NICs in   |
|           |           |           |           |           | the       |
|           |           |           |           |           | network   |
|           |           |           |           |           | that have |
|           |           |           |           |           | port      |
|           |           |           |           |           | isolation |
|           |           |           |           |           | enabled   |
+-----------+-----------+-----------+-----------+-----------+-----------+

nic: macvlan
^^^^^^^^^^^^

Supported instance types: container, VM

Selected using: ``nictype``, ``network``

Sets up a new network device based on an existing one but using a
different MAC address.

Device configuration properties:

+-----------+-----------+-----------+-----------+-----------+-----------+
| Key       | Type      | Default   | Required  | Managed   | Descripti |
|           |           |           |           |           | on        |
+===========+===========+===========+===========+===========+===========+
| parent    | string    | -         | yes       | yes       | The name  |
|           |           |           |           |           | of the    |
|           |           |           |           |           | host      |
|           |           |           |           |           | device    |
+-----------+-----------+-----------+-----------+-----------+-----------+
| network   | string    | -         | yes       | no        | The LXD   |
|           |           |           |           |           | network   |
|           |           |           |           |           | to link   |
|           |           |           |           |           | device to |
|           |           |           |           |           | (instead  |
|           |           |           |           |           | of        |
|           |           |           |           |           | parent)   |
+-----------+-----------+-----------+-----------+-----------+-----------+
| name      | string    | kernel    | no        | no        | The name  |
|           |           | assigned  |           |           | of the    |
|           |           |           |           |           | interface |
|           |           |           |           |           | inside    |
|           |           |           |           |           | the       |
|           |           |           |           |           | instance  |
+-----------+-----------+-----------+-----------+-----------+-----------+
| mtu       | integer   | parent    | no        | yes       | The MTU   |
|           |           | MTU       |           |           | of the    |
|           |           |           |           |           | new       |
|           |           |           |           |           | interface |
+-----------+-----------+-----------+-----------+-----------+-----------+
| hwaddr    | string    | randomly  | no        | no        | The MAC   |
|           |           | assigned  |           |           | address   |
|           |           |           |           |           | of the    |
|           |           |           |           |           | new       |
|           |           |           |           |           | interface |
+-----------+-----------+-----------+-----------+-----------+-----------+
| vlan      | integer   | -         | no        | no        | The VLAN  |
|           |           |           |           |           | ID to     |
|           |           |           |           |           | attach to |
+-----------+-----------+-----------+-----------+-----------+-----------+
| gvrp      | boolean   | false     | no        | no        | Register  |
|           |           |           |           |           | VLAN      |
|           |           |           |           |           | using     |
|           |           |           |           |           | GARP VLAN |
|           |           |           |           |           | Registrat |
|           |           |           |           |           | ion       |
|           |           |           |           |           | Protocol  |
+-----------+-----------+-----------+-----------+-----------+-----------+
| maas.subn | string    | -         | no        | yes       | MAAS IPv4 |
| et.ipv4   |           |           |           |           | subnet to |
|           |           |           |           |           | register  |
|           |           |           |           |           | the       |
|           |           |           |           |           | instance  |
|           |           |           |           |           | in        |
+-----------+-----------+-----------+-----------+-----------+-----------+
| maas.subn | string    | -         | no        | yes       | MAAS IPv6 |
| et.ipv6   |           |           |           |           | subnet to |
|           |           |           |           |           | register  |
|           |           |           |           |           | the       |
|           |           |           |           |           | instance  |
|           |           |           |           |           | in        |
+-----------+-----------+-----------+-----------+-----------+-----------+
| boot.prio | integer   | -         | no        | no        | Boot      |
| rity      |           |           |           |           | priority  |
|           |           |           |           |           | for VMs   |
|           |           |           |           |           | (higher   |
|           |           |           |           |           | boots     |
|           |           |           |           |           | first)    |
+-----------+-----------+-----------+-----------+-----------+-----------+

nic: sriov
^^^^^^^^^^

Supported instance types: container, VM

Selected using: ``nictype``, ``network``

Passes a virtual function of an SR-IOV enabled physical network device
into the instance.

Device configuration properties:

+-----------+-----------+-----------+-----------+-----------+-----------+
| Key       | Type      | Default   | Required  | Managed   | Descripti |
|           |           |           |           |           | on        |
+===========+===========+===========+===========+===========+===========+
| parent    | string    | -         | yes       | yes       | The name  |
|           |           |           |           |           | of the    |
|           |           |           |           |           | host      |
|           |           |           |           |           | device    |
+-----------+-----------+-----------+-----------+-----------+-----------+
| network   | string    | -         | yes       | no        | The LXD   |
|           |           |           |           |           | network   |
|           |           |           |           |           | to link   |
|           |           |           |           |           | device to |
|           |           |           |           |           | (instead  |
|           |           |           |           |           | of        |
|           |           |           |           |           | parent)   |
+-----------+-----------+-----------+-----------+-----------+-----------+
| name      | string    | kernel    | no        | no        | The name  |
|           |           | assigned  |           |           | of the    |
|           |           |           |           |           | interface |
|           |           |           |           |           | inside    |
|           |           |           |           |           | the       |
|           |           |           |           |           | instance  |
+-----------+-----------+-----------+-----------+-----------+-----------+
| mtu       | integer   | kernel    | no        | yes       | The MTU   |
|           |           | assigned  |           |           | of the    |
|           |           |           |           |           | new       |
|           |           |           |           |           | interface |
+-----------+-----------+-----------+-----------+-----------+-----------+
| hwaddr    | string    | randomly  | no        | no        | The MAC   |
|           |           | assigned  |           |           | address   |
|           |           |           |           |           | of the    |
|           |           |           |           |           | new       |
|           |           |           |           |           | interface |
+-----------+-----------+-----------+-----------+-----------+-----------+
| security. | boolean   | false     | no        | no        | Prevent   |
| mac_filte |           |           |           |           | the       |
| ring      |           |           |           |           | instance  |
|           |           |           |           |           | from      |
|           |           |           |           |           | spoofing  |
|           |           |           |           |           | another’s |
|           |           |           |           |           | MAC       |
|           |           |           |           |           | address   |
+-----------+-----------+-----------+-----------+-----------+-----------+
| vlan      | integer   | -         | no        | no        | The VLAN  |
|           |           |           |           |           | ID to     |
|           |           |           |           |           | attach to |
+-----------+-----------+-----------+-----------+-----------+-----------+
| maas.subn | string    | -         | no        | yes       | MAAS IPv4 |
| et.ipv4   |           |           |           |           | subnet to |
|           |           |           |           |           | register  |
|           |           |           |           |           | the       |
|           |           |           |           |           | instance  |
|           |           |           |           |           | in        |
+-----------+-----------+-----------+-----------+-----------+-----------+
| maas.subn | string    | -         | no        | yes       | MAAS IPv6 |
| et.ipv6   |           |           |           |           | subnet to |
|           |           |           |           |           | register  |
|           |           |           |           |           | the       |
|           |           |           |           |           | instance  |
|           |           |           |           |           | in        |
+-----------+-----------+-----------+-----------+-----------+-----------+
| boot.prio | integer   | -         | no        | no        | Boot      |
| rity      |           |           |           |           | priority  |
|           |           |           |           |           | for VMs   |
|           |           |           |           |           | (higher   |
|           |           |           |           |           | boots     |
|           |           |           |           |           | first)    |
+-----------+-----------+-----------+-----------+-----------+-----------+

nic: ovn
^^^^^^^^

Supported instance types: container, VM

Selected using: ``network``

Uses an existing OVN network and creates a virtual device pair to
connect the instance to it.

Device configuration properties:

+-----------+-----------+-----------+-----------+-----------+-----------+
| Key       | Type      | Default   | Required  | Managed   | Descripti |
|           |           |           |           |           | on        |
+===========+===========+===========+===========+===========+===========+
| network   | string    | -         | yes       | yes       | The LXD   |
|           |           |           |           |           | network   |
|           |           |           |           |           | to link   |
|           |           |           |           |           | device to |
+-----------+-----------+-----------+-----------+-----------+-----------+
| name      | string    | kernel    | no        | no        | The name  |
|           |           | assigned  |           |           | of the    |
|           |           |           |           |           | interface |
|           |           |           |           |           | inside    |
|           |           |           |           |           | the       |
|           |           |           |           |           | instance  |
+-----------+-----------+-----------+-----------+-----------+-----------+
| host_name | string    | randomly  | no        | no        | The name  |
|           |           | assigned  |           |           | of the    |
|           |           |           |           |           | interface |
|           |           |           |           |           | inside    |
|           |           |           |           |           | the host  |
+-----------+-----------+-----------+-----------+-----------+-----------+
| hwaddr    | string    | randomly  | no        | no        | The MAC   |
|           |           | assigned  |           |           | address   |
|           |           |           |           |           | of the    |
|           |           |           |           |           | new       |
|           |           |           |           |           | interface |
+-----------+-----------+-----------+-----------+-----------+-----------+
| ipv4.addr | string    | -         | no        | no        | An IPv4   |
| ess       |           |           |           |           | address   |
|           |           |           |           |           | to assign |
|           |           |           |           |           | to the    |
|           |           |           |           |           | instance  |
|           |           |           |           |           | through   |
|           |           |           |           |           | DHCP      |
+-----------+-----------+-----------+-----------+-----------+-----------+
| ipv6.addr | string    | -         | no        | no        | An IPv6   |
| ess       |           |           |           |           | address   |
|           |           |           |           |           | to assign |
|           |           |           |           |           | to the    |
|           |           |           |           |           | instance  |
|           |           |           |           |           | through   |
|           |           |           |           |           | DHCP      |
+-----------+-----------+-----------+-----------+-----------+-----------+
| ipv4.rout | string    | -         | no        | no        | Comma     |
| es        |           |           |           |           | delimited |
|           |           |           |           |           | list of   |
|           |           |           |           |           | IPv4      |
|           |           |           |           |           | static    |
|           |           |           |           |           | routes to |
|           |           |           |           |           | route to  |
|           |           |           |           |           | the NIC   |
+-----------+-----------+-----------+-----------+-----------+-----------+
| ipv6.rout | string    | -         | no        | no        | Comma     |
| es        |           |           |           |           | delimited |
|           |           |           |           |           | list of   |
|           |           |           |           |           | IPv6      |
|           |           |           |           |           | static    |
|           |           |           |           |           | routes to |
|           |           |           |           |           | route to  |
|           |           |           |           |           | the NIC   |
+-----------+-----------+-----------+-----------+-----------+-----------+
| ipv4.rout | string    | -         | no        | no        | Comma     |
| es.extern |           |           |           |           | delimited |
| al        |           |           |           |           | list of   |
|           |           |           |           |           | IPv4      |
|           |           |           |           |           | static    |
|           |           |           |           |           | routes to |
|           |           |           |           |           | route to  |
|           |           |           |           |           | the NIC   |
|           |           |           |           |           | and       |
|           |           |           |           |           | publish   |
|           |           |           |           |           | on uplink |
|           |           |           |           |           | network   |
+-----------+-----------+-----------+-----------+-----------+-----------+
| ipv6.rout | string    | -         | no        | no        | Comma     |
| es.extern |           |           |           |           | delimited |
| al        |           |           |           |           | list of   |
|           |           |           |           |           | IPv6      |
|           |           |           |           |           | static    |
|           |           |           |           |           | routes to |
|           |           |           |           |           | route to  |
|           |           |           |           |           | the NIC   |
|           |           |           |           |           | and       |
|           |           |           |           |           | publish   |
|           |           |           |           |           | on uplink |
|           |           |           |           |           | network   |
+-----------+-----------+-----------+-----------+-----------+-----------+
| boot.prio | integer   | -         | no        | no        | Boot      |
| rity      |           |           |           |           | priority  |
|           |           |           |           |           | for VMs   |
|           |           |           |           |           | (higher   |
|           |           |           |           |           | boots     |
|           |           |           |           |           | first)    |
+-----------+-----------+-----------+-----------+-----------+-----------+
| security. | string    | -         | no        | no        | Comma     |
| acls      |           |           |           |           | separated |
|           |           |           |           |           | list of   |
|           |           |           |           |           | Network   |
|           |           |           |           |           | ACLs to   |
|           |           |           |           |           | apply     |
+-----------+-----------+-----------+-----------+-----------+-----------+
| security. | string    | reject    | no        | no        | Action to |
| acls.defa |           |           |           |           | use for   |
| ult.ingre |           |           |           |           | ingress   |
| ss.action |           |           |           |           | traffic   |
|           |           |           |           |           | that      |
|           |           |           |           |           | doesn’t   |
|           |           |           |           |           | match any |
|           |           |           |           |           | ACL rule  |
+-----------+-----------+-----------+-----------+-----------+-----------+
| security. | string    | reject    | no        | no        | Action to |
| acls.defa |           |           |           |           | use for   |
| ult.egres |           |           |           |           | egress    |
| s.action  |           |           |           |           | traffic   |
|           |           |           |           |           | that      |
|           |           |           |           |           | doesn’t   |
|           |           |           |           |           | match any |
|           |           |           |           |           | ACL rule  |
+-----------+-----------+-----------+-----------+-----------+-----------+
| security. | boolean   | false     | no        | no        | Whether   |
| acls.defa |           |           |           |           | to log    |
| ult.ingre |           |           |           |           | ingress   |
| ss.logged |           |           |           |           | traffic   |
|           |           |           |           |           | that      |
|           |           |           |           |           | doesn’t   |
|           |           |           |           |           | match any |
|           |           |           |           |           | ACL rule  |
+-----------+-----------+-----------+-----------+-----------+-----------+
| security. | boolean   | false     | no        | no        | Whether   |
| acls.defa |           |           |           |           | to log    |
| ult.egres |           |           |           |           | egress    |
| s.logged  |           |           |           |           | traffic   |
|           |           |           |           |           | that      |
|           |           |           |           |           | doesn’t   |
|           |           |           |           |           | match any |
|           |           |           |           |           | ACL rule  |
+-----------+-----------+-----------+-----------+-----------+-----------+

nic: physical
^^^^^^^^^^^^^

Supported instance types: container, VM

Selected using: ``nictype``

Straight physical device passthrough from the host. The targeted device
will vanish from the host and appear in the instance.

Device configuration properties:

================ ======= ================= ========
===================================================
Key              Type    Default           Required Description
================ ======= ================= ========
===================================================
parent           string  -                 yes      The name of the host device
name             string  kernel assigned   no       The name of the interface inside the instance
mtu              integer parent MTU        no       The MTU of the new interface
hwaddr           string  randomly assigned no       The MAC address of the new interface
vlan             integer -                 no       The VLAN ID to attach to
gvrp             boolean false             no       Register VLAN using GARP VLAN Registration Protocol
maas.subnet.ipv4 string  -                 no       MAAS IPv4 subnet to register the instance in
maas.subnet.ipv6 string  -                 no       MAAS IPv6 subnet to register the instance in
boot.priority    integer -                 no       Boot priority for VMs (higher boots first)
================ ======= ================= ========
===================================================

nic: ipvlan
^^^^^^^^^^^

Supported instance types: container

Selected using: ``nictype``

Sets up a new network device based on an existing one using the same MAC
address but a different IP.

LXD currently supports IPVLAN in L2 and L3S mode.

In this mode, the gateway is automatically set by LXD, however IP
addresses must be manually specified using either one or both of
``ipv4.address`` and ``ipv6.address`` settings before instance is
started.

For DNS, the nameservers need to be configured inside the instance, as
these will not automatically be set.

It requires the following sysctls to be set:

If using IPv4 addresses:

::

   net.ipv4.conf.<parent>.forwarding=1

If using IPv6 addresses:

::

   net.ipv6.conf.<parent>.forwarding=1
   net.ipv6.conf.<parent>.proxy_ndp=1

Device configuration properties:

+-------------+-------------+-------------+-------------+-------------+
| Key         | Type        | Default     | Required    | Description |
+=============+=============+=============+=============+=============+
| parent      | string      | -           | yes         | The name of |
|             |             |             |             | the host    |
|             |             |             |             | device      |
+-------------+-------------+-------------+-------------+-------------+
| name        | string      | kernel      | no          | The name of |
|             |             | assigned    |             | the         |
|             |             |             |             | interface   |
|             |             |             |             | inside the  |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+
| mtu         | integer     | parent MTU  | no          | The MTU of  |
|             |             |             |             | the new     |
|             |             |             |             | interface   |
+-------------+-------------+-------------+-------------+-------------+
| mode        | string      | l3s         | no          | The IPVLAN  |
|             |             |             |             | mode        |
|             |             |             |             | (either     |
|             |             |             |             | ``l2`` or   |
|             |             |             |             | ``l3s``)    |
+-------------+-------------+-------------+-------------+-------------+
| hwaddr      | string      | randomly    | no          | The MAC     |
|             |             | assigned    |             | address of  |
|             |             |             |             | the new     |
|             |             |             |             | interface   |
+-------------+-------------+-------------+-------------+-------------+
| ipv4.addres | string      | -           | no          | Comma       |
| s           |             |             |             | delimited   |
|             |             |             |             | list of     |
|             |             |             |             | IPv4 static |
|             |             |             |             | addresses   |
|             |             |             |             | to add to   |
|             |             |             |             | the         |
|             |             |             |             | instance.   |
|             |             |             |             | In ``l2``   |
|             |             |             |             | mode these  |
|             |             |             |             | can be      |
|             |             |             |             | specified   |
|             |             |             |             | as CIDR     |
|             |             |             |             | values or   |
|             |             |             |             | singular    |
|             |             |             |             | addresses   |
|             |             |             |             | (if         |
|             |             |             |             | singular a  |
|             |             |             |             | subnet of   |
|             |             |             |             | /24 is      |
|             |             |             |             | used).      |
+-------------+-------------+-------------+-------------+-------------+
| ipv4.gatewa | string      | auto        | no          | In ``l3s``  |
| y           |             |             |             | mode,       |
|             |             |             |             | whether to  |
|             |             |             |             | add an      |
|             |             |             |             | automatic   |
|             |             |             |             | default     |
|             |             |             |             | IPv4        |
|             |             |             |             | gateway,    |
|             |             |             |             | can be      |
|             |             |             |             | ``auto`` or |
|             |             |             |             | ``none``.   |
|             |             |             |             | In ``l2``   |
|             |             |             |             | mode        |
|             |             |             |             | specifies   |
|             |             |             |             | the IPv4    |
|             |             |             |             | address of  |
|             |             |             |             | the         |
|             |             |             |             | gateway.    |
+-------------+-------------+-------------+-------------+-------------+
| ipv4.host_t | integer     | -           | no          | The custom  |
| able        |             |             |             | policy      |
|             |             |             |             | routing     |
|             |             |             |             | table ID to |
|             |             |             |             | add IPv4    |
|             |             |             |             | static      |
|             |             |             |             | routes to   |
|             |             |             |             | (in         |
|             |             |             |             | addition to |
|             |             |             |             | main        |
|             |             |             |             | routing     |
|             |             |             |             | table).     |
+-------------+-------------+-------------+-------------+-------------+
| ipv6.addres | string      | -           | no          | Comma       |
| s           |             |             |             | delimited   |
|             |             |             |             | list of     |
|             |             |             |             | IPv6 static |
|             |             |             |             | addresses   |
|             |             |             |             | to add to   |
|             |             |             |             | the         |
|             |             |             |             | instance.   |
|             |             |             |             | In ``l2``   |
|             |             |             |             | mode these  |
|             |             |             |             | can be      |
|             |             |             |             | specified   |
|             |             |             |             | as CIDR     |
|             |             |             |             | values or   |
|             |             |             |             | singular    |
|             |             |             |             | addresses   |
|             |             |             |             | (if         |
|             |             |             |             | singular a  |
|             |             |             |             | subnet of   |
|             |             |             |             | /64 is      |
|             |             |             |             | used).      |
+-------------+-------------+-------------+-------------+-------------+
| ipv6.gatewa | string      | auto (l3s), | no          | In ``l3s``  |
| y           |             | - (l2)      |             | mode,       |
|             |             |             |             | whether to  |
|             |             |             |             | add an      |
|             |             |             |             | automatic   |
|             |             |             |             | default     |
|             |             |             |             | IPv6        |
|             |             |             |             | gateway,    |
|             |             |             |             | can be      |
|             |             |             |             | ``auto`` or |
|             |             |             |             | ``none``.   |
|             |             |             |             | In ``l2``   |
|             |             |             |             | mode        |
|             |             |             |             | specifies   |
|             |             |             |             | the IPv6    |
|             |             |             |             | address of  |
|             |             |             |             | the         |
|             |             |             |             | gateway.    |
+-------------+-------------+-------------+-------------+-------------+
| ipv6.host_t | integer     | -           | no          | The custom  |
| able        |             |             |             | policy      |
|             |             |             |             | routing     |
|             |             |             |             | table ID to |
|             |             |             |             | add IPv6    |
|             |             |             |             | static      |
|             |             |             |             | routes to   |
|             |             |             |             | (in         |
|             |             |             |             | addition to |
|             |             |             |             | main        |
|             |             |             |             | routing     |
|             |             |             |             | table).     |
+-------------+-------------+-------------+-------------+-------------+
| vlan        | integer     | -           | no          | The VLAN ID |
|             |             |             |             | to attach   |
|             |             |             |             | to          |
+-------------+-------------+-------------+-------------+-------------+
| gvrp        | boolean     | false       | no          | Register    |
|             |             |             |             | VLAN using  |
|             |             |             |             | GARP VLAN   |
|             |             |             |             | Registratio |
|             |             |             |             | n           |
|             |             |             |             | Protocol    |
+-------------+-------------+-------------+-------------+-------------+

nic: p2p
^^^^^^^^

Supported instance types: container, VM

Selected using: ``nictype``

Creates a virtual device pair, putting one side in the instance and
leaving the other side on the host.

Device configuration properties:

+-------------+-------------+-------------+-------------+-------------+
| Key         | Type        | Default     | Required    | Description |
+=============+=============+=============+=============+=============+
| name        | string      | kernel      | no          | The name of |
|             |             | assigned    |             | the         |
|             |             |             |             | interface   |
|             |             |             |             | inside the  |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+
| mtu         | integer     | kernel      | no          | The MTU of  |
|             |             | assigned    |             | the new     |
|             |             |             |             | interface   |
+-------------+-------------+-------------+-------------+-------------+
| hwaddr      | string      | randomly    | no          | The MAC     |
|             |             | assigned    |             | address of  |
|             |             |             |             | the new     |
|             |             |             |             | interface   |
+-------------+-------------+-------------+-------------+-------------+
| host_name   | string      | randomly    | no          | The name of |
|             |             | assigned    |             | the         |
|             |             |             |             | interface   |
|             |             |             |             | inside the  |
|             |             |             |             | host        |
+-------------+-------------+-------------+-------------+-------------+
| limits.ingr | string      | -           | no          | I/O limit   |
| ess         |             |             |             | in bit/s    |
|             |             |             |             | for         |
|             |             |             |             | incoming    |
|             |             |             |             | traffic     |
|             |             |             |             | (various    |
|             |             |             |             | suffixes    |
|             |             |             |             | supported,  |
|             |             |             |             | see below)  |
+-------------+-------------+-------------+-------------+-------------+
| limits.egre | string      | -           | no          | I/O limit   |
| ss          |             |             |             | in bit/s    |
|             |             |             |             | for         |
|             |             |             |             | outgoing    |
|             |             |             |             | traffic     |
|             |             |             |             | (various    |
|             |             |             |             | suffixes    |
|             |             |             |             | supported,  |
|             |             |             |             | see below)  |
+-------------+-------------+-------------+-------------+-------------+
| limits.max  | string      | -           | no          | Same as     |
|             |             |             |             | modifying   |
|             |             |             |             | both        |
|             |             |             |             | limits.ingr |
|             |             |             |             | ess         |
|             |             |             |             | and         |
|             |             |             |             | limits.egre |
|             |             |             |             | ss          |
+-------------+-------------+-------------+-------------+-------------+
| ipv4.routes | string      | -           | no          | Comma       |
|             |             |             |             | delimited   |
|             |             |             |             | list of     |
|             |             |             |             | IPv4 static |
|             |             |             |             | routes to   |
|             |             |             |             | add on host |
|             |             |             |             | to NIC      |
+-------------+-------------+-------------+-------------+-------------+
| ipv6.routes | string      | -           | no          | Comma       |
|             |             |             |             | delimited   |
|             |             |             |             | list of     |
|             |             |             |             | IPv6 static |
|             |             |             |             | routes to   |
|             |             |             |             | add on host |
|             |             |             |             | to NIC      |
+-------------+-------------+-------------+-------------+-------------+
| boot.priori | integer     | -           | no          | Boot        |
| ty          |             |             |             | priority    |
|             |             |             |             | for VMs     |
|             |             |             |             | (higher     |
|             |             |             |             | boots       |
|             |             |             |             | first)      |
+-------------+-------------+-------------+-------------+-------------+

nic: routed
^^^^^^^^^^^

Supported instance types: container

Selected using: ``nictype``

This NIC type is similar in operation to IPVLAN, in that it allows an
instance to join an external network without needing to configure a
bridge and shares the host’s MAC address.

However it differs from IPVLAN because it does not need IPVLAN support
in the kernel and the host and instance can communicate with each other.

It will also respect netfilter rules on the host and will use the host’s
routing table to route packets which can be useful if the host is
connected to multiple networks.

IP addresses must be manually specified using either one or both of
``ipv4.address`` and ``ipv6.address`` settings before the instance is
started.

It sets up a veth pair between host and instance and then configures the
following link-local gateway IPs on the host end which are then set as
the default gateways in the instance:

169.254.0.1 fe80::1

It then configures static routes on the host pointing to the instance’s
veth interface for all of the instance’s IPs.

This nic can operate with and without a ``parent`` network interface
set.

With the ``parent`` network interface set proxy ARP/NDP entries of the
instance’s IPs are added to the parent interface allowing the instance
to join the parent interface’s network at layer 2.

For DNS, the nameservers need to be configured inside the instance, as
these will not automatically be set.

It requires the following sysctls to be set:

If using IPv4 addresses:

::

   net.ipv4.conf.<parent>.forwarding=1

If using IPv6 addresses:

::

   net.ipv6.conf.all.forwarding=1
   net.ipv6.conf.<parent>.forwarding=1
   net.ipv6.conf.all.proxy_ndp=1
   net.ipv6.conf.<parent>.proxy_ndp=1

Each NIC device can have multiple IP addresses added to them. However it
may be desirable to utilise multiple ``routed`` NIC interfaces. In these
cases one should set the ``ipv4.gateway`` and ``ipv6.gateway`` values to
“none” on any subsequent interfaces to avoid default gateway conflicts.
It may also be useful to specify a different host-side address for these
subsequent interfaces using ``ipv4.host_address`` and
``ipv6.host_address`` respectively.

Device configuration properties:

+-------------+-------------+-------------+-------------+-------------+
| Key         | Type        | Default     | Required    | Description |
+=============+=============+=============+=============+=============+
| parent      | string      | -           | no          | The name of |
|             |             |             |             | the host    |
|             |             |             |             | device to   |
|             |             |             |             | join the    |
|             |             |             |             | instance to |
+-------------+-------------+-------------+-------------+-------------+
| name        | string      | kernel      | no          | The name of |
|             |             | assigned    |             | the         |
|             |             |             |             | interface   |
|             |             |             |             | inside the  |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+
| host_name   | string      | randomly    | no          | The name of |
|             |             | assigned    |             | the         |
|             |             |             |             | interface   |
|             |             |             |             | inside the  |
|             |             |             |             | host        |
+-------------+-------------+-------------+-------------+-------------+
| mtu         | integer     | parent MTU  | no          | The MTU of  |
|             |             |             |             | the new     |
|             |             |             |             | interface   |
+-------------+-------------+-------------+-------------+-------------+
| hwaddr      | string      | randomly    | no          | The MAC     |
|             |             | assigned    |             | address of  |
|             |             |             |             | the new     |
|             |             |             |             | interface   |
+-------------+-------------+-------------+-------------+-------------+
| limits.ingr | string      | -           | no          | I/O limit   |
| ess         |             |             |             | in bit/s    |
|             |             |             |             | for         |
|             |             |             |             | incoming    |
|             |             |             |             | traffic     |
|             |             |             |             | (various    |
|             |             |             |             | suffixes    |
|             |             |             |             | supported,  |
|             |             |             |             | see below)  |
+-------------+-------------+-------------+-------------+-------------+
| limits.egre | string      | -           | no          | I/O limit   |
| ss          |             |             |             | in bit/s    |
|             |             |             |             | for         |
|             |             |             |             | outgoing    |
|             |             |             |             | traffic     |
|             |             |             |             | (various    |
|             |             |             |             | suffixes    |
|             |             |             |             | supported,  |
|             |             |             |             | see below)  |
+-------------+-------------+-------------+-------------+-------------+
| limits.max  | string      | -           | no          | Same as     |
|             |             |             |             | modifying   |
|             |             |             |             | both        |
|             |             |             |             | limits.ingr |
|             |             |             |             | ess         |
|             |             |             |             | and         |
|             |             |             |             | limits.egre |
|             |             |             |             | ss          |
+-------------+-------------+-------------+-------------+-------------+
| ipv4.addres | string      | -           | no          | Comma       |
| s           |             |             |             | delimited   |
|             |             |             |             | list of     |
|             |             |             |             | IPv4 static |
|             |             |             |             | addresses   |
|             |             |             |             | to add to   |
|             |             |             |             | the         |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+
| ipv4.gatewa | string      | auto        | no          | Whether to  |
| y           |             |             |             | add an      |
|             |             |             |             | automatic   |
|             |             |             |             | default     |
|             |             |             |             | IPv4        |
|             |             |             |             | gateway,    |
|             |             |             |             | can be      |
|             |             |             |             | “auto” or   |
|             |             |             |             | “none”      |
+-------------+-------------+-------------+-------------+-------------+
| ipv4.host_a | string      | 169.254.0.1 | no          | The IPv4    |
| ddress      |             |             |             | address to  |
|             |             |             |             | add to the  |
|             |             |             |             | host-side   |
|             |             |             |             | veth        |
|             |             |             |             | interface.  |
+-------------+-------------+-------------+-------------+-------------+
| ipv4.host_t | integer     | -           | no          | The custom  |
| able        |             |             |             | policy      |
|             |             |             |             | routing     |
|             |             |             |             | table ID to |
|             |             |             |             | add IPv4    |
|             |             |             |             | static      |
|             |             |             |             | routes to   |
|             |             |             |             | (in         |
|             |             |             |             | addition to |
|             |             |             |             | main        |
|             |             |             |             | routing     |
|             |             |             |             | table).     |
+-------------+-------------+-------------+-------------+-------------+
| ipv6.addres | string      | -           | no          | Comma       |
| s           |             |             |             | delimited   |
|             |             |             |             | list of     |
|             |             |             |             | IPv6 static |
|             |             |             |             | addresses   |
|             |             |             |             | to add to   |
|             |             |             |             | the         |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+
| ipv6.gatewa | string      | auto        | no          | Whether to  |
| y           |             |             |             | add an      |
|             |             |             |             | automatic   |
|             |             |             |             | default     |
|             |             |             |             | IPv6        |
|             |             |             |             | gateway,    |
|             |             |             |             | can be      |
|             |             |             |             | “auto” or   |
|             |             |             |             | “none”      |
+-------------+-------------+-------------+-------------+-------------+
| ipv6.host_a | string      | fe80::1     | no          | The IPv6    |
| ddress      |             |             |             | address to  |
|             |             |             |             | add to the  |
|             |             |             |             | host-side   |
|             |             |             |             | veth        |
|             |             |             |             | interface.  |
+-------------+-------------+-------------+-------------+-------------+
| ipv6.host_t | integer     | -           | no          | The custom  |
| able        |             |             |             | policy      |
|             |             |             |             | routing     |
|             |             |             |             | table ID to |
|             |             |             |             | add IPv6    |
|             |             |             |             | static      |
|             |             |             |             | routes to   |
|             |             |             |             | (in         |
|             |             |             |             | addition to |
|             |             |             |             | main        |
|             |             |             |             | routing     |
|             |             |             |             | table).     |
+-------------+-------------+-------------+-------------+-------------+
| vlan        | integer     | -           | no          | The VLAN ID |
|             |             |             |             | to attach   |
|             |             |             |             | to          |
+-------------+-------------+-------------+-------------+-------------+
| gvrp        | boolean     | false       | no          | Register    |
|             |             |             |             | VLAN using  |
|             |             |             |             | GARP VLAN   |
|             |             |             |             | Registratio |
|             |             |             |             | n           |
|             |             |             |             | Protocol    |
+-------------+-------------+-------------+-------------+-------------+

bridged, macvlan or ipvlan for connection to physical network
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

The ``bridged``, ``macvlan`` and ``ipvlan`` interface types can be used
to connect to an existing physical network.

``macvlan`` effectively lets you fork your physical NIC, getting a
second interface that’s then used by the instance. This saves you from
creating a bridge device and veth pairs and usually offers better
performance than a bridge.

The downside to this is that macvlan devices while able to communicate
between themselves and to the outside, aren’t able to talk to their
parent device. This means that you can’t use macvlan if you ever need
your instances to talk to the host itself.

In such case, a bridge is preferable. A bridge will also let you use mac
filtering and I/O limits which cannot be applied to a macvlan device.

``ipvlan`` is similar to ``macvlan``, with the difference being that the
forked device has IPs statically assigned to it and inherits the
parent’s MAC address on the network.

SR-IOV
^^^^^^

The ``sriov`` interface type supports SR-IOV enabled network devices.
These devices associate a set of virtual functions (VFs) with the single
physical function (PF) of the network device. PFs are standard PCIe
functions. VFs on the other hand are very lightweight PCIe functions
that are optimized for data movement. They come with a limited set of
configuration capabilities to prevent changing properties of the PF.
Given that VFs appear as regular PCIe devices to the system they can be
passed to instances just like a regular physical device. The ``sriov``
interface type expects to be passed the name of an SR-IOV enabled
network device on the system via the ``parent`` property. LXD will then
check for any available VFs on the system. By default LXD will allocate
the first free VF it finds. If it detects that either none are enabled
or all currently enabled VFs are in use it will bump the number of
supported VFs to the maximum value and use the first free VF. If all
possible VFs are in use or the kernel or card doesn’t support
incrementing the number of VFs LXD will return an error.

To create a ``sriov`` network device use:

::

   lxc config device add <instance> <device-name> nic nictype=sriov parent=<sriov-enabled-device>

To tell LXD to use a specific unused VF add the ``host_name`` property
and pass it the name of the enabled VF.

MAAS integration
^^^^^^^^^^^^^^^^

If you’re using MAAS to manage the physical network under your LXD host
and want to attach your instances directly to a MAAS managed network,
LXD can be configured to interact with MAAS so that it can track your
instances.

At the daemon level, you must configure ``maas.api.url`` and
``maas.api.key``, then set the ``maas.subnet.ipv4`` and/or
``maas.subnet.ipv6`` keys on the instance or profile’s ``nic`` entry.

This will have LXD register all your instances with MAAS, giving them
proper DHCP leases and DNS records.

If you set the ``ipv4.address`` or ``ipv6.address`` keys on the nic,
then those will be registered as static assignments in MAAS too.

Type: infiniband
~~~~~~~~~~~~~~~~

Supported instance types: container

LXD supports two different kind of network types for infiniband devices:

-  ``physical``: Straight physical device passthrough from the host. The
   targeted device will vanish from the host and appear in the instance.
-  ``sriov``: Passes a virtual function of an SR-IOV enabled physical
   network device into the instance.

Different network interface types have different additional properties,
the current list is:

+-----------+-----------+-----------+-----------+-----------+-----------+
| Key       | Type      | Default   | Required  | Used by   | Descripti |
|           |           |           |           |           | on        |
+===========+===========+===========+===========+===========+===========+
| nictype   | string    | -         | yes       | all       | The       |
|           |           |           |           |           | device    |
|           |           |           |           |           | type, one |
|           |           |           |           |           | of        |
|           |           |           |           |           | “physical |
|           |           |           |           |           | ”,        |
|           |           |           |           |           | or        |
|           |           |           |           |           | “sriov”   |
+-----------+-----------+-----------+-----------+-----------+-----------+
| name      | string    | kernel    | no        | all       | The name  |
|           |           | assigned  |           |           | of the    |
|           |           |           |           |           | interface |
|           |           |           |           |           | inside    |
|           |           |           |           |           | the       |
|           |           |           |           |           | instance  |
+-----------+-----------+-----------+-----------+-----------+-----------+
| hwaddr    | string    | randomly  | no        | all       | The MAC   |
|           |           | assigned  |           |           | address   |
|           |           |           |           |           | of the    |
|           |           |           |           |           | new       |
|           |           |           |           |           | interface |
|           |           |           |           |           | .         |
|           |           |           |           |           | Can be    |
|           |           |           |           |           | either    |
|           |           |           |           |           | full 20   |
|           |           |           |           |           | byte      |
|           |           |           |           |           | variant   |
|           |           |           |           |           | or short  |
|           |           |           |           |           | 8 byte    |
|           |           |           |           |           | variant   |
|           |           |           |           |           | (which    |
|           |           |           |           |           | will only |
|           |           |           |           |           | modify    |
|           |           |           |           |           | the last  |
|           |           |           |           |           | 8 bytes   |
|           |           |           |           |           | of the    |
|           |           |           |           |           | parent    |
|           |           |           |           |           | device)   |
+-----------+-----------+-----------+-----------+-----------+-----------+
| mtu       | integer   | parent    | no        | all       | The MTU   |
|           |           | MTU       |           |           | of the    |
|           |           |           |           |           | new       |
|           |           |           |           |           | interface |
+-----------+-----------+-----------+-----------+-----------+-----------+
| parent    | string    | -         | yes       | physical, | The name  |
|           |           |           |           | sriov     | of the    |
|           |           |           |           |           | host      |
|           |           |           |           |           | device or |
|           |           |           |           |           | bridge    |
+-----------+-----------+-----------+-----------+-----------+-----------+

To create a ``physical`` ``infiniband`` device use:

::

   lxc config device add <instance> <device-name> infiniband nictype=physical parent=<device>

SR-IOV with infiniband devices
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Infiniband devices do support SR-IOV but in contrast to other SR-IOV
enabled devices infiniband does not support dynamic device creation in
SR-IOV mode. This means users need to pre-configure the number of
virtual functions by configuring the corresponding kernel module.

To create a ``sriov`` ``infiniband`` device use:

::

   lxc config device add <instance> <device-name> infiniband nictype=sriov parent=<sriov-enabled-device>

Type: disk
~~~~~~~~~~

Supported instance types: container, VM

Disk entries are essentially mountpoints inside the instance. They can
either be a bind-mount of an existing file or directory on the host, or
if the source is a block device, a regular mount.

LXD supports the following additional source types:

-  Ceph-rbd: Mount from existing ceph RBD device that is externally
   managed. LXD can use ceph to manage an internal file system for the
   instance, but in the event that a user has a previously existing ceph
   RBD that they would like use for this instance, they can use this
   command. Example command

::

   lxc config device add <instance> ceph-rbd1 disk source=ceph:<my_pool>/<my-volume> ceph.user_name=<username> ceph.cluster_name=<username> path=/ceph

-  Ceph-fs: Mount from existing ceph FS device that is externally
   managed. LXD can use ceph to manage an internal file system for the
   instance, but in the event that a user has a previously existing ceph
   file sys that they would like use for this instancer, they can use
   this command. Example command.

::

   lxc config device add <instance> ceph-fs1 disk source=cephfs:<my-fs>/<some-path> ceph.user_name=<username> ceph.cluster_name=<username> path=/cephfs

-  VM cloud-init: Generate a cloud-init config ISO from the
   user.vendor-data, user.user-data and user.meta-data config keys and
   attach to the VM so that cloud-init running inside the VM guest will
   detect the drive on boot and apply the config. Only applicable to
   virtual-machine instances. Example command.

::

   lxc config device add <instance> config disk source=cloud-init:config

Currently only the root disk (path=/) and config drive
(source=cloud-init:config) are supported with virtual machines.

The following properties exist:

+-------------+-------------+-------------+-------------+-------------+
| Key         | Type        | Default     | Required    | Description |
+=============+=============+=============+=============+=============+
| limits.read | string      | -           | no          | I/O limit   |
|             |             |             |             | in byte/s   |
|             |             |             |             | (various    |
|             |             |             |             | suffixes    |
|             |             |             |             | supported,  |
|             |             |             |             | see below)  |
|             |             |             |             | or in iops  |
|             |             |             |             | (must be    |
|             |             |             |             | suffixed    |
|             |             |             |             | with        |
|             |             |             |             | “iops”)     |
+-------------+-------------+-------------+-------------+-------------+
| limits.writ | string      | -           | no          | I/O limit   |
| e           |             |             |             | in byte/s   |
|             |             |             |             | (various    |
|             |             |             |             | suffixes    |
|             |             |             |             | supported,  |
|             |             |             |             | see below)  |
|             |             |             |             | or in iops  |
|             |             |             |             | (must be    |
|             |             |             |             | suffixed    |
|             |             |             |             | with        |
|             |             |             |             | “iops”)     |
+-------------+-------------+-------------+-------------+-------------+
| limits.max  | string      | -           | no          | Same as     |
|             |             |             |             | modifying   |
|             |             |             |             | both        |
|             |             |             |             | limits.read |
|             |             |             |             | and         |
|             |             |             |             | limits.writ |
|             |             |             |             | e           |
+-------------+-------------+-------------+-------------+-------------+
| path        | string      | -           | yes         | Path inside |
|             |             |             |             | the         |
|             |             |             |             | instance    |
|             |             |             |             | where the   |
|             |             |             |             | disk will   |
|             |             |             |             | be mounted  |
|             |             |             |             | (only for   |
|             |             |             |             | containers) |
|             |             |             |             | .           |
+-------------+-------------+-------------+-------------+-------------+
| source      | string      | -           | yes         | Path on the |
|             |             |             |             | host,       |
|             |             |             |             | either to a |
|             |             |             |             | file/direct |
|             |             |             |             | ory         |
|             |             |             |             | or to a     |
|             |             |             |             | block       |
|             |             |             |             | device      |
+-------------+-------------+-------------+-------------+-------------+
| required    | boolean     | true        | no          | Controls    |
|             |             |             |             | whether to  |
|             |             |             |             | fail if the |
|             |             |             |             | source      |
|             |             |             |             | doesn’t     |
|             |             |             |             | exist       |
+-------------+-------------+-------------+-------------+-------------+
| readonly    | boolean     | false       | no          | Controls    |
|             |             |             |             | whether to  |
|             |             |             |             | make the    |
|             |             |             |             | mount       |
|             |             |             |             | read-only   |
+-------------+-------------+-------------+-------------+-------------+
| size        | string      | -           | no          | Disk size   |
|             |             |             |             | in bytes    |
|             |             |             |             | (various    |
|             |             |             |             | suffixes    |
|             |             |             |             | supported,  |
|             |             |             |             | see below). |
|             |             |             |             | This is     |
|             |             |             |             | only        |
|             |             |             |             | supported   |
|             |             |             |             | for the     |
|             |             |             |             | rootfs (/)  |
+-------------+-------------+-------------+-------------+-------------+
| size.state  | string      | -           | no          | Same as     |
|             |             |             |             | size above  |
|             |             |             |             | but applies |
|             |             |             |             | to the      |
|             |             |             |             | filesystem  |
|             |             |             |             | volume used |
|             |             |             |             | for saving  |
|             |             |             |             | runtime     |
|             |             |             |             | state in    |
|             |             |             |             | virtual     |
|             |             |             |             | machines.   |
+-------------+-------------+-------------+-------------+-------------+
| recursive   | boolean     | false       | no          | Whether or  |
|             |             |             |             | not to      |
|             |             |             |             | recursively |
|             |             |             |             | mount the   |
|             |             |             |             | source path |
+-------------+-------------+-------------+-------------+-------------+
| pool        | string      | -           | no          | The storage |
|             |             |             |             | pool the    |
|             |             |             |             | disk device |
|             |             |             |             | belongs to. |
|             |             |             |             | This is     |
|             |             |             |             | only        |
|             |             |             |             | applicable  |
|             |             |             |             | for storage |
|             |             |             |             | volumes     |
|             |             |             |             | managed by  |
|             |             |             |             | LXD         |
+-------------+-------------+-------------+-------------+-------------+
| propagation | string      | -           | no          | Controls    |
|             |             |             |             | how a       |
|             |             |             |             | bind-mount  |
|             |             |             |             | is shared   |
|             |             |             |             | between the |
|             |             |             |             | instance    |
|             |             |             |             | and the     |
|             |             |             |             | host. (Can  |
|             |             |             |             | be one of   |
|             |             |             |             | ``private`` |
|             |             |             |             | ,           |
|             |             |             |             | the         |
|             |             |             |             | default, or |
|             |             |             |             | ``shared``, |
|             |             |             |             | ``slave``,  |
|             |             |             |             | ``unbindabl |
|             |             |             |             | e``,        |
|             |             |             |             | ``rshared`` |
|             |             |             |             | ,           |
|             |             |             |             | ``rslave``, |
|             |             |             |             | ``runbindab |
|             |             |             |             | le``,       |
|             |             |             |             | ``rprivate` |
|             |             |             |             | `.          |
|             |             |             |             | Please see  |
|             |             |             |             | the Linux   |
|             |             |             |             | Kernel      |
|             |             |             |             | `shared     |
|             |             |             |             | subtree <ht |
|             |             |             |             | tps://www.k |
|             |             |             |             | ernel.org/d |
|             |             |             |             | oc/Document |
|             |             |             |             | ation/files |
|             |             |             |             | ystems/shar |
|             |             |             |             | edsubtree.t |
|             |             |             |             | xt>`__      |
|             |             |             |             | documentati |
|             |             |             |             | on          |
|             |             |             |             | for a full  |
|             |             |             |             | explanation |
|             |             |             |             | )           |
+-------------+-------------+-------------+-------------+-------------+
| shift       | boolean     | false       | no          | Setup a     |
|             |             |             |             | shifting    |
|             |             |             |             | overlay to  |
|             |             |             |             | translate   |
|             |             |             |             | the source  |
|             |             |             |             | uid/gid to  |
|             |             |             |             | match the   |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+
| raw.mount.o | string      | -           | no          | Filesystem  |
| ptions      |             |             |             | specific    |
|             |             |             |             | mount       |
|             |             |             |             | options     |
+-------------+-------------+-------------+-------------+-------------+
| ceph.user_n | string      | admin       | no          | If source   |
| ame         |             |             |             | is ceph or  |
|             |             |             |             | cephfs then |
|             |             |             |             | ceph        |
|             |             |             |             | user_name   |
|             |             |             |             | must be     |
|             |             |             |             | specified   |
|             |             |             |             | by user for |
|             |             |             |             | proper      |
|             |             |             |             | mount       |
+-------------+-------------+-------------+-------------+-------------+
| ceph.cluste | string      | ceph        | no          | If source   |
| r_name      |             |             |             | is ceph or  |
|             |             |             |             | cephfs then |
|             |             |             |             | ceph        |
|             |             |             |             | cluster_nam |
|             |             |             |             | e           |
|             |             |             |             | must be     |
|             |             |             |             | specified   |
|             |             |             |             | by user for |
|             |             |             |             | proper      |
|             |             |             |             | mount       |
+-------------+-------------+-------------+-------------+-------------+
| boot.priori | integer     | -           | no          | Boot        |
| ty          |             |             |             | priority    |
|             |             |             |             | for VMs     |
|             |             |             |             | (higher     |
|             |             |             |             | boots       |
|             |             |             |             | first)      |
+-------------+-------------+-------------+-------------+-------------+

Type: unix-char
~~~~~~~~~~~~~~~

Supported instance types: container

Unix character device entries simply make the requested character device
appear in the instance’s ``/dev`` and allow read/write operations to it.

The following properties exist:

+-------------+-------------+-------------+-------------+-------------+
| Key         | Type        | Default     | Required    | Description |
+=============+=============+=============+=============+=============+
| source      | string      | -           | no          | Path on the |
|             |             |             |             | host        |
+-------------+-------------+-------------+-------------+-------------+
| path        | string      | -           | no          | Path inside |
|             |             |             |             | the         |
|             |             |             |             | instance    |
|             |             |             |             | (one of     |
|             |             |             |             | “source”    |
|             |             |             |             | and “path”  |
|             |             |             |             | must be     |
|             |             |             |             | set)        |
+-------------+-------------+-------------+-------------+-------------+
| major       | int         | device on   | no          | Device      |
|             |             | host        |             | major       |
|             |             |             |             | number      |
+-------------+-------------+-------------+-------------+-------------+
| minor       | int         | device on   | no          | Device      |
|             |             | host        |             | minor       |
|             |             |             |             | number      |
+-------------+-------------+-------------+-------------+-------------+
| uid         | int         | 0           | no          | UID of the  |
|             |             |             |             | device      |
|             |             |             |             | owner in    |
|             |             |             |             | the         |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+
| gid         | int         | 0           | no          | GID of the  |
|             |             |             |             | device      |
|             |             |             |             | owner in    |
|             |             |             |             | the         |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+
| mode        | int         | 0660        | no          | Mode of the |
|             |             |             |             | device in   |
|             |             |             |             | the         |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+
| required    | boolean     | true        | no          | Whether or  |
|             |             |             |             | not this    |
|             |             |             |             | device is   |
|             |             |             |             | required to |
|             |             |             |             | start the   |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+

Type: unix-block
~~~~~~~~~~~~~~~~

Supported instance types: container

Unix block device entries simply make the requested block device appear
in the instance’s ``/dev`` and allow read/write operations to it.

The following properties exist:

+-------------+-------------+-------------+-------------+-------------+
| Key         | Type        | Default     | Required    | Description |
+=============+=============+=============+=============+=============+
| source      | string      | -           | no          | Path on the |
|             |             |             |             | host        |
+-------------+-------------+-------------+-------------+-------------+
| path        | string      | -           | no          | Path inside |
|             |             |             |             | the         |
|             |             |             |             | instance    |
|             |             |             |             | (one of     |
|             |             |             |             | “source”    |
|             |             |             |             | and “path”  |
|             |             |             |             | must be     |
|             |             |             |             | set)        |
+-------------+-------------+-------------+-------------+-------------+
| major       | int         | device on   | no          | Device      |
|             |             | host        |             | major       |
|             |             |             |             | number      |
+-------------+-------------+-------------+-------------+-------------+
| minor       | int         | device on   | no          | Device      |
|             |             | host        |             | minor       |
|             |             |             |             | number      |
+-------------+-------------+-------------+-------------+-------------+
| uid         | int         | 0           | no          | UID of the  |
|             |             |             |             | device      |
|             |             |             |             | owner in    |
|             |             |             |             | the         |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+
| gid         | int         | 0           | no          | GID of the  |
|             |             |             |             | device      |
|             |             |             |             | owner in    |
|             |             |             |             | the         |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+
| mode        | int         | 0660        | no          | Mode of the |
|             |             |             |             | device in   |
|             |             |             |             | the         |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+
| required    | boolean     | true        | no          | Whether or  |
|             |             |             |             | not this    |
|             |             |             |             | device is   |
|             |             |             |             | required to |
|             |             |             |             | start the   |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+

Type: usb
~~~~~~~~~

Supported instance types: container, VM

USB device entries simply make the requested USB device appear in the
instance.

The following properties exist:

+-------------+-------------+-------------+-------------+-------------+
| Key         | Type        | Default     | Required    | Description |
+=============+=============+=============+=============+=============+
| vendorid    | string      | -           | no          | The vendor  |
|             |             |             |             | id of the   |
|             |             |             |             | USB device  |
+-------------+-------------+-------------+-------------+-------------+
| productid   | string      | -           | no          | The product |
|             |             |             |             | id of the   |
|             |             |             |             | USB device  |
+-------------+-------------+-------------+-------------+-------------+
| uid         | int         | 0           | no          | UID of the  |
|             |             |             |             | device      |
|             |             |             |             | owner in    |
|             |             |             |             | the         |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+
| gid         | int         | 0           | no          | GID of the  |
|             |             |             |             | device      |
|             |             |             |             | owner in    |
|             |             |             |             | the         |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+
| mode        | int         | 0660        | no          | Mode of the |
|             |             |             |             | device in   |
|             |             |             |             | the         |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+
| required    | boolean     | false       | no          | Whether or  |
|             |             |             |             | not this    |
|             |             |             |             | device is   |
|             |             |             |             | required to |
|             |             |             |             | start the   |
|             |             |             |             | instance.   |
|             |             |             |             | (The        |
|             |             |             |             | default is  |
|             |             |             |             | false, and  |
|             |             |             |             | all devices |
|             |             |             |             | are         |
|             |             |             |             | hot-pluggab |
|             |             |             |             | le)         |
+-------------+-------------+-------------+-------------+-------------+

Type: gpu
~~~~~~~~~

GPU device entries simply make the requested gpu device appear in the
instance.

GPUs Available:
^^^^^^^^^^^^^^^

The following GPUs can be specified using the ``gputype`` property:

-  `physical <#gpu-physical>`__ Passes through an entire GPU. This is
   the default if ``gputype`` is unspecified.
-  `mdev <#gpu-mdev>`__ Creates and passes through a virtual GPU into
   the instance.
-  `mig <#gpu-mig>`__ Creates and passes through a MIG (Multi-Instance
   GPU) device into the instance.
-  `sriov <#gpu-sriov>`__ Passes a virtual function of an SR-IOV enabled
   GPU into the instance.

gpu: physical
^^^^^^^^^^^^^

Supported instance types: container, VM

Passes through an entire GPU.

The following properties exist:

========= ====== ======= ========
========================================================
Key       Type   Default Required Description
========= ====== ======= ========
========================================================
vendorid  string -       no       The vendor id of the GPU device
productid string -       no       The product id of the GPU device
id        string -       no       The card id of the GPU device
pci       string -       no       The pci address of the GPU device
uid       int    0       no       UID of the device owner in the instance (container only)
gid       int    0       no       GID of the device owner in the instance (container only)
mode      int    0660    no       Mode of the device in the instance (container only)
========= ====== ======= ========
========================================================

gpu: mdev
^^^^^^^^^

Supported instance types: VM

Creates and passes through a virtual GPU into the instance. A list of
available mdev profiles can be found by running
``lxc info --resources``.

The following properties exist:

========= ====== ======= ========
=============================================
Key       Type   Default Required Description
========= ====== ======= ========
=============================================
vendorid  string -       no       The vendor id of the GPU device
productid string -       no       The product id of the GPU device
id        string -       no       The card id of the GPU device
pci       string -       no       The pci address of the GPU device
mdev      string -       yes      The mdev profile to use (e.g. i915-GVTg_V5_4)
========= ====== ======= ========
=============================================

gpu: mig
^^^^^^^^

Supported instance types: container

Creates and passes through a MIG compute instance. This currently
requires NVIDIA MIG instances to be pre-created.

The following properties exist:

========= ====== ======= ======== =================================
Key       Type   Default Required Description
========= ====== ======= ======== =================================
vendorid  string -       no       The vendor id of the GPU device
productid string -       no       The product id of the GPU device
id        string -       no       The card id of the GPU device
pci       string -       no       The pci address of the GPU device
mig.ci    int    -       yes      Existing MIG compute instance ID
mig.gi    int    -       yes      Existing MIG GPU instance ID
========= ====== ======= ======== =================================

gpu: sriov
^^^^^^^^^^

Supported instance types: VM

Passes a virtual function of an SR-IOV enabled GPU into the instance.

The following properties exist:

========= ====== ======= ========
========================================
Key       Type   Default Required Description
========= ====== ======= ========
========================================
vendorid  string -       no       The vendor id of the parent GPU device
productid string -       no       The product id of the parent GPU device
id        string -       no       The card id of the parent GPU device
pci       string -       no       The pci address of the parent GPU device
========= ====== ======= ========
========================================

Type: proxy
~~~~~~~~~~~

Supported instance types: container (``nat`` and non-\ ``nat`` modes),
VM (``nat`` mode only)

Proxy devices allow forwarding network connections between host and
instance. This makes it possible to forward traffic hitting one of the
host’s addresses to an address inside the instance or to do the reverse
and have an address in the instance connect through the host.

The supported connection types are: \* ``tcp <-> tcp`` \*
``udp <-> udp`` \* ``unix <-> unix`` \* ``tcp <-> unix`` \*
``unix <-> tcp`` \* ``udp <-> tcp`` \* ``tcp <-> udp`` \*
``udp <-> unix`` \* ``unix <-> udp``

The proxy device also supports a ``nat`` mode where packets are
forwarded using NAT rather than being proxied through a separate
connection. This has benefit that the client address is maintained
without the need for the target destination to support the ``PROXY``
protocol (which is the only way to pass the client address through when
using the proxy device in non-nat mode).

When configuring a proxy device with ``nat=true``, you will need to
ensure that the target instance has a static IP configured in LXD on its
NIC device. E.g.

::

   lxc config device set <instance> <nic> ipv4.address=<ipv4.address> ipv6.address=<ipv6.address>

In order to define a static IPv6 address, the parent managed network
needs to have ``ipv6.dhcp.stateful`` enabled.

In NAT mode the supported connection types are:

-  ``tcp <-> tcp``
-  ``udp <-> udp``

When defining IPv6 addresses use square bracket notation, e.g.

::

   connect=tcp:[2001:db8::1]:80

You can specify that the connect address should be the IP of the
instance by setting the connect IP to the wildcard address (``0.0.0.0``
for IPv4 and ``[::]`` for IPv6).

The listen address can also use wildcard addresses when using non-NAT
mode. However when using ``nat`` mode you must specify an IP address on
the LXD host.

+-------------+-------------+-------------+-------------+-------------+
| Key         | Type        | Default     | Required    | Description |
+=============+=============+=============+=============+=============+
| listen      | string      | -           | yes         | The address |
|             |             |             |             | and port to |
|             |             |             |             | bind and    |
|             |             |             |             | listen      |
|             |             |             |             | (``<type>:< |
|             |             |             |             | addr>:<port |
|             |             |             |             | >[-<port>][ |
|             |             |             |             | ,<port>]``) |
+-------------+-------------+-------------+-------------+-------------+
| connect     | string      | -           | yes         | The address |
|             |             |             |             | and port to |
|             |             |             |             | connect to  |
|             |             |             |             | (``<type>:< |
|             |             |             |             | addr>:<port |
|             |             |             |             | >[-<port>][ |
|             |             |             |             | ,<port>]``) |
+-------------+-------------+-------------+-------------+-------------+
| bind        | string      | host        | no          | Which side  |
|             |             |             |             | to bind on  |
|             |             |             |             | (host/insta |
|             |             |             |             | nce)        |
+-------------+-------------+-------------+-------------+-------------+
| uid         | int         | 0           | no          | UID of the  |
|             |             |             |             | owner of    |
|             |             |             |             | the         |
|             |             |             |             | listening   |
|             |             |             |             | Unix socket |
+-------------+-------------+-------------+-------------+-------------+
| gid         | int         | 0           | no          | GID of the  |
|             |             |             |             | owner of    |
|             |             |             |             | the         |
|             |             |             |             | listening   |
|             |             |             |             | Unix socket |
+-------------+-------------+-------------+-------------+-------------+
| mode        | int         | 0644        | no          | Mode for    |
|             |             |             |             | the         |
|             |             |             |             | listening   |
|             |             |             |             | Unix socket |
+-------------+-------------+-------------+-------------+-------------+
| nat         | bool        | false       | no          | Whether to  |
|             |             |             |             | optimize    |
|             |             |             |             | proxying    |
|             |             |             |             | via NAT     |
|             |             |             |             | (requires   |
|             |             |             |             | instance    |
|             |             |             |             | NIC has     |
|             |             |             |             | static IP   |
|             |             |             |             | address)    |
+-------------+-------------+-------------+-------------+-------------+
| proxy_proto | bool        | false       | no          | Whether to  |
| col         |             |             |             | use the     |
|             |             |             |             | HAProxy     |
|             |             |             |             | PROXY       |
|             |             |             |             | protocol to |
|             |             |             |             | transmit    |
|             |             |             |             | sender      |
|             |             |             |             | information |
+-------------+-------------+-------------+-------------+-------------+
| security.ui | int         | 0           | no          | What UID to |
| d           |             |             |             | drop        |
|             |             |             |             | privilege   |
|             |             |             |             | to          |
+-------------+-------------+-------------+-------------+-------------+
| security.gi | int         | 0           | no          | What GID to |
| d           |             |             |             | drop        |
|             |             |             |             | privilege   |
|             |             |             |             | to          |
+-------------+-------------+-------------+-------------+-------------+

::

   lxc config device add <instance> <device-name> proxy listen=<type>:<addr>:<port>[-<port>][,<port>] connect=<type>:<addr>:<port> bind=<host/instance>

Type: unix-hotplug
~~~~~~~~~~~~~~~~~~

Supported instance types: container

Unix hotplug device entries make the requested unix device appear in the
instance’s ``/dev`` and allow read/write operations to it if the device
exists on the host system. Implementation depends on systemd-udev to be
run on the host.

The following properties exist:

+-------------+-------------+-------------+-------------+-------------+
| Key         | Type        | Default     | Required    | Description |
+=============+=============+=============+=============+=============+
| vendorid    | string      | -           | no          | The vendor  |
|             |             |             |             | id of the   |
|             |             |             |             | unix device |
+-------------+-------------+-------------+-------------+-------------+
| productid   | string      | -           | no          | The product |
|             |             |             |             | id of the   |
|             |             |             |             | unix device |
+-------------+-------------+-------------+-------------+-------------+
| uid         | int         | 0           | no          | UID of the  |
|             |             |             |             | device      |
|             |             |             |             | owner in    |
|             |             |             |             | the         |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+
| gid         | int         | 0           | no          | GID of the  |
|             |             |             |             | device      |
|             |             |             |             | owner in    |
|             |             |             |             | the         |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+
| mode        | int         | 0660        | no          | Mode of the |
|             |             |             |             | device in   |
|             |             |             |             | the         |
|             |             |             |             | instance    |
+-------------+-------------+-------------+-------------+-------------+
| required    | boolean     | false       | no          | Whether or  |
|             |             |             |             | not this    |
|             |             |             |             | device is   |
|             |             |             |             | required to |
|             |             |             |             | start the   |
|             |             |             |             | instance.   |
|             |             |             |             | (The        |
|             |             |             |             | default is  |
|             |             |             |             | false, and  |
|             |             |             |             | all devices |
|             |             |             |             | are         |
|             |             |             |             | hot-pluggab |
|             |             |             |             | le)         |
+-------------+-------------+-------------+-------------+-------------+

Type: tpm
~~~~~~~~~

Supported instance types: container, VM

TPM device entries enable access to a TPM emulator.

The following properties exist:

==== ====== ======= ========
===============================================
Key  Type   Default Required Description
==== ====== ======= ========
===============================================
path string -       yes      Path inside the instance (only for containers).
==== ====== ======= ========
===============================================

Type: pci
~~~~~~~~~

Supported instance types: VM

PCI device entries are used to pass raw PCI devices from the host into a
virtual machine.

The following properties exist:

======= ====== ======= ======== ==========================
Key     Type   Default Required Description
======= ====== ======= ======== ==========================
address string -       yes      PCI address of the device.
======= ====== ======= ======== ==========================

Units for storage and network limits
------------------------------------

Any value representing bytes or bits can make use of a number of useful
suffixes to make it easier to understand what a particular limit is.

Both decimal and binary (kibi) units are supported with the latter
mostly making sense for storage limits.

The full list of bit suffixes currently supported is:

-  bit (1)
-  kbit (1000)
-  Mbit (1000^2)
-  Gbit (1000^3)
-  Tbit (1000^4)
-  Pbit (1000^5)
-  Ebit (1000^6)
-  Kibit (1024)
-  Mibit (1024^2)
-  Gibit (1024^3)
-  Tibit (1024^4)
-  Pibit (1024^5)
-  Eibit (1024^6)

The full list of byte suffixes currently supported is:

-  B or bytes (1)
-  kB (1000)
-  MB (1000^2)
-  GB (1000^3)
-  TB (1000^4)
-  PB (1000^5)
-  EB (1000^6)
-  KiB (1024)
-  MiB (1024^2)
-  GiB (1024^3)
-  TiB (1024^4)
-  PiB (1024^5)
-  EiB (1024^6)

Instance types
--------------

LXD supports simple instance types. Those are represented as a string
which can be passed at instance creation time.

There are three allowed syntaxes:

-  ``<instance type>``
-  ``<cloud>:<instance type>``
-  ``c<CPU>-m<RAM in GB>``

For example, those 3 are equivalent:

-  t2.micro
-  aws:t2.micro
-  c1-m1

On the command line, this is passed like this:

.. code:: bash

   lxc launch ubuntu:20.04 my-instance -t t2.micro

The list of supported clouds and instance types can be found here:

https://github.com/dustinkirkland/instance-type

Hugepage limits via ``limits.hugepages.[size]``
-----------------------------------------------

LXD allows to limit the number of hugepages available to a container
through the ``limits.hugepage.[size]`` key. Limiting hugepages is done
through the hugetlb cgroup controller. This means the host system needs
to expose the hugetlb controller in the legacy or unified cgroup
hierarchy for these limits to apply. Note that architectures often
expose multiple hugepage sizes. In addition, architectures may expose
different hugepage sizes than other architectures.

Limiting hugepages is especially useful when LXD is configured to
intercept the mount syscall for the ``hugetlbfs`` filesystem in
unprivileged containers. When LXD intercepts a ``hugetlbfs`` mount
syscall, it will mount the ``hugetlbfs`` filesystem for a container with
correct ``uid`` and ``gid`` values as mount options. This makes it
possible to use hugepages from unprivileged containers. However, it is
recommended to limit the number of hugepages available to the container
through ``limits.hugepages.[size]`` to stop the container from being
able to exhaust the hugepages available to the host.

Resource limits via ``limits.kernel.[limit name]``
--------------------------------------------------

LXD exposes a generic namespaced key ``limits.kernel.*`` which can be
used to set resource limits for a given instance. It is generic in the
sense that LXD will not perform any validation on the resource that is
specified following the ``limits.kernel.*`` prefix. LXD cannot know
about all the possible resources that a given kernel supports. Instead,
LXD will simply pass down the corresponding resource key after the
``limits.kernel.*`` prefix and its value to the kernel. The kernel will
do the appropriate validation. This allows users to specify any
supported limit on their system. Some common limits are:

+-----------+---------------+-------------------------------------------+
| Key       | Resource      | Description                               |
+===========+===============+===========================================+
| limits.ke | RLIMIT_AS     | Maximum size of the process’s virtual     |
| rnel.as   |               | memory                                    |
+-----------+---------------+-------------------------------------------+
| limits.ke | RLIMIT_CORE   | Maximum size of the process’s coredump    |
| rnel.core |               | file                                      |
+-----------+---------------+-------------------------------------------+
| limits.ke | RLIMIT_CPU    | Limit in seconds on the amount of cpu     |
| rnel.cpu  |               | time the process can consume              |
+-----------+---------------+-------------------------------------------+
| limits.ke | RLIMIT_DATA   | Maximum size of the process’s data        |
| rnel.data |               | segment                                   |
+-----------+---------------+-------------------------------------------+
| limits.ke | RLIMIT_FSIZE  | Maximum size of files the process may     |
| rnel.fsiz |               | create                                    |
| e         |               |                                           |
+-----------+---------------+-------------------------------------------+
| limits.ke | RLIMIT_LOCKS  | Limit on the number of file locks that    |
| rnel.lock |               | this process may establish                |
| s         |               |                                           |
+-----------+---------------+-------------------------------------------+
| limits.ke | RLIMIT_MEMLOC | Limit on the number of bytes of memory    |
| rnel.meml | K             | that the process may lock in RAM          |
| ock       |               |                                           |
+-----------+---------------+-------------------------------------------+
| limits.ke | RLIMIT_NICE   | Maximum value to which the process’s nice |
| rnel.nice |               | value can be raised                       |
+-----------+---------------+-------------------------------------------+
| limits.ke | RLIMIT_NOFILE | Maximum number of open files for the      |
| rnel.nofi |               | process                                   |
| le        |               |                                           |
+-----------+---------------+-------------------------------------------+
| limits.ke | RLIMIT_NPROC  | Maximum number of processes that can be   |
| rnel.npro |               | created for the user of the calling       |
| c         |               | process                                   |
+-----------+---------------+-------------------------------------------+
| limits.ke | RLIMIT_RTPRIO | Maximum value on the real-time-priority   |
| rnel.rtpr |               | that maybe set for this process           |
| io        |               |                                           |
+-----------+---------------+-------------------------------------------+
| limits.ke | RLIMIT_SIGPEN | Maximum number of signals that maybe      |
| rnel.sigp | DING          | queued for the user of the calling        |
| ending    |               | process                                   |
+-----------+---------------+-------------------------------------------+

A full list of all available limits can be found in the manpages for the
``getrlimit(2)``/``setrlimit(2)`` system calls. To specify a limit
within the ``limits.kernel.*`` namespace use the resource name in
lowercase without the ``RLIMIT_`` prefix, e.g. \ ``RLIMIT_NOFILE``
should be specified as ``nofile``. A limit is specified as two colon
separated values which are either numeric or the word ``unlimited``
(e.g. ``limits.kernel.nofile=1000:2000``). A single value can be used as
a shortcut to set both soft and hard limit (e.g.
``limits.kernel.nofile=3000``) to the same value. A resource with no
explicitly configured limitation will be inherited from the process
starting up the instance. Note that this inheritance is not enforced by
LXD but by the kernel.

Snapshot scheduling
-------------------

LXD supports scheduled snapshots which can be created at most once every
minute. There are three configuration options. ``snapshots.schedule``
takes a shortened cron expression:
``<minute> <hour> <day-of-month> <month> <day-of-week>``. If this is
empty (default), no snapshots will be created.
``snapshots.schedule.stopped`` controls whether or not stopped instance
are to be automatically snapshotted. It defaults to ``false``.
``snapshots.pattern`` takes a pongo2 template string, and the pongo2
context contains the ``creation_date`` variable. Be aware that you
should format the date (e.g. use
``{{ creation_date|date:"2006-01-02_15-04-05" }}``) in your template
string to avoid forbidden characters in your snapshot name. Another way
to avoid name collisions is to use the placeholder ``%d``. If a snapshot
with the same name (excluding the placeholder) already exists, all
existing snapshot names will be taken into account to find the highest
number at the placeholders position. This number will be incremented by
one for the new name. The starting number if no snapshot exists will be
``0``.
