.. _config:

Syncthing Configuration
=======================

Synopsis
--------

::

    $HOME/.config/syncthing
    $HOME/Library/Application Support/Syncthing
    %AppData%/Syncthing
    %localappdata%/Syncthing

Description
-----------

Syncthing uses a single directory to store configuration, crypto keys
and index caches. The location defaults to ``$HOME/.config/syncthing``
(Unix-like), ``$HOME/Library/Application Support/Syncthing`` (Mac),
``%AppData%/Syncthing`` (Windows XP) or ``%LocalAppData%/Syncthing``
(Windows 7+). It can be changed at runtime using the ``-home`` flag. In this
directory the following files are located:

:file:`config.xml`
    The configuration file, in XML format.

:file:`cert.pem`, :file:`key.pem`
    The device's RSA public and private key. These form the basis for the
    device ID. The key must be kept private.

:file:`https-cert.pem`, :file:`https-key.pem`
    The certificate and key for HTTPS GUI connections. These may be replaced
    with a custom certificate for HTTPS as desired.

:file:`index-{*}.db`
    A directory holding the database with metadata and hashes of the files
    currently on disk and available from peers.

:file:`csrftokens.txt`
    A list of recently issued CSRF tokens (for protection against browser cross
    site request forgery).

Config File Format
------------------

The following shows an example of the default configuration file (IDs will differ):

.. code-block:: xml

    <configuration version="26">
        <folder id="zj2AA-q55a7" label="Default Folder" path="/Users/jb/Sync/" type="readwrite" rescanIntervalS="60" fsWatcherEnabled="false" fsWatcherDelayS="10" ignorePerms="false" autoNormalize="true">
            <device id="3LT2GA5-CQI4XJM-WTZ264P-MLOGMHL-MCRLDNT-MZV4RD3-KA745CL-OGAERQZ"></device>
            <filesystemType>basic</filesystemType>
            <minDiskFree unit="%">1</minDiskFree>
            <versioning></versioning>
            <copiers>0</copiers>
            <pullers>0</pullers>
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
        </folder>
        <device id="3LT2GA5-CQI4XJM-WTZ264P-MLOGMHL-MCRLDNT-MZV4RD3-KA745CL-OGAERQZ" name="syno" compression="metadata" introducer="false" skipIntroductionRemovals="false" introducedBy="">
            <address>dynamic</address>
            <paused>false</paused>
        </device>
        <gui enabled="true" tls="false" debugging="false">
            <address>127.0.0.1:8384</address>
            <apikey>k1dnz1Dd0rzTBjjFFh7CXPnrF12C49B1</apikey>
            <theme>default</theme>
        </gui>
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
            <urUniqueID>LFWe2vn3</urUniqueID>
            <urURL>https://data.syncthing.net/newdata</urURL>
            <urPostInsecurely>false</urPostInsecurely>
            <urInitialDelayS>1800</urInitialDelayS>
            <restartOnWakeup>true</restartOnWakeup>
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
            <trafficClass>0</trafficClass>
            <weakHashSelectionMethod>auto</weakHashSelectionMethod>
            <stunServer>default</stunServer>
            <stunKeepaliveSeconds>24</stunKeepaliveSeconds>
            <kcpNoDelay>false</kcpNoDelay>
            <kcpUpdateIntervalMs>25</kcpUpdateIntervalMs>
            <kcpFastResend>false</kcpFastResend>
            <kcpCongestionControl>true</kcpCongestionControl>
            <kcpSendWindowSize>128</kcpSendWindowSize>
            <kcpReceiveWindowSize>128</kcpReceiveWindowSize>
            <defaultFolderPath>~</defaultFolderPath>
            <minHomeDiskFreePct>0</minHomeDiskFreePct>
            <setLowPriority>true</setLowPriority>
        </options>
    </configuration>

Configuration Element
---------------------

.. code-block:: xml

    <configuration version="26">
        <folder></folder>
        <device></device>
        <gui></gui>
        <options></options>
        <ignoredDevice>5SYI2FS-LW6YAXI-JJDYETS-NDBBPIO-256MWBO-XDPXWVG-24QPUM4-PDW4UQU</ignoredDevice>
        <ignoredFolder>bd7q3-zskm5</ignoredDevice>
    </configuration>

This is the root element. It has one attribute:

version
    The config version. Increments whenever a change is made that requires
    migration from previous formats.

It contains the elements described in the following sections and these two
additional child elements:

