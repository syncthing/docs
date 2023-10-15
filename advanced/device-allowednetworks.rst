.. _allowed-network:

allowedNetwork
==============

``allowedNetwork`` is an advanced device setting that affects connection
handling. The default is for this setting is the XML node is not present,
indicating that there is no restriction on the allowed networks for a device.

By setting this to a list of networks, connections to the
given device will be limited to those networks. The networks refer to the
address of the *remote* device, not the network that the local device is
presently on.

Given a list of:
```<device ...>
  <allowedNetwork>192.168.0.0/16</allowedNetwork>
  <allowedNetwork>172.16.0.0/12</allowedNetwork>
  <allowedNetwork>2001:db8::/32</allowedNetwork>
</device>```

Syncthing will:

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
