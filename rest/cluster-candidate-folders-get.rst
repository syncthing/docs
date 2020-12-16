GET /rest/cluster/candidate/folders
===================================

.. versionadded:: 1.FIXME

Lists folder IDs that we do not currently share with a known remote
device directly, but which they already have indirect access to.  That
information was collected as part of other remote devices'
``ClusterConfig`` messages, which share the specific folder with both
our instance and the candidate device.  See :ref:`clusterconfig` for
what information is available about such devices.  They share a folder
with us already, but only indirectly over at least one extra "hop",
the introducing device.  Adding the missing direct link densifies the
cluster's mesh structure without giving additional access to anyone.

Given a known device ID in the ``device`` parameter, the list is
limited to only folder IDs for which that device is a candidate.

.. code-block:: json

    {
      "abcde-fghij": {
	"UYGDMA4-TPHOFO5-2VQYDCC-7CWX7XW-INZINQT-LE4B42N-4JUZTSM-IWCSXA4": {
	  "YZJBJFX-RDBL7WY-6ZGKJ2D-4MJB4E7-ZATSDUY-LD6Y3L3-MLFUYWE-AEMXJAC": {
	    "label": "Mighty nice folder",
	    "time": "2020-03-18T11:43:07+01:00"
	  }
	}
      },
      "cpkn4-57ysy": {
	"DOVII4U-SQEEESM-VZ2CVTC-CJM4YN5-QNV7DCU-5U3ASRL-YVFG6TH-W5DV5AA": {
	  "YZJBJFX-RDBL7WY-6ZGKJ2D-4MJB4E7-ZATSDUY-LD6Y3L3-MLFUYWE-AEMXJAC": {
	    "label": "Pictures from Joe",
	    "time": "2020-03-18T11:43:07+01:00"
	  }
	},
	"UYGDMA4-TPHOFO5-2VQYDCC-7CWX7XW-INZINQT-LE4B42N-4JUZTSM-IWCSXA4": {
	  "AIBAEAQ-CAIBAEC-AQCAIBA-EAQCAIA-BAEAQCA-IBAEAQC-CAIBAEA-QCAIBA7": {
	    "label": "Family pics",
	    "time": "2020-11-22T14:56:00+01:00"
	  },
	  "YZJBJFX-RDBL7WY-6ZGKJ2D-4MJB4E7-ZATSDUY-LD6Y3L3-MLFUYWE-AEMXJAC": {
	    "label": "Family pics",
	    "time": "2020-03-18T11:43:07+01:00"
	  }
	}
      }
    }
