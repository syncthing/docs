Debug Endpoints
================

These endpoints require the :stconf:opt:`gui.debugging` configuration option to
be enabled and yield an access denied error code otherwise.


GET /rest/debug/peerCompletion
------------------------------

Summarizes the completion percentage for each remote device.  Returns an object
with device IDs as keys and an integer percentage as values.


GET /rest/debug/httpmetrics
---------------------------

Returns statistics about each served REST API endpoint, to diagnose how much
time was spent generating the responses.


GET /rest/debug/cpuprof
-----------------------

Used to capture a profile of what Syncthing is doing on the CPU.  See
:doc:`/users/profiling`.


GET /rest/debug/heapprof
------------------------

Used to capture a profile of what Syncthing is doing with the heap memory.  See
:doc:`/users/profiling`.


GET /rest/debug/support
-----------------------

Collects information about the running instance for troubleshooting purposes.
Returns a "support bundle" as a zipped archive, which should be sent to the
developers after verifying it contains no sensitive personal information.
Credentials for the web GUI and the API key are automatically redacted already.


GET /rest/debug/file
--------------------

Shows diagnostics about a certain file in a shared folder.  Takes the ``folder``
(folder ID) and ``file`` (folder relative path) parameters.

.. code-block:: bash

    $ curl -H X-API-Key:... "http://localhost:8384/rest/debug/file?folder=default&file=foo/bar"

The returned object contains the same info as :doc:`db-file-get`, plus a summary
of ``globalVersions``.
