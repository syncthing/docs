insecureAllowOldTLSVersions
===========================

.. versionadded:: 1.17.0

Syncthing by default allows only TLS 1.3 or higher for sync connections.
Setting this option makes Syncthing also allow TLS 1.2. Use only for
compatibility with very old Syncthing instances, or other implementations as
required.

This option does not affect the GUI/API connections; those always allow TLS
1.2 or higher.

The default value is ``false``.
