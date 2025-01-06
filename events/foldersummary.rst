FolderSummary
-------------

The FolderSummary event is emitted when folder contents have changed
locally. This can be used to calculate the current local completion
state.

.. code-block:: json

    {
        "id": 16,
        "type": "FolderSummary",
        "time": "2024-01-31T08:27:30.777875+01:00",
        "data": {
            "folder": "default",
            "summary": {
                "error": "",
                "errors": 0,
                "globalBytes": 4,
                "globalDeleted": 18664,
                "globalDirectories": 0,
                "globalFiles": 1,
                "globalSymlinks": 0,
                "globalTotalItems": 18665,
                "ignorePatterns": false,
                "inSyncBytes": 4,
                "inSyncFiles": 1,
                "invalid": "",
                "localBytes": 4,
                "localDeleted": 18664,
                "localDirectories": 0,
                "localFiles": 1,
                "localSymlinks": 0,
                "localTotalItems": 18665,
                "needBytes": 0,
                "needDeletes": 0,
                "needDirectories": 0,
                "needFiles": 0,
                "needSymlinks": 0,
                "needTotalItems": 0,
                "pullErrors": 0,
                "receiveOnlyChangedBytes": 0,
                "receiveOnlyChangedDeletes": 0,
                "receiveOnlyChangedDirectories": 0,
                "receiveOnlyChangedFiles": 0,
                "receiveOnlyChangedSymlinks": 0,
                "receiveOnlyTotalItems": 0,
                "remoteSequence": {
                    "MRIW7OK-NETT3M4-N6SBWME-N25O76W-YJKVXPH-FUMQJ3S-P57B74J-GBITBAC": 37329
                },
                "sequence": 37329,
                "state": "idle",
                "stateChanged": "2024-01-31T08:27:24+01:00",
                "version": 37329,
                "watchError": ""
            }
        }
    }
