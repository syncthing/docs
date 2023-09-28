POST /rest/db/deleteignored
====================

.. versionadded:: 0.14.50

Request remove all ignored files from local filesystem. This API call does nothing if the folder is not a receive
only folder.

Takes the mandatory parameter `folder` (folder ID).

.. code-block:: bash

    curl -X POST -H X-API-Key:... http://127.0.0.1:8384/rest/db/deleteignored?folder=default
