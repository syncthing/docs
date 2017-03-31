.. _local-change-detected:

LocalChangeDetected
-------------------

Generated upon scan whenever the local disk has discovered an updated file from the
previous scan.  This does *not* include events that are discovered and copied from
other devices (:ref:`remote-change-detected`), only files that were changed on the
local filesystem.

.. code-block:: json

  {
    "id": 7,
    "globalID": 59,
    "time": "2016-09-26T22:07:10.7189141-04:00",
    "type": "LocalChangeDetected",
    "data": {
      "action": "deleted",
      "folderID": "vitwy-zjxqt",
      "label": "TestSync",
      "path": "C:\\Users\\Nate\\Sync\\testfolder\\test file.rtf",
      "type": "file"
    }
  }
