numConnections
==============

.. versionadded:: 1.25.0

``numConnections`` is a device setting that affects connection handling. A
zero value means to use the Syncthing default. As of version 1.25.0 the
default is to use one connection, like earlier versions of Syncthing. This
may change in the future.

Multiple connections will be maintained to the device if you set this to a
value greater than one. Multiple connections can yield improved performance
by load-balancing traffic over multiple physical links or in other
scenarios.

A simple form of negotiation is used to decide how many connections to use
between a device pair. It goes like this:

- If either side is configured to use a single connection, then a single
  connection is used. Since the default is to use a single connection this
  means that to use more than one connection both sides must be configured
  to do so.
- If both sides are configured to use multiple connections, then the larger
  of the two values is used. That is, if one side is configured to use three
  connections and the other is set to use eight connections, eight
  connections will be used.
- A maximum of 128 connections will be used under all circumstances. It is
  likely that the "return on investment" in further connections is
  negligible above about 10 to 20 connections, so this limit should be
  sufficient for all realistic use cases.

.. note::

    Additional connections are established over time, roughly at the rate of
    one per minute when Syncthing is in a steady state, so you may not see
    the expected number of connections immediately after changing this
    setting.

Load Balancing
--------------

When there are multiple connections between two devices, one connection is
dedicated to metadata transmission: index updates, changes to folder pause
status, etc. Requests and responses are sent over the other connections
randomly. The number of connections in the GUI is represented as `1 + n` for
this reason, e.g. if you configure four connections, the GUI will show `1 +
3` to indicate one metadata connection and three data connections.

Rate Limiting
-------------

Device rate limiting applies to the sum of traffic on all connections,
regardless of the number of connections. The limit is not per connection.

Connection Types
----------------

Both TCP and QUIC connections are supported for multiple connections.
Syncthing will, however, only keep connections with the best priority; by
default, TCP has better priority than QUIC, so establishing a TCP connection
will cause existing QUIC connections to be closed. Connection priorities can
be configured.

Multiple connections cannot be established over relays.
