.. default-domain:: stconf

Configuration Tuning
====================

Syncthing ships with a set of default values designed to work well for most
users under common circumstances. However, there are a number of
configuration options that can be tweaked to improve performance or to
reduce resource consumption. This article makes recommendations, but you
will need to experiment to find the best settings for your particular setup.

Tuning for High Performance
------------------------------------------------

You have resources to spare and want to optimize for performance. To begin
with, make sure you are running a 64 bit build and that the index database
is on an SSD. First some general options:

- :opt:`progressUpdateIntervalS`
    Set to -1 to disable progress updates. Time spent on progress
    updates is time not spent on syncing.

- :opt:`setLowPriority`
    Set to false to run Syncthing with normal priority. This may allow
    Syncthing to use more CPU time, which can improve performance.

- :opt:`maxFolderConcurrency`
    Find an appropriate setting for the number of folders to sync
    concurrently, taking into account the inherent concurrency of the
    underlying storage system. If two folders are on the same underlying
    spinning disk, syncing them concurrently will be a strict loss of
    performance. If the underlying storage is a large array of disk,
    then syncing many folders concurrently may be beneficial.

- :opt:`databaseTuning`
    Set to ``large``. Regardless of the size of the actual database, this
    increases a number of buffers and settings to optimize for higher
    throughput.

- :opt:`maxConcurrentIncomingRequestKiB`
    This sets the maximum amount of data being processed (loaded from
    disk and transmitted over the network) concurrently at any given
    time. This is a global limiter, not per folder. If you have a lot of
    folders and/or a lot of devices and memory to spare you may want to
    increase this value. The default is 256 MiB, consider values of 1024
    MiB or more.

These options are folder specific and should be set on each folder:

- :opt:`fsWatcherEnabled`
    When possible, using filesystem notifications avoids unnecessary I/O in
    large periodic scans. Changes are detected faster using notifications
    than scans.

- :opt:`copiers`
    The number of routines used for the copy stage of file syncing. Similar
    to other concurrency options, if there are a lot of files to sync and if
    the I/O system can handle it it you may see increased performance by
    increasing this value. The default is system dependent, somewhere
    between 1 and the number of CPU cores available.

- :opt:`hashers`
    When hashing locally changed files, the number of hashing routines to
    use. Higher values mean more I/O and CPU load and may increase
    performance when there are a lot of files to hash, assuming I/O
    bandwidth and CPU are available. The default is system dependent,
    somewhere between 1 and the number of CPU cores available.

- :opt:`pullerMaxPendingKiB`
    The maximum amount of data to have outstanding requests for at any given
    time. Higher values may improve performance, especially if the network
    or I/O latency is high. The default is 32 MiB.

- :opt:`scanProgressIntervalS`
    Providing the GUI with scan progress updates is not very expensive, but
    is effort that could be spent on something more productive. If you don't
    need to see scan progress, set this to -1 to disable it.

- :opt:`weakHashThresholdPct`
    Syncthing will by default look for rolling (weak) hash matches to detect
    data shifted in a file if a lot of data has changed in the file. If your
    use case doesn't cause data to shift in a file, and if the files are
    large (movies, VM images, ...) it is unnecessary to spend time checking
    weak hashes. Set the threshold to 101% to disable use of weak hashes.

- :opt:`maxConcurrentWrites`
    Synchting limits the number of outstanding write system calls at any
    given time to avoid overloading the I/O system. If you increased
    copiers, outstanding network requests, or other settings that increase
    the number of concurrent writes, you may need to increase this value.
    The default is 2.

- :opt:`disableFsync`
    Syncthing calls ``fsync()`` on files and directories after syncing them
    to ensure they are safe and sound on stable storage. This is a good
    thing, but it can be expensive. If you have a lot of files to sync
    and/or a lot of I/O bandwidth available, you may see a performance
    increase by disabling ``fsync()``. This is not recommended for most
    setups, as you are increasing the risk of data loss in case of a power
    outage or system crash.

