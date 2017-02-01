.. _debugging:

.. todo::
    This page mostly duplicates syncthing(1). Needs merging.

Debugging Syncthing
===================

There's a lot that happens behind the covers, and Syncthing is generally
quite silent about it. A number of environment variables can be used to
set the logging to verbose for various parts of the program, and to
enable profiling.

Environment Variables
---------------------

STTRACE
~~~~~~~

The environment variable ``STTRACE`` can be set to a comma separated
list of "facilities", to enable extra debugging information for said
facility. A facility generally maps to a Go package, although there are
a few extra that map to parts of the ``main`` package. Currently, the
following facilities are supported (an up to date list is always printed
by ``syncthing --help``):

-  ``beacon`` (Multicast and broadcast discovery)
-  ``config`` (Configuration loading and saving)
-  ``connections`` (Connection handling)
-  ``db`` (The database layer)
-  ``dialer`` (Dialing connections)
-  ``discover`` (Remote device discovery)
-  ``events`` (Event generation and logging)
-  ``http`` (REST API)
-  ``main`` (Main package)
-  ``model`` (The root hub)
-  ``nat`` (NAT discovery and port mapping)
-  ``pmp`` (NAT-PMP discovery and port mapping)
-  ``protocol`` (The BEP protocol)
-  ``scanner`` (File change detection and hashing)
-  ``sha256`` (SHA256 hashing package)
-  ``stats`` (Persistent device and folder statistics)
-  ``sync`` (Mutexes)
-  ``upgrade`` (Binary upgrades)
-  ``upnp`` (UPnP discovery and port mapping)
-  ``versioner`` (File versioning)
-  ``all`` (all of the above)

The debug output is often of the kind that it doesn't make much sense
without looking at the code. The purpose of some of the different packages /
facilities are something like this:

-  ``beacon`` sends and receives UDP broadcasts used by the local
   discovery system. Debugging here will show which interfaces and
   addresses are selected for broadcasts, etc.
-  ``discover`` sends and receives local discovery packets. Debugging
   here will output the parsed packets, devices that get registered etc.
-  ``model`` is the largest chunk of the system; this is where pulling
   of out of date files happens, indexes are sent and received, and incoming
   requests for file chunks are logged.
-  ``scanner`` is the local filesystem scanner. Debugging here will
   output information about changed and unchanged files.

Enabling any of the facilities will also change the log format to
include microsecond timestamps and file names plus line numbers. This
can be used to enable this extra information on the normal logging
level, without enabling any debugging: ``STTRACE=somethingnonexistent``
for example.

Under Unix (including Mac) the easiest way to run Syncthing with an
environment variable set is to prepend the variable to the command line.
I.e:

``$ STTRACE=model syncthing``

On windows, it needs to be set prior to running Syncthing.

::

    C:\> set STTRACE=model
    C:\> syncthing

STPROFILER
~~~~~~~~~~

The ``STPROFILER`` environment variable sets the listen address for the
HTTP profiler. If set to for example ``:9090`` the profiler will start
and listen on port 9090. http://localhost:9090/debug/pprof is then the
address to the profiler. Se ``go tool pprof`` for more information.

STGUIASSETS
~~~~~~~~~~~

Directory to load GUI assets from. Overrides compiled in assets. Useful
for developing webgui, commonly use ``STGUIASSETS=gui bin/syncthing``

STCPUPROFILE
~~~~~~~~~~~~

Write a CPU profile to ``cpu-$pid.pprof`` on exit.

STHEAPPROFILE
~~~~~~~~~~~~~

Write heap profiles to ``heap-$pid-$timestamp.pprof`` each time
heap usage increases.

STBLOCKPROFILE
~~~~~~~~~~~~~~

Write block profiles to ``block-$pid-$timestamp.pprof`` every 20
seconds.

STPERFSTATS
~~~~~~~~~~~

Write running performance statistics to ``perf-$pid.csv``. Not supported on
Windows.

STNOUPGRADE
~~~~~~~~~~~

Disable automatic upgrades.

GOMAXPROCS
~~~~~~~~~~

Set the maximum number of CPU cores to use. Defaults to all available
CPU cores.

GOGC
~~~~

Percentage of heap growth at which to trigger GC. Default is 100. Lower
numbers keep peak memory usage down, at the price of CPU usage (ie.
performance)
