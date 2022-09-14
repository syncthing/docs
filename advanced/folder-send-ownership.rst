sendOwnership
=============

.. versionadded:: 1.22.0

When enabled, Syncthing will record ownership for local files when they are
scanned and send this information to peer devices. Peer devices configured
to :doc:`sync ownership <folder-sync-ownership>` will use this information.

On Windows, scanning ownership information has a fairly significant
performance impact on scan times.

.. seealso:: :doc:`folder-sync-ownership`
