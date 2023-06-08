.. _folder-types:

Folder Types
============

Send & Receive Folder
---------------------

This is the standard folder type. Under this setting, a folder will both send
changes to and receive changes from remote devices.

.. _folder-sendonly:

Send Only Folder
----------------

A folder can be set in "send only mode" among the folder settings.

.. image:: foldersendonly.png

The intention is for this to be used on devices where a "reference copy" of
files are kept - where the files are not expected to be changed on other
devices or where such changes would be undesirable.

In send only mode, all changes from other devices in the cluster are
ignored. Changes are still *received* so the folder may become "out of
sync", but no changes will be applied.

When a send only folder becomes out of sync, a red "Override Changes"
button is shown at the bottom of the folder details.

.. image:: override.png

Clicking this button will enforce this host's current state on the
rest of the cluster. Any changes made to files will be overwritten by
the version on this host, any files that don't exist on this host will
be deleted, and so on.

.. _folder-recvonly:

Receive Only Folder
-------------------

.. versionadded:: 0.14.50

The receive only folder is the logical opposite of the send only folder. In
this mode, all changes from the cluster are applied and redistributed to
other devices, as they are in the default send-receive mode. Local changes
are however not distributed to other devices. This mode is useful for
replication mirrors, backup destinations, or other cases where no local
modifications are expected or allowed.

Much like a send-receive folder, any local modifications are preserved and
do not cause the folder to become "out of sync". The device will however
look out of sync on *other* devices, as it no longer has the
latest/expected version of the modified file.

When local changes have been detected Syncthing will show a red "Revert Local
Changes" button on the folder. Activating this will cause the local
modifications to be undone - added files will be deleted, modified or
deleted files will be re-synced from the cluster.

In normal operation, a locally modified file that is subsequently modified
by the cluster will cause a sync conflict. The conflict will be resolved
with the cluster version as the winner. Being a receive-only folder, the
sync conflict copy will not be sent to the cluster - and will be deleted if
"Revert Local Changes" is used.
