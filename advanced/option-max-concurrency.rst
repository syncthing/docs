maxFolderConcurrency
====================

.. versionadded:: 1.4.0

The ``maxFolderConcurrency`` option controls how many folders may
concurrently be in I/O-intensive operations such as syncing or scanning. The
default value is the same as the number of logical CPU cores in the system.
Folders waiting for their turn to scan or sync will show up as "Waiting to
Scan" or "Waiting to Sync" until the total number of ongoing such operations
is low enough to let them proceed.

Valid Values
------------

**Zero (0)**:
    The default, means the number of logical CPUs in the system (i.e., 2, 4,
    8, etc.), autodetected.

**Negative (< 0)**:
    No limit on the number of concurrent operations. This was the default in
    versions < 1.4.0.

**A positive integer (> 0)**:
    Use this specific limit.
