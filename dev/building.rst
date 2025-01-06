.. _building:

Building Syncthing
==================

.. note::
    You probably only need to go through the build process if you are going
    to do development on Syncthing or if you need to do a special packaging
    of it. For all other purposes we recommend using the official binary
    releases instead.

    If you still prefer to build Syncthing from source for your own use, be
    aware that there is a built-in automatic upgrade mechanism that will
    overwrite your built binary with a downloaded version. To avoid this, you
    can use the ``--no-upgrade`` build flag (see below), but you will be
    responsible for your own upgrades.

Branches and Tags
-----------------

You should base your work on the ``main`` branch when doing your
development. This branch is usually what will be going into the next
release and always what pull requests should be based on.

If you're looking to build and package a release of Syncthing you should
instead use the latest tag (``vX.Y.Z``) as the contents of ``main``
may be unstable and unsuitable for general consumption.

Prerequisites
-------------

-  The latest stable version of Go. The previous stable version should also
   work; older versions will likely not work. This largely follows Go's
   `Release Policy <https://go.dev/doc/devel/release#policy>`__.
-  Git
-  If you want to build Debian packages FPM is required. See FPM's
   `installation information <https://fpm.readthedocs.io/en/latest/installation.html>`__.
-  To build Windows executables, installing `goversioninfo
   <https://github.com/josephspurrier/goversioninfo>`__ is recommended
   in order to add file properties and icon to the compiled binaries.
-  Building Android binaries requires `Android NDK <https://developer.android.com/ndk>`__.

If you're not already a Go developer, the easiest way to get going
is to download the latest version of Go as instructed in
https://go.dev/doc/install.

.. note::
        Because Syncthing uses Go modules you do not need to set or care about "GOPATH".
        However, the GOPATH still defaults to ``~/go`` and you'd be best to *not*
        put your Syncthing source in there, for now.

Building (Unix)
---------------

- Install the prerequisites.
- Open a terminal.
- Run the commands below.

.. code-block:: bash

    # Pick a place for your Syncthing source.
    $ mkdir -p ~/dev
    $ cd ~/dev

    # Grab the code.
    $ git clone https://github.com/syncthing/syncthing.git

    # Now we have the source. Time to build!
    $ cd syncthing

    # You should be inside ~/dev/syncthing right now.
    $ go run build.go

Unless something goes wrong, you will have a ``syncthing`` binary built
and ready in ``~/dev/syncthing/bin``.

Building (Windows)
------------------

- Install the prerequisites.
- Open a ``cmd`` Window.
- Run the commands below.

.. code-block:: batch

    # Pick a place for your Syncthing source.
    > md "%USERPROFILE%\dev"
    > cd /d "%USERPROFILE%\dev"

    # Grab the code.
    > git clone https://github.com/syncthing/syncthing.git

    # Now we have the source. Time to build!
    > cd syncthing
    > go run build.go

Unless something goes wrong, you will have a ``syncthing.exe`` binary
built and ready in ``%USERPROFILE%\dev\syncthing\bin``.

Subcommands and Options
-----------------------

The following ``build.go`` subcommands and options exist.

``go run build.go install``
  Installs binaries in ``./bin`` (default command, this is what happens when
  build.go is run without any commands or parameters).

``go run build.go build``
  Builds just the named target, or ``syncthing`` by default, to the current
  directory. Use this when cross compiling, with parameters for what to cross
  compile to: ``go run build.go --goos linux --goarch 386 build``.

``go run build.go test``
  Runs the tests.

``go run build.go deb``
  Creates a Debian package in the current directory. Requires FPM
  and a Unixy build.

``go run build.go tar``
  Creates a Syncthing tar.gz dist file in the current directory. Assumes a
  Unixy build.

``go run build.go zip``
  Creates a Syncthing zip dist file in the current directory. Assumes a
  Windows build.

The options ``--no-upgrade``, ``--goos`` and ``--goarch`` can be given to
influence ``build``, ``tar`` and ``zip``. Examples:

``go run build.go --goos linux --goarch 386 tar``
  Builds a tar.gz distribution of Syncthing for linux-386.

``go run build.go --goos windows --no-upgrade zip``
  Builds a zip distribution of Syncthing for Windows (current architecture) with
  upgrading disabled.

.. _versiontagging:

Version Tagging
---------------

The binaries are "tagged" with a version derived from the current Git commit
(or the ``RELEASE`` file, see below) and the current username and hostname.
The username and hostname can be overridden by the ``BUILD_USER`` and
``BUILD_HOST`` environment variables, for example::

  $ BUILD_USER=builder BUILD_HOST=buildhost.local go run build.go
  $ ./bin/syncthing --version
  syncthing v1.8.0 ... builder@buildhost.local 2020-07-30 11:49:14 UTC

In addition the timestamp (by default taken from the current Git commit) can
be overridden by the ``SOURCE_DATE_EPOCH`` variable, in Unix epoch seconds.

Building without Git
--------------------

Syncthing can be built perfectly fine from a source tarball of course.
If the tarball is from our build server it contains a file called
``RELEASE`` that informs the build system of the version being
built. If you're building from a different source package, for example
one automatically generated by GitHub, you must instead pass the
``--version`` flag to ``build.go``.

If you are building something that will be installed as a package
(Debian, RPM, ...) you almost certainly want to use ``--no-upgrade`` as
well to prevent the built in upgrade system from being activated.

``go run build.go --version v0.10.26 --no-upgrade tar``
  Builds a tar.gz distribution of Syncthing for the current OS/arch, tagged as
  ``v0.10.26``, with upgrades disabled.
