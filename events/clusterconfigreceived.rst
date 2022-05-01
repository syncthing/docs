ClusterConfigReceived
---------------------

.. versionadded:: 1.20.0

The ``ClusterConfigReceived`` event is emitted after processing such a protocol
message received from a remote device.  It is mainly used for internal purposes.

.. code-block:: json

    {
        "id": 84,
        "globalID": 84,
        "type": "ClusterConfigReceived",
        "time": "2022-04-27T14:14:27.043576583+09:00",
        "data": {
            "device": "I6KAH76-66SLLLB-5PFXSOA-UFJCDZC-YAOMLEK-CP2GB32-BV5RQST-3PSROAU"
        }
    }
