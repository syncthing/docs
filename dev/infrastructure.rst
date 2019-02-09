Syncthing Infrastructure
========================

This is a list of the infrastructure that powers the Syncthing project.
Unless otherwise noted, the default is that it's a VM hosted by :user:`calmh`.

GitHub
------

All repos, issue trackers and binary releases are hosted at `GitHub <https://github.com/syncthing>`__.

Main & Documentation Websites
------------------------------

Static HTML, served by Caddy.

- `syncthing.net <https://syncthing.net/>`__
- `apt.syncthing.net <https://apt.syncthing.net>`__ (Debian packages)
- `docs.syncthing.net <https://docs.syncthing.net/>`__ (Sphinx for site generation)
- `upgrades.syncthing.net <https://upgrades.syncthing.net/meta.json>`__ (upgrade metadata)

Forum Website
-------------

Powered by Discourse.

- `forum.syncthing.net <https://forum.syncthing.net/>`__

Global Discovery Servers
------------------------

Runs the ``stdiscosrv`` instances that serve global discovery requests. The
discovery setup is a load balanced cluster and the members can change
without prior notice. As of the time of writing they are all hosted at
DigitalOcean.

- discovery.syncthing.net (multiple A and AAAA records, for queries)
- discovery-v4.syncthing.net (multiple A records, for IPv4 announces)
- discovery-v6.syncthing.net (multiple AAAA records, for IPv6 announces)

Relay Pool Server
-----------------

Runs the `strelaypoolsrv <https://github.com/syncthing/syncthing/tree/master/cmd/strelaypoolsrv>`__
daemon to handle dynamic registration and announcement of relays.

- `relays.syncthing.net <http://relays.syncthing.net>`__

Relay Servers
-------------

Hosted by friendly people on the internet.

Usage Reporting Server
----------------------

Runs the `ursrv <https://github.com/syncthing/syncthing/tree/master/cmd/ursrv>`__
daemon with PostgreSQL and Nginx.

- `data.syncthing.net <https://data.syncthing.net/>`__

Build Servers
-------------

Runs TeamCity and does the core builds.

- `build.syncthing.net <https://build.syncthing.net/>`__

There are various build agents; Linux, Windows, and Mac. These are currently
provided by :user:`calmh` or Kastelo.


Signing Server
--------------

Signs and uploads the release bundles to GitHub.

- secure.syncthing.net

Monitoring
----------

The infrastructure is monitored and its status is publicly accessible on the following urls:

- `status.syncthing.net <http://status.syncthing.net>`__ (Apex Ping)
- `monitor.syncthing.net <https://monitor.syncthing.net>`__ (Grafana)

