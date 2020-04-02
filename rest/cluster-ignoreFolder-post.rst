POST /rest/cluster/ignoreFolder
===============================

Configure a folder to be ignored when offered by a certain device.
Only pending folders are acted upon, as listed by
:ref:`rest-cluster-pendingFolders`.

Takes the mandatory parameter `device` (device ID).  If the `folder`
parameter is not given, all pending folders offered by that device are
ignored.  Otherwise just the given folder ID.

.. code-block:: bash

    curl -X POST -H "X-API-key: ..." "http://localhost:8384/rest/cluster/ignoreFolder?device=...&folder=abcde-vwxyz"
