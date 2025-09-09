STTRACE
    Used to increase the debugging verbosity in specific facilities,
    generally mapping to a Go package. Enter a comma-separated string of
    facilities to trace: ``api,beacon``. Optionally, a log level can be
    given per facility to specify something other than DEBUG:
    ``api:WARN,beacon:ERR``, potentially overriding a global ``--log-level``
    adjustment.

    The valid facility strings are listed below; additionally, ``syncthing
    serve --help`` always outputs the most up-to-date list.

    api
        REST API
    beacon
        Multicast and broadcast discovery
    config
        Configuration loading and saving
    connections
        Connection handling
    db/sqlite
        SQLite database
    dialer
        Dialing connections
    discover
        Remote device discovery
    events
        Event generation and logging
    fs
        Filesystem access
    main
        Main package
    model
        The root hub
    nat
        NAT discovery and port mapping
    pmp
        NAT-PMP discovery and port mapping
    protocol
        The BEP protocol
    relay/client
        Relay client
    scanner
        File change detection and hashing
    stun
        STUN functionality
    syncthing
        Main run facility
    upgrade
        Binary upgrades
    upnp
        UPnP discovery and port mapping
    ur
        Usage reporting
    versioner
        File versioning
    watchaggregator
        Filesystem event watcher

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
