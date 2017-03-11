GET /rest/system/browse
===================

Returns a list of directories matching the path given by the optional parameter
``current``. The path can use `patterns as described in Go's filepath package
<https://golang.org/pkg/path/filepath/#Match>`_. It is also possible to list a
directory's contents by adding a trailing path separator.

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
