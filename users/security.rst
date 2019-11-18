Security Principles
===================

Security is one of the primary project goals. This means that it should not be
possible for an attacker to join a cluster uninvited, and it should not be
possible to extract private information from intercepted traffic. Currently this
is implemented as follows.

All device to device traffic is protected by TLS. To prevent uninvited devices
from joining a cluster, the certificate fingerprint of each device is compared
to a preset list of acceptable devices at connection establishment. The
fingerprint is computed as the SHA-256 hash of the certificate and displayed
in a human-friendly encoding, called Device ID.

Incoming requests for file data are verified to the extent that the requested
file name must exist in the local index and the global model.

For information about ensuring you are running the code you think you are and
for reporting security vulnerabilities, please see the official `security page
<https://syncthing.net/security.html>`__.

Information Leakage
-------------------

Global Discovery
~~~~~~~~~~~~~~~~

When global discovery is enabled, Syncthing sends an announcement every 30
minutes to the global discovery servers so that they can keep a mapping
between your device ID and external IP. The announcement contain the device
ID and listening port(s). Also, when connecting to other devices that have
not been seen on the local network, a query is sent to the global discovery
servers containing the device ID of the requested device. The connection to
the discovery server is encrypted using TLS and the discovery server
certificate is verified, so the contents of the query should be considered
private between the device and the discovery server. The discovery servers
are currently hosted by :user:`calmh`. Global discovery defaults to **on**.

When turned off, devices with dynamic addresses not on the local network cannot
be found and connected to.

An eavesdropper on the Internet can deduce which machines are running
Syncthing with global discovery enabled, and what their device IDs are.

The operator of the discovery server can map arbitrary device addresses to
IP addresses, and deduce which devices are connected to each other.

If a different global discovery server is configured, no data is sent to the
default global discovery servers.

Local Discovery
~~~~~~~~~~~~~~~

When local discovery is enabled, Syncthing sends broadcast (IPv4) and multicast
(IPv6) packets to the local network every 30 seconds. The packets contain the
device ID and listening port. Local discovery defaults to **on**.

An eavesdropper on the local network can deduce which machines are running
Syncthing with local discovery enabled, and what their device IDs are.

When turned off, devices with dynamic addresses on the local network cannot be
found and connected to.

Upgrade Checks
~~~~~~~~~~~~~~

When automatic upgrades are enabled, Syncthing checks for a new version at
startup and then once every twelve hours. This is by an HTTPS request to the
download site for releases, currently hosted by :user:`calmh`.
Automatic upgrades default to **on** (unless Syncthing was compiled with
upgrades disabled).

Even when automatic upgrades are disabled in the configuration, an upgrade check
as above is done when the GUI is loaded, in order to show the "Upgrade to ..."
button when necessary. This can be disabled only by compiling Syncthing with
upgrades disabled.

The actual download, should an upgrade be available, is done from
**GitHub**, thus exposing the user to them.

The upgrade check (or download) requests *do not* contain any identifiable
information about the user or device.

Usage Reporting
~~~~~~~~~~~~~~~

When usage reporting is enabled, Syncthing reports usage data at startup and
then every 24 hours. The report is sent as an HTTPS POST to the usage reporting
server, currently hosted by :user:`calmh`. The contents of the usage report can
be seen behind the "Preview" link in settings. Usage reporting defaults to
**off** but the GUI will ask once about enabling it, shortly after the first
install.

The reported data is protected from eavesdroppers, but the connection to the
usage reporting server itself may expose the client as running Syncthing.

Sync Connections (BEP)
~~~~~~~~~~~~~~~~~~~~~~

Sync connections are attempted to all configured devices, when the address is
possible to resolve. The sync connection is based on TLS 1.2 or TLS 1.3. The TLS
certificates can be obtained by an eavesdropper, altough it is more difficult to do so in TLS 1.3. This means that the contents of the certificate are visible, which includes certificate Common Name (by default ``syncthing``).

An eavesdropper can deduce that this is a Syncthing connection and under certain circumstances calculate the
device IDs involved based on the hashes of the sent certificates.

Likewise, if the sync port (default 22000) is accessible from the internet, a
port scanner may discover it, attempt a TLS negotiation and thus obtain the
device certificate. This provides the same information as in the eavesdropper
case.

Relay Connections
~~~~~~~~~~~~~~~~~

When relaying is enabled, Syncthing will look up the pool of public relays
and establish a connection to one of them (the best, based on an internal
heuristic). The selected relay server will learn the connecting device's
device ID. Relay servers can be run by **anyone in the general public**.
Relaying defaults to **on**. Syncthing can be configured to disable
relaying, or only use specific relays.

If a relay connections is required between two devices, the relay will learn
the other device's device ID as well.

Any data exchanged between the two devices is encrypted as usual and not
subject to inspection by the relay.

Web GUI
~~~~~~~

If the web GUI is accessible, it exposes the device as running Syncthing. The
web GUI defaults to being reachable from the **local host only**.

In Short
--------

Parties doing surveillance on your network (whether that be corporate IT, the
NSA or someone else) will be able to see that you use Syncthing, and your device
IDs `are OK to share anyway
<https://docs.syncthing.net/users/faq.html#should-i-keep-my-device-ids-secret>`__,
but the actual transmitted data is protected as well as we can. Knowing your
device ID can expose your IP address, using global discovery.

Protecting your Syncthing keys and identity
-------------------------------------------

Anyone who can access the Syncthing TLS keys and config file on your device can
impersonate your device, connect to your peers, and then have access to your
synced files. Here are some general principles to protect your files:

#. If a device of yours is lost, make sure to revoke its access from your other
   devices.
#. If you're syncing confidential data on an encrypted disk to guard against
   device theft, put the Syncthing config folder on the same encrypted disk to
   avoid leaking keys and metadata. Or, use whole disk encryption.
