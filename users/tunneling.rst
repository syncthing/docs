SSH Tunneling
===========================================

SSH tunneling can serve two purposes. One is a point-to-point secure tunnel between two machines (e.g., for Syncthing purposes), and another is to make use of a port you have available when you don't control the firewall. Here we'll assume **22/TCP** is open from ``hosta`` to ``hostb`` and we want to Syncthing directories between these two hosts. There is no need for ports **22000/TCP** or **22001/TCP** to be open between the machines for this to work. Port **21027/UDP** isn't needed because there is no discovery, we'll explicitly tell each member where to find the other.

SSH server config is not in scope, but we'll also assume the server, ``hostb``, is configured to allow SSH connections as user ``syncguy`` with RSA key ``somekey.pem``, and port forwarding is allowed. Meanwhile, ``hosta`` is the SSH client. For Syncthing they are symmetrical peers.

Create the SSH Tunnel
---------------------
First open a tunnel from ``hosta`` to ``hostb`` by running the SSH client on ``hosta``, such that **localhost:22001/TCP** on each machine redirects to **localhost:22000/TCP** on the other (for **syncthing** to use)::

 #/bin/bash
 ssh -i ~/.keys/somekey.pem \
    -L 127.0.0.1:22001:127.0.0.1:22000 \
    -R 127.0.0.1:22001:127.0.0.1:22000 \
    syncguy@hostb

127.0.0.1 is explicitly used throughout the example so the tunnels and Syncthing do NOT listen on externally exposed interfaces, for better security.

Listen on localhost
-------------------
Now in Syncthing on both sides of the tunnel (``hosta`` and ``hostb``) in Settings | Connections, you can disable/uncheck all options: Enable NAT Traversal, Local Discovery, Global Discovery, and Enable Relaying. Also configure Syncthing to listen only on localhost by setting Sync Protocol Listen Addresses to::

 tcp://127.0.0.1:22000

Provide Address for Remote Device
---------------------------------
Next add the remote device and use Edit | Advanced to assign the Addresses of::

 tcp://127.0.0.1:22001

Port **22001/TCP** is the SSH tunnel that will redirect to localhost port **22000/TCP** on the other machine. This same configuration is done on both ``hosta`` and ``hostb``, and then they can find each other through the tunnel.

Troubleshooting
---------------
To ensure all is working, run netstat, on both ``hosta`` and ``hostb``, with elevated privilege to confirm **ssh** and **syncthing** are listening, and that connections are established between **ssh** and **syncthing** on port 22001::

 [syncguy@hosta ~]$ sudo netstat -tupna | grep 2200

A typical correct result looks like this::

 tcp        0      0 127.0.0.1:22000         0.0.0.0:*               LISTEN      16035/bin/syncthing 
 tcp        0      0 127.0.0.1:22001         0.0.0.0:*               LISTEN      16112/ssh           
 tcp        0      0 127.0.0.1:22001         127.0.0.1:45042         ESTABLISHED 16112/ssh           
 tcp        0      0 127.0.0.1:45042         127.0.0.1:22001         ESTABLISHED 16035/bin/syncthing

Common mistakes when SSH tunneling include misconfiguration of the SSH client and/or server, wrong permissions on keys, forgetting about selinux policy, and being blocked by firewall, i.e. even **22/TCP** is blocked. If you want the tunnel to establish after reboot and re-establish upon failure, consider **autossh**.
