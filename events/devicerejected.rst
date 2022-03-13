DeviceRejected (DEPRECATED)
---------------------------

.. deprecated:: v1.13.0
   This event is still emitted for compatibility, but deprecated.  Use
   the replacement :doc:`pendingdeviceschanged` event instead.

Emitted when there is a connection from a device we are not configured
to talk to.

.. code-block:: json

    {
        "id": 24,
        "globalID": 24,
        "type": "DeviceRejected",
        "time": "2014-08-19T10:43:00.562821045+02:00",
        "data": {
            "address": "127.0.0.1:51807",
            "name": "My dusty computer",
            "device": "EJHMPAQ-OGCVORE-ISB4IS3-SYYVJXF-TKJGLTU-66DIQPF-GJ5D2GX-GQ3OWQK"
        }
    }
