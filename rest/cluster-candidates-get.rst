GET /rest/cluster/candidates
============================

.. versionadded:: 1.FIXME

Lists remote devices which were introduced as part of another remote
device's ``ClusterConfig`` message, but do not directly share a
specific folder with our instance.  See :ref:`clusterconfig` for what
information is available about such devices.  They share a folder with
us already, but only indirectly over at least one extra "hop", the
introducing device.  Adding the missing direct link densifies the
cluster's mesh structure without giving additional access to anyone.

Given an ID in the ``folder`` parameter, candidate devices are
enumerated where it is only shared indirectly.  Any device which is
already known locally is referenced only by its device ID.  Unknown
device entries are supplemented with metadata from the introducing
devices.

Alternatively, a device ID can be given in the ``device`` parameter,
in which case only folder IDs are enumerated for which that is a
candidate device.

.. code-block:: json

    [
      {
	"deviceID": "...",
	"certName": "",
	"addresses": [
	  "192.168.1.2:22000",
	  "[2a02:8070:::ff34:1234::aabb]:22000"
	],
	"introducedBy": [
	  {
	    "time": "2020-03-18T11:43:07+01:00",
	    "deviceID": "YZJBJFX-RDBL7WY-6ZGKJ2D-4MJB4E7-ZATSDUY-LD6Y3L3-MLFUYWE-AEMXJAC",
	    "suggestedName": "Jane's Laptop"
	  },
	  {
	    "time": "2020-03-01T10:12:13+01:00",
	    "deviceID": "DOVII4U-SQEEESM-VZ2CVTC-CJM4YN5-QNV7DCU-5U3ASRL-YVFG6TH-W5DV5AA",
	    "suggestedName": "Jane"
	  },
	]
      },
      {
	"deviceID": "UYGDMA4-TPHOFO5-2VQYDCC-7CWX7XW-INZINQT-LE4B42N-4JUZTSM-IWCSXA4"
      }
    ]
