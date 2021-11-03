Events
======

Introduction
------------

Events are messages about actions that have occurred over LXD. Using the
API endpoint ``/1.0/events`` directly or via ``lxc monitor`` will
connect to a WebSocket through which logs and lifecycle messages will be
streamed.

Event types
-----------

LXD Currently supports three event types. - **Logging**: Shows all
logging messages regardless of the server logging level. -
**Operation**: Shows all ongoing operations from creation to completion
(including updates to their state and progress metadata). -
**Lifecycle**: Shows an audit trail for specific actions occurring over
LXD.

Event structure
---------------

Example:
~~~~~~~~

.. code:: yaml

   location: cluster_name
   metadata:
     action: network-updated
     requestor:
       protocol: unix
       username: root
     source: /1.0/networks/lxdbr0
   timestamp: "2021-03-14T00:00:00Z"
   type: lifecycle

-  ``location``: The cluster member name (if clustered).
-  ``timestamp``: Time that the event occurred in RFC3339 format.
-  ``type``: The type of event this is (one of ``logging``,
   ``operation``, or ``lifecycle``).
-  ``metadata``: Information about the specific event type.

Logging event structure
~~~~~~~~~~~~~~~~~~~~~~~

-  ``message``: The log message.
-  ``level``: The log-level of the log.
-  ``context``: Additional information included in the event.

Operation event structure
~~~~~~~~~~~~~~~~~~~~~~~~~

-  ``id``: The UUID of the operation.
-  ``class``: The type of operation (task, token, or websocket).
-  ``description``: A description of the operation.
-  ``created_at``: The operation’s creation date.
-  ``updated_at``: The operation’s date of last change.
-  ``status``: The current state of the operation.
-  ``status_code``: The operation status code.
-  ``resources``: Resources affected by this operation.
-  ``metadata``: Operation specific metadata.
-  ``may_cancel``: Whether the operation may be cancelled.
-  ``err``: Error message of the operation.
-  ``location``: The cluster member name (if clustered).

Lifecycle event structure
~~~~~~~~~~~~~~~~~~~~~~~~~

-  ``action``: The lifecycle action that occurred.
-  ``requestor``: Information about who is making the request (if
   applicable).
-  ``source``: Path to what is being acted upon.
-  ``context``: Additional information included in the event.

Supported lifecycle events
--------------------------

