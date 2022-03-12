PendingDevicesChanged
---------------------

.. versionadded:: 1.14.0

Emitted when pending devices were added / updated (connection from
unknown ID) or removed (device is ignored, dismissed or added).

.. code-block:: json

    {
      "id": 87,
      "type": "PendingDevicesChanged",
      "time": "2020-12-22T22:24:37.578586718+01:00",
      "data": {
	"added": [
	  {
	    "address": "127.0.0.1:51807",
	    "deviceID": "EJHMPAQ-OGCVORE-ISB4IS3-SYYVJXF-TKJGLTU-66DIQPF-GJ5D2GX-GQ3OWQK",
	    "name": "My dusty computer"
	  }
	],
	"removed": [
	  {
	    "deviceID": "P56IOI7-MZJNU2Y-IQGDREY-DM2MGTI-MGL3BXN-PQ6W5BM-TBBZ4TJ-XZWICQ2"
	  }
	]
      }
    }
