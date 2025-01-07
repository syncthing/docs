.. default-domain:: stconf

Syncthing Configuration
=======================

Synopsis
--------

::

    $XDG_STATE_HOME/syncthing
    $HOME/.local/state/syncthing
    $HOME/Library/Application Support/Syncthing
    %LOCALAPPDATA%\Syncthing


.. _config-locations:

Description
-----------

.. versionchanged:: 1.27.0

    The default location of the configuration and database directory on
    Unix-like systems was changed to ``$XDG_STATE_HOME/syncthing`` or
    ``$HOME/.local/state/syncthing``. Previously the default config location
    was ``$XDG_CONFIG_HOME/syncthing`` or ``$HOME/.config/syncthing``. The
    database directory was previously ``$HOME/.config/syncthing`` or, if the
    environment variable was set, ``$XDG_DATA_HOME/syncthing``. Existing
    installations may still use these directories instead of the newer
    defaults.

.. versionadded:: 1.5.0

    Database and config can now be set separately. Previously the database was
    always located in the same directory as the config.

Syncthing uses a single directory to store configuration and crypto keys.
Syncthing also keeps an index database with file metadata which is by
default stored in the same directory, though this can be overridden.

The location defaults to ``$XDG_STATE_HOME/syncthing`` or
``$HOME/.local/state/syncthing`` (Unix-like), ``$HOME/Library/Application
Support/Syncthing`` (Mac), or ``%LOCALAPPDATA%\Syncthing`` (Windows). It can
be changed at runtime using the ``--config`` or ``--home`` flags or the
corresponding environment variables (``$STCONFDIR`` or ``STHOMEDIR``). The
following files are located in this directory:

:file:`config.xml`
    The configuration file, in XML format.

:file:`cert.pem`, :file:`key.pem`
    The device's ECDSA public and private key. These form the basis for the
    device ID. The key must be kept private.

:file:`https-cert.pem`, :file:`https-key.pem`
    The certificate and key for HTTPS GUI connections. These may be replaced
    with a custom certificate for HTTPS as desired.

The database is by default stored in the same directory as the config, but
the location may be overridden by the ``--data`` or ``--home`` flags or the
corresponding environment variables (``$STDATADIR`` or ``STHOMEDIR``).

The database directory contains the following files, among others:

:file:`index-{*}.db`
    A directory holding the database with metadata and hashes of the files
    currently on disk and available from peers.

:file:`syncthing.log`
    Log output, on some systems.

:file:`audit-{*}.log`
    Audit log data, when enabled.

:file:`panic-{*}.log`
    Crash log data, when required.


Config File Format
------------------

The following shows an example of a default configuration file (IDs will differ):


.. note::
   The config examples are present for illustration. Do **not** copy them
   entirely to use as your config. They are likely out-of-date and the values
   may no longer correspond to the defaults.


