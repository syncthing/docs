LocalChangeDetected
-------------------

Generated upon scan whenever the local disk has discovered an updated file from the
previous scan.  This does *not* include events that are discovered and copied from
other devices (:doc:`remotechangedetected`), only files that were changed on the
local filesystem.

.. note:: This event is not included in :doc:`/rest/events-get` endpoint without
   a mask specified, but needs to be selected explicitly.

.. code-block:: json

  {
    "id": 7,
    "globalID": 59,
    "time": "2016-09-26T22:07:10.7189141-04:00",
    "type": "LocalChangeDetected",
    "data": {
      "action": "deleted",
      "folder": "vitwy-zjxqt",
      "folderID": "vitwy-zjxqt",
      "label": "TestSync",
      "path": "test file.rtf",
      "type": "file"
    }
  }

.. deprecated:: v1.1.2
  The ``folderID`` field is a legacy name kept only for compatibility.  Use the
  ``folder`` field with identical content instead.
