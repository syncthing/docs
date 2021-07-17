ConfigSaved
-----------

Emitted after the config has been saved by the user or by Syncthing
itself.

.. code-block:: json

    {
        "id": 50,
        "globalID": 50,
        "type": "ConfigSaved",
        "time": "2014-12-13T00:09:13.5166486Z",
        "data": {
            "version": 7,
            "folders": [{"..."}],
            "devices": [{"..."}],
            "gui": {"..."},
            "ldap": {"..."},
            "options": {"..."},
            "remoteIgnoredDevices": [{"..."}],
            "defaults": {"..."}
        }
    }