.. code-block:: xml

    <configuration version="37">
        <folder id="default" label="Default Folder" path="/Users/jb/Sync/" type="sendreceive" rescanIntervalS="3600" fsWatcherEnabled="true" fsWatcherDelayS="10" fsWatcherTimeoutS="0" ignorePerms="false" autoNormalize="true">
            <filesystemType>basic</filesystemType>
            <device id="S7UKX27-GI7ZTXS-GC6RKUA-7AJGZ44-C6NAYEB-HSKTJQK-KJHU2NO-CWV7EQW" introducedBy="">
                <encryptionPassword></encryptionPassword>
            </device>
            <minDiskFree unit="%">1</minDiskFree>
            <versioning>
                <cleanupIntervalS>3600</cleanupIntervalS>
                <fsPath></fsPath>
                <fsType>basic</fsType>
            </versioning>
            <copiers>0</copiers>
            <pullerMaxPendingKiB>0</pullerMaxPendingKiB>
            <hashers>0</hashers>
            <order>random</order>
            <ignoreDelete>false</ignoreDelete>
            <scanProgressIntervalS>0</scanProgressIntervalS>
            <pullerPauseS>0</pullerPauseS>
            <maxConflicts>-1</maxConflicts>
            <disableSparseFiles>false</disableSparseFiles>
            <disableTempIndexes>false</disableTempIndexes>
            <paused>false</paused>
            <weakHashThresholdPct>25</weakHashThresholdPct>
            <markerName>.stfolder</markerName>
            <copyOwnershipFromParent>false</copyOwnershipFromParent>
            <modTimeWindowS>0</modTimeWindowS>
            <maxConcurrentWrites>2</maxConcurrentWrites>
            <disableFsync>false</disableFsync>
            <blockPullOrder>standard</blockPullOrder>
            <copyRangeMethod>standard</copyRangeMethod>
            <caseSensitiveFS>false</caseSensitiveFS>
            <junctionsAsDirs>false</junctionsAsDirs>
            <syncOwnership>false</syncOwnership>
            <sendOwnership>false</sendOwnership>
            <syncXattrs>false</syncXattrs>
            <sendXattrs>false</sendXattrs>
        </folder>
        <device id="S7UKX27-GI7ZTXS-GC6RKUA-7AJGZ44-C6NAYEB-HSKTJQK-KJHU2NO-CWV7EQW" name="syno" compression="metadata" introducer="false" skipIntroductionRemovals="false" introducedBy="">
            <address>dynamic</address>
            <paused>false</paused>
            <autoAcceptFolders>false</autoAcceptFolders>
            <maxSendKbps>0</maxSendKbps>
            <maxRecvKbps>0</maxRecvKbps>
            <ignoredFolder time="2022-01-09T19:09:52Z" id="br63e-wyhb7" label="Foo"></ignoredFolder>
            <maxRequestKiB>0</maxRequestKiB>
            <untrusted>false</untrusted>
            <remoteGUIPort>0</remoteGUIPort>
        </device>
        <gui enabled="true" tls="false" debugging="false">
            <address>127.0.0.1:8384</address>
            <apikey>k1dnz1Dd0rzTBjjFFh7CXPnrF12C49B1</apikey>
            <theme>default</theme>
        </gui>
        <ldap></ldap>
        <options>
            <listenAddress>default</listenAddress>
            <globalAnnounceServer>default</globalAnnounceServer>
            <globalAnnounceEnabled>true</globalAnnounceEnabled>
            <localAnnounceEnabled>true</localAnnounceEnabled>
            <localAnnouncePort>21027</localAnnouncePort>
            <localAnnounceMCAddr>[ff12::8384]:21027</localAnnounceMCAddr>
            <maxSendKbps>0</maxSendKbps>
            <maxRecvKbps>0</maxRecvKbps>
            <reconnectionIntervalS>60</reconnectionIntervalS>
            <relaysEnabled>true</relaysEnabled>
            <relayReconnectIntervalM>10</relayReconnectIntervalM>
            <startBrowser>true</startBrowser>
            <natEnabled>true</natEnabled>
            <natLeaseMinutes>60</natLeaseMinutes>
            <natRenewalMinutes>30</natRenewalMinutes>
            <natTimeoutSeconds>10</natTimeoutSeconds>
            <urAccepted>0</urAccepted>
            <urSeen>0</urSeen>
            <urUniqueID></urUniqueID>
            <urURL>https://data.syncthing.net/newdata</urURL>
            <urPostInsecurely>false</urPostInsecurely>
            <urInitialDelayS>1800</urInitialDelayS>
            <autoUpgradeIntervalH>12</autoUpgradeIntervalH>
            <upgradeToPreReleases>false</upgradeToPreReleases>
            <keepTemporariesH>24</keepTemporariesH>
            <cacheIgnoredFiles>false</cacheIgnoredFiles>
            <progressUpdateIntervalS>5</progressUpdateIntervalS>
            <limitBandwidthInLan>false</limitBandwidthInLan>
            <minHomeDiskFree unit="%">1</minHomeDiskFree>
            <releasesURL>https://upgrades.syncthing.net/meta.json</releasesURL>
            <overwriteRemoteDeviceNamesOnConnect>false</overwriteRemoteDeviceNamesOnConnect>
            <tempIndexMinBlocks>10</tempIndexMinBlocks>
            <unackedNotificationID>authenticationUserAndPassword</unackedNotificationID>
            <trafficClass>0</trafficClass>
            <setLowPriority>true</setLowPriority>
            <maxFolderConcurrency>0</maxFolderConcurrency>
            <crashReportingURL>https://crash.syncthing.net/newcrash</crashReportingURL>
            <crashReportingEnabled>true</crashReportingEnabled>
            <stunKeepaliveStartS>180</stunKeepaliveStartS>
            <stunKeepaliveMinS>20</stunKeepaliveMinS>
            <stunServer>default</stunServer>
            <databaseTuning>auto</databaseTuning>
            <maxConcurrentIncomingRequestKiB>0</maxConcurrentIncomingRequestKiB>
            <announceLANAddresses>true</announceLANAddresses>
            <sendFullIndexOnUpgrade>false</sendFullIndexOnUpgrade>
            <connectionLimitEnough>0</connectionLimitEnough>
            <connectionLimitMax>0</connectionLimitMax>
            <insecureAllowOldTLSVersions>false</insecureAllowOldTLSVersions>
        </options>
        <remoteIgnoredDevice time="2022-01-09T20:02:01Z" id="5SYI2FS-LW6YAXI-JJDYETS-NDBBPIO-256MWBO-XDPXWVG-24QPUM4-PDW4UQU" name="bugger" address="192.168.0.20:22000"></remoteIgnoredDevice>
        <defaults>
            <folder id="" label="" path="~" type="sendreceive" rescanIntervalS="3600" fsWatcherEnabled="true" fsWatcherDelayS="10" fsWatcherTimeoutS="0" ignorePerms="false" autoNormalize="true">
                <filesystemType>basic</filesystemType>
                <device id="S7UKX27-GI7ZTXS-GC6RKUA-7AJGZ44-C6NAYEB-HSKTJQK-KJHU2NO-CWV7EQW" introducedBy="">
                    <encryptionPassword></encryptionPassword>
                </device>
                <minDiskFree unit="%">1</minDiskFree>
                <versioning>
                    <cleanupIntervalS>3600</cleanupIntervalS>
                    <fsPath></fsPath>
                    <fsType>basic</fsType>
                </versioning>
                <copiers>0</copiers>
                <pullerMaxPendingKiB>0</pullerMaxPendingKiB>
                <hashers>0</hashers>
                <order>random</order>
                <ignoreDelete>false</ignoreDelete>
                <scanProgressIntervalS>0</scanProgressIntervalS>
                <pullerPauseS>0</pullerPauseS>
                <maxConflicts>10</maxConflicts>
                <disableSparseFiles>false</disableSparseFiles>
                <disableTempIndexes>false</disableTempIndexes>
                <paused>false</paused>
                <weakHashThresholdPct>25</weakHashThresholdPct>
                <markerName>.stfolder</markerName>
                <copyOwnershipFromParent>false</copyOwnershipFromParent>
                <modTimeWindowS>0</modTimeWindowS>
                <maxConcurrentWrites>2</maxConcurrentWrites>
                <disableFsync>false</disableFsync>
                <blockPullOrder>standard</blockPullOrder>
                <copyRangeMethod>standard</copyRangeMethod>
                <caseSensitiveFS>false</caseSensitiveFS>
                <junctionsAsDirs>false</junctionsAsDirs>
                <syncOwnership>false</syncOwnership>
                <sendOwnership>false</sendOwnership>
                <syncXattrs>false</syncXattrs>
                <sendXattrs>false</sendXattrs>
            </folder>
            <device id="" compression="metadata" introducer="false" skipIntroductionRemovals="false" introducedBy="">
                <address>dynamic</address>
                <paused>false</paused>
                <autoAcceptFolders>false</autoAcceptFolders>
                <maxSendKbps>0</maxSendKbps>
                <maxRecvKbps>0</maxRecvKbps>
                <maxRequestKiB>0</maxRequestKiB>
                <untrusted>false</untrusted>
                <remoteGUIPort>0</remoteGUIPort>
            </device>
        </defaults>
    </configuration>


Configuration Element
---------------------

.. code-block:: xml

    <configuration version="37">
        <folder></folder>
        <device></device>
        <gui></gui>
        <ldap></ldap>
        <options></options>
        <remoteIgnoredDevice></remoteIgnoredDevice>
        <defaults></defaults>
    </configuration>

This is the root element. It has one attribute:

.. option:: configuration.version

    The config version. Increments whenever a change is made that requires
    migration from previous formats.

It contains the elements described in the following sections and any number of
this additional child element:

.. option:: configuration.remoteIgnoredDevice

    Contains the ID of the device that should be ignored. Connection attempts
    from this device are logged to the console but never displayed in the web
    GUI.


Folder Element
--------------

.. code-block:: xml

    <folder id="default" label="Default Folder" path="/Users/jb/Sync/" type="sendreceive" rescanIntervalS="3600" fsWatcherEnabled="true" fsWatcherDelayS="10" fsWatcherTimeoutS="0" ignorePerms="false" autoNormalize="true">
        <filesystemType>basic</filesystemType>
        <device id="S7UKX27-GI7ZTXS-GC6RKUA-7AJGZ44-C6NAYEB-HSKTJQK-KJHU2NO-CWV7EQW" introducedBy="">
            <encryptionPassword></encryptionPassword>
        </device>
        <minDiskFree unit="%">1</minDiskFree>
        <versioning>
            <cleanupIntervalS>3600</cleanupIntervalS>
            <fsPath></fsPath>
            <fsType>basic</fsType>
        </versioning>
        <copiers>0</copiers>
        <pullerMaxPendingKiB>0</pullerMaxPendingKiB>
        <hashers>0</hashers>
        <order>random</order>
        <ignoreDelete>false</ignoreDelete>
        <scanProgressIntervalS>0</scanProgressIntervalS>
        <pullerPauseS>0</pullerPauseS>
        <maxConflicts>-1</maxConflicts>
        <disableSparseFiles>false</disableSparseFiles>
        <disableTempIndexes>false</disableTempIndexes>
        <paused>false</paused>
        <weakHashThresholdPct>25</weakHashThresholdPct>
        <markerName>.stfolder</markerName>
        <copyOwnershipFromParent>false</copyOwnershipFromParent>
        <modTimeWindowS>0</modTimeWindowS>
        <maxConcurrentWrites>2</maxConcurrentWrites>
        <disableFsync>false</disableFsync>
        <blockPullOrder>standard</blockPullOrder>
        <copyRangeMethod>standard</copyRangeMethod>
        <caseSensitiveFS>false</caseSensitiveFS>
        <junctionsAsDirs>false</junctionsAsDirs>
        <syncOwnership>false</syncOwnership>
        <sendOwnership>false</sendOwnership>
        <syncXattrs>false</syncXattrs>
        <sendXattrs>false</sendXattrs>
    </folder>

