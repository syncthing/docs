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

You have resources to spare and want to optimize for performance.
First some general options:

- :opt:`progressUpdateIntervalS`:
    Set to -1 to disable progress updates. Time spent on progress
    updates is time not spent on syncing.

- :opt:`setLowPriority`:
    Set to false to run Syncthing with normal priority. This may allow
    Syncthing to use more CPU time, which can improve performance.

- :opt:`maxFolderConcurrency`:
    Find an appropriate setting for the number of folders to sync
    concurrently, taking into account the inherent concurrency of the
    underlying storage system. If two folders are on the same underlying
    spinning disk, syncing them concurrently will be a strict loss of
    performance. If the underlying storage is a large array of disk,
    then syncing many folders concurrently may be beneficial.

- :opt:`databaseTuning`:
    Set to `large`. Regardless of the size of the actual database, this
    increases a number of buffers and settings to optimize for higher
    throughput.

- :opt:`maxConcurrentIncomingRequestKiB`:
    This sets the maximum amount of data being processed (loaded from
    disk and transmitted over the network) concurrently at any given
    time. This is a global limited, not per folder. If you have a lot of
    folders and/or a lot of devices and memory to spare you may want to
    increase this value. The default is 256 MiB, consider values of 1024
    MiB or more.

These options are folder specific and should be set on each folder:

- :opt:`fsWatcherEnabled`
- :opt:`copiers, hashers`
- :opt:`pullerMaxPendingKiB`
- :opt:`order`
- :opt:`scanProgressIntervalS`
- :opt:`weakHashThresholdPct`
- :opt:`maxConcurrentWrites`
- :opt:`disableFsync`
- :opt:`blockPullOrder`
- :opt:`copyRangeMethod`
- :opt:`caseSensitiveFS`
- :opt:`syncOwnership/syncXattrs, sendOwnership/sendXattrs`

Other things:

- ``GOMEMLIMIT`` and ``GOGC``: These environment variables can be used to
  control the garbage collector. For large setups, setting ``GOMEMLIMIT`` to
  the desired max amount of memory Syncthing should use can improve
  performance. The reason is that this reduces garbage collector frequency
  during lower memory usage.

Tuning for Low Resources
------------------------

You have limited resources and want Syncthing to use as few as possible. You
care less about performance.

- General Options:

    - :option:`progressUpdateIntervalS`: Set to -1 to disable progress updates. Progress updates aren't absolutely essential and consume some amount of CPU and memory.
    - :option:`maxFolderConcurrency`
    - :option:`databaseTuning`
    - :option:`maxConcurrentIncomingRequestKiB`

- Folders:

    - :option:`fsWatcherEnabled`
    - :option:`copiers, hashers`
    - :option:`pullerMaxPendingKiB`
    - :option:`order`
    - :option:`scanProgressIntervalS`
    - :option:`weakHashThresholdPct`
    - :option:`maxConcurrentWrites`
    - :option:`disableFsync`
    - :option:`blockPullOrder`
    - :option:`copyRangeMethod`
    - :option:`caseSensitiveFS`
    - :option:`syncOwnership/syncXattrs, sendOwnership/sendXattrs`
