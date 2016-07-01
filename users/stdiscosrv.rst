Synthing Discovery Server
=========================

Synopsis
--------

::

    stdiscosrv [-cert=<file>] [-db-backend=<string>] [-db-dsn=<string>] [-debug] [-http] [-key=<string>]
               [-limit-avg=<int>] [-limit-burst=<int>] [-limit-cache=<int>] [-listen=<address>]
               [-stats-file=<file>]

Description
-----------

Syncthing relies on a discovery server to find peers. Anyone can run a discovery
server and point its syncthing installations to it.

Options
-------

.. cmdoption:: -cert=<file>

    Certificate file (default "cert.pem").

.. cmdoption:: -db-backend=<string>

    Database backend to use (default "ql").

.. cmdoption:: -db-dsn=<string>

    Database DSN (default "memory://stdiscosrv").

.. cmdoption:: -debug

    Enable debug output.

.. cmdoption:: -http

    Listen on HTTP (behind an HTTPS proxy).

.. cmdoption:: -key=<file>

    Key file (default "key.pem").

.. cmdoption:: -limit-avg=<int>

    Allowed average package rate, per 10 s (default 5).

.. cmdoption:: -limit-burst=<int>

    Allowed burst size, packets (default 20).

.. cmdoption:: -limit-cache=<int>

    Limiter cache entries (default 10240).

.. cmdoption:: -listen=<address>

    Listen address (default ":8443").

.. cmdoption:: -stats-file=<file>

    File to write periodic operation stats to.

See Also
--------

:manpage:`syncthing-networking(7)`, :manpage:`syncthing-faq(7)`
