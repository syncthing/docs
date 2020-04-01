GET /rest/cluster/pendingFolders
================================

.. versionadded:: 1.FIXME

Lists folders which remote devices have offered to us, but are not yet
shared from our instance to them.  Takes the optional ``device``
parameter to only return folders offered by a specific remote device.

.. code-block:: json

    [
      {
	"id": "cpkn4-57ysy",
	"offeredBy": [
	  {
	    "deviceID": "...",
	    "time": "2020-03-18T11:43:07+01:00",
	    "label": "Joe's folder"
	  },
	  {
	    "deviceID": "...",
	    "time": "2020-03-01T10:12:13+01:00",
	    "label": "Jane's and Joe's folder"
	  }
	]
      },
      {
	"id": "abcde-fghij",
	"offeredBy": [
	  {
	    "deviceID": "...",
	    "time": "2020-03-18T11:43:07+01:00",
	    "label": "MyPics"
	  }
	]
      }
    ]