- :opt:`blockPullOrder`
    Syncthing by default uses a pseudo-random block order when pulling in
    order to distribute load over multiple devices better. If you are
    generally downloading files from only one device, and if you have
    spinning disks as the underlying storage, you may see a performance
    increase by setting this to ``inOrder``.

- :opt:`copyRangeMethod`
    If your underlying filesystem supports it, you may see a performance
    increase by enabling a copy-on-write method, as it reduces the amount of
    data actually copied on disk when syncing files.

- :opt:`caseSensitiveFS`
    If your underlying filesystem is case sensitive, you may see a
    performance increase by enabling this option. This disables a number of
    safety checks that are required for case insensitive filesystems, and
    can cause data loss if your underlying filesystem is *not* in fact case
    sensitive.

- :opt:`syncOwnership`/:opt:`syncXattrs`, :opt:`sendOwnership`/:opt:`sendXattrs`
    Use these if they are required for your use case, but keep in mind they
    have a fairly high performance cost.

For devices, consider the following:

- :opt:`numConnections`
    Set at or above the number of CPU cores available. This allows maximum
    concurrency for TLS connections and may improve performance.

Other things:

- ``GOMEMLIMIT`` and ``GOGC``: These environment variables can be used to
  control the garbage collector. For large setups, setting ``GOMEMLIMIT`` to
  the desired max amount of memory Syncthing should use can improve
  performance. The reason is that this reduces garbage collector frequency
  during lower memory usage. Read more in the `Go
  GC guide <https://golang.org/doc/gc-guide>`__.

Tuning for Low Resources
------------------------

You have limited resources and want Syncthing to use as few as possible. You
care less about performance.

General options:

- :opt:`progressUpdateIntervalS`
    Set to -1 to disable progress updates. Progress updates aren't
    absolutely essential and consume some amount of CPU and memory.

- :opt:`maxFolderConcurrency`
    Set to 1 to sync folders sequentially, reducing the peak memory usage.

- :opt:`databaseTuning`
    Set to ``small``. Regardless of the size of the actual database size,
    this reduces the size of a number of buffers to optimize for reduced
    memory usage.

- :opt:`maxConcurrentIncomingRequestKiB`
    Set to 32 MiB to reduce the amount of memory used for buffering
    responses to incoming requests.

Folders options:

- :opt:`fsWatcherEnabled`
    If possible, using the filesystem notifications is more efficient than
    doing full periodic scans.

- :opt:`copiers`, :opt:`hashers`
    Set to 1 to reduce the amount of concurrency when syncing and hashing a
    folder, reducing peak memory usage.

- :opt:`pullerMaxPendingKiB`
    Set to 16 MiB to reduce the amount of memory used for buffering
    while syncing.

- :opt:`scanProgressIntervalS`
    Set to -1 to disable scan progress updates. Keeping track of scan progress
    uses memory and CPU.

- :opt:`weakHashThresholdPct`
    Set to 101% to disable use of weak hashes. Using weak hashes has a
    memory cost.

- :opt:`copyRangeMethod`
    If your underlying filesystem supports it, using copyrange is more
    efficient than having Syncthing do the data copying.

- :opt:`caseSensitiveFS`
    If your underlying filesystem is case sensitive, set this to skip a
    number of checks which have a memory cost due to caching. These checks
    are required for case insensitive filesystems, and disabling them can
    cause data loss if your underlying filesystem is *not* in fact case
    sensitive.

Device options:

- :opt:`numConnections`
    Set to 1 to reduce the amount of overhead per device, as each connection
    has a memory and CPU cost.

Other things:

- ``GOMEMLIMIT`` and ``GOGC``
    These environment variables can be used to control the garbage
    collector. For small setups, setting ``GOMEMLIMIT`` to the desired max
    amount of memory Syncthing should use can make the garbage collector
    adhere more closely to the desired limit. Read more in the `Go GC guide
    <https://golang.org/doc/gc-guide>`__.

- ``GOMAXPROCS``
    This environment variable can be used to control the maximum number
    number concurrently running threads Syncthing uses. Setting it to 1 (or
    any number lower than your actual number of cores) will reduce the
    amount of CPU used by Syncthing at any given moment.