One or more ``folder`` elements must be present in the file. Each element
describes one folder. The following attributes may be set on the ``folder``
element:

.. option:: folder.id
    :mandatory:

    The folder ID, which must be unique.

.. option:: folder.label

    The label of a folder is a human readable and descriptive local name. May
    be different on each device, empty, and/or identical to other folder
    labels. (optional)

.. option:: folder.filesystemType

    The internal file system implementation used to access this folder, detailed
    in a :doc:`separate chapter </advanced/folder-filesystem-type>`.

.. option:: folder.path
    :mandatory:

    The path to the directory where the folder is stored on this
    device; not sent to other devices.

.. option:: folder.type

    Controls how the folder is handled by Syncthing. Possible values are:

    ``sendreceive``
        The folder is in default mode. Sending local and accepting remote changes.
        Note that this type was previously called "readwrite" which is deprecated
        but still accepted in incoming configs.

    ``sendonly``
        The folder is in "send only" mode -- it will not be modified by
        Syncthing on this device.
        Note that this type was previously called "readonly" which is deprecated
        but still accepted in incoming configs.

    ``receiveonly``
        The folder is in "receive only" mode -- it will not propagate
        changes to other devices.

    ``receiveencrypted``
        Must be used on untrusted devices, where the data cannot be decrypted
        because no folder password was entered.  See :doc:`untrusted`.

.. option:: folder.rescanIntervalS

    The rescan interval, in seconds. Can be set to ``0`` to disable when external
    plugins are used to trigger rescans.

.. option:: folder.fsWatcherEnabled

    If set to ``true``, this detects changes to files in the folder and scans them.

.. option:: folder.fsWatcherDelayS

    The duration during which changes detected are accumulated, before a scan is
    scheduled (only takes effect if :opt:`fsWatcherEnabled` is set to ``true``).

.. option:: folder.fsWatcherTimeoutS

    The maximum delay before a scan is triggered when a file is continuously
    changing. If unset or zero a default value is calculated based on
    :opt:`fsWatcherDelayS`.

.. option:: folder.ignorePerms

    If ``true``, files originating from this folder will be announced to remote
    devices with the "no permission bits" flag.  The remote devices will use
    whatever their default permission setting is when creating the files.  The
    primary use case is for file systems that do not support permissions, such
    as FAT, or environments where changing permissions is impossible.

.. option:: folder.autoNormalize

    Automatically correct UTF-8 normalization errors found in file names.  The
    mechanism and how to set it up is described in a :doc:`separate chapter
    </advanced/folder-autonormalize>`.

The following child elements may exist:

.. option:: folder.device
    :aliases: folder.devices

    These must have the ``id`` attribute and can have an ``introducedBy``
    attribute, identifying the device that introduced us to share this folder
    with the given device.  If the original introducer unshares this folder with
    this device, our device will follow and unshare the folder (subject to
    :opt:`skipIntroductionRemovals` being ``false`` on the introducer device).

    All mentioned devices are those that will be sharing the folder in question.
    Each mentioned device must have a separate ``device`` element later in the file.
    It is customary that the local device ID is included in all folders.
    Syncthing will currently add this automatically if it is not present in
    the configuration file.

    The ``encryptionPassword`` sub-element contains the secret needed to decrypt
    this folder's data on the remote device.  If left empty, the data is plainly
    accessible (but still protected by the transport encryption).  The mechanism
    and how to set it up is described in a :doc:`separate chapter <untrusted>`.

.. option:: folder.minDiskFree

    The minimum required free space that should be available on the disk this
    folder resides.  The folder will be stopped when the value drops below the
    threshold.  The element content is interpreted according to the given
    ``unit`` attribute.  Accepted ``unit`` values are ``%`` (percent of the disk
    / volume size), ``kB``, ``MB``, ``GB`` and ``TB``.  Set to zero to disable.

.. option:: folder.versioning

    Specifies a versioning configuration.

    .. seealso::
        :doc:`versioning`

.. option:: folder.copiers
            folder.hashers

    The number of copier and hasher routines to use, or ``0`` for the
    system determined optimums. These are low-level performance options for
    advanced users only; do not change unless requested to or you've actually
    read and understood the code yourself. :)

.. option:: folder.pullerMaxPendingKiB

    Controls when we stop sending requests to other devices once we’ve got this
    much unserved requests.  The number of pullers is automatically adjusted
    based on this desired amount of outstanding request data.

.. option:: folder.order

    The order in which needed files should be pulled from the cluster.  It has
    no effect when the folder type is "send only".  The possibles values are:

    ``random`` (default)
        Pull files in random order. This optimizes for balancing resources among
        the devices in a cluster.

    ``alphabetic``
        Pull files ordered by file name alphabetically.

    ``smallestFirst``, ``largestFirst``
        Pull files ordered by file size; smallest and largest first respectively.

    ``oldestFirst``, ``newestFirst``
        Pull files ordered by modification time; oldest and newest first
        respectively.

    Note that the scanned files are sent in batches and the sorting is applied
    only to the already discovered files. This means the sync might start with
    a 1 GB file even if there is 1 KB file available on the source device until
    the 1 KB becomes known to the pulling device.

.. option:: folder.ignoreDelete

    .. warning::
        Enabling this is highly discouraged - use at your own risk. You have been warned.

    When set to ``true``, this device will pretend not to see instructions to
    delete files from other devices.  The mechanism is described in a
    :doc:`separate chapter </advanced/folder-ignoredelete>`.

.. option:: folder.scanProgressIntervalS

    The interval in seconds with which scan progress information is sent to the GUI. Setting to ``0``
    will cause Syncthing to use the default value of two.

.. option:: folder.pullerPauseS

    Tweak for rate limiting the puller when it retries pulling files. Don't
    change this unless you know what you're doing.

.. option:: folder.maxConflicts

    The maximum number of conflict copies to keep around for any given file.
    The default, ``-1``, means an unlimited number. Setting this to ``0`` disables
    conflict copies altogether.

.. option:: folder.disableSparseFiles

    By default, blocks containing all zeros are not written, causing files
    to be sparse on filesystems that support this feature. When set to ``true``,
    sparse files will not be created.

.. option:: folder.disableTempIndexes

    By default, devices exchange information about blocks available in
    transfers that are still in progress, which allows other devices to
    download parts of files that are not yet fully downloaded on your own
    device, essentially making transfers more torrent like. When set to
    ``true``, such information is not exchanged for this folder.

.. option:: folder.paused

    True if this folder is (temporarily) suspended.

.. option:: folder.weakHashThresholdPct

    Use weak hash if more than the given percentage of the file has changed. Set
    to ``-1`` to always use weak hash. Default is ``25``.

