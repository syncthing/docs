GET /rest/cluster/pending/folders
=================================

.. versionadded:: 1.13.0

Lists folders which remote devices have offered to us, but are not yet
shared from our instance to them.  Takes the optional ``device``
parameter to only return folders offered by a specific remote device.
Other offering devices are also omitted from the result.

.. code-block:: json

    {
      "cpkn4-57ysy": {
        "offeredBy": {
          "P56IOI7-MZJNU2Y-IQGDREY-DM2MGTI-MGL3BXN-PQ6W5BM-TBBZ4TJ-XZWICQ2": {
            "time": "2020-03-18T11:43:07Z",
            "label": "Joe's folder",
            "receiveEncrypted": true,
            "remoteEncrypted": false
          },
          "DOVII4U-SQEEESM-VZ2CVTC-CJM4YN5-QNV7DCU-5U3ASRL-YVFG6TH-W5DV5AA": {
            "time": "2020-03-01T10:12:13Z",
            "label": "Jane's and Joe's folder",
            "receiveEncrypted": false,
            "remoteEncrypted": false
          }
        }
      },
      "abcde-fghij": {
        "offeredBy": {
          "P56IOI7-MZJNU2Y-IQGDREY-DM2MGTI-MGL3BXN-PQ6W5BM-TBBZ4TJ-XZWICQ2": {
            "time": "2020-03-18T11:43:07Z",
            "label": "MyPics",
            "receiveEncrypted": false,
            "remoteEncrypted": false
          }
        }
      }
    }
