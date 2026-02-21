Usage
=====

What if there is a conflict?
----------------------------

.. seealso::
    :ref:`conflict-handling`

.. _marker-faq:

How do I serve a folder from a read only filesystem?
----------------------------------------------------

Syncthing requires a "folder marker" to indicate that the folder is present
and healthy. By default this is a directory called ``.stfolder`` that is
created by Syncthing when the folder is added. If this folder can't be
created (you are serving files from a CD or something) you can instead set
the advanced config ``Marker Name`` to the name of some file or folder that
you know will always exist in the folder.

I really hate the ``.stfolder`` directory, can I remove it?
-----------------------------------------------------------

See the previous question.

Am I able to nest shared folders in Syncthing?
----------------------------------------------

Sharing a folder that is within an already shared folder is possible, but it has
its caveats. What you must absolutely avoid are circular shares. This is just
one example, there may be other undesired effects. Nesting shared folders is not
supported, recommended or coded for, but it can be done successfully when you
know what you're doing - you have been warned.

How do I rename/move a synced folder?
-------------------------------------

Syncthing doesn't have a direct way to do this, as it's potentially
dangerous to do so if you're not careful - it may result in data loss if
something goes wrong during the move and is synchronized to your other
devices.

The easy way to rename or move a synced folder on the local system is to
remove the folder in the Syncthing UI, move it on disk, then re-add it using
the new path.

It's important to do this when the folder is already in sync between your
devices, as it is otherwise unpredictable which changes will "win" after the
move. Changes made on other devices may be overwritten, or changes made
locally may be overwritten by those on other devices.

An alternative way is to shut down Syncthing, move the folder on disk (including
the ``.stfolder`` marker), edit the path directly in ``config.xml`` in the
configuration folder (see :doc:`/users/config`) and then start Syncthing again.

How do I configure multiple users on a single machine?
------------------------------------------------------

Each user should run their own Syncthing instance. Be aware that you might need
to configure listening ports such that they do not overlap (see :doc:`/users/config`).

Does Syncthing support syncing between folders on the same system?
------------------------------------------------------------------

No. Syncthing is not designed to sync locally and the overhead involved in
doing so using Syncthing's method would be wasteful. There are better
programs to achieve this such as `rsync <https://rsync.samba.org/>`__ or
`Unison <https://www.cis.upenn.edu/~bcpierce/unison>`__.

When I do have two distinct Syncthing-managed folders on two hosts, how does Syncthing handle moving files between them?
------------------------------------------------------------------------------------------------------------------------

Syncthing does not specially handle this case, and most files will most likely get
re-downloaded.

In detail, the behavior depends on the scan order. If you have folders A and B,
and move files from A to B, if A gets scanned first, it will announce the removal of
the files to others who will then remove the files. As you rescan B, B will
announce the addition of new files, and other peers will have nowhere to get
them from apart from re-downloading them.

If B gets rescanned first, B will announce additions first, and remote
peers will then reconstruct the files (not rename, more like copying block by
block) from A, and then as A gets rescanned, it will remove the files from A.

A workaround would be to copy first from A to B, rescan B, wait for B to
copy the files on the remote side, and then delete from A.

Can I help initial sync by copying files manually?
--------------------------------------------------

If you have a large folder that you want to keep in sync over a not-so-fast network, and you have the possibility to move all files to the remote device in a faster manner, here is a procedure to follow:

- Create the folder on the local device, but don't share it with the remote device yet.
- Copy the files from the local device to the remote device using regular file copy. If this takes a long time (perhaps requiring travelling there physically), it may be a good idea to make sure that the files on the local device are not updated while you are doing this.
- Create the folder on the remote device, and copy the Folder ID from the folder on the local device, as we want the folders to be considered the same. Then wait until scanning the folder is done.

- Now share the folder with the other device, on both sides. Syncthing will exchange file information, updating the database, but existing files will not be transferred. This may still take a while initially, be patient and wait until it settled.

Is Syncthing my ideal backup application?
-----------------------------------------

No. Syncthing is not a great backup application because all changes to your
files (modifications, deletions, etc.) will be propagated to all your
devices. You can enable versioning, but we encourage you to use other tools
to keep your data safe from your (or our) mistakes.

How can I exclude files with brackets (``[]``) in the name?
-----------------------------------------------------------

The patterns in .stignore are glob patterns, where brackets are used to
denote character ranges. That is, the pattern ``q[abc]x`` will match the
files ``qax``, ``qbx`` and ``qcx``.

To match an actual file *called* ``q[abc]x`` the pattern needs to "escape"
the brackets, like so: ``q\[abc\]x``.

On Windows, use the ``|`` character to escape the brackets (``q|[abc|]x``), as the ``\``
character is used as a path separator.

How do I access the web GUI from another computer?
--------------------------------------------------

The default listening address is 127.0.0.1:8384, so you can only access the GUI
from the same machine.  This is for security reasons.  To access it from another
computer, change the ``GUI listen address`` option in the web GUI from
``127.0.0.1:8384`` to ``0.0.0.0:8384``, or change the ``config.xml``:

.. code-block:: xml

    <gui enabled="true" tls="false">
      <address>127.0.0.1:8384</address>

to

.. code-block:: xml

    <gui enabled="true" tls="true">
      <address>0.0.0.0:8384</address>

Then the GUI is accessible from everywhere.  There is no filtering based on
e.g. source address (use a firewall for that).  You should set a password and
enable HTTPS with this configuration.  You can do this from inside the GUI.

If both your computers are Unix-like (Linux, Mac, etc.) you can also leave the
GUI settings at default and use an SSH port forward to access it.  For example,

