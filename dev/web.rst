Working on the Web GUI
======================

Syncthing has a Bootstrap & Javascript based frontend that is built from the
``gui/`` directory in the source distribution. The first level of
directories underneath correspond to the currently selected theme. By
default Syncthing will therefore use the files inside ``gui/default/`` for
the GUI, although the other themes (``dark``, ``black``, ``light``) can be
selected in the settings dialog. When Syncthing is compiled all of these
files are packed into the binary.

When working on the GUI it's not convenient to have to rebuild and restart
the binary to see a change. Instead an environment variable can be set
pointing to the raw files, and Syncthing will use these directly from disk
when serving the GUI. Assuming that you have checked out the source code in
``~/dev/syncthing``, you can run Syncthing like this::

    $ export STGUIASSETS=~/dev/syncthing/gui
    $ syncthing

When you visit the Syncthing GUI on http://localhost:8384/ the GUI will be
served directly from the source code assets. You can change these and simply
reload in the browser to see your changes applied.

On Windows, assuming a checkout in ``C:\dev\syncthing`` the corresponding
sequence is::

    C:\> SET STGUIASSETS=C:\dev\syncthing\gui
    C:\> syncthing
