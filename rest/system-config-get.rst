GET /rest/system/config (DEPRECATED)
====================================

.. deprecated:: v1.12.0
   This endpoint still works as before but is deprecated. Use :ref:`rest-config`
   instead.

Returns the current configuration.

.. code-block:: json

    {
      "version": 35,
      "folders": [
	{
	  "id": "default",
	  "label": "Default Folder",
	  "filesystemType": "basic",
	  "path": "...",
	  "type": "sendreceive",
	  "devices": [
	    {
	      "deviceID": "...",
	      "introducedBy": "",
	      "encryptionPassword": ""
	    }
	  ],
	  "rescanIntervalS": 3600,
	  "fsWatcherEnabled": true,
	  "fsWatcherDelayS": 10,
	  "ignorePerms": false,
	  "autoNormalize": true,
	  "minDiskFree": {
	    "value": 1,
	    "unit": "%"
	  },
	  "versioning": {
	    "type": "",
	    "params": {},
	    "cleanupIntervalS": 3600,
	    "fsPath": "",
	    "fsType": "basic"
	  },
	  "copiers": 0,
	  "pullerMaxPendingKiB": 0,
	  "hashers": 0,
	  "order": "random",
	  "ignoreDelete": false,
	  "scanProgressIntervalS": 0,
	  "pullerPauseS": 0,
	  "maxConflicts": -1,
	  "disableSparseFiles": false,
	  "disableTempIndexes": false,
	  "paused": false,
	  "weakHashThresholdPct": 25,
	  "markerName": ".stfolder",
	  "copyOwnershipFromParent": false,
	  "modTimeWindowS": 0,
	  "maxConcurrentWrites": 2,
	  "disableFsync": false,
	  "blockPullOrder": "standard",
	  "copyRangeMethod": "standard",
	  "caseSensitiveFS": false,
	  "junctionsAsDirs": true
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
	  "ignoredFolders": [
	    {
	      "time": "2022-01-09T19:09:52Z",
	      "id": "br63e-wyhb7",
	      "label": "Foo"
	    }
	  ],
	  "maxRequestKiB": 0,
	  "untrusted": false,
	  "remoteGUIPort": 0
	}
      ],
      "gui": {
	"enabled": true,
	"address": "127.0.0.1:8384",
	"unixSocketPermissions": "",
	"user": "Username",
	"password": "$2a$10$ZFws69T4FlvWwsqeIwL.TOo5zOYqsa/.TxlUnsGYS.j3JvjFTmxo6",
	"authMode": "static",
	"useTLS": false,
	"apiKey": "k1dnz1Dd0rzTBjjFFh7CXPnrF12C49B1",
	"insecureAdminAccess": false,
	"theme": "default",
	"debugging": false,
	"insecureSkipHostcheck": false,
	"insecureAllowFrameLoading": false
      },
      "ldap": {
	"address": "",
	"bindDN": "",
	"transport": "plain",
	"insecureSkipVerify": false,
	"searchBaseDN": "",
	"searchFilter": ""
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
	"startBrowser": true,
	"natEnabled": true,
	"natLeaseMinutes": 60,
	"natRenewalMinutes": 30,
	"natTimeoutSeconds": 10,
	"urAccepted": 0,
	"urSeen": 0,
	"urUniqueId": "...",
	"urURL": "https://data.syncthing.net/newdata",
	"urPostInsecurely": false,
	"urInitialDelayS": 1800,
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
	"unackedNotificationIDs": [
	  "authenticationUserAndPassword"
	],
	"trafficClass": 0,
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
	"maxConcurrentIncomingRequestKiB": 0,
	"announceLANAddresses": true,
	"sendFullIndexOnUpgrade": false,
	"featureFlags": [],
	"connectionLimitEnough": 0,
	"connectionLimitMax": 0,
	"insecureAllowOldTLSVersions": false
      },
      "remoteIgnoredDevices": [
	{
	  "time": "2022-01-09T20:02:01Z",
	  "deviceID": "...",
	  "name": "...",
	  "address": "192.168.0.20:22000"
	}
      ],
      "defaults": {
	"folder": {
	  "id": "",
	  "label": "",
	  "filesystemType": "basic",
	  "path": "~",
	  "type": "sendreceive",
	  "devices": [
	    {
	      "deviceID": "...",
	      "introducedBy": "",
	      "encryptionPassword": ""
	    }
	  ],
	  "rescanIntervalS": 3600,
	  "fsWatcherEnabled": true,
	  "fsWatcherDelayS": 10,
	  "ignorePerms": false,
	  "autoNormalize": true,
	  "minDiskFree": {
	    "value": 1,
	    "unit": "%"
	  },
	  "versioning": {
	    "type": "",
	    "params": {},
	    "cleanupIntervalS": 3600,
	    "fsPath": "",
	    "fsType": "basic"
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
	  "modTimeWindowS": 0,
	  "maxConcurrentWrites": 2,
	  "disableFsync": false,
	  "blockPullOrder": "standard",
	  "copyRangeMethod": "standard",
	  "caseSensitiveFS": false,
	  "junctionsAsDirs": false
	},
	"device": {
	  "deviceID": "",
	  "name": "",
	  "addresses": [
	    "dynamic"
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
	  "maxRequestKiB": 0,
	  "untrusted": false,
	  "remoteGUIPort": 0
	}
      }
    }
