FolderCompletion
----------------

The ``FolderCompletion`` event is emitted when the local or remote
contents for a folder changes. It contains the completion percentage for
a given remote device and is emitted once per currently connected remote
device.

.. versionadded:: 1.20.0

  Indication whether the remote device has accepted the folder (shares it with
  us) as well, and whether it is paused.  The ``remoteState`` field is
  ``unknown`` when the remote device is not connected.  Otherwise it can be
  either ``paused``, ``notSharing``, or ``valid`` if the remote is sharing back.

.. code-block:: json

    {
        "id": 84,
        "globalID": 84,
        "type": "FolderCompletion",
        "time": "2022-04-27T14:14:27.043576583+09:00",
        "data": {
            "completion": 100,
            "device": "I6KAH76-66SLLLB-5PFXSOA-UFJCDZC-YAOMLEK-CP2GB32-BV5RQST-3PSROAU",
            "folder": "default",
            "globalBytes": 17,
            "globalItems": 4,
            "needBytes": 0,
            "needDeletes": 0,
            "needItems": 0,
            "remoteState": "valid",
            "sequence": 12
        }
    }
