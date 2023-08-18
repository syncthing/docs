Prometheus-Style Metrics
========================

Syncthing provides an endpoint for Prometheus-style metrics. Metrics are
served on the ``/metrics`` path on the GUI / API address. The metrics endpoint
requires authentication when the GUI / API is configured to require
authentication; see :doc:`/dev/rest` for details.

Metrics
-------

The following metrics are available.

.. include:: ../includes/metrics-list.rst
