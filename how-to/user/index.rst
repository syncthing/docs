Users
=====

How to rename or move a synced folder
-------------------------------------

Syncthing doesn't have a direct way to do this, as it's potentially
dangerous to do so if you're not careful - it may result in data loss if
something goes wrong during the move and is synchronized to your other
devices.

The easy way to rename or move a synced folder on the local system is to
remove the folder in the Syncthing UI, move it on disk, then re-add it using
the new path.

It's best to do this when the folder is already in sync between your
devices, as it is otherwise unpredictable which changes will "win" after the
move. Changes made on other devices may be overwritten, or changes made
locally may be overwritten by those on other devices.

An alternative way is to shut down Syncthing, move the folder on disk (including
the ``.stfolder`` marker), edit the path directly in ``config.xml`` in the
configuration folder (see :ref:`config`) and then start Syncthing again.
