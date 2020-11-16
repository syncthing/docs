Environment variable USE_BADGER
===============================

.. versionadded:: 1.7.0

Syncthing traditionally uses the LevelDB database (`syndtr/goleveldb
<https://github.com/syndtr/goleveldb>`__). While this has served us well for
many years we are looking at potential alternatives. One of these is the
Badger database (`dgraph-io/goleveldb
<https://github.com/dgraph-io/goleveldb>`__). Syncthing currently supports
either with identical functionality. The default remains LevelDB but setting
the environment variable ``USE_BADGER=1`` causes Syncthing to instead use
Badger.

In order to support experimentation Syncthing performs automatic conversion
between the database formats without data loss. That is, if there is an
existing LevelDB database and ``USE_BADGER`` is set, the LevelDB database is
converted to a Badger database and archived. Similarly, if ``USE_BADGER`` is
not set and there is no LevelDB database but there is a Badger database, the
conversion is made in the opposite direction. If there is no existing
database at all a new one will be created of the correct type.

Migration Process
-----------------

To migrate to Badger, start Syncthing with the environment variable
``USE_BADGER`` set to ``1`` or any other non-empty value. The current
LevelDB database will be migrated to Badger format and archived.

To revert the migration, start Syncthing without the environment variable
set. The current Badger database will be migrated back into LevelDB format.
If you're migrating back due to some Badger related disaster (Syncthing
can't even migrate it back), the simply rename the
``index-v0.14.0.db.migrated.20200529135506`` (or similar) back to
``index-v0.14.0.db`` and start Syncthing without ``USE_BADGER`` set.

Directory Names
---------------

The following names are used for the various databases, always inside the
Syncthing database dir (set with ``-home`` or ``--data``):

``index-v0.14.0.db``
    A current, active LevelDB database.

``indexdb.badger``
    A current, active Badger database.

``index-v0.14.0.db.migrated.20200529135506``
    An archived, migrated LevelDB database after switching to Badger (date stamp will differ).

``indexdb.badger.migrated.20200529135506``
    An archived, migrated Badger database after switching to LevelDB (date stamp will differ).

Considerations
--------------

We have limited experience with how Badger behaves in various situations.
We're interested in both stability and performance data. As such, please
have usage reporting and/or (at the very least) :ref:`crashRep` enabled so
that we might hear of any serious issues.

Currently the :ref:`option-databaseTuning` option has no effect on Badger.
It's possible that we might implement the tuning options in Badger as well
after gathering more data on how it behaves.