.. option:: folder.markerName

    Name of a directory or file in the folder root to be used as
    :ref:`marker-faq`. Default is ``.stfolder``.

.. option:: folder.copyOwnershipFromParent

    On Unix systems, tries to copy file/folder ownership from the parent directory (the directory it's located in).
    Requires running Syncthing as a privileged user, or granting it additional capabilities (e.g. CAP_CHOWN on Linux).

.. option:: folder.modTimeWindowS

    Allowed modification timestamp difference when comparing files for
    equivalence. To be used on file systems which have unstable
    modification timestamps that might change after being recorded
    during the last write operation. Default is ``2`` on Android when the
    folder is located on a FAT partition, and ``0`` otherwise.

.. option:: folder.maxConcurrentWrites

    Maximum number of concurrent write operations while syncing. Increasing this might increase or
    decrease disk performance, depending on the underlying storage. Default is ``2``.

.. option:: folder.disableFsync

    .. warning::
        This is a known insecure option - use at your own risk.

    Disables committing file operations to disk before recording them in the
    database.  Disabling fsync can lead to data corruption.  The mechanism is
    described in a :doc:`separate chapter </advanced/folder-disable-fsync>`.

.. option:: folder.blockPullOrder

    Order in which the blocks of a file are downloaded. This option controls how quickly different parts of the
    file spread between the connected devices, at the cost of causing strain on the storage.

    Available options:

    ``standard`` (default)
        The blocks of a file are split into N equal continuous sequences, where N is the number of connected
        devices. Each device starts downloading its own sequence, after which it picks other devices
        sequences at random. Provides acceptable data distribution and minimal spinning disk strain.

    ``random``
        The blocks of a file are downloaded in a random order. Provides great data distribution, but very taxing on
        spinning disk drives.

    ``inOrder``
        The blocks of a file are downloaded sequentially, from start to finish. Spinning disk drive friendly, but provides
        no improvements to data distribution.

.. option:: folder.copyRangeMethod

    Provides a choice of method for copying data between files.  This can be
    used to optimise copies on network filesystems, improve speed of large
    copies or clone the data using copy-on-write functionality if the underlying
    filesystem supports it.  The mechanism is described in a :doc:`separate
    chapter </advanced/folder-copyrangemethod>`.

.. option:: folder.caseSensitiveFS

    Affects performance by disabling the extra safety checks for case
    insensitive filesystems.  The mechanism and how to set it up is described in
    a :doc:`separate chapter </advanced/folder-caseSensitiveFS>`.

.. option:: folder.junctionsAsDirs

    NTFS directory junctions are treated as ordinary directories, if this is set
    to ``true``.

.. option:: folder.syncOwnership

    File and directory ownership is synced when this is set to ``true``. See
    :doc:`/advanced/folder-sync-ownership` for more information.

.. option:: folder.sendOwnership

    File and directory ownership information is scanned when this is set to
    ``true``. See :doc:`/advanced/folder-send-ownership` for more information.

.. option:: folder.syncXattrs

    File and directory extended attributes are synced when this is set to
    ``true``. See :doc:`/advanced/folder-sync-xattrs` for more information.

.. option:: folder.sendXattrs

    File and directory extended attributes are scanned and sent to other
    devices when this is set to ``true``. See
    :doc:`/advanced/folder-send-xattrs` for more information.


Device Element
--------------

.. code-block:: xml

    <device id="S7UKX27-GI7ZTXS-GC6RKUA-7AJGZ44-C6NAYEB-HSKTJQK-KJHU2NO-CWV7EQW" name="syno" compression="metadata" introducer="false" skipIntroductionRemovals="false" introducedBy="2CYF2WQ-AKZO2QZ-JAKWLYD-AGHMQUM-BGXUOIS-GYILW34-HJG3DUK-LRRYQAR">
        <address>dynamic</address>
        <paused>false</paused>
        <autoAcceptFolders>false</autoAcceptFolders>
        <maxSendKbps>0</maxSendKbps>
        <maxRecvKbps>0</maxRecvKbps>
        <ignoredFolder time="2022-01-09T19:09:52Z" id="br63e-wyhb7" label="Foo"></ignoredFolder>
        <maxRequestKiB>0</maxRequestKiB>
        <untrusted>false</untrusted>
        <remoteGUIPort>0</remoteGUIPort>
        <numConnections>0</numConnections>
    </device>
    <device id="2CYF2WQ-AKZO2QZ-JAKWLYD-AGHMQUM-BGXUOIS-GYILW34-HJG3DUK-LRRYQAR" name="syno local" compression="metadata" introducer="true" skipIntroductionRemovals="false" introducedBy="">
        <address>tcp://192.0.2.1:22001</address>
        <paused>true</paused>
        <allowedNetwork>192.168.0.0/16</allowedNetwork>
        <autoAcceptFolders>false</autoAcceptFolders>
        <maxSendKbps>100</maxSendKbps>
        <maxRecvKbps>100</maxRecvKbps>
        <maxRequestKiB>65536</maxRequestKiB>
        <untrusted>false</untrusted>
        <remoteGUIPort>8384</remoteGUIPort>
        <numConnections>0</numConnections>
    </device>

One or more ``device`` elements must be present in the file. Each element
describes a device participating in the cluster. It is customary to include a
``device`` element for the local device; Syncthing will currently add one if
it is not present. The following attributes may be set on the ``device``
element:

.. option:: device.id
    :mandatory:
    :aliases: device.deviceID

    The :ref:`device ID <device-ids>`.

.. option:: device.name

    A friendly name for the device. (optional)

.. option:: device.compression

    Whether to use protocol compression when sending messages to this device.
    The possible values are:

    ``metadata``
        Compress metadata packets, such as index information. Metadata is
        usually very compression friendly so this is a good default.

    ``always``
        Compress all packets, including file data. This is recommended if the
        folders contents are mainly compressible data such as documents or
        text files.

    ``never``
        Disable all compression.

.. option:: device.introducer

    Set to true if this device should be trusted as an introducer, i.e. we
    should copy their list of devices per folder when connecting.

    .. seealso::
        :doc:`introducer`

.. option:: device.skipIntroductionRemovals

    Set to true if you wish to follow only introductions and not de-introductions.
    For example, if this is set, we would not remove a device that we were introduced
    to even if the original introducer is no longer listing the remote device as known.

.. option:: device.introducedBy

    Defines which device has introduced us to this device. Used only for following de-introductions.

.. option:: device.certName

    The device certificate's common name, if it is not the default "syncthing".

From the following child elements at least one ``address`` child must exist.

.. option:: device.address
    :mandatory: At least one must be present.
    :aliases: device.addresses

    Contains an address or host name to use when attempting to connect to this device.
    Entries other than ``dynamic`` need a protocol specific prefix. For the TCP protocol
    the prefixes ``tcp://`` (dual-stack), ``tcp4://`` (IPv4 only) or ``tcp6://`` (IPv6 only) can be used.
    The prefixes for the QUIC protocol are analogous: ``quic://``, ``quic4://`` and ``quic6://``
    Note that IP addresses need not use IPv4 or IPv6 prefixes; these are optional. Accepted formats are:

    IPv4 address (``tcp://192.0.2.42``)
        The default port (22000) is used.

    IPv4 address and port (``tcp://192.0.2.42:12345``)
        The address and port is used as given.

    IPv6 address (``tcp://[2001:db8::23:42]``)
        The default port (22000) is used. The address must be enclosed in
        square brackets.

    IPv6 address and port (``tcp://[2001:db8::23:42]:12345``)
        The address and port is used as given. The address must be enclosed in
        square brackets.

    Host name (``tcp6://fileserver``)
        The host name will be used on the default port (22000) and connections
        will be attempted only via IPv6.

    Host name and port (``tcp://fileserver:12345``)
        The host name will be used on the given port and connections will be
        attempted via both IPv4 and IPv6, depending on name resolution.

    ``dynamic``
        The word ``dynamic`` (without any prefix) means to use local and
        global discovery to find the device.

    You can set multiple addresses *and* combine it with the ``dynamic`` keyword
    for example:

    .. code-block:: xml

        <device id="...">
            <address>tcp://192.0.2.1:22001</address>
            <address>quic://192.0.1.254:22000</address>
            <address>dynamic</address>
        </device>

    In the GUI, multiple values are separated by commas.

.. option:: device.paused

    True if synchronization with this devices is (temporarily) suspended.

.. option:: device.allowedNetwork
    :aliases: device.allowedNetworks

    If given, this restricts connections to this device to only this network.
    The mechanism is described in detail in a :doc:`separate chapter
    </advanced/device-allowednetworks>`).  To configure multiple networks, you
    can either: repeat ``<allowedNetwork>`` tags in the configuration file or
    enter several networks separated by commas in the GUI.

.. option:: device.autoAcceptFolders

    If ``true``, folders shared from this remote device are automatically added
    and synced locally under the :opt:`default path <defaults.folder>`.  For the
    folder name, Syncthing tries to use the label from the remote device, and if
    the same label already exists, it then tries to use the folder's ID.  If
    that exists as well, the folder is just offered to accept manually.  A local
    folder already added with the same ID will just be shared rather than
    created separately.

.. option:: device.maxSendKbps

    Maximum send rate to use for this device. Unit is kibibytes/second, despite
    the config name looking like kilobits/second.

.. option:: device.maxRecvKbps

    Maximum receive rate to use for this device. Unit is kibibytes/second,
    despite the config name looking like kilobits/second.

.. option:: device.ignoredFolder
    :aliases: device.ignoredFolders

    Contains the ID of the folder that should be ignored. This folder will
    always be skipped when advertised from the containing remote device,
    i.e. this will be logged, but there will be no dialog shown in the web GUI.
    Multiple ignored folders are represented by repeated ``<ignoredFolder>``
    tags in the configuration file.

.. option:: device.maxRequestKiB

    Maximum amount of data to have outstanding in requests towards this device.
    Unit is kibibytes.

.. option:: device.remoteGUIPort

    If set to a positive integer, the GUI will display an HTTP link to the IP
    address which is currently used for synchronization.  Only the TCP port is
    exchanged for the value specified here.  Note that any port forwarding or
    firewall settings need to be done manually and the link will probably not
    work for link-local IPv6 addresses because of modern browser limitations.

.. option:: device.untrusted

    This boolean value marks a particular device as untrusted, which disallows
    ever sharing any unencrypted data with it.  Every folder shared with that
    device then needs an encryption password set, or must already be of the
    "receive encrypted" type locally.  Refer to the detailed explanation under
    :doc:`untrusted`.

.. option:: device.numConnections

    The number of connections to this device. See
    :doc:`/advanced/device-numconnections` for more information.


GUI Element
-----------

.. code-block:: xml

    <gui enabled="true" tls="false" debugging="false">
        <address>127.0.0.1:8384</address>
        <apikey>k1dnz1Dd0rzTBjjFFh7CXPnrF12C49B1</apikey>
        <theme>default</theme>
    </gui>


There must be exactly one ``gui`` element. The GUI configuration is also used by
the :doc:`/dev/rest` and the :doc:`/dev/events`. The following attributes may be
set on the ``gui`` element:

.. option:: gui.enabled

    If not ``true``, the GUI and API will not be started.

.. option:: gui.tls
    :aliases: gui.useTLS

    If set to ``true``, TLS (HTTPS) will be enforced. Non-HTTPS requests will
    be redirected to HTTPS. When set to ``false``, TLS connections are
    still possible but not required.

.. option:: gui.debugging

    This enables :doc:`/users/profiling` and additional endpoints in the REST
    API, see :doc:`/rest/debug`.

The following child elements may be present:

.. option:: gui.address
    :mandatory: Exactly one element must be present.

    Set the listen address.  Allowed address formats are:

    IPv4 address and port (``127.0.0.1:8384``)
        The address and port are used as given.

    IPv6 address and port (``[::1]:8384``)
        The address and port are used as given. The address must be enclosed in
        square brackets.

    Wildcard and port (``0.0.0.0:12345``, ``[::]:12345``, ``:12345``)
        These are equivalent and will result in Syncthing listening on all
        interfaces via both IPv4 and IPv6.

    UNIX socket location (``/var/run/st.sock``)
        If the address is an absolute path it is interpreted as the path to a UNIX socket.

.. option:: gui.unixSocketPermissions

    When ``address`` is set to a UNIX socket location, set this to an octal value
    to override the default permissions of the socket.

.. option:: gui.user

    Set to require authentication.

.. option:: gui.password

    Contains the bcrypt hash of the real password.

.. option:: gui.apikey

    If set, this is the API key that enables usage of the REST interface.

.. option:: gui.insecureAdminAccess

    If true, this allows access to the web GUI from outside (i.e. not localhost)
    without authorization. A warning will displayed about this setting on startup.

.. option:: gui.insecureSkipHostcheck

    When the GUI / API is bound to localhost, we enforce that the ``Host``
    header looks like localhost.  This option bypasses that check.

.. option:: gui.insecureAllowFrameLoading

    Allow rendering the GUI within an ``<iframe>``, ``<frame>`` or ``<object>``
    by not setting the ``X-Frame-Options: SAMEORIGIN`` HTTP header.  This may be
    needed for serving the Syncthing GUI as part of a website through a proxy.

.. option:: gui.theme

    The name of the theme to use.

.. option:: gui.authMode

    Authentication mode to use. If not present, the authentication mode (static)
    is controlled by the presence of user/password fields for backward compatibility.

    ``static``
        Authentication using user and password.

    ``ldap``
        LDAP authentication. Requires ldap top level config section to be present.

.. option:: gui.sendBasicAuthPrompt

    .. versionadded:: 1.26.0

    Prior to version 1.26.0 the GUI used HTTP Basic Authorization for login, but
    starting in version 1.26.0 it uses an HTML form by default. Basic
    Authorization is still supported when the ``Authorization`` request header
    is present in a request, but some browsers don't send the header unless
    prompted by a 401 response.

    When this setting is enabled, the GUI will respond to unauthenticated
    requests with a 401 response prompting for Basic Authorization, so that
    ``https://user:pass@localhost`` style URLs continue to work in standard
    browsers. Other clients that always send the ``Authorization`` request
    header do not need this setting.

    When this setting is disabled, the GUI will not send 401 responses so users
    won't see browser popups prompting for username and password.


LDAP Element
------------

.. code-block:: xml

    <ldap>
        <address>localhost:389</address>
        <bindDN>cn=%s,ou=users,dc=syncthing,dc=net</bindDN>
        <transport>nontls</transport>
        <insecureSkipVerify>false</insecureSkipVerify>
    </ldap>

The ``ldap`` element contains LDAP configuration options.  The mechanism is
described in detail under :doc:`ldap`.

.. option:: ldap.address
   :mandatory:

    LDAP server address (server:port).

.. option:: ldap.bindDN
   :mandatory:

    BindDN for user authentication.
    Special ``%s`` variable should be used to pass username to LDAP.

.. option:: ldap.transport

    ``nontls``
        Non secure connection.

    ``tls``
        TLS secured connection.

    ``starttls``
        StartTLS connection mode.

.. option:: ldap.insecureSkipVerify

    Skip verification (``true`` or ``false``).

.. option:: ldap.searchBaseDN

    Base DN for user searches.

.. option:: ldap.searchFilter

    Search filter for user searches.


Options Element
---------------

.. code-block:: xml

    <options>
        <listenAddress>default</listenAddress>
        <globalAnnounceServer>default</globalAnnounceServer>
        <globalAnnounceEnabled>true</globalAnnounceEnabled>
        <localAnnounceEnabled>true</localAnnounceEnabled>
        <localAnnouncePort>21027</localAnnouncePort>
        <localAnnounceMCAddr>[ff12::8384]:21027</localAnnounceMCAddr>
        <maxSendKbps>0</maxSendKbps>
        <maxRecvKbps>0</maxRecvKbps>
        <reconnectionIntervalS>60</reconnectionIntervalS>
        <relaysEnabled>true</relaysEnabled>
        <relayReconnectIntervalM>10</relayReconnectIntervalM>
        <startBrowser>true</startBrowser>
        <natEnabled>true</natEnabled>
        <natLeaseMinutes>60</natLeaseMinutes>
        <natRenewalMinutes>30</natRenewalMinutes>
        <natTimeoutSeconds>10</natTimeoutSeconds>
        <urAccepted>0</urAccepted>
        <urSeen>0</urSeen>
        <urUniqueID></urUniqueID>
        <urURL>https://data.syncthing.net/newdata</urURL>
        <urPostInsecurely>false</urPostInsecurely>
        <urInitialDelayS>1800</urInitialDelayS>
        <autoUpgradeIntervalH>12</autoUpgradeIntervalH>
        <upgradeToPreReleases>false</upgradeToPreReleases>
        <keepTemporariesH>24</keepTemporariesH>
        <cacheIgnoredFiles>false</cacheIgnoredFiles>
        <progressUpdateIntervalS>5</progressUpdateIntervalS>
        <limitBandwidthInLan>false</limitBandwidthInLan>
        <minHomeDiskFree unit="%">1</minHomeDiskFree>
        <releasesURL>https://upgrades.syncthing.net/meta.json</releasesURL>
        <overwriteRemoteDeviceNamesOnConnect>false</overwriteRemoteDeviceNamesOnConnect>
        <tempIndexMinBlocks>10</tempIndexMinBlocks>
        <unackedNotificationID>authenticationUserAndPassword</unackedNotificationID>
        <trafficClass>0</trafficClass>
        <setLowPriority>true</setLowPriority>
        <maxFolderConcurrency>0</maxFolderConcurrency>
        <crashReportingURL>https://crash.syncthing.net/newcrash</crashReportingURL>
        <crashReportingEnabled>true</crashReportingEnabled>
        <stunKeepaliveStartS>180</stunKeepaliveStartS>
        <stunKeepaliveMinS>20</stunKeepaliveMinS>
        <stunServer>default</stunServer>
        <databaseTuning>auto</databaseTuning>
        <maxConcurrentIncomingRequestKiB>0</maxConcurrentIncomingRequestKiB>
        <announceLANAddresses>true</announceLANAddresses>
        <sendFullIndexOnUpgrade>false</sendFullIndexOnUpgrade>
        <connectionLimitEnough>0</connectionLimitEnough>
        <connectionLimitMax>0</connectionLimitMax>
        <insecureAllowOldTLSVersions>false</insecureAllowOldTLSVersions>
    </options>

The ``options`` element contains all other global configuration options.

.. option:: options.listenAddress
    :aliases: options.listenAddresses

    The listen address for incoming sync connections. See
    :ref:`listen-addresses` for the allowed syntax.  To configure multiple
    addresses, you can either: repeat ``<listenAddress>`` tags in the
    configuration file or enter several addresses separated by commas in the
    GUI.

.. option:: options.globalAnnounceServer
    :aliases: options.globalAnnounceServers

    A URI to a global announce (discovery) server, or the word ``default`` to
    include the default servers. Any number of globalAnnounceServer elements
    may be present. The syntax for non-default entries is that of an HTTP or
    HTTPS URL. A number of options may be added as query options to the URL:
    ``insecure`` to prevent certificate validation (required for HTTP URLs)
    and ``id=<device ID>`` to perform certificate pinning. The device ID to
    use is printed by the discovery server on startup.  To configure multiple
    servers, you can either: repeat ``<globalAnnounceServer>`` tags in the
    configuration file or enter several servers separated by commas in the
    GUI.

.. option:: options.globalAnnounceEnabled

    Whether to announce this device to the global announce (discovery) server,
    and also use it to look up other devices.

.. option:: options.localAnnounceEnabled

    Whether to send announcements to the local LAN, also use such
    announcements to find other devices.

.. option:: options.localAnnouncePort

    The port on which to listen and send IPv4 broadcast announcements to.

.. option:: options.localAnnounceMCAddr

    The group address and port to join and send IPv6 multicast announcements on.

.. option:: options.maxSendKbps

    Outgoing data rate limit, in kibibytes per second.

.. option:: options.maxRecvKbps

    Incoming data rate limits, in kibibytes per second.

.. option:: options.reconnectionIntervalS

    The number of seconds to wait between each attempt to connect to currently
    unconnected devices.

.. option:: options.relaysEnabled

    When ``true``, relays will be connected to and potentially used for device to device connections.

.. option:: options.relayReconnectIntervalM

    Sets the interval, in minutes, between relay reconnect attempts.

.. option:: options.startBrowser

    Whether to attempt to start a browser to show the GUI when Syncthing starts.

.. option:: options.natEnabled

    Whether to attempt to perform a UPnP and NAT-PMP port mapping for
    incoming sync connections.

.. option:: options.natLeaseMinutes

    Request a lease for this many minutes; zero to request a permanent lease.

.. option:: options.natRenewalMinutes

    Attempt to renew the lease after this many minutes.

.. option:: options.natTimeoutSeconds

    When scanning for UPnP devices, wait this long for responses.

.. option:: options.urAccepted

    Whether the user has accepted to submit anonymous usage data. The default,
    ``0``, mean the user has not made a choice, and Syncthing will ask at some
    point in the future. ``-1`` means no, a number above zero means that that
    version of usage reporting has been accepted.

.. option:: options.urSeen

    The highest usage reporting version that has already been shown in the web GUI.

.. option:: options.urUniqueID

    The unique ID sent together with the usage report. Generated when usage
    reporting is enabled.

.. option:: options.urURL

    The URL to post usage report data to, when enabled.

.. option:: options.urPostInsecurely

    When true, the UR URL can be http instead of https, or have a self-signed
    certificate. The default is ``false``.

.. option:: options.urInitialDelayS

    The time to wait from startup for the first usage report to be sent. Allows
    the system to stabilize before reporting statistics.

.. option:: options.autoUpgradeIntervalH

    Check for a newer version after this many hours. Set to ``0`` to disable
    automatic upgrades.

.. option:: options.upgradeToPreReleases

    If ``true``, automatic upgrades include release candidates (see
    :ref:`releases`).

.. option:: options.keepTemporariesH

    Keep temporary failed transfers for this many hours. While the temporaries
    are kept, the data they contain need not be transferred again.

.. option:: options.cacheIgnoredFiles

    Whether to cache the results of ignore pattern evaluation. Performance
    at the price of memory. Defaults to ``false`` as the cost for evaluating
    ignores is usually not significant.

.. option:: options.progressUpdateIntervalS

    How often in seconds the progress of ongoing downloads is made available to
    the GUI. Set to ``-1`` to disable. Note that when disabled, the detailed
    sync progress for Out of Sync Items which shows how much of each file has
    been reused, copied, and downloaded will not work.

.. option:: options.limitBandwidthInLan

    Whether to apply bandwidth limits to devices in the same broadcast domain
    as the local device.

.. option:: options.minHomeDiskFree

    The minimum required free space that should be available on the partition
    holding the configuration and index.  The element content is interpreted
    according to the given ``unit`` attribute.  Accepted ``unit`` values are
    ``%`` (percent of the disk / volume size), ``kB``, ``MB``, ``GB`` and
    ``TB``.  Set to zero to disable.

.. option:: options.releasesURL

    The URL from which release information is loaded, for automatic upgrades.

.. option:: options.alwaysLocalNet
    :aliases: options.alwaysLocalNets

    Network that should be considered as local given in CIDR notation.  To
    configure multiple networks, you can either: repeat ``<alwaysLocalNet>``
    tags in the configuration file or enter several networks separated by
    commas in the GUI.

.. option:: options.overwriteRemoteDeviceNamesOnConnect

    If set, device names will always be overwritten with the name given by
    remote on each connection. By default, the name that the remote device
    announces will only be adopted when a name has not already been set.

.. option:: options.tempIndexMinBlocks

    When exchanging index information for incomplete transfers, only take
    into account files that have at least this many blocks.

.. option:: options.unackedNotificationID
    :aliases: options.unackedNotificationIDs

    ID of a notification to be displayed in the web GUI. Will be removed once
    the user acknowledged it (e.g. a transition notice on an upgrade).  Multiple
    IDs are represented by repeated ``<unackedNotificationID>`` tags in the
    configuration file.

.. option:: options.trafficClass

    Specify an IPv4 type of service (TOS)/IPv6 traffic class for outgoing
    packets. To specify a differentiated services code point (DSCP) the value
    must be bit shifted to the left by two to take the two least significant
    ECN bits into account.

.. option:: options.stunServer
    :aliases: options.stunServers

    Server to be used for STUN, given as ip:port. The keyword ``default`` gets
    expanded to
    ``stun.callwithus.com:3478``, ``stun.counterpath.com:3478``,
    ``stun.counterpath.net:3478``, ``stun.ekiga.net:3478``,
    ``stun.hitv.com:3478``, ``stun.ideasip.com:3478``,
    ``stun.internetcalls.com:3478``, ``stun.miwifi.com:3478``,
    ``stun.schlund.de:3478``,``stun.sipgate.net:10000``,
    ``stun.sipgate.net:3478``, ``stun.voip.aebc.com:3478``,
    ``stun.voiparound.com:3478``, ``stun.voipbuster.com:3478``,
    ``stun.voipstunt.com:3478`` and ``stun.xten.com:3478`` (this is the default).

    To configure multiple servers, you can either: repeat ``<stunServer>`` tags
    in the configuration file or enter several servers separated by commas in
    the GUI.

.. option:: options.stunKeepaliveStartS

    Interval in seconds between contacting a STUN server to maintain NAT
    mapping. Default is ``24`` and you can set it to ``0`` to disable contacting
    STUN servers.  The interval is automatically reduced if needed, down to a
    minimum of :opt:`stunKeepaliveMinS`.

.. option:: options.stunKeepaliveMinS

    Minimum for the :opt:`stunKeepaliveStartS` interval, in seconds.

.. option:: options.setLowPriority

    Syncthing will attempt to lower its process priority at startup.
    Specifically: on Linux, set itself to a separate process group, set the
    niceness level of that process group to nine and the I/O priority to
    best effort level five; on other Unixes, set the process niceness level
    to nine; on Windows, set the process priority class to below normal. To
    disable this behavior, for example to control process priority yourself
    as part of launching Syncthing, set this option to ``false``.

.. option:: options.maxFolderConcurrency

    This option controls how many folders may concurrently be in I/O-intensive
    operations such as syncing or scanning.  The mechanism is described in
    detail in a :doc:`separate chapter </advanced/option-max-concurrency>`.

.. option:: options.crashReportingURL
    :aliases: options.crURL

    Server URL where :doc:`automatic crash reports <crashrep>` will be sent if
    enabled.

.. option:: options.crashReportingEnabled

    Switch to opt out from the :doc:`automatic crash reporting <crashrep>`
    feature. Set ``false`` to keep Syncthing from sending panic logs on serious
    troubles.  Defaults to ``true``, to help the developers troubleshoot.

.. option:: options.databaseTuning

    Controls how Syncthing uses the backend key-value database that stores the
    index data and other persistent data it needs.  The available options and
    implications are explained in a :doc:`separate chapter
    </advanced/option-database-tuning>`.

.. option:: options.maxConcurrentIncomingRequestKiB

    This limits how many bytes we have "in the air" in the form of response data
    being read and processed.

.. option:: options.announceLANAddresses

    Enable (the default) or disable announcing private (RFC1918) LAN IP
    addresses to global discovery.

.. option:: options.sendFullIndexOnUpgrade

    Controls whether all index data is resent when an upgrade has happened,
    equivalent to starting Syncthing with :option:`--reset-deltas`.  This used
    to be the default behavior in older versions, but is mainly useful as a
    troubleshooting step and causes high database churn. The default is now
    ``false``.

.. option:: options.featureFlag
    :aliases: options.featureFlags

    Feature flags are simple strings that, when added to the configuration, may
    unleash unfinished or still-in-development features to allow early user
    testing.  Any supported value will be separately announced with the feature,
    so that regular users do not enable it by accident.  To configure multiple
    flags, you can either: repeat ``<featureFlag>`` tags in the configuration
    file or enter several flags separated by commas in the GUI.

.. option:: options.connectionLimitEnough

    The number of connections at which we stop trying to connect to more
    devices, zero meaning no limit.  Does not affect incoming connections.  The
    mechanism is described in detail in a :doc:`separate chapter
    </advanced/option-connection-limits>`.

.. option:: options.connectionLimitMax

    The maximum number of connections which we will allow in total, zero meaning
    no limit.  Affects incoming connections and prevents attempting outgoing
    connections.  The mechanism is described in detail in a :doc:`separate
    chapter </advanced/option-connection-limits>`.

.. option:: options.insecureAllowOldTLSVersions

    Only for compatibility with old versions of Syncthing on remote devices, as
    detailed in :doc:`/advanced/option-insecure-allow-old-tls-versions`.


Defaults Element
----------------

.. code-block:: xml

    <defaults>
        <folder id="" label="" path="~" type="sendreceive" rescanIntervalS="3600" fsWatcherEnabled="true" fsWatcherDelayS="10" fsWatcherTimeoutS="0" ignorePerms="false" autoNormalize="true">
            <filesystemType>basic</filesystemType>
            <device id="S7UKX27-GI7ZTXS-GC6RKUA-7AJGZ44-C6NAYEB-HSKTJQK-KJHU2NO-CWV7EQW" introducedBy="">
                <encryptionPassword></encryptionPassword>
            </device>
            <minDiskFree unit="%">1</minDiskFree>
            <versioning>
                <cleanupIntervalS>3600</cleanupIntervalS>
                <fsPath></fsPath>
                <fsType>basic</fsType>
            </versioning>
            <copiers>0</copiers>
            <pullerMaxPendingKiB>0</pullerMaxPendingKiB>
            <hashers>0</hashers>
            <order>random</order>
            <ignoreDelete>false</ignoreDelete>
            <scanProgressIntervalS>0</scanProgressIntervalS>
            <pullerPauseS>0</pullerPauseS>
            <maxConflicts>10</maxConflicts>
            <disableSparseFiles>false</disableSparseFiles>
            <disableTempIndexes>false</disableTempIndexes>
            <paused>false</paused>
            <weakHashThresholdPct>25</weakHashThresholdPct>
            <markerName>.stfolder</markerName>
            <copyOwnershipFromParent>false</copyOwnershipFromParent>
            <modTimeWindowS>0</modTimeWindowS>
            <maxConcurrentWrites>2</maxConcurrentWrites>
            <disableFsync>false</disableFsync>
            <blockPullOrder>standard</blockPullOrder>
            <copyRangeMethod>standard</copyRangeMethod>
            <caseSensitiveFS>false</caseSensitiveFS>
            <junctionsAsDirs>false</junctionsAsDirs>
        </folder>
        <device id="" compression="metadata" introducer="false" skipIntroductionRemovals="false" introducedBy="">
            <address>dynamic</address>
            <paused>false</paused>
            <autoAcceptFolders>false</autoAcceptFolders>
            <maxSendKbps>0</maxSendKbps>
            <maxRecvKbps>0</maxRecvKbps>
            <maxRequestKiB>0</maxRequestKiB>
            <untrusted>false</untrusted>
            <remoteGUIPort>0</remoteGUIPort>
            <numConnections>0</numConnections>
        </device>
        <ignores>
            <line>!foo2</line>
            <line>// comment</line>
            <line>(?d).DS_Store</line>
            <line>*2</line>
            <line>qu*</line>
        </ignores>
    </defaults>

The ``defaults`` element describes a template for newly added device and folder
options.  These will be used when adding a new remote device or folder, either
through the GUI or the command line interface.  The following child elements can
be present in the ``defaults`` element:

.. option:: defaults.device

    Template for a ``device`` element, with the same internal structure.  Any
    fields here will be used for a newly added remote device.  The ``id``
    attribute is meaningless in this context.

.. option:: defaults.folder

    Template for a ``folder`` element, with the same internal structure.  Any
    fields here will be used for a newly added shared folder.  The ``id``
    attribute is meaningless in this context.

    The UI will propose to create new folders at the path given in the ``path``
    attribute (used to be ``defaultFolderPath`` under ``options``).  It also
    applies to folders automatically accepted from a remote device.

    Even sharing with other remote devices can be done in the template by
    including the appropriate :opt:`folder.device` element underneath.

.. option:: defaults.ignores
    :aliases: defaults.ignores.lines

    .. versionadded:: 1.19.0

    Template for the :ref:`ignore patterns <ignoring-files>` applied to new
    folders.  These are copied to the ``.stignore`` file when a folder is
    automatically accepted from a remote device.  The GUI uses them to pre-fill
    the respective field when adding a new folder as well.  In XML, each pattern
    line is represented as by a ``<line>`` element.


.. _listen-addresses:

Listen Addresses
^^^^^^^^^^^^^^^^

The following address types are accepted in sync protocol listen addresses.
If you want Syncthing to listen on multiple addresses, you can either: repeat
``<listenAddress>`` tags in the configuration file or enter several addresses
separated by commas in the GUI.

Default listen addresses (``default``)
    This is equivalent to ``tcp://0.0.0.0:22000``, ``quic://0.0.0.0:22000``
    and ``dynamic+https://relays.syncthing.net/endpoint``.

TCP wildcard and port (``tcp://0.0.0.0:22000``, ``tcp://:22000``)
    These are equivalent and will result in Syncthing listening on all
    interfaces, IPv4 and IPv6, on the specified port.

TCP IPv4 wildcard and port (``tcp4://0.0.0.0:22000``, ``tcp4://:22000``)
    These are equivalent and will result in Syncthing listening on all
    interfaces via IPv4 only.

TCP IPv4 address and port (``tcp4://192.0.2.1:22000``)
    This results in Syncthing listening on the specified address and port, IPv4
    only.

TCP IPv6 wildcard and port (``tcp6://[::]:22000``, ``tcp6://:22000``)
    These are equivalent and will result in Syncthing listening on all
    interfaces via IPv6 only.

TCP IPv6 address and port (``tcp6://[2001:db8::42]:22000``)
    This results in Syncthing listening on the specified address and port, IPv6
    only.

QUIC address and port (e.g. ``quic://0.0.0.0:22000``)
    Syntax is the same as for TCP, also ``quic4`` and ``quic6`` can be used.

Static relay address (``relay://192.0.2.42:22067?id=abcd123...``)
    Syncthing will connect to and listen for incoming connections via the
    specified relay address.

    .. todo:: Document available URL parameters.

Dynamic relay pool (``dynamic+https://192.0.2.42/relays``)
    Syncthing will fetch the specified HTTPS URL, parse it for a JSON payload
    describing relays, select a relay from the available ones and listen via
    that as if specified as a static relay above.

    .. todo:: Document available URL parameters.


Syncing Configuration Files
---------------------------

Syncing configuration files between devices (such that multiple devices are
using the same configuration files) can cause issues. This is easy to do
accidentally if you sync your home folder between devices. A common symptom
of syncing configuration files is two devices ending up with the same Device ID.

If you want to use Syncthing to backup your configuration files, it is recommended
that the files you are backing up are in a :ref:`folder-sendonly` to prevent other
devices from overwriting the per device configuration. The folder on the remote
device(s) should not be used as configuration for the remote devices.

If you'd like to sync your home folder in non-send only mode, you may add the
folder that stores the configuration files to the :ref:`ignore list <ignoring-files>`.
If you'd also like to backup your configuration files, add another folder in
send only mode for just the configuration folder.