ignoredDevice
    Contains the ID of the device that should be ignored. Connection attempts
    from this device are logged to the console but never displayed in the web
    GUI.

ignoredFolder
    Contains the ID of the folder that should be ignored. This folder will
    always be skipped when advertised from a remote device, i.e. this will be
    logged, but there will be no dialog about it in the web GUI.


Folder Element
--------------

.. code-block:: xml

    <folder id="zj2AA-q55a7" label="Default Folder" path="/Users/jb/Sync/" type="readwrite" rescanIntervalS="60" fsWatcherEnabled="false" fsWatcherDelayS="10" ignorePerms="false" autoNormalize="true">
        <device id="3LT2GA5-CQI4XJM-WTZ264P-MLOGMHL-MCRLDNT-MZV4RD3-KA745CL-OGAERQZ"></device>
        <filesystemType>basic</filesystemType>
        <minDiskFree unit="%">1</minDiskFree>
        <versioning></versioning>
        <copiers>0</copiers>
        <pullers>0</pullers>
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
    </folder>

One or more ``folder`` elements must be present in the file. Each element
describes one folder. The following attributes may be set on the ``folder``
element:

id
    The folder ID, must be unique. (mandatory)

label
    The label of a folder is a human readable and descriptive local name. May
    be different on each device, empty, and/or identical to other folder
    labels. (optional)

path
    The path to the directory where the folder is stored on this
    device; not sent to other devices. (mandatory)

type
    Controls how the folder is handled by Syncthing. Possible values are:

    readwrite
        The folder is in default mode. Sending local and accepting remote changes.

    readonly
        The folder is in "send-only" mode -- it will not be modified by
        Syncthing on this device.

rescanIntervalS
    The rescan interval, in seconds. Can be set to zero to disable when external
    plugins are used to trigger rescans.

fsWatcherEnabled
    If enabled this detects changes to files in the folder and scans them.

.. _fsWatcherDelayS:

fsWatcherDelayS
    The duration during which changes detected are accumulated, before a scan is
    scheduled (only takes effect if ``fsWatcherEnabled`` is true).

ignorePerms
    True if the folder should ignore permissions.

autoNormalize
    Automatically correct UTF-8 normalization errors found in file names.

The following child elements may exist:

device
    These must have the ``id`` attribute and can have an ``introducedBy`` attribute,
    identifying the device that introduced us to share this folder with the given device.
    If the original introducer unshares this folder with this device, our device will follow
    and unshare the folder (subject to skipIntroductionRemovals being false on the introducer device).
    All mentioned devices are those that will be sharing the folder in question.
    Each mentioned device must have a separate ``device`` element later in the file.
    It is customary that the local device ID is included in all folders.
    Syncthing will currently add this automatically if it is not present in
    the configuration file.

minDiskFree
    The minimum required free space that should be available on the disk this folder
    resides. The folder will be stopped when the value drops below the threshold. Accepted units are
    ``%``, ``kB``, ``MB``, ``GB`` and ``TB``. Set to zero to disable.

versioning
    Specifies a versioning configuration.

.. seealso::
    :ref:`versioning`

copiers, pullers, hashers
    The number of copier, puller and hasher routines to use, or zero for the
    system determined optimum. These are low level performance options for
    advanced users only; do not change unless requested to or you've actually
    read and understood the code yourself. :)

order
    The order in which needed files should be pulled from the cluster.
    The possibles values are:

    random
        Pull files in random order. This optimizes for balancing resources among
        the devices in a cluster.

    alphabetic
        Pull files ordered by file name alphabetically.

    smallestFirst, largestFirst
        Pull files ordered by file size; smallest and largest first respectively.

    oldestFirst, newestFirst
        Pull files ordered by modification time; oldest and newest first
        respectively.

ignoreDelete
    When set to true, this device will pretend not to see instructions to
    delete files from other devices.

scanProgressIntervalS
    The interval with which scan progress information is sent to the GUI. Zero
    means the default value (two seconds).

pullerPauseS
    Tweak for rate limiting the puller when it retries pulling files. Don't
    change these unless you know what you're doing.

maxConflicts
    The maximum number of conflict copies to keep around for any given file.
    The default, -1, means an unlimited number. Setting this to zero disables
    conflict copies altogether.

disableSparseFiles
    By default, blocks containing all zeroes are not written, causing files
    to be sparse on filesystems that support the concept. When set to true,
    sparse files will not be created.

