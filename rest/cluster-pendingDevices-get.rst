GET /rest/cluster/pendingDevices
================================

.. versionadded:: 1.FIXME

Lists remote devices which have tried to connect, but are not yet
configured in our instance.

.. code-block:: json

    [
      {
	"deviceID": "...",
	"time": "2020-03-18T11:43:07+01:00",
	"name": "Friend Joe",
	"address": "tcp://192.168.1.2:22000"
      }
    ]
