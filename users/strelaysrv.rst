.. _strelaysrv:

Syncthing Relay Server
======================

Synopsis
--------

::

    strelaysrv [-debug] [-ext-address=<address>] [-global-rate=<bytes/s>] [-keys=<dir>] [-listen=<listen addr>]
               [-message-timeout=<duration>] [-nat] [-nat-lease=<duration> [-nat-renewal=<duration>]
               [-nat-timeout=<duration>] [-network-timeout=<duration>] [-per-session-rate=<bytes/s>]
               [-ping-interval=<duration>] [-pools=<pool addresses>] [-protocol=<string>] [-provided-by=<string>]
               [-status-srv=<listen addr>]

Description
-----------

Syncthing relies on a network of community-contributed relay servers. Anyone
can run a relay server, and it will automatically join the relay pool and be
available to Syncthing users. The current list of relays can be found at
http://relays.syncthing.net/.

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

.. cmdoption:: -nat

    Use UPnP/NAT-PMP to acquire external port mapping

..cmdoption:: -nat-lease=<duration>

    NAT lease length in minutes (default 60)

..cmdoption:: -nat-renewal=<duration>

    NAT renewal frequency in minutes (default 30)

..cmdoption:: -nat-timeout=<duration>

    NAT discovery timeout in seconds (default 10)

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
    "http://relays.syncthing.net/endpoint"). Blank to disable announcement to
    a pool, thereby remaining a private relay.

.. cmdoption:: -protocol=<string>

    Protocol used for listening. 'tcp' for IPv4 and IPv6, 'tcp4' for IPv4, 'tcp6' for IPv6 (default "tcp").

.. cmdoption:: -provided-by=<string>

    An optional description about who provides the relay.

.. cmdoption:: -status-srv=<listen addr>

    Listen address for status service (blank to disable) (default ":22070").


Setting Up
----------

Primarily, you need to decide on a directory to store the TLS key and
certificate and a listen port. The default listen port of 22067 works, but for
optimal compatibility a well known port for encrypted traffic such as 443 is
recommended. This may require additional setup to work without running
as root or a privileged user, see `Running on port 443 as an unprivileged user`_
below. In principle something similar to this should work on a Linux/Unix
system::

    $ sudo useradd relaysrv
    $ sudo mkdir /etc/relaysrv
    $ sudo chown relaysrv /etc/relaysrv
    $ sudo -u relaysrv /usr/local/bin/relaysrv -keys /etc/relaysrv

This creates a user ``relaysrv`` and a directory ``/etc/relaysrv`` to store
the keys. The keys are generated on first startup. The relay will join the
global relay pool, unless a ``-pools=""`` argument is given.

To make the relay server start automatically at boot, use the recommended
procedure for your operating system.

Client configuration
~~~~~~~~~~~~~~~~~~~~

Syncthing can be configured to use specific relay servers (exclusively of the public pool) by adding the required servers to the Sync Protocol Listen Address field, under Actions and Settings. The format is as follows:

  relay://<host name|IP>[:port]/?id=<relay device ID>

For example:

  relay://private-relay-1.example.com:443/?id=ITZRNXE-YNROGBZ-HXTH5P7-VK5NYE5-QHRQGE2-7JQ6VNJ-KZUEDIU-5PPR5AM
The relay's device ID is output on start-up.

Running on port 443 as an unprivileged user
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

It is recommended that you run the relay on port 443 (or another port which is
commonly allowed through corporate firewalls), in order to maximise the chances
that people are able to connect. However, binding to ports below 1024 requires
root privileges, and running a relay as root is not recommended. Thankfully
there are a couple of approaches available to you.

One option is to run the relay on port 22067, and use an ``iptables`` rule
to forward traffic from port 443 to port 22067, for example::

    iptables -t nat -A PREROUTING -i eth0 -p tcp --dport 443 -j REDIRECT --to-port 22067

Or, if you're using ``ufw``, add the following to ``/etc/ufw/before.rules``::

    *nat
    :PREROUTING ACCEPT [0:0]
    :POSTROUTING ACCEPT [0:0]

    -A PREROUTING -i eth0 -p tcp --dport 443 -j REDIRECT --to-port 22067

    COMMIT

You will need to start ``relaysrv`` with ``-ext-address ":443"``. This tells
``relaysrv`` that it can be contacted on port 443, even though it is listening
on port 22067. You will also need to let both port 443 and 22067 through your
firewall.

Another option is `described here <https://wiki.apache.org/httpd/NonRootPortBinding>`__,
although your milage may vary.

Firewall Considerations
-----------------------

The relay server listens on two ports by default.  One for data connections and the other
for providing public statistics at http://relays.syncthing.net/.  The firewall, such as
``iptables``, must permit incoming TCP connetions to the following ports:

* Data port:  ``22067/tcp`` overriden with ``-listen`` and advertised with ``-ext-address``
* Status port: ``22070/tcp`` overriden with ``-status-srv`` 

Runtime ``iptables`` rules to allow access to the default ports::

    iptables -I INPUT -p tcp --dport 22067 -j ACCEPT
    iptables -I INPUT -p tcp --dport 22070 -j ACCEPT
    
Please consult Linux distribution documentation to persist firewall rules.

See Also
--------

:manpage:`syncthing-relay(7)`, :manpage:`syncthing-faq(7)`,
:manpage:`syncthing-networking(7)`
