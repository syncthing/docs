.. _case-sensitive-fs:

caseSensitiveFS
===============

.. versionadded:: 1.9.0

``caseSensitiveFS`` is an advanced folder setting that affects file name
handling. 

With ``caseSensitiveFS`` set to  ``false`` (the default setting)
Syncthing's case sensitivity safety checks are enabled. 
Syncthing will then attempt to detect and prevent case-only file
name collisions that can occur on case insensitive systems such as Windows
and macOS, or other systems with case insensitive file systems.

When set to ``true`` the extra safety checks for case insensitive
filesystems are disabled. This will provide a small improvement in
performance when the underlying filesystem is positively known to be
case-sensitive already. This was the behavior of Syncthing 1.8.0 and earlier.

.. note:: This option should normally be set to ``false``. It is
	  **not** meant to change the basic principles of how Syncthing
	  :ref:`handles case-sensitivity <case-sensitivity>`.
