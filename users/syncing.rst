.. _syncing:

Understanding Synchronization
=============================

This article describes the mechanisms Syncthing uses to bring files in sync
on a high level.

Blocks
------

Files are divided into *blocks*. The blocks making up a file are all the
same size (except the last one in the file which may be smaller). The block
size is dependent on the file size and varies from 128 KiB up to 16 MiB.
Each file is sliced into a number of these blocks, and the SHA256 hash of
each block is computed. This results in a *block list* containing the
offset, size and hash of all blocks in the file.

To update a file, Syncthing compares the block list of the current version
of the file to the block list of the desired version of the file. It then
tries to find a source for each block that differs. This might be locally,
if another file already has a block with the same hash, or it may be from
another device in the cluster. In the first case the block is simply copied
on disk, in the second case it is requested over the network from the other
device.

When a block is copied or received from another device, its SHA256 hash is
computed and compared with the expected value. If it matches the block is
written to a temporary copy of the file, otherwise it is discarded and
Syncthing tries to find another source for the block.

.. _scanning:

Scanning
--------

There are two methods how Syncthing detects changes: By regular full scans and
by notifications received from the filesystem ("watcher"). By default the
watcher is enabled and full scans are done once per hour. This behaviour can be
changed by folder. Increasing the full scan interval uses less resources and is
useful for example on large folders that change infrequently. To make sure that
not all folders are rescanned at the same time, the actual scan interval is a
random time between 3/4 and 5/4 of the given scan interval. Even with watcher
enabled it is advised to keep regular full scans enabled, as it is possible that
some changes aren't picked up by it.

During a rescan (regardless whether full or from watcher) the existing files are
checked for changes to their modification time, size or permission bits. The
file is "rehashed" if a change is detected based on those attributes, that is a
new block list is calculated for the file. It is not possible to know which
parts of a file have changed without reading the file and computing new SHA256
hashes for each block.

The watcher does not immediately schedule a scan when a change is detected. It
accumulates changes for 10s (adjustable by :stconf:opt:`fsWatcherDelayS`) and deleted files
are further delayed for 1min. Therefore it is expected that you experience a
slight delay between making the change and it appearing on another device.

Changes that were detected and hashed are transmitted to the other devices
after each rescan.

Syncing
-------

Syncthing keeps track of several versions of each file - the version that it
currently has on disk, called the *local* version, the versions announced by
all other connected devices, and the "best" (usually the most recent)
version of the file. This version is called the *global* version and is the
one that each device strives to be up to date with.

This information is kept in the *index database*, which is stored in the
configuration directory and called ``index-vx.y.z.db`` (for some version
x.y.z which may not be exactly the version of Syncthing you're running).

When new index data is received from other devices Syncthing recalculates
which version for each file should be the global version, and compares this
to the current local version. When the two differ, Syncthing needs to
synchronize the file. The block lists are compared to build a list of needed
blocks, which are then requested from the network or copied locally, as
described above.

.. _conflict-handling:

Conflicting Changes
-------------------

Syncthing does recognize conflicts.  When a file has been modified on two
devices simultaneously and the content actually differs, one of the files will
be renamed to ``<filename>.sync-conflict-<date>-<time>-<modifiedBy>.<ext>``.
The file with the older modification time will be marked as the conflicting file
and thus be renamed.  If the modification times are equal, the file originating
from the device which has the larger value of the first 63 bits for its device
ID will be marked as the conflicting file.  If the conflict is between a
modification and a deletion of the file, the modified file always wins and is
resurrected without renaming on the device where it was deleted.

Beware that the ``<filename>.sync-conflict-<date>-<time>-<modifiedBy>.<ext>``
files are treated as normal files after they are created, so they are propagated
between devices.  We do this because the conflict is detected and resolved on
one device, creating the ``sync-conflict`` file, but it's just as much of a
conflict everywhere else and we don't know which of the conflicting files is the
"best" from the user point of view.

.. _case-sensitivity:

Case Sensitivity in File Names
------------------------------

In principle, Syncthing works with *case-sensitive* paths, meaning
that ``file.txt`` and ``FILE.txt`` denote two independent things.
Consequently, it never considers both as if they were somehow related
for synchronizing their contents.  Some operating systems
(e.g. Windows, Mac and Android to an extent) assume the opposite,
treating them as the same file (or directory) and can never have both
names simultaneously.

Thus, if a remote device shares a file which would clash with an
existing local file, it cannot be synchronized to such a system and
will be reported as a case conflict by Syncthing.  Similarly, if two
remote devices share differing case variants to your local device,
only one of them will be pulled, and the other one marked with an
appropriate error message.  Which one "wins" is currently not
predictable.

In order to resolve such a case conflict situation, you need to decide
on a consistent file name and **manually** enforce that across all
involved, case-insensitive devices.

This cautious behavior tries to save you from possible data loss
caused by different files overwriting each other's contents.  That
could have happened before version 1.9.0, where the same file would
erroneously be accessed under two case-differing file names.

All this does not concern the folder root path, but only relative
paths within each shared folder.

.. _temporary-files:

Temporary Files
---------------

Syncthing never writes directly to a destination file. Instead all changes
are made to a temporary copy which is then moved in place over the old
version. If an error occurs during the copying or syncing, such as a
necessary block not being available, the temporary file is kept around for
up to a day. This is to avoid needlessly requesting data over the network.

The temporary files are named ``.syncthing.original-filename.ext.tmp`` or,
on Windows, ``~syncthing~original-filename.ext.tmp`` where
``original-filename.ext`` is the destination filename. The temporary file is
normally hidden. If the temporary file name would be too long due to the addition of the prefix and extra extension, the temporary files are named ``.syncthing.<hash>.tmp`` or, on Windows, ``~syncthing~<hash>.tmp`` where ``<hash>`` is a SHA-256 hash of the original filename.

.. note::

    Note that the two prefixes ``.syncthing.`` and ``~syncthing~`` are
    considered Syncthing namespace, meaning that any files whose names
    start with them will automatically be ignored and excluded from
    synchronisation by Syncthing. Please avoid using these prefixes in
    your filenames.