.. code-block:: bash

    $ ssh -L 9090:127.0.0.1:8384 user@othercomputer.example.com

will log you into ``othercomputer.example.com``, and present the *remote*
Syncthing GUI on http://localhost:9090 on your *local* computer.

If you only want to access the remote GUI and don't want the terminal session,
use this example:

.. code-block:: bash

    $ ssh -N -L 9090:127.0.0.1:8384 user@othercomputer.example.com

If only your remote computer is Unix-like, you can still access it with SSH from
Windows.  Under Windows 10 or later you can use the same ``ssh`` command if you
`install the OpenSSH Client <https://learn.microsoft.com/windows-server/administration/openssh/openssh_install_firstuse>`__.

I don't like the GUI or the theme. Can it be changed?
-----------------------------------------------------

You can change the theme in the settings. Syncthing ships with other themes
than the default.

If you want a custom theme or a completely different GUI, you can add your
own.
By default, Syncthing will look for a directory ``gui`` inside the Syncthing
home folder. To change the directory to look for themes, you need to set the
STGUIASSETS environment variable. To get the concrete directory, run
syncthing with the ``--paths`` parameter. It will print all the relevant paths,
including the "GUI override directory".

To add e.g. a red theme, you can create the file ``red/assets/css/theme.css``
inside the GUI override directory to override the default CSS styles.

To create a whole new GUI, you should checkout the files at
https://github.com/syncthing/syncthing/tree/main/gui/default
to get an idea how to do that.


How do I upgrade Syncthing?
---------------------------

If you use a package manager such as Debian's apt-get, you should upgrade
using the package manager. If you use the binary packages linked from
Syncthing.net, you can use Syncthing's built-in automatic upgrade functionality.

- If automatic upgrades is enabled (which is the default), Syncthing will
  upgrade itself automatically within 24 hours of a new release.

- The upgrade button appears in the web GUI when a new version has been
  released. Pressing it will perform an upgrade.

- To force an upgrade from the command line, run ``syncthing --upgrade``.

Note that your system should have CA certificates installed which allows a
secure connection to GitHub (e.g. FreeBSD requires ``sudo pkg install
ca_root_nss``). If ``curl`` or ``wget`` works with normal HTTPS sites, then
so should Syncthing.

Where do I find the latest release?
-----------------------------------

We release new versions through GitHub. The latest release is always found
`on the release page
<https://github.com/syncthing/syncthing/releases/latest>`_. Unfortunately
GitHub does not provide a single URL to automatically download the latest
version. We suggest to use the `GitHub API <https://api.github.com/repos/syncthing/syncthing/releases/latest>`__ and parsing
the JSON response.


How do I run Syncthing as a daemon process on Linux?
----------------------------------------------------

If you're using systemd, runit, or upstart, we ship `example configurations <https://github.com/syncthing/syncthing/tree/main/etc>`__.

If however you're not using one of these tools, you have a couple of options.
If your system has a tool called ``start-stop-daemon`` installed (that's the name
of the command, not the package), look into the local documentation for that, it
will almost certainly cover 100% of what you want to do.  If you don't have
``start-stop-daemon``, there are a bunch of other software packages you could use
to do this.  The most well known is called daemontools, and can be found in the
standard package repositories for almost every modern Linux distribution.
Other popular tools with similar functionality include S6 and the aforementioned
runit.

.. _inotify-limits:

How do I increase the inotify limit to get my filesystem watcher to work?
-------------------------------------------------------------------------

You are probably reading this because you encountered the following error with
the filesystem watcher on linux:

    Failed to start filesystem watcher for folder yourLabel (yourID): failed to
    set up inotify handler. Please increase inotify limits, see https://docs.syncthing.net/users/faq.html#inotify-limits

Linux typically restricts the number of watches per user (usually 8192). If
you have many directories, you will need to adjust that number.

On many Linux distributions you can run the following to fix it::

    echo "fs.inotify.max_user_watches=204800" | sudo tee -a /etc/sysctl.conf

On Arch Linux and potentially others it is preferred to write this line into a
separate file, i.e. you should run::

    echo "fs.inotify.max_user_watches=204800" | sudo tee -a /etc/sysctl.d/90-inotify-max-user-watches.conf

This only takes effect after a reboot. To adjust the limit immediately, run::

    echo 204800 | sudo tee /proc/sys/fs/inotify/max_user_watches

How do I reset the GUI password?
--------------------------------

If you've forgotten / lost the GUI password, you can reset it using the
:option:`--gui-password` (and possibly :option:`--gui-user`) options to the
``syncthing generate`` subcommand.  This should be done while Syncthing is not
running.

1. Stop Syncthing: ``syncthing cli operations shutdown``
2. ``syncthing generate --gui-password=myNewPassword --gui-user=newUserName``
3. Restart Syncthing as usual.

*Alternatively, in step 2*, you can manually delete the :stconf:opt:`<user>
<gui.user>` and :stconf:opt:`<password> <gui.password>` XML tags from the
``<gui>`` block in file ``config.xml``.  The location of the file depends on the
OS and is described in the :doc:`configuration documentation </users/config>`.

For example, the two emphasized lines below would be removed from the file.

.. code-block:: text
    :emphasize-lines: 3,4

    <gui enabled="true" tls="false">
       <address>127.0.0.1:8384</address>
       <user>syncguy</user>
       <password>$2a$10$s9wWHOQe...Cq7GPye69</password>
       <apikey>9RCKohqCAyrj5RjpyZdR2wXmQ9PyQFeN</apikey>
       <theme>default</theme>
    </gui>
