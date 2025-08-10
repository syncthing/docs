GET /rest/system/loglevels
==========================

.. versionadded:: 2.0.0

Returns the set of log facilities and their current log level.

.. code-block:: json

    {
      "levels": {
        "api": "INFO",
        "beacon": "INFO",
        ...
        "versioner": "INFO",
        "watchaggregator": "INFO"
      },
      "packages": {
        "api": "REST API",
        "beacon": "Multicast and broadcast discovery",
        ...
        "versioner": "File versioning",
        "watchaggregator": "Filesystem event watcher"
      }
    }
