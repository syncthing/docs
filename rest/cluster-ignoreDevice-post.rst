POST /rest/cluster/ignoreDevice
===============================

Configure a device to be ignored when it's trying to connect.  Only
pending devices are acted upon, as listed by
:ref:`rest-cluster-pendingDevices`.

Takes the mandatory parameter `device` (device ID).

.. code-block:: bash

    curl -X POST -H "X-API-key: ..." "http://localhost:8384/rest/cluster/ignoreDevice?device=..."
