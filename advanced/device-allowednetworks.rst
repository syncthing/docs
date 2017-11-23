.. _allowed-networks:

allowedNetworks
===============

``allowedNetworks`` is an advanced device setting that affects connection
handling. The default is for this setting to be empty, indicating that there
is no restriction on the allowed networks for a device.

By setting this to a comma separated list of networks, connections to the
given device will be limited to those networks. The networks refer to the
address of the *remote* device, not the network that the local device is
presently on.

Given a value of ``192.168.0.0/16, 172.16.0.0/12, 2001:db8::/32`` Syncthing will:

 - Allow connections from the device from addresses in the specified
   networks.

 - Reject connections from the device from addresses outside the specified
   networks.

 - Attempt connections to addresses in the specified networks (manually
   configured or discovered).

 - Not attempt connections to addresses outside the specified networks,
   regardless of whether manually configured or automatically discovered.

Allowed values are numeric IPv4 and IPv6 prefixes in CIDR format, as in the
example. Hostnames, netmasks in octet format, etc., are not supported.

If the value is not empty it will be enforced for all connections.
Mentioning only an IPv4 prefix will thus deny all IPv6 connections and vice
versa. Use ``0.0.0.0/0`` (IPv4) or ``::/0`` (IPv6) to allow all connections on
that address family.
