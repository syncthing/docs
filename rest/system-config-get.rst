GET /rest/system/config
=======================

Returns the current configuration.

.. code-block:: bash

    {
        {
          "version": 15,
          "folders": [
            {
              "id": "GXWxf-3zgnU",
              "label": "MyFolder",
              "path": "...",
              "type": "readwrite",
              "devices": [
                {
                  "deviceID": "..."
                }
              ],
              "rescanIntervalS": 60,
              "ignorePerms": false,
              "autoNormalize": true,
              "minDiskFreePct": 1,
              "versioning": {
                "type": "simple",
                "params": {
                  "keep": "5"
                }
              },
              "copiers": 0,
              "pullers": 0,
              "hashers": 0,
              "order": "random",
              "ignoreDelete": false,
              "scanProgressIntervalS": 0,
              "pullerSleepS": 0,
              "pullerPauseS": 0,
              "maxConflicts": 10,
              "disableSparseFiles": false,
              "disableTempIndexes": false,
              "fsync": false,
              "invalid": ""
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
              "introducer": false
            }
          ],
          "gui": {
            "enabled": true,
            "address": "127.0.0.1:8384",
            "user": "Username",
            "password": "$2a$10$ZFws69T4FlvWwsqeIwL.TOo5zOYqsa/.TxlUnsGYS.j3JvjFTmxo6",
            "useTLS": false,
            "apiKey": "pGahcht56664QU5eoFQW6szbEG6Ec2Cr",
            "insecureAdminAccess": false,
            "theme": "default"
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
            "urUniqueId": "",
            "urURL": "https://data.syncthing.net/newdata",
            "urPostInsecurely": false,
            "urInitialDelayS": 1800,
            "restartOnWakeup": true,
            "autoUpgradeIntervalH": 12,
            "keepTemporariesH": 24,
            "cacheIgnoredFiles": false,
            "progressUpdateIntervalS": 5,
            "limitBandwidthInLan": false,
            "minHomeDiskFreePct": 1,
            "releasesURL": "https://upgrades.syncthing.net/meta.json",
            "alwaysLocalNets": [],
            "overwriteRemoteDeviceNamesOnConnect": false,
            "tempIndexMinBlocks": 10
          },
          "ignoredDevices": [],
          "ignoredFolders": []
        }
    }
