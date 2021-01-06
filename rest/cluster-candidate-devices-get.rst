GET /rest/cluster/candidate/devices
===================================

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

.. code-block:: json

    {
      "DOVII4U-SQEEESM-VZ2CVTC-CJM4YN5-QNV7DCU-5U3ASRL-YVFG6TH-W5DV5AA": {
	"introducedBy": {
	  "YZJBJFX-RDBL7WY-6ZGKJ2D-4MJB4E7-ZATSDUY-LD6Y3L3-MLFUYWE-AEMXJAC": {
	    "commonFolders": {
	      "cpkn4-57ysy": "Pictures from Joe"
	    },
	    "time": "2020-03-18T11:43:07+01:00"
	  }
	}
      },
      "P56IOI7-MZJNU2Y-IQGDREY-DM2MGTI-MGL3BXN-PQ6W5BM-TBBZ4TJ-XZWICQ2": {
	"addresses": [
	  "tcp://192.168.1.2:22000",
	  "tcp://[2a02:8070::ff34:1234::aabb]:22000",
	  "tcp://janes.laptop.example.com:22000",
	  "dynamic"
	],
	"certName": "",
	"introducedBy": {
	  "DOVII4U-SQEEESM-VZ2CVTC-CJM4YN5-QNV7DCU-5U3ASRL-YVFG6TH-W5DV5AA": {
	    "commonFolders": {
	      "cpkn4-57ysy": "Pics from Jane"
	    },
	    "suggestedName": "Jane",
	    "time": "2020-03-01T10:12:13+01:00"
	  },
	  "YZJBJFX-RDBL7WY-6ZGKJ2D-4MJB4E7-ZATSDUY-LD6Y3L3-MLFUYWE-AEMXJAC": {
	    "commonFolders": {
	      "cpkn4-57ysy": "Pics of J & J"
	    },
	    "suggestedName": "Jane's Laptop",
	    "time": "2020-03-18T11:43:07+01:00"
	  }
	}
      },
      "UYGDMA4-TPHOFO5-2VQYDCC-7CWX7XW-INZINQT-LE4B42N-4JUZTSM-IWCSXA4": {
	"introducedBy": {
	  "AIBAEAQ-CAIBAEC-AQCAIBA-EAQCAIA-BAEAQCA-IBAEAQC-CAIBAEA-QCAIBA7": {
	    "commonFolders": {
	      "cpkn4-57ysy": "Family pics"
	    },
	    "time": "2020-02-22T14:56:00+01:00"
	  },
	  "YZJBJFX-RDBL7WY-6ZGKJ2D-4MJB4E7-ZATSDUY-LD6Y3L3-MLFUYWE-AEMXJAC": {
	    "commonFolders": {
	      "abcde-fghij": "Mighty nice folder",
	      "cpkn4-57ysy": "Family pics"
	    },
	    "time": "2020-03-18T11:43:07+01:00"
	  }
	}
      }
    }
