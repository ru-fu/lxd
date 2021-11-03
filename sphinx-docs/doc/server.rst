Server configuration
====================

The server configuration is a simple set of key and values.

The key/value configuration is namespaced with the following namespaces
currently supported:

-  ``backups`` (backups configuration)
-  ``candid`` (External user authentication through Candid)
-  ``cluster`` (cluster configuration)
-  ``core`` (core daemon configuration)
-  ``images`` (image configuration)
-  ``maas`` (MAAS integration)
-  ``rbac`` (Role Based Access Control through external Candid +
   Canonical RBAC)

+------+--------+-----------+---------------+-------------------------+
| Key  | Type   | Scope     | Default       | Description             |
+======+========+===========+===============+=========================+
| back | string | global    | gzip          | Compression algorithm   |
| ups. |        |           |               | to use for new images   |
| comp |        |           |               | (bzip2, gzip, lzma, xz  |
| ress |        |           |               | or none)                |
| ion_ |        |           |               |                         |
| algo |        |           |               |                         |
| rith |        |           |               |                         |
| m    |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| cand | string | global    | -             | Public key of the       |
| id.a |        |           |               | candid server (required |
| pi.k |        |           |               | for HTTP-only servers)  |
| ey   |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| cand | string | global    | -             | URL of the the external |
| id.a |        |           |               | authentication endpoint |
| pi.u |        |           |               | using Candid            |
| rl   |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| cand | string | global    | -             | Comma-separated list of |
| id.d |        |           |               | allowed Candid domains  |
| omai |        |           |               | (empty string means all |
| ns   |        |           |               | domains are valid)      |
+------+--------+-----------+---------------+-------------------------+
| cand | intege | global    | 3600          | Candid macaroon expiry  |
| id.e | r      |           |               | in seconds              |
| xpir |        |           |               |                         |
| y    |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| clus | string | local     | -             | Address to use for      |
| ter. |        |           |               | clustering traffic      |
| http |        |           |               |                         |
| s_ad |        |           |               |                         |
| dres |        |           |               |                         |
| s    |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| clus | intege | global    | 3             | Minimal numbers of      |
| ter. | r      |           |               | cluster members with a  |
| imag |        |           |               | copy of a particular    |
| es_m |        |           |               | image (set 1 for no     |
| inim |        |           |               | replication, -1 for all |
| al_r |        |           |               | members)                |
| epli |        |           |               |                         |
| ca   |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| clus | intege | global    | 2             | Maximum number of       |
| ter. | r      |           |               | cluster members that    |
| max_ |        |           |               | will be assigned the    |
| stan |        |           |               | database stand-by role  |
| dby  |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| clus | intege | global    | 3             | Maximum number of       |
| ter. | r      |           |               | cluster members that    |
| max_ |        |           |               | will be assigned the    |
| vote |        |           |               | database voter role     |
| rs   |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| clus | intege | global    | 20            | Number of seconds after |
| ter. | r      |           |               | which an unresponsive   |
| offl |        |           |               | node is considered      |
| ine_ |        |           |               | offline                 |
| thre |        |           |               |                         |
| shol |        |           |               |                         |
| d    |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| core | string | local     | -             | Address to bind the BGP |
| .bgp |        |           |               | server to (BGP)         |
| _add |        |           |               |                         |
| ress |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| core | string | global    | -             | The BGP Autonomous      |
| .bgp |        |           |               | System Number to use    |
| _asn |        |           |               | for the local server    |
+------+--------+-----------+---------------+-------------------------+
| core | string | local     | -             | A unique identifier for |
| .bgp |        |           |               | this BGP server         |
| _rou |        |           |               | (formatted as an IPv4   |
| teri |        |           |               | address)                |
| d    |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| core | string | local     | -             | Address to bind the     |
| .deb |        |           |               | pprof debug server to   |
| ug_a |        |           |               | (HTTP)                  |
| ddre |        |           |               |                         |
| ss   |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| core | string | local     | -             | Address to bind for the |
| .htt |        |           |               | remote API (HTTPS)      |
| ps_a |        |           |               |                         |
| ddre |        |           |               |                         |
| ss   |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| core | boolea | global    | -             | Whether to set          |
| .htt | n      |           |               | Access-Control-Allow-Cr |
| ps_a |        |           |               | edentials               |
| llow |        |           |               | http header value to    |
| ed_c |        |           |               | “true”                  |
| rede |        |           |               |                         |
| ntia |        |           |               |                         |
| ls   |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| core | string | global    | -             | Access-Control-Allow-He |
| .htt |        |           |               | aders                   |
| ps_a |        |           |               | http header value       |
| llow |        |           |               |                         |
| ed_h |        |           |               |                         |
| eade |        |           |               |                         |
| rs   |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| core | string | global    | -             | Access-Control-Allow-Me |
| .htt |        |           |               | thods                   |
| ps_a |        |           |               | http header value       |
| llow |        |           |               |                         |
| ed_m |        |           |               |                         |
| etho |        |           |               |                         |
| ds   |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| core | string | global    | -             | Access-Control-Allow-Or |
| .htt |        |           |               | igin                    |
| ps_a |        |           |               | http header value       |
| llow |        |           |               |                         |
| ed_o |        |           |               |                         |
| rigi |        |           |               |                         |
| n    |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| core | string | global    | -             | Comma-separated list of |
| .htt |        |           |               | IP addresses of trusted |
| ps_t |        |           |               | servers to provide the  |
| rust |        |           |               | client’s address        |
| ed_p |        |           |               | through the proxy       |
| roxy |        |           |               | connection header       |
+------+--------+-----------+---------------+-------------------------+
| core | string | global    | -             | Address to bind the     |
| .met |        |           |               | metrics server to       |
| rics |        |           |               | (HTTPS)                 |
| _add |        |           |               |                         |
| ress |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| core | string | global    | -             | https proxy to use, if  |
| .pro |        |           |               | any (falls back to      |
| xy_h |        |           |               | HTTPS_PROXY environment |
| ttps |        |           |               | variable)               |
+------+--------+-----------+---------------+-------------------------+
| core | string | global    | -             | http proxy to use, if   |
| .pro |        |           |               | any (falls back to      |
| xy_h |        |           |               | HTTP_PROXY environment  |
| ttp  |        |           |               | variable)               |
+------+--------+-----------+---------------+-------------------------+
| core | string | global    | -             | hosts which don’t need  |
| .pro |        |           |               | the proxy for use       |
| xy_i |        |           |               | (similar format to      |
| gnor |        |           |               | NO_PROXY,               |
| e_ho |        |           |               | e.g. 1.2.3.4,1.2.3.5,   |
| sts  |        |           |               | falls back to NO_PROXY  |
|      |        |           |               | environment variable)   |
+------+--------+-----------+---------------+-------------------------+
| core | intege | global    | 5             | Number of minutes to    |
| .shu | r      |           |               | wait for running        |
| tdow |        |           |               | operations to complete  |
| n_ti |        |           |               | before LXD server shut  |
| meou |        |           |               | down                    |
| t    |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| core | boolea | global    | -             | Whether to              |
| .tru | n      |           |               | automatically trust     |
| st_c |        |           |               | clients signed by the   |
| a_ce |        |           |               | CA                      |
| rtif |        |           |               |                         |
| icat |        |           |               |                         |
| es   |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| core | string | global    | -             | Password to be provided |
| .tru |        |           |               | by clients to setup a   |
| st_p |        |           |               | trust                   |
| assw |        |           |               |                         |
| ord  |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| imag | boolea | global    | true          | Whether to              |
| es.a | n      |           |               | automatically update    |
| uto_ |        |           |               | any image that LXD      |
| upda |        |           |               | caches                  |
| te_c |        |           |               |                         |
| ache |        |           |               |                         |
| d    |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| imag | intege | global    | 6             | Interval in hours at    |
| es.a | r      |           |               | which to look for       |
| uto_ |        |           |               | update to cached images |
| upda |        |           |               | (0 disables it)         |
| te_i |        |           |               |                         |
| nter |        |           |               |                         |
| val  |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| imag | string | global    | gzip          | Compression algorithm   |
| es.c |        |           |               | to use for new images   |
| ompr |        |           |               | (bzip2, gzip, lzma, xz  |
| essi |        |           |               | or none)                |
| on_a |        |           |               |                         |
| lgor |        |           |               |                         |
| ithm |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| imag | string | -         | -             | Default architecture    |
| es.d |        |           |               | which should be used in |
| efau |        |           |               | mixed architecture      |
| lt_a |        |           |               | cluster                 |
| rchi |        |           |               |                         |
| tect |        |           |               |                         |
| ure  |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| imag | intege | global    | 10            | Number of days after    |
| es.r | r      |           |               | which an unused cached  |
| emot |        |           |               | remote image will be    |
| e_ca |        |           |               | flushed                 |
| che_ |        |           |               |                         |
| expi |        |           |               |                         |
| ry   |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| maas | string | global    | -             | API key to manage MAAS  |
| .api |        |           |               |                         |
| .key |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| maas | string | global    | -             | URL of the MAAS server  |
| .api |        |           |               |                         |
| .url |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| maas | string | local     | hostname      | Name of this LXD host   |
| .mac |        |           |               | in MAAS                 |
| hine |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| netw | string | global    | br-int        | OVS integration bridge  |
| ork. |        |           |               | to use for OVN networks |
| ovn. |        |           |               |                         |
| inte |        |           |               |                         |
| grat |        |           |               |                         |
| ion_ |        |           |               |                         |
| brid |        |           |               |                         |
| ge   |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| netw | string | global    | unix:/var/run | OVN northbound database |
| ork. |        |           | /ovn/ovnnb_db | connection string       |
| ovn. |        |           | .sock         |                         |
| nort |        |           |               |                         |
| hbou |        |           |               |                         |
| nd_c |        |           |               |                         |
| onne |        |           |               |                         |
| ctio |        |           |               |                         |
| n    |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| rbac | string | global    | -             | The Candid agent        |
| .age |        |           |               | private key as provided |
| nt.p |        |           |               | during RBAC             |
| riva |        |           |               | registration            |
| te_k |        |           |               |                         |
| ey   |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| rbac | string | global    | -             | The Candid agent public |
| .age |        |           |               | key as provided during  |
| nt.p |        |           |               | RBAC registration       |
| ubli |        |           |               |                         |
| c_ke |        |           |               |                         |
| y    |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| rbac | string | global    | -             | The Candid agent url as |
| .age |        |           |               | provided during RBAC    |
| nt.u |        |           |               | registration            |
| rl   |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| rbac | string | global    | -             | The Candid agent        |
| .age |        |           |               | username as provided    |
| nt.u |        |           |               | during RBAC             |
| sern |        |           |               | registration            |
| ame  |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| rbac | intege | global    | -             | RBAC macaroon expiry in |
| .api | r      |           |               | seconds                 |
| .exp |        |           |               |                         |
| iry  |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| rbac | string | global    | -             | Public key of the RBAC  |
| .api |        |           |               | server (required for    |
| .key |        |           |               | HTTP-only servers)      |
+------+--------+-----------+---------------+-------------------------+
| rbac | string | global    | -             | URL of the external     |
| .api |        |           |               | RBAC server             |
| .url |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| stor | string | local     | -             | Volume to use to store  |
| age. |        |           |               | the backup tarballs     |
| back |        |           |               | (syntax is POOL/VOLUME) |
| ups_ |        |           |               |                         |
| volu |        |           |               |                         |
| me   |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+
| stor | string | local     | -             | Volume to use to store  |
| age. |        |           |               | the image tarballs      |
| imag |        |           |               | (syntax is POOL/VOLUME) |
| es_v |        |           |               |                         |
| olum |        |           |               |                         |
| e    |        |           |               |                         |
+------+--------+-----------+---------------+-------------------------+

