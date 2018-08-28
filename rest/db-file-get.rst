GET /rest/db/file
=================

Returns most data available about a given file, including version and
availability. Takes ``folder`` and ``file`` parameters.

.. code-block:: json

    {
      "availability": [
        {
          "id": "ITZRNXE-YNROGBZ-HXTH5P7-VK5NYE5-QHRQGE2-7JQ6VNJ-KZUEDIU-5PPR5AM",
          "fromTemporary": false
        }
      ],
      "global": {
        "deleted": false,
        "ignored": false,
        "invalid": false,
        "localFlags": 0,
        "modified": "2018-08-18T12:21:13.836784059+02:00",
        "modifiedBy": "SYNO4VL",
        "mustRescan": false,
        "name": "testfile",
        "noPermissions": false,
        "numBlocks": 1,
        "permissions": "0755",
        "sequence": 107499,
        "size": 1234,
        "type": 0,
        "version": [
          "SYNO4VL:1"
        ]
      },
      "local": {
        "deleted": false,
        "ignored": false,
        "invalid": false,
        "localFlags": 0,
        "modified": "2018-08-18T12:21:13.836784059+02:00",
        "modifiedBy": "SYNO4VL",
        "mustRescan": false,
        "name": "testfile",
        "noPermissions": false,
        "numBlocks": 1,
        "permissions": "0755",
        "sequence": 111038,
        "size": 1234,
        "type": 0,
        "version": [
          "SYNO4VL:1"
        ]
      }
    }

