GET /rest/system/config
=======================

Returns the current configuration.

.. code-block:: json

    {
      "version": 30,
      "folders": [
	{
	  "id": "GXWxf-3zgnU",
	  "label": "MyFolder",
	  "filesystemType": "basic",
	  "path": "...",
	  "type": "sendreceive",
	  "devices": [
	    {
	      "deviceID": "...",
	      "introducedBy": ""
	    }
	  ],
	  "rescanIntervalS": 60,
	  "fsWatcherEnabled": false,
	  "fsWatcherDelayS": 10,
	  "ignorePerms": false,
	  "autoNormalize": true,
	  "minDiskFree": {
	    "value": 1,
	    "unit": "%"
	  },
	  "versioning": {
	    "type": "simple",
	    "params": {
	      "keep": "5"
	    }
	  },
	  "copiers": 0,
	  "pullerMaxPendingKiB": 0,
	  "hashers": 0,
	  "order": "random",
	  "ignoreDelete": false,
	  "scanProgressIntervalS": 0,
	  "pullerPauseS": 0,
	  "maxConflicts": 10,
	  "disableSparseFiles": false,
	  "disableTempIndexes": false,
	  "paused": false,
	  "weakHashThresholdPct": 25,
	  "markerName": ".stfolder",
	  "copyOwnershipFromParent": false,
	  "modTimeWindowS": 0
	}
      ],
      "devices": [
	{
	  "deviceID": "...",
	  "name": "Laptop",
	  "addresses": [
	    "dynamic",
	    "tcp://192.168.1.2:22000"
	  ],
	  "compression": "metadata",
	  "certName": "",
	  "introducer": false,
	  "skipIntroductionRemovals": false,
	  "introducedBy": "",
	  "paused": false,
	  "allowedNetworks": [],
	  "autoAcceptFolders": false,
	  "maxSendKbps": 0,
	  "maxRecvKbps": 0,
	  "ignoredFolders": [],
	  "pendingFolders": [
	    {
	      "time": "2019-06-05T10:21:22+02:00",
	      "id": "cpkn4-57ysy",
	      "label": "SomeonesFolder"
	    }
	  ],
	  "maxRequestKiB": 0
	}
      ],
      "gui": {
	"enabled": true,
	"address": "127.0.0.1:8384",
	"user": "Username",
	"password": "$2a$10$ZFws69T4FlvWwsqeIwL.TOo5zOYqsa/.TxlUnsGYS.j3JvjFTmxo6",
	"authMode": "static",
	"useTLS": false,
	"apiKey": "pGahcht56664QU5eoFQW6szbEG6Ec2Cr",
	"insecureAdminAccess": false,
	"theme": "default",
	"debugging": false,
	"insecureSkipHostcheck": false,
	"insecureAllowFrameLoading": false
      },
      "ldap": {
	"addresd": "",
	"bindDN": "",
	"transport": "plain",
	"insecureSkipVerify": false
      },
      "options": {
	"listenAddresses": [
	  "default"
	],
	"globalAnnounceServers": [
	  "default"
	],
	"globalAnnounceEnabled": true,
	"localAnnounceEnabled": true,
	"localAnnouncePort": 21027,
	"localAnnounceMCAddr": "[ff12::8384]:21027",
	"maxSendKbps": 0,
	"maxRecvKbps": 0,
	"reconnectionIntervalS": 60,
	"relaysEnabled": true,
	"relayReconnectIntervalM": 10,
	"startBrowser": false,
	"natEnabled": true,
	"natLeaseMinutes": 60,
	"natRenewalMinutes": 30,
	"natTimeoutSeconds": 10,
	"urAccepted": -1,
	"urSeen": 2,
	"urUniqueId": "",
	"urURL": "https://data.syncthing.net/newdata",
	"urPostInsecurely": false,
	"urInitialDelayS": 1800,
	"restartOnWakeup": true,
	"autoUpgradeIntervalH": 12,
	"upgradeToPreReleases": false,
	"keepTemporariesH": 24,
	"cacheIgnoredFiles": false,
	"progressUpdateIntervalS": 5,
	"limitBandwidthInLan": false,
	"minHomeDiskFree": {
	  "value": 1,
	  "unit": "%"
	},
	"releasesURL": "https://upgrades.syncthing.net/meta.json",
	"alwaysLocalNets": [],
	"overwriteRemoteDeviceNamesOnConnect": false,
	"tempIndexMinBlocks": 10,
	"unackedNotificationIDs": [],
	"trafficClass": 0,
	"defaultFolderPath": "~",
	"setLowPriority": true,
	"maxFolderConcurrency": 0,
	"crURL": "https://crash.syncthing.net/newcrash",
	"crashReportingEnabled": true,
	"stunKeepaliveStartS": 180,
	"stunKeepaliveMinS": 20,
	"stunServers": [
	  "default"
	],
	"databaseTuning": "auto",
	"maxConcurrentIncomingRequestKiB": 0
      },
      "remoteIgnoredDevices": [],
      "pendingDevices": []
    }
