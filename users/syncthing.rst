Syncthing
=========

Synopsis
--------

::

    syncthing [-audit] [-browser-only] [-generate=<dir>] [-gui-address=<address>] [-gui-apikey=<key>]
              [-home=<dir>] [-logfile=<filename>] [-logflags=<flags>] [-no-browser]
              [-no-console] [-no-restart] [-paths] [-paused] [-reset-database] [-reset-deltas] [-upgrade] [-upgrade-check]
              [-upgrade-to=<url>] [-verbose] [-version]

Description
-----------

Syncthing is an application that lets you synchronize your files across multiple
devices. This means the creation, modification or deletion of files on one
machine will automatically be replicated to your other devices. We believe your
data is your data alone and you deserve to choose where it is stored. Therefore
Syncthing does not upload your data to the cloud but exchanges your data across
your machines as soon as they are online at the same time.

Options
-------

.. cmdoption:: -audit

    Write events to audit file.

.. cmdoption:: -generate=<dir>

    Generate key and config in specified dir, then exit.

.. cmdoption:: -gui-address=<address>

    Override GUI address.

.. cmdoption:: -home=<dir>

    Set configuration directory. The default configuration directory is:
    ``$HOME/.config/syncthing``.

.. cmdoption:: -logfile=<filename>

    Set destination filename for logging (use ``"-"`` for stdout, which is the default option).

.. cmdoption:: -logflags=<flags>

    Select information in log line prefix, default 2. The ``-logflags`` value is
    a sum of the following:

    -  1: Date
    -  2: Time
    -  4: Microsecond time
    -  8: Long filename
    - 16: Short filename

    To prefix each log line with date and time, set ``-logflags=3`` (1 + 2 from
    above). The value 0 is used to disable all of the above. The default is to
    show time only (2).

.. cmdoption:: -no-browser

    Do not start a browser.

.. cmdoption:: -no-console

    Hide the console window. (On Windows only)

.. cmdoption:: -no-restart

    Do not restart; just exit.

.. cmdoption:: -paths

    Print the paths used for configuration, keys, database, GUI overrides, default sync folder and the log file.

.. cmdoption:: -reset-database

    Reset the database, forcing a full rescan and resync.

.. cmdoption:: -reset-deltas

    Reset delta index IDs, forcing a full index exchange.

.. cmdoption:: -upgrade

    Perform upgrade.

.. cmdoption:: -upgrade-check

    Check for available upgrade.

.. cmdoption:: -upgrade-to=<url>

    Force upgrade directly from specified URL.

.. cmdoption:: -verbose

    Print verbose log output.

.. cmdoption:: -version

    Show version.

Deprecated Options
------------------

.. cmdoption:: -gui-apikey=<key>

    Override GUI API key (deprecated as of v0.13, to be removed in the future - do not use).

Exit Codes
----------

0
    Success / Shutdown
1
    Error
2
    Upgrade not available
3
    Restarting
4
    Upgrading

Some of these exit codes are only returned when running without a monitor
process (with environment variable ``STNORESTART`` set). Exit codes over 125 are
usually returned by the shell/binary loader/default signal handler. Exit codes
over 128+N on Unix usually represent the signal which caused the process to
exit. For example, ``128 + 9 (SIGKILL) = 137``.

Development Settings
--------------------

The following environment variables modify Syncthing's behavior in ways that
are mostly useful for developers. Use with care.
If you start Syncthing from within service managers like systemd or supervisor,
path expansion may not be supported.

STNODEFAULTFOLDER
    Don't create a default folder when starting for the first time. This
    variable will be ignored anytime after the first run.
STGUIASSETS
    Directory to load GUI assets from. Overrides compiled in assets.
STTRACE
    A comma separated string of facilities to trace. The valid facility strings
    are:

    main
        Main package
    model
        The root hub
    config
        Configuration loading and saving
    db
        The database layer
    scanner
        File change detection and hashing
    versioner
        File versioning

    beacon
        Multicast and broadcast discovery
    connections
        Connection handling
    dialer
        Dialing connections
    discover
        Remote device discovery
    relay
        Relay interaction
    protocol
        The BEP protocol
    nat
        NAT discovery and port mapping
    pmp
        NAT-PMP discovery and port mapping
    upnp
        UPnP discovery and port mapping

    events
        Event generation and logging
    http
       REST API
    sha256
        SHA256 hashing package (debug currently unused)
    stats
        Persistent device and folder statistics
    sync
        Mutexes
    upgrade
        Binary upgrades

all
        All of the above

STPROFILER
    Set to a listen address such as "127.0.0.1:9090" to start the profiler with
    HTTP access.
STCPUPROFILE
    Write a CPU profile to cpu-$pid.pprof on exit.
STHEAPPROFILE
    Write heap profiles to ``heap-$pid-$timestamp.pprof`` each time heap usage
    increases.
STBLOCKPROFILE
    Write block profiles to ``block-$pid-$timestamp.pprof`` every 20 seconds.
STPERFSTATS
    Write running performance statistics to ``perf-$pid.csv``. Not supported on
    Windows.
STDEADLOCK
    
STDEADLOCKTIMEOUT
    
STDEADLOCKTHRESHOLD
    
STNOUPGRADE
    Disable automatic upgrades.
STHASHING
    Specifiy which hashing package to use. Defaults to automatic based on
    peformance. Specify "minio" (compatibility) or "standard" for the default Go implementation.
GOMAXPROCS
    Set the maximum number of CPU cores to use. Defaults to all available CPU
    cores.
GOGC
    Percentage of heap growth at which to trigger GC. Default is 100. Lower
    numbers keep peak memory usage down, at the price of CPU usage
    (i.e. performance).

See Also
--------

:manpage:`syncthing-config(5)`, :manpage:`syncthing-stignore(5)`,
:manpage:`syncthing-device-ids(7)`, :manpage:`syncthing-security(7)`,
:manpage:`syncthing-networking(7)`, :manpage:`syncthing-versioning(7)`,
:manpage:`syncthing-faq(7)`
