.. _gui-listen:

The GUI Listen Address
======================

The GUI (and, together with it, the REST API) listens on a single TCP port
for HTTP and HTTPS connections. By default this address is ``127.0.0.1:8384``.
The ``127.0.0.1`` part means "localhost" which means it only listens for
connections from the same computer Syncthing is running on. This is a
reasonably safe default as it means you need to be logged in on the computer
to access the GUI - it cannot be accessed from the network.

To allow access from the network, change the listen address to
``0.0.0.0:8384``, where "0.0.0.0" means roughly "anywhere". You should then
also set a username and a strong password for authentication and check the
option to use HTTPS. You are otherwise, potentially, opening up your
Syncthing installation for the world.

Note that specifying your computer's LAN address (e.g. ``192.168.0.123:8384``)
will **NOT** restrict access to only devices on your local network!  Connections
with that address *as destination* will then be accepted, regardless of their
origin.  Proper network configuration and security (especially a firewall) is
required to enforce such filtering, as it cannot be done reliably by Syncthing
itself.

Unix sockets are supported by specifying an absolute path
(``/run/syncthing/syncthing.socket``).

Port Numbers
------------

The default port number is 8384. It's traditional for custom HTTP services
to live somewhere in the 8xxx-series and it's an unusual enough port that
it's usually free. Syncthing will however choose another, random, port if
port 8384 is taken by something else at the time of installation. The port
used is always displayed on the console when starting up.

You can change the port number to something else if you prefer, keeping in
mind the following restrictions:

- You can use port numbers in the unprivileged range, 1024 to 65535.

- The port should not already be used by something else.

Note that changing the port number is a somewhat dangerous operation. If the
port number you select is unusable for whatever reason, Syncthing will not
be able to present its GUI and you will need to locate and manually edit the
configuration file to rectify the situation. Changing the port number on a
remote Syncthing installation is not recommended, unless you have other
means of access as well.

To use a port number lower than 1024, you will need to:

- Ensure that Syncthing has the required privilege to open the port. How to
  accomplish this depends on your operating system - please refer to the
  relevant operation system documentation. Keep in mind that Syncthing should
  not, in general, run as a privileged user (``root``, ``SYSTEM``, etc).

- Use the advanced config editor or edit the configuration file to set the
  port number.

We do not recommend using a port number lower than 1024.
