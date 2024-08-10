STTRACE
    Used to increase the debugging verbosity in specific or all facilities,
    generally mapping to a Go package. Enabling any of these also enables
    microsecond timestamps, file names plus line numbers. Enter a
    comma-separated string of facilities to trace. ``syncthing --help`` always
    outputs an up-to-date list. The valid facility strings are:

    Main and operational facilities:
        config
            Configuration loading and saving.
        db
            The database layer.
        main
            Main package.
        model
            The root hub; the largest chunk of the system. File pulling, index
            transmission and requests for chunks.
        scanner
            File change detection and hashing.
        versioner
            File versioning.

    Networking facilities:
        beacon
            Multicast and broadcast UDP discovery packets: Selected interfaces
            and addresses.
        connections
            Connection handling.
        dialer
            Dialing connections.
        discover
            Remote device discovery requests, replies and registration of
            devices.
        nat
            NAT discovery and port mapping.
        pmp
            NAT-PMP discovery and port mapping.
        protocol
            The BEP protocol.
        relay
            Relay interaction (``strelaysrv``).
        upnp
            UPnP discovery and port mapping.

    Other facilities:
        fs
            Filesystem access.
        events
            Event generation and logging.
        http
           REST API.
        sha256
            SHA256 hashing package (this facility currently unused).
        stats
            Persistent device and folder statistics.
        sync
            Mutexes. Used for debugging race conditions and deadlocks.
        upgrade
            Binary upgrades.
        walkfs
            Filesystem access while walking.

        all
            All of the above.

STBLOCKPROFILE
    Write block profiles to ``block-$pid-$timestamp.pprof`` every 20 seconds.

STCPUPROFILE
    Write a CPU profile to ``cpu-$pid.pprof`` on exit.

STDEADLOCKTIMEOUT
    Used for debugging internal deadlocks; sets debug sensitivity. Use only
    under direction of a developer.

STLOCKTHRESHOLD
    Used for debugging internal deadlocks; sets debug sensitivity. Use only
    under direction of a developer.

STGUIADDRESS
    Override GUI listen address.  Equivalent to passing :option:`--gui-address`.

STGUIAPIKEY
    Override the API key needed to access the GUI / REST API.  Equivalent to
    passing :option:`--gui-apikey`.

STGUIASSETS
    Directory to load GUI assets from. Overrides compiled in assets. Useful for
    developing webgui, commonly use ``STGUIASSETS=gui bin/syncthing``.

STHEAPPROFILE
    Write heap profiles to ``heap-$pid-$timestamp.pprof`` each time heap usage
    increases.

STNODEFAULTFOLDER
    Don't create a default folder when starting for the first time. This
    variable will be ignored anytime after the first run.  Equivalent to the
    :option:`--no-default-folder` flag.

STNORESTART
    Equivalent to the :option:`--no-restart` flag.

STNOUPGRADE
    Disable automatic upgrades.  Equivalent to the :option:`--no-upgrade` flag.

STPROFILER
    Set to a listen address such as "127.0.0.1:9090" to start the profiler with
    HTTP access, which then can be reached at
    http://localhost:9090/debug/pprof. See ``go tool pprof`` for more
    information.

STPERFSTATS
    Write running performance statistics to ``perf-$pid.csv``. Not supported on
    Windows.

STRECHECKDBEVERY
    Time before folder statistics (file, dir, ... counts) are recalculated from
    scratch. The given duration must be parseable by Go's ``time.ParseDuration``. If
    missing or not parseable, the default value of 1 month is used. To force
    recalculation on every startup, set it to ``1s``.

STGCINDIRECTEVERY
    Sets the time interval in between database garbage collection runs.  The
    given duration must be parseable by Go's ``time.ParseDuration``.

GOMAXPROCS
    Set the maximum number of CPU cores to use. Defaults to all available CPU
    cores.

GOGC
    Percentage of heap growth at which to trigger GC. Default is 100. Lower
    numbers keep peak memory usage down, at the price of CPU usage
    (i.e. performance).

LOGGER_DISCARD
    Hack to completely disable logging, for example when running benchmarks.
    Set to any nonempty value to use it.
