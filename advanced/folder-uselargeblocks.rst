.. versionadded:: 1.2.0

    As of Syncthing 1.2.0 large blocks are always enabled and this configuration
    option has been removed.

.. versionadded:: 1.1.0

    Syncthing version 1.1.0 and newer have large blocks enabled by default for
    new folders.

.. versionadded:: 0.14.48

    Large blocks can be enabled in Syncthing version 0.14.48 and newer.

useLargeBlocks
==============

``useLargeBlocks`` is an advanced folder setting that affects the handling
of blocks for files larger than 256 MiB. When enabled, the file will be
indexed and transferred using blocks larger than the standard 128 KiB. This
results in a smaller block list and thus lower overhead. The larger block
sizes are powers of two from 128 KiB up to 16 MiB. Syncthing will
attempt to select a block size to keep the number of blocks in a file
between 1000 and 2000, using the largest and smallest block size accordingly
at either extreme.

Compatibility
-------------

Syncthing version 0.14.46 and newer will accept and handle files with large
blocks, regardless of whether large blocks are enabled on that device.

Syncthing version 0.14.45 and older will initially appear to accept files
scanned with large blocks, but will later panic during some internal file
operations. Do not enable large blocks in clusters with devices still on
v0.14.45 or older.

When large blocks are *not* enabled, local changes to any given file are
indexed in standard (small) blocks - regardless of whether other devices are
using large blocks for the same folder.

When two devices do not agree on the desired block size for a file, the
entire file must be transferred instead of only the changed blocks. To avoid
frequent block size changes for files straddling a threshold boundary there
is a certain elasticity or hysteresis built into the system. The block size
of an existing file is only changed when the difference in block size
exceeds one level, i.e., from 256 KiB to 1 MiB, but not from 256 KiB to 512
KiB.

