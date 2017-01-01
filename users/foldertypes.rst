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

A folder can be set in "send-only mode" among the folder settings.

.. image:: foldersendonly.png

The intention is for this to be used on devices where a "master copy" of
files are kept - where the files are not expected to be changed on other
devices or where such changes would be undesirable.

In send-only mode, all changes from other devices in the cluster are
ignored. Changes are still *received* so the folder may become "out of
sync", but no changes will be applied.

When a send-only folder becomes out of sync, a red "Override Changes"
button is shown at the bottom of the folder details.

.. image:: override.png

Clicking this button will enforce this host's current state on the
rest of the cluster. Any changes made to files will be overwritten by
the version on this host, any files that don't exist on this host will
be deleted, and so on.

Receive Only Folder
-------------------

A folder can be set in "receive-only mode" among the folder settings.

.. image:: folderreceiveonly.png

In this mode, a device will never send any local updates to the cluster.
It will only ever receive changes.

If an existing file gets modified locally, it will automatically be replaced 
with the cluster version on the next scan.

The "Delete local changes" will remove any files or directories that 
have been added locally. **Use with care!**

Enable :ref:`versioning` or set maxConflicts, to avoid losing any local 
changes.
