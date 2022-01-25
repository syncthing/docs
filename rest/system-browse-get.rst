GET /rest/system/browse
=======================

Returns a list of directories matching the path given by the optional parameter
``current``. The path can use `patterns as described in Go's filepath package
<https://pkg.go.dev/path/filepath#Match>`_. A '*' will always be appended
to the given path (e.g. ``/tmp/`` matches all its subdirectories). If the option
``current`` is not given, filesystem root paths are returned.

.. code-block:: bash

    $ curl -H "X-API-Key: yourkey" localhost:8384/rest/system/browse | json_pp
    [
        "/"
    ]

    $ curl -H "X-API-Key: yourkey" localhost:8384/rest/system/browse?current=/var/ | json_pp
    [
        "/var/backups/",
        "/var/cache/",
        "/var/lib/",
        "/var/local/",
        "/var/lock/",
        "/var/log/",
        "/var/mail/",
        "/var/opt/",
        "/var/run/",
        "/var/spool/",
        "/var/tmp/"
    ]

    $ curl -H "X-API-Key: yourkey" localhost:8384/rest/system/browse?current=/var/*o | json_pp
    [
        "/var/local/",
        "/var/lock/",
        "/var/log/",
        "/var/opt/",
        "/var/spool/"
    ]
