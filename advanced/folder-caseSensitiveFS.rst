caseSensitiveFS
===============

.. versionadded:: 1.9.0

``caseSensitiveFS`` is an advanced folder setting that affects file name
handling. When set to ``true`` the extra safety checks for case insensitive
filesystems are disabled, reverting the behavior to that of Syncthing 1.8.0
and earlier. With the safety checks enabled (``caseSensitiveFS = false``,
the default) Syncthing will attempt to detect and prevent case-only file
name collisions that can occur on case insensitive systems such as Windows
and macOS, or other systems with case insensitive file systems.

.. note:: This option should normally be set to ``false``.
