GET /rest/system/status
=======================

Returns information about current system status and resource usage. The CPU percent value has been deprected from the API and will always report 0.

.. code-block:: json

    {
      "alloc": 30618136,
      "connectionServiceStatus": {
        "dynamic+https://relays.syncthing.net/endpoint": {
          "error": null,
          "lanAddresses": [
            "relay://23.92.71.120:443/?id=53STGR7-YBM6FCX-PAZ2RHM-YPY6OEJ-WYHVZO7-PCKQRCK-PZLTP7T-434XCAD&pingInterval=1m0s&networkTimeout=2m0s&sessionLimitBps=0&globalLimitBps=0&statusAddr=:22070&providedBy=canton7"
          ],
          "wanAddresses": [
            "relay://23.92.71.120:443/?id=53STGR7-YBM6FCX-PAZ2RHM-YPY6OEJ-WYHVZO7-PCKQRCK-PZLTP7T-434XCAD&pingInterval=1m0s&networkTimeout=2m0s&sessionLimitBps=0&globalLimitBps=0&statusAddr=:22070&providedBy=canton7"
          ]
        },
        "tcp://0.0.0.0:22000": {
          "error": null,
          "lanAddresses": [
            "tcp://0.0.0.0:22000"
          ],
          "wanAddresses": [
            "tcp://0.0.0.0:22000"
          ]
        }
      },
      "cpuPercent": 0,
      "discoveryEnabled": true,
      "discoveryErrors": {
        "global@https://discovery-v4-1.syncthing.net/v2/": "500 Internal Server Error",
        "global@https://discovery-v4-2.syncthing.net/v2/": "Post https://discovery-v4-2.syncthing.net/v2/: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)",
        "global@https://discovery-v4-3.syncthing.net/v2/": "Post https://discovery-v4-3.syncthing.net/v2/: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)",
        "global@https://discovery-v6-1.syncthing.net/v2/": "Post https://discovery-v6-1.syncthing.net/v2/: dial tcp [2001:470:28:4d6::5]:443: connect: no route to host",
        "global@https://discovery-v6-2.syncthing.net/v2/": "Post https://discovery-v6-2.syncthing.net/v2/: dial tcp [2604:a880:800:10::182:a001]:443: connect: no route to host",
        "global@https://discovery-v6-3.syncthing.net/v2/": "Post https://discovery-v6-3.syncthing.net/v2/: dial tcp [2400:6180:0:d0::d9:d001]:443: connect: no route to host"
      },
      "discoveryMethods": 8,
      "goroutines": 49,
      "lastDialStatus": {
          "tcp://10.20.30.40": {
            "when": "2019-05-16T07:41:23Z",
            "error": "dial tcp 10.20.30.40:22000: i/o timeout"
          },
          "tcp://172.16.33.3:22000": {
            "when": "2019-05-16T07:40:43Z",
            "ok": true
          },
          "tcp://83.233.120.221:22000": {
            "when": "2019-05-16T07:41:13Z",
            "error": "dial tcp 83.233.120.221:22000: connect: connection refused"
          }
      },
      "myID": "P56IOI7-MZJNU2Y-IQGDREY-DM2MGTI-MGL3BXN-PQ6W5BM-TBBZ4TJ-XZWICQ2",
      "pathSeparator": "/",
      "startTime": "2016-06-06T19:41:43.039284753+02:00",
      "sys": 42092792,
      "themes": [
        "default",
        "dark"
      ],
      "tilde": "/Users/jb",
      "uptime": 2635
    }

.. versionadded:: 1.2.0

  The ``lastDialStatus`` dictionary contains the last error (or ``null`` for
  success) for each peer address that Syncthing has attempted to contact.
  The ``connectionServiceStatus`` entries gained ``"error": null``
  attributes where previously there would be no ``error`` attribute at all
  in the success case.
