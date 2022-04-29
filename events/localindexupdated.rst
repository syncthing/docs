LocalIndexUpdated
-----------------

Generated when the local index information has changed, due to
synchronizing one or more items from the cluster or discovering local
changes during a scan.

.. code-block:: json

    {
        "id": 59,
        "globalID": 59,
        "type": "LocalIndexUpdated",
        "time": "2014-07-17T13:27:28.051369434+02:00",
        "data": {
            "folder": "default",
            "items": 1000,
            "filenames": [
                "foo",
                "bar",
                "baz"
            ],
            "sequence": 12345,
            "version": 12345
        }
    }

.. deprecated:: v1.10.0
  The ``version`` field is a legacy name kept only for compatibility.  Use the
  ``sequence`` field with identical content instead.