disableTempIndexes
    By default, devices exchange information about blocks available in
    transfers that are still in progress. When set to true, such information
    is not exchanged for this folder.

paused
    True if this folder is (temporarily) suspended.

weakHashThresholdPct
    Use weak hash if more than the given percentage of the file has changed. Set
    to -1 to always use weak hash. Default value is 25.

markerName
    Name of a directory or file in the folder root to be used as
    :ref:`marker-faq`. Default is ".stfolder".

fsync
    .. deprecated:: v0.14.37

    Transfer updated (from other devices) files to permanent storage before
    committing the changes to the internal database.

pullerSleepS
    .. deprecated:: v0.14.41

    Tweak for rate limiting the puller. Don't change these unless you know
    what you're doing.


Device Element
--------------

.. code-block:: xml

    <device id="5SYI2FS-LW6YAXI-JJDYETS-NDBBPIO-256MWBO-XDPXWVG-24QPUM4-PDW4UQU" name="syno" compression="metadata" introducer="false" skipIntroductionRemovals="false" introducedBy="2CYF2WQ-AKZO2QZ-JAKWLYD-AGHMQUM-BGXUOIS-GYILW34-HJG3DUK-LRRYQAR">
        <address>dynamic</address>
    </device>
    <device id="2CYF2WQ-AKZO2QZ-JAKWLYD-AGHMQUM-BGXUOIS-GYILW34-HJG3DUK-LRRYQAR" name="syno local" compression="metadata" introducer="false">
        <address>tcp://192.0.2.1:22001</address>
        <paused>true<paused>
        <allowedNetwork>192.168.0.0/16<allowedNetwork>
    </device>

One or more ``device`` elements must be present in the file. Each element
describes a device participating in the cluster. It is customary to include a
``device`` element for the local device; Syncthing will currently add one if
it is not present. The following attributes may be set on the ``device``
element:

id
    The device ID. This must be written in canonical form, that is without any
    spaces or dashes. (mandatory)

name
    A friendly name for the device. (optional)

compression
    Whether to use protocol compression when sending messages to this device.
    The possible values are:

    metadata
        Compress metadata packets, such as index information. Metadata is
        usually very compression friendly so this is a good default.

    always
        Compress all packets, including file data. This is recommended if the
        folders contents are mainly compressible data such as documents or
        text files.

    never
        Disable all compression.

introducer
    Set to true if this device should be trusted as an introducer, i.e. we
    should copy their list of devices per folder when connecting.

.. seealso::
    :ref:`introducer`

skipIntroductionRemovals
    Set to true if you wish to follow only introductions and not de-introductions.
    For example, if this is set, we would not remove a device that we were introduced
    to even if the original introducer is no longer listing the remote device as known.

introducedBy
    Defines which device has introduced us to this device. Used only for following de-introductions.

From following child elements at least one ``address`` child must exist.

