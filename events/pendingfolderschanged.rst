PendingFoldersChanged
---------------------

.. versionadded:: 1.14.0

Emitted when pending folders were added / updated (offered by some
device, but not shared to them) or removed (folder ignored, dismissed
or added or no longer offered from the remote device).  A removed
entry without a ``deviceID`` attribute means that the folder is no
longer pending for any device.

.. code-block:: json

    {
      "id": 101,
      "type": "PendingFoldersChanged",
      "time": "2020-12-22T22:36:55.66744317+01:00",
      "data": {
	"added": [
	  {
	    "deviceID": "EJHMPAQ-OGCVORE-ISB4IS3-SYYVJXF-TKJGLTU-66DIQPF-GJ5D2GX-GQ3OWQK",
	    "folderID": "GXWxf-3zgnU",
	    "folderLabel": "My Pictures"
	    "receiveEncrypted": "false"
	    "remoteEncrypted": "false"
	  }
	],
	"removed": [
	  {
	    "deviceID": "P56IOI7-MZJNU2Y-IQGDREY-DM2MGTI-MGL3BXN-PQ6W5BM-TBBZ4TJ-XZWICQ2",
	    "folderID": "neyfh-sa2nu"
	  },
	  {
	    "folderID": "abcde-fghij"
	  }
	]
      }
    }
