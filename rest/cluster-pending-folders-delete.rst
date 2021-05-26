DELETE /rest/cluster/pending/folders
====================================

.. versionadded:: 1.18.0

Remove records about a pending folder announced from a remote device.
Valid values for the ``folder`` and ``device`` parameters are those
from the corresponding :doc:`/rest/cluster-pending-folders-get`
endpoint.  The ``device`` parameter is optional and affects
announcements of this folder from the given device, or from *any*
device if omitted.

.. code-block:: bash

    $ curl -X DELETE -H "X-API-Key: abc123" http://localhost:8384/rest/cluster/pending/folders?folder=cpkn4-57ysy&device=P56IOI7-MZJNU2Y-IQGDREY-DM2MGTI-MGL3BXN-PQ6W5BM-TBBZ4TJ-XZWICQ2

Returns status 200 and no content upon success, or status 500 and a
plain text error on failure.  A :doc:`/events/pendingfolderschanged`
event will be generated in response.

For a more permanent effect, also for future announcements of the same
folder ID, the folder should be ignored in the :doc:`configuration
</users/config>` instead.
