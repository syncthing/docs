syncOwnership
=============

.. versionadded:: 1.21.0

When enabled, Syncthing attempts to also synchronise ownership between
devices. Ownership is divided into two kinds:

- "Unix" ownership, which is the Unix style UID and GID, and
- "Windows" ownership, which the Windows file ownership concept.

The two are not mixed -- that is, ownership information is not synchronised
between POSIX and Windows systems, only POSIX-to-POSIX and
Windows-to-Windows. Nonetheless Syncthing attempts to leave unknown
ownership data intact: POSIX systems will not overwrite Windows ownership
data but pass it on unmodified, and vice versa.

.. note::
  In order for there to be ownership information to apply, the peer device
  must have either ``syncOwnership`` or :doc:`folder-send-ownership` enabled.

Unix implementation
-------------------

Syncthing records both the numerical UID and GID for a file and the
corresponding user and group names, when they are known. When applying
ownership Syncthing will first attempt to look up a local user or group with
the given name, and if that fails it will fall back to the numerical UID and
GID.

Elevated permissions
~~~~~~~~~~~~~~~~~~~~

Syncthing, when running as a normal user account, doesn't have permission to
alter file ownership. There are several reasonable ways of running Syncthing
with elevated permissions to enable ownership sync:

- As root, in a Docker container, with the synced data mounted as a
  volume.
- As a normal user, with extra capabilities granted to the executable.

To grant extra capabilities, the following steps must be taken:

- The executable must be owned by root and not writable by normal users.
- The executable must be granted the CHOWN and FOWNER capabilities. The
  CHOWN capability is required to be able to change ownership on the file.
  However, once that has been done Syncthing may no longer have permission
  to act on the file in other ways while running as anonymous user. The
  FOWNER capability overrides this.

Example commands of setting Syncthing up in this manner::

    % sudo chown root /usr/local/bin/syncthing
    % sudo chmod 755 /usr/local/bin/syncthing
    % sudo setcap CAP_CHOWN,CAP_FOWNER=pe /usr/local/bin/syncthing

.. note:: Note that automated upgrades cannot be used with Syncthing elevated
   in this manner as any automated upgrade would undo the capabilities granted.

When using systemd to start the service automatically, the capabilities can be
set in the unit file instead of touching the executable, see
:ref:`autostart-systemd-permissions`.

Windows implementation
----------------------

Syncthing records the account name of the owner, and whether it is a group
or user. On the receiving side a user or group with the corresponding name
is looked up and set as the owner.

On Windows, syncing ownership has a fairly significant performance impact on
scan times.

.. seealso:: :doc:`folder-send-ownership`
