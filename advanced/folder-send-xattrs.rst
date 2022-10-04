sendXattrs
==========

.. versionadded:: 1.22.0

When enabled, Syncthing will record extended attributes for local files when
they are scanned and send this information to peer devices. Peer devices
configured to :doc:`sync extended attributes <folder-sync-xattrs>` will use
this information.

Scanning extended attributes may have a performance impact on scan times.

.. seealso:: :doc:`folder-sync-xattrs`
