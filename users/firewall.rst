.. _firewall-setup:

Firewall Setup
==============

Router Setup
------------

Port Forwards
~~~~~~~~~~~~~

If you have a NAT router which supports UPnP, the easiest way to get a working
port forward is to make sure UPnP setting is enabled on both Syncthing and the
router – Syncthing will try to handle the rest. If it succeeds you will see a
message in the console saying::

    Created UPnP port mapping for external port XXXXX on UPnP device YYYYY.

If this is not possible or desirable, you should set up a port forwarding for ports
**22000/TCP** and **22000/UDP** (or whichever port is set in the *Sync Protocol Listen Address* setting).
The external forwarded ports and the internal destination ports have to be the same
(e.g. 22000/TCP).

Communication in Syncthing works both ways. Therefore if you set up port
forwards for one device, other devices will be able to connect to it even when
they are behind a NAT network or firewall.

In the absence of port forwarding, :ref:`relaying` may work well enough to get
devices connected and synced, but will perform poorly in comparison to a
direct connection.

Local Discovery
~~~~~~~~~~~~~~~

The router needs to allow/forward broad-/multicasts for local discovery to work.
Usually these are allowed by default in a single local subnet, but may be
blocked between different subnets or even between a bridged Wi-Fi and LAN.

If you are unable to set up your router thus or your firewall as shown below,
and your devices have static IP addresses, you can specify them directly by
changing the default ``dynamic`` setting for *Addresses* to something like:
``tcp://192.168.1.xxx:22000, dynamic``.

Local Firewall
--------------

If your PC has a local firewall, you will need to open the following ports for
incoming and outgoing traffic:

-  Port **22000/TCP**: TCP based sync protocol traffic
-  Port **22000/UDP**: QUIC based sync protocol traffic
-  Port **21027/UDP**: for discovery broadcasts on IPv4 and multicasts on IPv6

If you configured a custom port in the *Sync Protocol Listen Address* setting,
you have to adapt the firewall rules accordingly.

Uncomplicated Firewall (ufw)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~
If you're using ``ufw`` on Linux and have installed the `Syncthing package
<https://apt.syncthing.net/>`__, you can allow the necessary ports by running::

    sudo ufw allow syncthing

If you also want to allow external access to the Syncthing web GUI, run::

    sudo ufw allow syncthing-gui

Allowing external access is **not**  necessary for a typical installation.

You can then verify that the ports mentioned above are allowed::

    sudo ufw status verbose

In case you installed Syncthing manually you can follow the `instructions to manually add the syncthing preset
<https://github.com/syncthing/syncthing/tree/main/etc/firewall-ufw>`__ to ufw.

Firewalld
~~~~~~~~~
If you are using `Firewalld <https://firewalld.org/>`__ it has included
support for syncthing (since version 0.5.0, January 2018), and you can enable
it with::

    sudo firewall-cmd --zone=public --add-service=syncthing --permanent
    sudo firewall-cmd --reload

Similarly there is also a ``syncthing-gui`` service.

nftables
~~~~~~~~
For hosts using a somewhat standard setup of ``nftables``, placing the
following content in ``/etc/nftables.d/syncthing.nft`` should allow syncthing
to be discovered via local discovery and receive direct connections.

    table inet filter {
    	chain input {
    		udp dport 21027 accept comment "Allow syncthing discovery"
    		udp dport 22000 accept comment "Allow syncthing peers"
    	}
    }

Remote Web GUI
--------------

To be able to access the web GUI from other computers, you need to change the
*GUI Listen Address* setting from the default ``127.0.0.1:8384`` to
``0.0.0.0:8384``. You also need to open the port in your local firewall if you
have one.

Tunneling via SSH
~~~~~~~~~~~~~~~~~

If you have SSH access to the machine running Syncthing but would rather not
open the web GUI port to the outside world, you can access it through a SSH
tunnel instead. You can start a tunnel with a command like the following::

    ssh -L 9999:localhost:8384 machine

This will bind to your local port 9999 and forward all connections from there to
port 8384 on the target machine. This still works even if Syncthing is bound to
listen on localhost only.

Via a Proxy
-----------

Syncthing can use a SOCKS5 proxy for outbound connections. Please see :ref:`proxying`.
