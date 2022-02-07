GET /rest/folder/versions
=========================

.. versionadded:: 0.14.44

Takes one mandatory parameter, ``folder``, and returns the list of archived
files that could be recovered.  How many versions are available depends on the
:doc:`/users/versioning` configuration.  Each entry specifies when the file
version was archived as the ``versionTime``, the ``modTime`` when it was last
modified before being archived, and the size in bytes.

.. code-block:: json

    {
      "dir1/dir2/bar": [
        {
          "versionTime": "2022-02-06T20:44:12+01:00",
          "modTime": "2021-01-14T13:21:22+01:00",
          "size": 4
        }
      ],
      "baz": [
        {
          "versionTime": "2022-02-06T20:44:20+01:00",
          "modTime": "2021-01-14T13:23:49+01:00",
          "size": 4
        }
      ],
      "foo": [
        {
          "versionTime": "2022-02-06T20:55:31+01:00",
          "modTime": "2022-02-06T20:44:13+01:00",
          "size": 4
        },
        {
          "versionTime": "2022-02-06T20:44:20+01:00",
          "modTime": "2021-01-14T13:21:16+01:00",
          "size": 4
        }
      ]
    }
