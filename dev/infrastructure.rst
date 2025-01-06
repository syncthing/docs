Syncthing Infrastructure
========================

This is a list of the infrastructure that powers the Syncthing project.

GitHub
------

All repos, issue trackers and binary releases are hosted at `GitHub <https://github.com/syncthing>`__.
GitHub Actions powers our builds and tests (CI).

Main & Documentation Websites
------------------------------

Websites are published to GitHub Pages.

- `syncthing.net <https://syncthing.net/>`__ (Hugo for site generation)
- `docs.syncthing.net <https://docs.syncthing.net/>`__ (Sphinx for site generation)

General Services
----------------

Several services run in a Kubernetes cluster at Scaleway:

- Crash reporting server
- `Debian/APT packages <https://apt.syncthing.net>`__, with the actual
  packages served from object storage at Scaleway
- Global discovery
- `Relay pool server <https://relays.syncthing.net>`__
- `Roadmap voting site <https://roadmap.syncthing.net>`__
- `Service monitoring (Mimir/Loki/Grafana) <https://mon.syncthing.net>`__
- Upgrade server

The `forum <https://forum.syncthing.net/>`__ is a separate VM, though also at Scaleway.

Relay Servers
-------------

Hosted by friendly people on the internet.

Usage Reporting Server
----------------------

Runs the `ursrv <https://github.com/syncthing/syncthing/tree/main/cmd/ursrv>`__
daemon with PostgreSQL and Nginx.

- `data.syncthing.net <https://data.syncthing.net/>`__

Signing Server
--------------

Signs and uploads the release bundles to GitHub.

- secure.syncthing.net

External Monitoring
-------------------

The infrastructure is monitored and its status is publicly accessible on the following urls:

- `status.syncthing.net <https://status.syncthing.net>`__ (updown.io service)

