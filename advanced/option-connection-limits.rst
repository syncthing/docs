Connection Limits
=================

.. versionadded:: 1.13.0

Connection limits can be set to prevent Syncthing from establishing
connections under some circumstances. For the most part you want Syncthing
to connect to all its configured peers, thus you should usually *not* use
this option. Connection limits are useful in specific scenarios concerning
large deployments only, and care must be taken when selecting limits to work
in that scenario. The general recommendation is to leave these settings at
their default of zero, meaning unlimited.

Mechanism
---------

There are two limits, called ``enough`` and ``max``. In short, once there
are *enough* connections Syncthing will stop trying to connect to other
devices; when the *max* is reached Syncthing will also refuse incoming
connections.

Either can be set individually, leaving the other at zero, or both can be
set. When setting both values, ``enough`` should be smaller than ``max`` or
it will have no effect (Syncthing will also stop connecting outwards once
``max`` is reached).

Scenarios
---------

Load Balancing
~~~~~~~~~~~~~~

Consider a setup with a handful "servers" and many "clients". The servers
are fully connected amongst each other and essentially equivalent. The
clients should connect to one server each in order to receive updates. One
way of accomplishing this is to divide the clients into (static) groups and
configure each group to connect to a specific server.

Another way is using connection limits, configuring each client identically
for all servers but setting the ``max`` connection limit to ``1``. This has
the advantage that if one server becomes unavailable the clients will
migrate to other servers.

When establishing new connections Syncthing will preferentially connect to
devices it was recently connected to, thus clients will usually stay on the
same server over a restart.
