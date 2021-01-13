.. _pending-folders-changed:

PendingFoldersChanged
---------------------

.. versionadded:: 1.14.0

Emitted when pending folders were added / updated (offered by some
device, but not shared to them) or removed (folder ignored or added or
no longer offered from the remote device).

.. code-block:: json

    {
      "id": 101,
      "type": "PendingFoldersChanged",
      "time": "2020-12-22T22:36:55.66744317+01:00",
      "data": {
	"added": [
	  {
	    "device": "EJHMPAQ-OGCVORE-ISB4IS3-SYYVJXF-TKJGLTU-66DIQPF-GJ5D2GX-GQ3OWQK",
	    "folder": "GXWxf-3zgnU",
	    "folderLabel": "My Pictures"
	  }
	],
	"removed": [
	  {
	    "device": "P56IOI7-MZJNU2Y-IQGDREY-DM2MGTI-MGL3BXN-PQ6W5BM-TBBZ4TJ-XZWICQ2",
	    "folder": "neyfh-sa2nu"
	  }
	]
      }
    }
