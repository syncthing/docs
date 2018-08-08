POST /rest/db/override
======================

Request override of a send only folder. Override means to make the local
version latest, overriding changes made on other devices. This API call does
nothing if the folder is not a send only folder.

Takes the mandatory parameter `folder` (folder ID).

.. code-block:: bash

    curl -X POST -H X-API-key:... http://127.0.0.1:8384/rest/db/override?folder=default
