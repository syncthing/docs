GET /rest/events
================

Returns general events that have occured since startup of the client.
Events include things like 'StateChanged', 'DeviceDiscovered',
'LocalIndexUpdated', 'Ping', etc...

.. code-block:: bash

    $ curl -s http://localhost:8384/rest/events | json
  {
    "id": 9,
    "globalID": 9,
    "time": "2016-09-26T22:46:02.8266126-04:00",
    "type": "DeviceDiscovered",
    "data": {
      "addrs": [
        "tcp://10.1.3.9:22000"
      ],
      "device": "XXXXXXX-XXXXXXX-XXXXXXX-XXXXXXX-XXXXXXX-XXXXXXX-XXXXXXX-XXXXXXX"
    }
  },
  {
    "id": 10,
    "globalID": 10,
    "time": "2016-09-26T22:46:04.8907271-04:00",
    "type": "DeviceConnected",
    "data": {
      "addr": "10.150.30.9:22000",
      "clientName": "syncthing",
      "clientVersion": "v0.14.7",
      "deviceName": "hostnamehere",
      "id": "XXXXXXX-XXXXXXX-XXXXXXX-XXXXXXX-XXXXXXX-XXXXXXX-XXXXXXX-XXXXXXX",
      "type": "TCP (Client)"
    }
  },
  {
    "id": 11,
    "globalID": 11,
    "time": "2016-09-26T22:46:04.9267302-04:00",
    "type": "StateChanged",
    "data": {
      "duration": 3.8412221,
      "folder": "vitwy-zuxqt",
      "from": "idle",
      "to": "syncing"
    }
  },
...more events
