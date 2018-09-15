Syncthing Infrastructure
========================

This is a list of the infrastructure that powers the Syncthing project.
Unless otherwise noted, the default is that it's a VM hosted by :user:`calmh`.

GitHub
------

All repos, issue trackers and binary releases are hosted at `GitHub <https://github.com/syncthing>`__.

Main & Documentation Websites
------------------------------

Static HTML, served by Nginx.

- `syncthing.net <https://syncthing.net/>`__ (SmartOS container, 1024 MB)
- `docs.syncthing.net <https://docs.syncthing.net/>`__ (Sphinx for site generation)

Forum Website
-------------

Powered by Discourse.

- `forum.syncthing.net <https://forum.syncthing.net/>`__ (Ubuntu Linux, 2048 MB)

Upgrades Server
---------------

The ``upgrades.syncthing.net`` server is a virtualhost on ``syncthing.net``.
And the ``meta.json`` file is just a regular synchronized
static copy of github releases. The actual signed binary releases are hosted on github.

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

Runs the `relaypoolsrv <https://github.com/syncthing/syncthing/tree/master/cmd/strelaypoolsrv>`__
daemon to handle dynamic registration and announcement of relays.

- `relays.syncthing.net <http://relays.syncthing.net>`__ (SmartOS container, 256 MB)

Relay Servers
-------------

Hosted by friendly people on the internet.

Usage Reporting Server
----------------------

Runs the `ursrv <https://github.com/syncthing/usage-reporting/tree/master/cmd/ursrv>`__
daemon with PostgreSQL and Nginx.

- `data.syncthing.net <https://data.syncthing.net/>`__ (Ubuntu Linux, 512 MB)

Build Servers, Core
-------------------

Runs TeamCity and does the core builds, Ubuntu Linux.

- `build.syncthing.net <https://build.syncthing.net/>`__ (TeamCity frontend, SmartOS container, 2048 MB)

Build Servers, Android and Other
--------------------------------

Runs a Gradle daemon building releases for Android.

- `Travis CI <https://travis-ci.org/syncthing/syncthing-android>`__

Runs a Jenkins slave and builds release builds for their respective OS.

- native-windows-amd64 (Windows 2012 R2, Amazon EC2 t2.micro)
- native-darwin-amd64 (Mac OS X 10.8.5, MacBookPro5,1 in calmh's garage)
- native-ubuntu-amd64 (Ubuntu 14.04, 2048 MB, hosted at DigitalOcean)
- native-ubuntu-386 (Ubuntu 14.04, 512 MB, hosted at DigitalOcean)

APT Server
----------

Serves the APT repository for Debian/Ubuntu users. Runs Nginx.

- `apt.syncthing.net <https://apt.syncthing.net>`__ (SmartOS container, 256 MB)

Signing Server
--------------

Signs and uploads the release bundles to GitHub.

- secure.syncthing.net (SmartOS container, 2048 MB)

Monitoring
----------

The infrastructure is monitored and its status is publicly accessible on the following urls:

- `status.syncthing.net <https://status.syncthing.net>`__
- `monitor.syncthing.net <https://monitor.syncthing.net>`__ (Grafana)
