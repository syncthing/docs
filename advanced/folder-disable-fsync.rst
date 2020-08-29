disableFsync
============

``disableFsync`` is an advanced folder setting that affects file
modifications. Normally, when a file has been modified Syncthing calls
``fsync()`` on that file and the containing directory. This forces file data
that is cached in RAM to be flushed to disk. This ensures that data is
safely stored on disk and thus prevents data loss in the case of a power
failure soon after file modification.

There is however a performance cost to doing this, especially on rotating
disks or network filesystems, especially syncing many small files. Disabling
``fsync()`` improves performance at the price of risking data loss in a
power failure situation.

.. note:: This option should normally be set to ``false``.
