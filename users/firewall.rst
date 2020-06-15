.. _firewall-setup:

Firewall Setup
==============

Port Forwards
-------------

If you have a NAT router which supports UPnP, the easiest way to get a working
port forward is to make sure UPnP setting is enabled on both Syncthing and the
router – Syncthing will try to handle the rest. If it succeeds you will see a
message in the console saying::

    Created UPnP port mapping for external port XXXXX on UPnP device YYYYY.

If this is not possible or desirable you should set up a port forward for port
**22000/TCP**, or the port set in the *Sync Protocol Listen Address* setting.
The external forwarded port and the internal destination port has to be the same
(i.e. 22000/TCP).

Communication in Syncthing works both ways. Therefore if you set up port
forwards for one device, other devices will be able to connect to it even when
they are behind a NAT network or firewall.

In the absence of port forwarding, :ref:`relaying` may work well enough to get
devices connected and synced, but will perform poorly in comparison to a
direct connection.

Local Firewall
--------------

If your PC has a local firewall, you will need to open the following ports for
incoming and outgoing traffic:

-  Port **22000/TCP** (or the actual listening port if you have changed
   the *Sync Protocol Listen Address* setting.)
-  Port **21027/UDP** (for discovery broadcasts on IPv4 and multicasts on IPv6)

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
If you are using [Firewalld](https://www.firewalld.org) it has included
support for syncthing (since version 0.5.0, January 2018), and you can enable
it with

    sudo firewall-cmd --zone=public --add-service=syncthing --permanent
    sudo firewall-cmd --reload

Similarly there is also a syncthing-gui service.


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
