GET /rest/db/file
=================

Returns most data available about a given file, including version and
availability. Takes ``folder`` and ``file`` parameters.

.. code-block:: json

    {
      "availability": [{
        "id": "I6KAH76-66SLLLB-5PFXSOA-UFJCDZC-YAOMLEK-CP2GB32-BV5RQST-3PSROAU",
        "fromTemporary": false
      }],
      "global": {
        "deleted": false,
        "invalid": false,
        "flags": "0644",
        "sequence": 3,
        "modified": "2015-04-20T22:20:45+09:00",
        "modifiedBy": "I6KAH76",
        "noPermissions": false,
        "name": "util.go",
        "numBlocks": 1,
        "size": 9642,
        "type": 0,
        "version": [
          "I6KAH76:1"
        ]
      },
      "local": {
        "deleted": false,
        "invalid": false,
        "flags": "0644",
        "sequence": 4,
        "modified": "2015-04-20T22:20:45+09:00",
        "modifiedBy": "I6KAH76",
        "noPermissions": false,
        "name": "util.go",
        "numBlocks": 1,
        "size": 9642,
        "type": 0,
        "version": [
          "I6KAH76:1"
        ]
      }
    }
