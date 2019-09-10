databaseTuning
==============

.. versionadded:: 1.3.0

The ``databaseTuning`` option controls how Syncthing uses the backend
key-value database that stores the index data and other persistent data
Syncthing needs. In most cases this database is fairly small (hundred
megabytes or less) and the default tuning is optimized for this. However in
cases with large folders or many devices the database may grow, and updates
may be more frequent. In these cases it's better to use larger buffers,
allow more memory to be used for cache, allow a larger amount of overhead on
disk in order to improve compaction performance, and so on.

The ``databaseTuning`` option can have one of three values:

- ``small``: This is the old set of tuning parameters, recommended for small
  databases.
- ``large``: This is a new set of tuning parameters, recommended for large
  databases.

- ``auto``: Syncthing will automatically use either the large or small
  tuning parameters depending on the database size on disk. A database that is
  200 MiB or more in size is considered "large" for the purposes of this
  option.

The default value is ``auto``.

