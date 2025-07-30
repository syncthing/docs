STTRACE
    Used to increase the debugging verbosity in specific or all facilities,
    generally mapping to a Go package. Enabling any of these also enables
    microsecond timestamps, file names plus line numbers. Enter a
    comma-separated string of facilities to trace. ``syncthing --help`` always
    outputs an up-to-date list. The valid facility strings are:

    Main and operational facilities:
        config
            Configuration loading and saving.
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
        sqlite
            The SQLite database
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

STLOCKTHRESHOLD
    Used for debugging internal deadlocks; sets debug sensitivity. Use only
    under direction of a developer.

STVERSIONEXTRA
    Add extra information to the version string in logs and the version line
    in the GUI. Can be set to the name of a wrapper or tool controlling
    syncthing to communicate this to the end user.

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
