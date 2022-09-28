syncXattrs
==========

.. versionadded:: 1.22.0

When enabled, Syncthing attempts to also synchronize filesystem extended
attributes between devices.

Extended attributes are recorded and synced on a per-operating system basis.
That is, if a file has an extended attribute on one operating system, it
will not be synced to another operating system. For example, if a file has
an extended attribute on Linux, it will not be synced to macOS. This is
because the extended attribute names, values and interpretation is different
on different operating systems.

Syncthing will attempt to preserve information about foreign extended
attributes when a file is manipulated. That is, if a file is modified on
Linux it's expected that macOS extended attributes will be preserved in the
metadata and hence applied again on macOS. This is not always possible - for
example if a file is *moved* on Linux, the macOS extended attributes for
that file will be lost.

As of Syncthing 1.22.0 the supported operating systems are Linux, macOS,
FreeBSD and NetBSD.

.. note::
  In order for there to be extended attributes to apply, the peer device
  must have either ``syncXattrs`` or :doc:`folder-send-xattrs` enabled.

Elevated permissions
~~~~~~~~~~~~~~~~~~~~

Syncthing, when running as a normal user account, may not have permission to
access or manipulate all extended attributes. See
:doc:`folder-sync-ownership` for more information of granting appropriate
permissions.

.. seealso:: :doc:`folder-send-xattrs`