+------------+-----------------------+---------------------------------+
| Name       | Description           | Additional Information          |
+============+=======================+=================================+
| ``certific | A new certificate has |                                 |
| ate-create | been added to the     |                                 |
| d``        | server trust store.   |                                 |
+------------+-----------------------+---------------------------------+
| ``certific | The certificate has   |                                 |
| ate-delete | been deleted from the |                                 |
| d``        | trust store.          |                                 |
+------------+-----------------------+---------------------------------+
| ``certific | The certificate’s     |                                 |
| ate-update | configuration has     |                                 |
| d``        | been updated.         |                                 |
+------------+-----------------------+---------------------------------+
| ``cluster- | The certificate for   |                                 |
| certificat | the whole cluster has |                                 |
| e-updated` | changed.              |                                 |
| `          |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``cluster- | Clustering has been   |                                 |
| disabled`` | disabled for this     |                                 |
|            | machine.              |                                 |
+------------+-----------------------+---------------------------------+
| ``cluster- | Clustering has been   |                                 |
| enabled``  | enabled for this      |                                 |
|            | machine.              |                                 |
+------------+-----------------------+---------------------------------+
| ``cluster- | A new machine has     |                                 |
| member-add | joined the cluster.   |                                 |
| ed``       |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``cluster- | The cluster member    |                                 |
| member-rem | has been removed from |                                 |
| oved``     | the cluster.          |                                 |
+------------+-----------------------+---------------------------------+
| ``cluster- | The cluster member    | ``old_name``: the previous      |
| member-ren | has been renamed.     | name.                           |
| amed``     |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``cluster- | The cluster member’s  |                                 |
| member-upd | configuration been    |                                 |
| ated``     | edited.               |                                 |
+------------+-----------------------+---------------------------------+
| ``cluster- | A join token for      |                                 |
| token-crea | adding a cluster      |                                 |
| ted``      | member has been       |                                 |
|            | created.              |                                 |
+------------+-----------------------+---------------------------------+
| ``config-u | The server            |                                 |
| pdated``   | configuration has     |                                 |
|            | changed.              |                                 |
+------------+-----------------------+---------------------------------+
| ``image-al | An alias has been     | ``target``: the original        |
| ias-create | created for an        | instance.                       |
| d``        | existing image.       |                                 |
+------------+-----------------------+---------------------------------+
| ``image-al | An alias has been     | ``target``: the original        |
| ias-delete | deleted for an        | instance.                       |
| d``        | existing image.       |                                 |
+------------+-----------------------+---------------------------------+
| ``image-al | The alias for an      | ``old_name``: the previous      |
| ias-rename | existing image has    | name.                           |
| d``        | been renamed.         |                                 |
+------------+-----------------------+---------------------------------+
| ``image-al | The configuration for | ``target``: the original        |
| ias-update | an image alias has    | instance.                       |
| d``        | changed.              |                                 |
+------------+-----------------------+---------------------------------+
| ``image-cr | A new image has been  | ``type``: container or vm.      |
| eated``    | added to the image    |                                 |
|            | store.                |                                 |
+------------+-----------------------+---------------------------------+
| ``image-de | The image has been    |                                 |
| leted``    | deleted from the      |                                 |
|            | image store.          |                                 |
+------------+-----------------------+---------------------------------+
| ``image-re | The local image copy  |                                 |
| freshed``  | has updated to the    |                                 |
|            | current source image  |                                 |
|            | version.              |                                 |
+------------+-----------------------+---------------------------------+
| ``image-re | The raw image file    | ``target``: destination server. |
| trieved``  | has been downloaded   |                                 |
|            | from the server.      |                                 |
+------------+-----------------------+---------------------------------+
| ``image-se | A one-time key to     |                                 |
| cret-creat | fetch this image has  |                                 |
| ed``       | been created.         |                                 |
+------------+-----------------------+---------------------------------+
| ``image-up | The image’s           |                                 |
| dated``    | configuration has     |                                 |
|            | changed.              |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | A backup of the       |                                 |
| -backup-cr | instance has been     |                                 |
| eated``    | created.              |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The instance backup   |                                 |
| -backup-de | has been deleted.     |                                 |
| leted``    |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The instance backup   | ``old_name``: the previous      |
| -backup-re | has been renamed.     | name.                           |
| named``    |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The raw instance      |                                 |
| -backup-re | backup file has been  |                                 |
| trieved``  | downloaded.           |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | Connected to the      | ``type``: console or vga.       |
| -console`` | console of the        |                                 |
|            | instance.             |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The console buffer    |                                 |
| -console-r | has been reset.       |                                 |
| eset``     |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The console log has   |                                 |
| -console-r | been downloaded.      |                                 |
| etrieved`` |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | A new instance has    |                                 |
| -created`` | been created.         |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The instance has been |                                 |
| -deleted`` | deleted.              |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | A command has been    | ``command``: the command to be  |
| -exec``    | executed on the       | executed.                       |
|            | instance.             |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | A file on the         | ``file``: path to the file.     |
| -file-dele | instance has been     |                                 |
| ted``      | deleted.              |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The file has been     | ``file-source``: local file     |
| -file-push | pushed to the         | path. ``file-destination``:     |
| ed``       | instance.             | destination file path.          |
|            |                       | ``info``: file information.     |
+------------+-----------------------+---------------------------------+
| ``instance | The file has been     | ``file-source``: instance file  |
| -file-retr | downloaded from the   | path. ``file-destination``:     |
| ieved``    | instance.             | destination file path.          |
+------------+-----------------------+---------------------------------+
| ``instance | The instance’s        |                                 |
| -log-delet | specified log file    |                                 |
| ed``       | has been deleted.     |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The instance’s        |                                 |
| -log-retri | specified log file    |                                 |
| eved``     | has been downloaded.  |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The instance’s image  |                                 |
| -metadata- | metadata has been     |                                 |
| retrieved` | downloaded.           |                                 |
| `          |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The instance’s image  |                                 |
| -metadata- | metadata has changed. |                                 |
| updated``  |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | A new image template  | ``path``: relative file path.   |
| -metadata- | file for the instance |                                 |
| template-c | has been created.     |                                 |
| reated``   |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The image template    | ``path``: relative file path.   |
| -metadata- | file for the instance |                                 |
| template-d | has been deleted.     |                                 |
| eleted``   |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The image template    | ``path``: relative file path.   |
| -metadata- | file for the instance |                                 |
| template-r | has been downloaded.  |                                 |
| etrieved`` |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The instance has been |                                 |
| -paused``  | put in a paused       |                                 |
|            | state.                |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The instance has been | ``old_name``: the previous      |
| -renamed`` | renamed.              | name.                           |
+------------+-----------------------+---------------------------------+
| ``instance | The instance has      |                                 |
| -restarted | restarted.            |                                 |
| ``         |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The instance has been | ``snapshot``: name of the       |
| -restored` | restored from a       | snapshot being restored.        |
| `          | snapshot.             |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The instance has      |                                 |
| -resumed`` | resumed after being   |                                 |
|            | paused.               |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The instance has shut |                                 |
| -shutdown` | down.                 |                                 |
| `          |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The instance has      |                                 |
| -started`` | started.              |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The instance has      |                                 |
| -stopped`` | stopped.              |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The instance’s        |                                 |
| -updated`` | configuration has     |                                 |
|            | changed.              |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | A snapshot of the     |                                 |
| -snapshot- | instance has been     |                                 |
| created``  | created.              |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The instance snapshot |                                 |
| -snapshot- | has been deleted.     |                                 |
| deleted``  |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The instance snapshot | ``old_name``: the previous      |
| -snapshot- | has been renamed.     | name.                           |
| renamed``  |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``instance | The instance          |                                 |
| -snapshot- | snapshot’s            |                                 |
| updated``  | configuration has     |                                 |
|            | changed.              |                                 |
+------------+-----------------------+---------------------------------+
| ``network- | A new network acl has |                                 |
| acl-create | been created.         |                                 |
| d``        |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``network- | The network acl has   |                                 |
| acl-delete | been deleted.         |                                 |
| d``        |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``network- | The network acl has   | ``old_name``: the previous      |
| acl-rename | been renamed.         | name.                           |
| d``        |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``network- | The network acl       |                                 |
| acl-update | configuration has     |                                 |
| d``        | changed.              |                                 |
+------------+-----------------------+---------------------------------+
| ``network- | A network device has  |                                 |
| created``  | been created.         |                                 |
+------------+-----------------------+---------------------------------+
| ``network- | The network device    |                                 |
| deleted``  | has been deleted.     |                                 |
+------------+-----------------------+---------------------------------+
| ``network- | The network device    | ``old_name``: the previous      |
| renamed``  | has been renamed.     | name.                           |
+------------+-----------------------+---------------------------------+
| ``network- | The network device’s  |                                 |
| updated``  | configuration has     |                                 |
|            | changed.              |                                 |
+------------+-----------------------+---------------------------------+
| ``operatio | The operation has     |                                 |
| n-cancelle | been cancelled.       |                                 |
| d``        |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``profile- | A new profile has     |                                 |
| created``  | been created.         |                                 |
+------------+-----------------------+---------------------------------+
| ``profile- | The profile has been  |                                 |
| deleted``  | deleted.              |                                 |
+------------+-----------------------+---------------------------------+
| ``profile- | The profile has been  | ``old_name``: the previous      |
| renamed``  | renamed .             | name.                           |
+------------+-----------------------+---------------------------------+
| ``profile- | The profile’s         |                                 |
| updated``  | configuration has     |                                 |
|            | changed.              |                                 |
+------------+-----------------------+---------------------------------+
| ``project- | A new project has     |                                 |
| created``  | been created.         |                                 |
+------------+-----------------------+---------------------------------+
| ``project- | The project has been  |                                 |
| deleted``  | deleted.              |                                 |
+------------+-----------------------+---------------------------------+
| ``project- | The project has been  | ``old_name``: the previous      |
| renamed``  | renamed.              | name.                           |
+------------+-----------------------+---------------------------------+
| ``project- | The project’s         |                                 |
| updated``  | configuration has     |                                 |
|            | changed.              |                                 |
+------------+-----------------------+---------------------------------+
| ``storage- | A new storage pool    | ``target``: cluster member      |
| pool-creat | has been created.     | name.                           |
| ed``       |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``storage- | The storage pool has  |                                 |
| pool-delet | been deleted.         |                                 |
| ed``       |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``storage- | The storage pool’s    | ``target``: cluster member      |
| pool-updat | configuration has     | name.                           |
| ed``       | changed.              |                                 |
+------------+-----------------------+---------------------------------+
| ``storage- | A new backup for the  | ``type``: container,            |
| volume-bac | storage volume has    | virtual-machine, image, or      |
| kup-create | been created.         | custom.                         |
| d``        |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``storage- | The storage volume’s  |                                 |
| volume-bac | backup has been       |                                 |
| kup-delete | deleted.              |                                 |
| d``        |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``storage- | The storage volume’s  | ``old_name``: the previous      |
| volume-bac | backup has been       | name.                           |
| kup-rename | renamed.              |                                 |
| d``        |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``storage- | The storage volume’s  |                                 |
| volume-bac | backup has been       |                                 |
| kup-retrie | downloaded.           |                                 |
| ved``      |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``storage- | A new storage volume  | ``type``: container,            |
| volume-cre | has been created.     | virtual-machine, image, or      |
| ated``     |                       | custom.                         |
+------------+-----------------------+---------------------------------+
| ``storage- | The storage volume    |                                 |
| volume-del | has been deleted.     |                                 |
| eted``     |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``storage- | The storage volume    | ``old_name``: the previous      |
| volume-ren | has been renamed.     | name.                           |
| amed``     |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``storage- | The storage volume    | ``snapshot``: name of the       |
| volume-res | has been restored     | snapshot being restored.        |
| tored``    | from a snapshot.      |                                 |
+------------+-----------------------+---------------------------------+
| ``storage- | The storage volume’s  |                                 |
| volume-upd | configuration has     |                                 |
| ated``     | changed.              |                                 |
+------------+-----------------------+---------------------------------+
| ``storage- | A new storage volume  | ``type``: container,            |
| volume-sna | snapshot has been     | virtua-machine, image, or       |
| pshot-crea | created.              | custom.                         |
| ted``      |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``storage- | The storage volume’s  |                                 |
| volume-sna | snapshot has been     |                                 |
| pshot-dele | deleted.              |                                 |
| ted``      |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``storage- | The storage volume’s  | ``old_name``: the previous      |
| volume-sna | snapshot has been     | name.                           |
| pshot-rena | renamed.              |                                 |
| med``      |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``storage- | The configuration for |                                 |
| volume-sna | the storage volume’s  |                                 |
| pshot-upda | snapshot has changed. |                                 |
| ted``      |                       |                                 |
+------------+-----------------------+---------------------------------+
| ``warning- | The warning’s status  |                                 |
| acknowledg | has been set to       |                                 |
| ed``       | “acknowledged”.       |                                 |
+------------+-----------------------+---------------------------------+
| ``warning- | The warning has been  |                                 |
| deleted``  | deleted.              |                                 |
+------------+-----------------------+---------------------------------+
| ``warning- | The warning’s status  |                                 |
| reset``    | has been set to       |                                 |
|            | “new”.                |                                 |
+------------+-----------------------+---------------------------------+
