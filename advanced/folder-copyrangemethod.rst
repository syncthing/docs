.. _folder-copyRangeMethod:

copyRangeMethod
===============

.. versionadded:: 1.8.0

Provides a choice of method for copying data between files. This can be used
to optimise copies on network filesystems, improve speed of large copies or
clone the data using copy-on-write functionality if the underlying
filesystem supports it.

The following values are accepted:

    ``standard`` (default)
        Reads the data from source file into application memory, writes the
        data from application memory into the destination file. (This is the
        method used by Syncthing prior to this option being introduced.)

        *Available on: All platforms*

    ``copy_file_range``
        Uses the Linux ``copy_file_range`` syscall which, if the underlying
        filesystem supports it, uses copy-on-write semantics to clone the
        data. Introduced in Linux 4.5 and tested on XFS and BTRFS. Some
        network filesystems might use this to perform server-side copies.

        | *Tested on: BTRFS, EXT4, XFS, ZFS*
        | *Available on: Linux*

    ``ioctl``
        Uses the ``ioctl`` syscall with ``FICLONERANGE`` option which, if
        the underlying filesystem supports it, uses copy-on-write semantics
        to clone the data. Officially introduced in Linux 4.5, but was
        previously known as ``BTRFS_IOC_CLONE_RANGE``, which was used to
        provide copy-on-write semantics to BTRFS filesystems since Linux
        2.6.29. Some network filesystems might use this to perform
        server-side copies. Will fail if not supported by the underlying
        filesystem.

        | *Tested on: BTRFS*
        | *Available on: Linux*

    ``sendfile``
        Uses the ``sendfile`` syscall which performs in-kernel copy,
        avoiding having to copy the data into application memory.

        | *Tested on: BTRFS, XFS, EXT4*
        | *Available on: Linux, Solaris*

    ``duplicate_extents``
        Uses Windows Block Cloning via ``FSCTL_DUPLICATE_EXTENTS_TO_FILE``,
        which provides copy-on-write semantics to clone the data. Requires
        Windows 10 v1607 / Windows Server 2016 or later, and a compatible
        filesystem (ReFS, SMB 3.1.1, CsvFS). Will fail if not supported
        by the underlying filesystem.

        | *Tested on: ReFS*
        | *Available on: Windows*

    ``all``
        Tries all of the copy methods in the following order: ``ioctl``,
        ``copy_file_range``, ``sendfile``, ``duplicate_extents``,
        ``standard``.

        *Available on: All platforms*
