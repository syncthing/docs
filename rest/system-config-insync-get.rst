GET /rest/system/config/insync (DEPRECATED)
===========================================

.. deprecated:: v1.12.0
   This endpoint still works as before but is deprecated. Use
   :ref:`rest-config-insync` instead.

Returns whether the config is in sync, i.e. whether the running
configuration is the same as that on disk.

.. code-block:: json

    {
      "configInSync": true
    }
