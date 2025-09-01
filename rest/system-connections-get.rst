GET /rest/system/connections
============================


.. note:: Return format changed in versions 0.13.0, 0.14.14, 1.2.0, 1.19.0, 1.23.0 and 1.25.0.


Returns the list of configured devices and some metadata associated
with them.

The connection types are ``tcp-client``, ``tcp-server``, ``relay-client``, ``relay-server``, ``quic-client`` and ``quic-server``.

.. code-block:: json

    {
      "connections": {
	"DOVII4U-SQEEESM-VZ2CVTC-CJM4YN5-QNV7DCU-5U3ASRL-YVFG6TH-W5DV5AA": {
	  "address": "",
	  "at": "0001-01-01T00:00:00Z",
	  "clientVersion": "",
	  "connected": false,
	  "inBytesTotal": 0,
	  "isLocal": false,
	  "outBytesTotal": 0,
	  "paused": false,
	  "startedAt": "0001-01-01T00:00:00Z",
	  "type": ""
	},
	"UYGDMA4-TPHOFO5-2VQYDCC-7CWX7XW-INZINQT-LE4B42N-4JUZTSM-IWCSXA4": {
	  "address": "",
	  "at": "0001-01-01T00:00:00Z",
	  "clientVersion": "",
	  "connected": false,
	  "inBytesTotal": 0,
	  "isLocal": false,
	  "outBytesTotal": 0,
	  "paused": false,
	  "startedAt": "0001-01-01T00:00:00Z",
	  "type": ""
	},
	"YZJBJFX-RDBL7WY-6ZGKJ2D-4MJB4E7-ZATSDUY-LD6Y3L3-MLFUYWE-AEMXJAC": {
	  "address": "127.0.0.1:22002",
	  "at": "2015-11-07T17:29:47.691548971+01:00",
	  "clientVersion": "v0.12.1",
	  "connected": true,
	  "inBytesTotal": 556,
	  "isLocal": true,
	  "outBytesTotal": 550,
	  "paused": false,
	  "startedAt": "2015-11-07T00:09:47Z",
	  "type": "tcp-client"
	}
      },
      "total": {
	"at": "2015-11-07T17:29:47.691637262+01:00",
	"inBytesTotal": 1479,
	"outBytesTotal": 1318
      }
    }
