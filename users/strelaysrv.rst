Syncthing Relay Server
======================

Synopsis
--------

::

    strelaysrv [-debug] [-ext-address=<address>] [-global-rate=<bytes/s>] [-keys=<dir>] [-listen=<listen addr>]
               [-message-timeout=<duration>] [-network-timeout=<duration>] [-per-session-rate=<bytes/s>]
               [-ping-interval=<duration>] [-pools=<pool addresses>] [-provided-by=<string>] [-status-srv=<listen addr>]

Description
-----------

Syncthing relies on a network of community-contributed relay servers. Anyone
can run a relay server, and it will automatically join the relay pool and be
available to Syncthing users. The current list of relays can be found at
https://relays.syncthing.net.

Options
-------

.. cmdoption:: -debug

    Enable debug output.

.. cmdoption:: -ext-address=<address>

    An optional address to advertising as being available on. Allows listening
    on an unprivileged port with port forwarding from e.g. 443, and be
    connected to on port 443.

.. cmdoption:: -global-rate=<bytes/s>

    Global rate limit, in bytes/s.

.. cmdoption:: -keys=<dir>

    Directory where cert.pem and key.pem is stored (default ".").

.. cmdoption:: -listen=<listen addr>

    Protocol listen address (default ":22067").

.. cmdoption:: -message-timeout=<duration>

    Maximum amount of time we wait for relevant messages to arrive (default 1m0s).

.. cmdoption:: -network-timeout=<duration>

    Timeout for network operations between the client and the relay. If no data
    is received between the client and the relay in this period of time, the
    connection is terminated. Furthermore, if no data is sent between either
    clients being relayed within this period of time, the session is also
    terminated. (default 2m0s)

.. cmdoption:: -per-session-rate=<bytes/s>

    Per session rate limit, in bytes/s.

.. cmdoption:: -ping-interval=<duration>

    How often pings are sent (default 1m0s).

.. cmdoption:: -pools=<pool addresses>

    Comma separated list of relay pool addresses to join (default
    "https://relays.syncthing.net/endpoint"). Blank to disable announcement to
    a pool, thereby remaining a private relay.

.. cmdoption:: -provided-by=<string>

    An optional description about who provides the relay.

.. cmdoption:: -status-srv=<listen addr>

    Listen address for status service (blank to disable) (default ":22070").

See Also
--------

:manpage:`syncthing-relay(7)`, :manpage:`syncthing-faq(7)`,
:manpage:`syncthing-networking(7)`
