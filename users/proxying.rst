.. _proxying:

Using Proxies
=============

.. versionadded:: 0.12.0

Syncthing can run behind a SOCKS5, HTTP, or HTTPS proxy. This enables use behind some corporate
firewalls, tunneling via SSH, and over Tor. The Syncthing instance that is
behind the proxy is limited to outbound connections - it can not listen for
incoming connections via the proxy. It is however possible to receive incoming
connections via :ref:`relaying`.

.. versionadded:: 2.1.0

    Support for HTTP and HTTPS proxies via ``all_proxy``.

There is no graphical configuration to enable proxy usage in Syncthing.
Instead, an environment variable ``all_proxy`` must be set that points to the
proxy. The value of this variable is the proxy URL. For example, on a
Linux-like system::

  $ export all_proxy=socks5://10.20.30.40:5060
  $ syncthing
  [monitor] 11:25:39 INFO: Starting syncthing
  ...
  [monitor] 11:25:40 INFO: Proxy settings detected

Or for an HTTP/HTTPS proxy::

  $ export all_proxy=http://10.20.30.40:8080
  $ # Or with authentication and encryption
  $ export all_proxy=https://user:password@proxy.example.com:8443
  $ syncthing

HTTP proxies requiring basic authentication can be specified by embedding
credentials in the proxy URL. When credentials are used over a cleartext HTTP
proxy, a warning will be logged.

Note that this environment variable is *not* named with capital letters - it
must be exactly ``all_proxy``. The "Proxy settings detected" log message
indicates that Syncthing is using the proxy configuration.

Disabling Fallback
------------------

.. versionadded:: 0.13.0

By default, Syncthing will attempt a direct connection if a connection via the
proxy fails. This is desirable when moving frequently between a proxied and
non-proxied environment. However it may be undesirable if the intention is to
force all connections through a Tor proxy or similar. In that case, setting the
environment variable ``ALL_PROXY_NO_FALLBACK`` (with capital letters) will
prevent the fallback behavior. For example::

  $ export all_proxy=socks5://10.20.30.40:5060
  $ export ALL_PROXY_NO_FALLBACK=1
  $ syncthing
  [monitor] 11:33:13 INFO: Starting syncthing
  ...
  [monitor] 11:33:13 INFO: Proxy settings detected
  [monitor] 11:33:13 INFO: Proxy fallback disabled

Caveats
-------

Syncthing will use the ``http_proxy`` / ``https_proxy`` settings in combination with
``all_proxy``. Proxies specified with ``http_`` variants will be used behind the
configured proxy.

If this is not desired, make sure to undefine the ``http_proxy`` / ``https_proxy``
variables for syncthing.

When using an HTTP proxy via ``all_proxy``, only outbound TCP connections are
supported. Non-TCP transports and listening sockets are not supported through
HTTP or HTTPS proxies.
