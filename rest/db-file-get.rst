GET /rest/db/file
=================

Returns data about a given file, including version and
availability.

Request
-------

The call requires parameters for `folder` and `file` in the query string:

- `folder` is the folder ID which you can find in the Synthing Web GUI, e.g. `5camp-slpa8`

- `file` is the full path starting from the folder root to the file you are interested in. The path and filename must be correctly URL-encoded.


Example Request
^^^^^^^^^^^^^^^

  curl --silent --get --header "X-API-KEY: YOUR_API_TOKEN" "http://localhost:8080/rest/db/file?folder=YOUR_FOLDER_ID" --data-urlencode "file=path/to/file.pdf"

Response
--------


.. code-block::

    {
      "availability": [
        {
          "id": "ITZRNXE-YNROGBZ-HXTH5P7-VK5NYE5-QHRQGE2-7JQ6VNJ-KZUEDIU-5PPR5AM",
          "fromTemporary": false
        }
      ],
      "global": { /* a file entry */ },
      "local": { /* a file entry */ }
    }

``local`` and ``global`` refer to the current file on disk and the globally newest file, respectively.

A file entry looks like this::

    {
      {
        "deleted": false,
        "ignored": false,
        "inodeChange": "1970-01-01T01:00:00+01:00",
        "invalid": false,
        "localFlags": 0,
        "modified": "2022-09-28T08:07:19.979723+02:00",
        "modifiedBy": "523ITIE",
        "mustRescan": false,
        "name": "img",
        "noPermissions": false,
        "numBlocks": 0,
        "permissions": "0755",
        "platform": { /* platform specific data */ },
        "sequence": 914,
        "size": 128,
        "type": "FILE_INFO_TYPE_DIRECTORY",
        "version": [
          "523ITIE:1664345275"
        ]
      },
      "mtime": {
        "err": null,
        "value": {
          "real": "0001-01-01T00:00:00Z",
          "virtual": "0001-01-01T00:00:00Z"
        }
      }
    }

Platform specific data may be ownership, extended attributes, etc. and is
divided into entries per operating system / platform. An example platform
entry containing ownership information for Unix systems and an extended
attribute for macOS ("darwin") looks as follows::

    {
      "darwin": {
        "xattrs": [
          {
            "name": "net.kastelo.xattrtest",
            "value": "aGVsbG8="
          }
        ]
      },
      "freebsd": null,
      "linux": null,
      "netbsd": null,
      "unix": {
        "gid": 20,
        "groupName": "staff",
        "ownerName": "jb",
        "uid": 501
      },
      "windows": null
    }
