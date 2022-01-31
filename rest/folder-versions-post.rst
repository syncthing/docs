POST /rest/folder/versions
==========================

Restore archived versions of a given set of files.  Expects an object with
attributes named after the relative file paths, with timestamps as values
matching valid ``versionTime`` entries in the corresponding
:doc:`folder-versions-get` response object.

Takes the mandatory parameter ``folder`` (folder ID).  Returns an object
containing any error messages that occurred during restoration of the file, with
the file path as attribute name.

.. code-block:: bash

    curl -X POST -H X-API-key:... http://127.0.0.1:8384/rest/folder/versions?folder=default -d '{
      "dir1/dir2/bar": "2022-02-06T20:44:12+01:00",
      "baz": "2022-02-06T20:44:20+01:00"
    }'