address
    Contains an address or host name to use when attempting to connect to this device.
    Entries other than ``dynamic`` must be prefixed with ``tcp://`` (dual-stack),
    ``tcp4://`` (IPv4 only) or ``tcp6://`` (IPv6 only). Note that IP addresses need
    not use tcp4/tcp6; these are optional. Accepted formats are:

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
        The word ``dynamic`` (without ``tcp://`` prefix) means to use local and
        global discovery to find the device.

paused
    True if synchronization with this devices is (temporarily) suspended.

allowedNetwork
    If given, this restricts connections to this device to only this network
    (see :ref:`allowed-networks`).

GUI Element
-----------

.. code-block:: xml

    <gui enabled="true" tls="false" debugging="false">
        <address>127.0.0.1:8384</address>
        <apikey>l7jSbCqPD95JYZ0g8vi4ZLAMg3ulnN1b</apikey>
        <theme>default</theme>
    </gui>


There must be exactly one ``gui`` element. The GUI configuration is also used
by the :ref:`rest-api` and the :ref:`event-api`. The following attributes may
be set on the ``gui`` element:

enabled
    If not ``true``, the GUI and API will not be started.

tls
    If set to ``true``, TLS (HTTPS) will be enforced. Non-HTTPS requests will
    be redirected to HTTPS. When this is set to ``false``, TLS connections are
    still possible but it is not mandatory.

debugging
    This enables :ref:`profiling` and additional debugging endpoints in the :ref:`rest-api`.

The following child elements may be present:

address
    Set the listen address. One address element must be present. Allowed address formats are:

    IPv4 address and port (``127.0.0.1:8384``)
        The address and port is used as given.

    IPv6 address and port (``[::1]:8384``)
        The address and port is used as given. The address must be enclosed in
        square brackets.

    Wildcard and port (``0.0.0.0:12345``, ``[::]:12345``, ``:12345``)
        These are equivalent and will result in Syncthing listening on all
        interfaces via both IPv4 and IPv6.

user
    Set to require authentication.

password
    Contains the bcrypt hash of the real password.

apikey
    If set, this is the API key that enables usage of the REST interface.

insecureAdminAccess
    If true, this allows access to the web GUI from outside (i.e. not localhost)
    without authorization. A warning will displayed about this setting on startup.

theme
    The name of the theme to use.

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
        <urUniqueID></urUniqueID>
        <urURL>https://data.syncthing.net/newdata</urURL>
        <urPostInsecurely>false</urPostInsecurely>
        <urInitialDelayS>1800</urInitialDelayS>
        <restartOnWakeup>true</restartOnWakeup>
        <autoUpgradeIntervalH>12</autoUpgradeIntervalH>
        <keepTemporariesH>24</keepTemporariesH>
        <cacheIgnoredFiles>false</cacheIgnoredFiles>
        <progressUpdateIntervalS>5</progressUpdateIntervalS>
        <limitBandwidthInLan>false</limitBandwidthInLan>
        <minHomeDiskFree unit="%">1</minHomeDiskFree>
        <releasesURL>https://api.github.com/repos/syncthing/syncthing/releases?per_page=30</releasesURL>
        <overwriteRemoteDeviceNamesOnConnect>false</overwriteRemoteDeviceNamesOnConnect>
        <tempIndexMinBlocks>10</tempIndexMinBlocks>
        <defaultFolderPath>~</defaultFolderPath>
    </options>

The ``options`` element contains all other global configuration options.

listenAddress
    The listen address for incoming sync connections. See
    :ref:`listen-addresses` for allowed syntax.

globalAnnounceServer
    A URI to a global announce (discovery) server, or the word ``default`` to
    include the default servers. Any number of globalAnnounceServer elements
    may be present. The syntax for non-default entries is that of a HTTP or
    HTTPS URL. A number of options may be added as query options to the URL:
    ``insecure`` to prevent certificate validation (required for HTTP URLs)
    and ``id=<device ID>`` to perform certificate pinning. The device ID to
    use is printed by the discovery server on startup.

globalAnnounceEnabled
    Whether to announce this device to the global announce (discovery) server,
    and also use it to look up other devices.

localAnnounceEnabled
    Whether to send announcements to the local LAN, also use such
    announcements to find other devices.

localAnnouncePort
    The port on which to listen and send IPv4 broadcast announcements to.

localAnnounceMCAddr
    The group address and port to join and send IPv6 multicast announcements on.

maxSendKbps
    Outgoing data rate limit, in kibibytes per second.

maxRecvKbps
    Incoming data rate limits, in kibibytes per second.

reconnectionIntervalS
    The number of seconds to wait between each attempt to connect to currently
    unconnected devices.

relaysEnabled
    When true, relays will be connected to and potentially used for device to device connections.

relayReconnectIntervalM
    Sets the interval, in minutes, between relay reconnect attempts.

startBrowser
    Whether to attempt to start a browser to show the GUI when Syncthing starts.

natEnabled
    Whether to attempt to perform a UPnP and NAT-PMP port mapping for
    incoming sync connections.

natLeaseMinutes
    Request a lease for this many minutes; zero to request a permanent lease.

natRenewalMinutes
    Attempt to renew the lease after this many minutes.

natTimeoutSeconds
    When scanning for UPnP devices, wait this long for responses.

urAccepted
    Whether the user has accepted to submit anonymous usage data. The default,
    ``0``, mean the user has not made a choice, and Syncthing will ask at some
    point in the future. ``-1`` means no, a number above zero means that that
    version of usage reporting has been accepted.

urSeen
    The highest usage reporting version that has already been shown in the web GUI.

urUniqueID
    The unique ID sent together with the usage report. Generated when usage
    reporting is enabled.

urURL
    The URL to post usage report data to, when enabled.

urPostInsecurely
    When true, the UR URL can be http instead of https, or have a self-signed
    certificate. The default is ``false``.

urInitialDelayS
    The time to wait from startup to the first usage report being sent. Allows
    the system to stabilize before reporting statistics.

restartOnWakeup
    Whether to perform a restart of Syncthing when it is detected that we are
    waking from sleep mode (i.e. a folded up laptop).

autoUpgradeIntervalH
    Check for a newer version after this many hours. Set to zero to disable
    automatic upgrades.

upgradeToPreReleases
    If true, automatic upgrades include release candidates (see
    :ref:`release-channels`).

keepTemporariesH
    Keep temporary failed transfers for this many hours. While the temporaries
    are kept, the data they contain need not be transferred again.

cacheIgnoredFiles
    Whether to cache the results of ignore pattern evaluation. Performance
    at the price of memory. Defaults to ``false`` as the cost for evaluating
    ignores is usually not significant.

progressUpdateIntervalS
    How often in seconds the progress of ongoing downloads is made available to
    the GUI.

limitBandwidthInLan
    Whether to apply bandwidth limits to devices in the same broadcast domain
    as the local device.

minHomeDiskFree
    The minimum required free space that should be available on the
    partition holding the configuration and index. Accepted units are ``%``, ``kB``,
    ``MB``, ``GB`` and ``TB``.

releasesURL
    The URL from which release information is loaded, for automatic upgrades.

alwaysLocalNet
    Network that should be considered as local given in CIDR notation.

overwriteRemoteDeviceNamesOnConnect
    If set, device names will always be overwritten with the name given by
    remote on each connection. By default, the name that the remote device
    announces will only be adopted when a name has not already been set.

tempIndexMinBlocks
    When exchanging index information for incomplete transfers, only take
    into account files that have at least this many blocks.

unackedNotificationID
    ID of a notification to be displayed in the web GUI. Will be removed once
    the user acknowledged it (e.g. an transition notice on an upgrade).

trafficClass
    Specify a type of service (TOS)/traffic class of outgoing packets.

weakHashSelectionMethod
    Specify whether weak hashing is used, possible options are
    ``WeakHashAlways``, ``WeakHashNever`` and ``WeakHashAuto``. Deciding
    automatically means running benchmarks at startup to decide whether the
    performance impact is acceptable (this is the default).

stunServer
    Server to be used for STUN, given as ip:port. The keyword ``default`` gets
    expanded to
    ``stun.callwithus.com:3478``, ``stun.counterpath.com:3478``,
    ``stun.counterpath.net:3478``, ``stun.ekiga.net:3478``,
    ``stun.ideasip.com:3478``, ``stun.internetcalls.com:3478``,
    ``stun.schlund.de:3478``, ``stun.sipgate.net:10000``,
    ``stun.sipgate.net:3478``, ``stun.voip.aebc.com:3478``,
    ``stun.voiparound.com:3478``, ``stun.voipbuster.com:3478``,
    ``stun.voipstunt.com:3478``, ``stun.voxgratia.org:3478`` and
    ``stun.xten.com:3478`` (this is the default).

stunKeepaliveSeconds
    Interval in seconds between contacting a STUN server to
    maintain NAT mapping. Default is ``24`` and you can set it to ``0`` to
    disable contacting STUN servers.

kcpNoDelay, kcpUpdateIntervalMs, kcpFastResend, kcpCongestionControl, kcpSendWindowSize, kcpReceiveWindowSize
    Various KCP tweaking parameters.

defaultFolderPath
    The UI will propose to create new folders at this path. This can be disabled by
    setting this to an empty string.

.. _set-low-priority:

setLowPriority
    Syncthing will attempt to lower its process priority at startup.
    Specifically: on Linux, set itself to a separate process group, set the
    niceness level of that process group to nine and the I/O priority to
    best effort level five; on other Unixes, set the process niceness level
    to nine; on Windows, set the process priority class to below normal. To
    disable this behavior, for example to control process priority yourself
    as part of launching Syncthing, set this option to ``false``.

.. _listen-addresses:

Listen Addresses
^^^^^^^^^^^^^^^^

The following address types are accepted in sync protocol listen addresses:

Default listen addresses (``default``)
    This is equivalent to ``tcp://0.0.0.0:22000`` and
    ``dynamic+https://relays.syncthing.net/endpoint``.

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

If you'd like to sync your home folder in non-send-only mode, you may add the
folder that stores the configuration files to the :ref:`ignore list <ignoring-files>`.
If you'd also like to backup your configuration files, add another folder in
send-only mode for just the configuration folder.
