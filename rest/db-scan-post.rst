POST /rest/db/scan
==================

Request immediate scan. Takes the optional parameters ``folder`` (folder ID),
``sub`` (path relative to the folder root) and ``next`` (time in seconds). If
``folder`` is omitted or empty all folders are scanned. If ``sub`` is given,
only this path (and children, in case it's a directory) is scanned. The ``next``
argument delays Syncthing's automated rescan interval for a given amount of
seconds.

Requesting scan of a path that no longer exists, but previously did, is
valid and will result in Syncthing noticing the deletion of the path in
question.

Returns status 200 and no content upon success, or status 500 and a
plain text error if an error occurred during scanning.

.. code-block:: bash

    curl -X POST http://127.0.0.1:8384/rest/db/scan?folder=default&sub=foo/bar