Those keys can be set using the lxc tool with:

.. code:: bash

   lxc config set <key> <value>

When operating as part of a cluster, the keys marked with a ``global``
scope will immediately be applied to all the cluster members. Those keys
with a ``local`` scope must be set on a per member basis using the
``--target`` option of the command line tool.

Exposing LXD to the network
---------------------------

By default, LXD can only be used by local users through a UNIX socket.

To expose LXD to the network, you’ll need to set ``core.https_address``.
All remote clients can then connect to LXD and access any image which
was marked for public use.

Trusted clients can be manually added to the trust store on the server
with ``lxc config trust add`` or the ``core.trust_password`` key can be
set allowing for clients to self-enroll into the trust store at
connection time by providing the confgiured password.

More details about authentication can be found `here <security.md>`__.

External authentication
-----------------------

LXD when accessed over the network can be configured to use external
authentication through `Candid <https://github.com/canonical/candid>`__.

Setting the ``candid.*`` configuration keys above to the values matching
your Candid deployment will allow users to authenticate through their
web browsers and then get trusted by LXD.

For those that have a Canonical RBAC server in front of their Candid
server, they can instead set the ``rbac.*`` configuration keys which are
a superset of the ``candid.*`` ones and allow for LXD to integrate with
the RBAC service.

When integrated with RBAC, individual users and groups can be granted
various level of access on a per-project basis. All of this is driven
externally through the RBAC service.

More details about authentication can be found `here <security.md>`__.
