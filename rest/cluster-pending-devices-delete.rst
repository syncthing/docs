DELETE /rest/cluster/pending/devices
====================================

.. versionadded:: 1.18.0

Remove records about a pending remote device which tried to connect.
Valid values for the ``device`` parameter are those from the
corresponding :doc:`/rest/cluster-pending-devices-get` endpoint.

.. code-block:: bash

    $ curl -X DELETE -H "X-API-Key: abc123" http://localhost:8384/rest/cluster/pending/devices?device=P56IOI7-MZJNU2Y-IQGDREY-DM2MGTI-MGL3BXN-PQ6W5BM-TBBZ4TJ-XZWICQ2

Returns status 200 and no content upon success, or status 500 and a
plain text error on failure.  A :doc:`/events/pendingdeviceschanged`
event will be generated in response.

For a more permanent effect, also for future connections from the same
device ID, the device should be ignored in the :doc:`configuration
</users/config>` instead.
