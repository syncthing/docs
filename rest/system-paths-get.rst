GET /rest/system/paths
======================

.. versionadded:: 1.21.0

Returns the path locations used internally for storing configuration, database,
and others.

.. code-block:: json

    {
      "auditLog": "/home/user/.local/share/syncthing/audit-${timestamp}.log",
      "baseDir-config": "/home/user/.config/syncthing",
      "baseDir-data": "/home/user/.local/share/syncthing",
      "baseDir-userHome": "/home/user",
      "certFile": "/home/user/.config/syncthing/cert.pem",
      "config": "/home/user/.config/syncthing/config.xml",
      "csrfTokens": "/home/user/.config/syncthing/csrftokens.txt",
      "database": "/home/user/.local/share/syncthing/index-v0.14.0.db",
      "defFolder": "/home/user/Sync",
      "guiAssets": "/home/user/src/syncthing/gui",
      "httpsCertFile": "/home/user/.config/syncthing/https-cert.pem",
      "httpsKeyFile": "/home/user/.config/syncthing/https-key.pem",
      "keyFile": "/home/user/.config/syncthing/key.pem",
      "logFile": "-",
      "panicLog": "/home/user/.local/share/syncthing/panic-${timestamp}.log"
    }
