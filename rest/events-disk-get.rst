GET /rest/events/disk
=====================

Returns local disk events that occur when the scanner detects local file system
changes (:ref:`local-change-detected`) or when files are pulled from a remote
device (:ref:`remote-change-detected`).

Optional GET parameters:
 - since (events starting after the given ID)
 - timeout (fail after given seconds if no event is available, 2s by default)
 - limit (return last x number of events)

.. code-block:: bash

    $ curl -s http://localhost:8384/rest/events/disk?limit=4 | json
   {
    "id": 4,
    "globalID": 45,
    "time": "2016-09-26T22:06:10.4734536-04:00",
    "type": "LocalChangeDetected",
    "data": {
      "action": "added",
      "folderID": "vitwy-zxuqt",
      "label": "TestSync",
      "path": "C:\\Users\\Nate\\Sync\\testfolder",
      "type": "dir"
    }
  },
  {
    "id": 5,
    "globalID": 46,
    "time": "2016-09-26T22:06:10.4754548-04:00",
    "type": "LocalChangeDetected",
    "data": {
      "action": "added",
      "folderID": "vitwy-zxuqt",
      "label": "TestSync",
      "path": "C:\\Users\\Nate\\Sync\\dfghdfj\\test file.rtf",
      "type": "file"
    }
  },
  {
    "id": 6,
    "globalID": 58,
    "time": "2016-09-26T22:07:10.7189141-04:00",
    "type": "LocalChangeDetected",
    "data": {
      "action": "deleted",
      "folderID": "vitwy-zxuqt",
      "label": "TestSync",
      "path": "C:\\Users\\Nate\\Sync\\testfolder",
      "type": "dir"
    }
  },
  {
    "id": 7,
    "globalID": 59,
    "time": "2016-09-26T22:07:10.7189141-04:00",
    "type": "LocalChangeDetected",
    "data": {
      "action": "deleted",
      "folderID": "vitwy-zxuqt",
      "label": "TestSync",
      "path": "C:\\Users\\Nate\\Sync\\dfghdfj\\test file.rtf",
      "type": "file"
    }
  }
