.. _debugging:

Debugging Syncthing
===================

A file isn't syncing
--------------------

If you wonder why a given file is out of sync, or you suspect it's caused by
a bug, we need to collect some information. You'll need *the name of the file*
and *the folder ID* it's in.

To begin with, check the GUI on the *receiving* side. It may be listed under
"Failed items", with a cause for the failure. Secondly, check the logs on
the *receiving* side. If the file has failed to sync this will be mentioned
in the logs even if it is not currently visible in the GUI. (If you are
debugging this together with someone on the forum, post screenshots of the
GUI and the logs in question.)

If nothing relevant showed up so far, or the question is why the file is
considered out of sync to begin with in a receive-only setup, etc., we can
look closer at what Syncthing knows about the file. You'll need to do these
steps on *both* the receiving side and sending side. We will need to use the
Syncthing CLI for this.

1. An extract of database metadata for the file, taked by the command
   ``syncthing debug database-file $folderID $fileName``, where
   ``$folderID`` represents the folder ID (e.g. ``abcd-1234``) and
   ``fileName`` is the name of the file, including any directories, relative
   to the folder root. (Syncthing v2 only; but we're probably not debugging
   this on v1 any more.)

2. Details of the file information for the file, taked by the command
   ``syncthing cli debug file $folderID $fileName`` (same folder ID and file
   name as above).

Post the output from both of these, verbatim, indicating which is the
receiving side and which is the sending side.

Debug logs
----------

There's a lot that happens behind the covers, and Syncthing is generally
quite silent about it. A number of environment variables can be used to
set the logging to verbose for various parts of the program, and to
enable profiling.

Under Unix (including Mac) the easiest way to run Syncthing with an
environment variable set is to prepend the variable to the command line.
I.e:

``$ STTRACE=model syncthing``

On windows, it needs to be set prior to running Syncthing.

::

    C:\> set STTRACE=model
    C:\> syncthing

Environment variables
---------------------

.. include:: ../includes/env-vars.rst

Stepping with breakpoints
-------------------------

If you like to step through the running program, build a non-optimized binary and run with  https://github.com/derekparker/delve.

Follow these steps:

::

    $ go run build.go -debug-binary build
    $ STNODEFAULTFOLDER=1 STNOUPGRADE=1  STNORESTART=1 dlv --listen=:2345 --headless=true --api-version=2 exec ./syncthing -- --home=./_test_config --no-browser

For installing and using delve itself see:

-  VSCode (Microsoft): https://github.com/golang/vscode-go/blob/master/docs/debugging.md

-  GoLand (JetBrains): create remote run configuration and follow the two steps displayed
