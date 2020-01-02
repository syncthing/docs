.. _download-progress:

DownloadProgress
----------------

Emitted during file downloads for each folder for each file. By default
only a single file in a folder is handled at the same time, but custom
configuration can cause multiple files to be shown.

.. code-block:: json

    {
        "id": 221,
        "globalID": 221,
        "type": "DownloadProgress",
        "time": "2014-12-13T00:26:12.9876937Z",
        "data": {
            "folder1": {
                "file1": {
                    "total": 800,
                    "pulling": 2,
                    "copiedFromOrigin": 0,
                    "reused": 633,
                    "copiedFromElsewhere": 0,
                    "pulled": 38,
                    "bytesTotal": 104792064,
                    "bytesDone": 87883776
                },
                "dir\\file2": {
                    "total": 80,
                    "pulling": 2,
                    "copiedFromOrigin": 0,
                    "reused": 0,
                    "copiedFromElsewhere": 0,
                    "pulled": 32,
                    "bytesTotal": 10420224,
                    "bytesDone": 4128768
                }
            },
            "folder2": {
                "file3": {
                    "total": 800,
                    "pulling": 2,
                    "copiedFromOrigin": 0,
                    "reused": 633,
                    "copiedFromElsewhere": 0,
                    "pulled": 38,
                    "bytesTotal": 104792064,
                    "bytesDone": 87883776
                },
                "dir\\file4": {
                    "total": 80,
                    "pulling": 2,
                    "copiedFromOrigin": 0,
                    "reused": 0,
                    "copiedFromElsewhere": 0,
                    "pulled": 32,
                    "bytesTotal": 10420224,
                    "bytesDone": 4128768
                }
            }
        }
    }

-  ``total`` - total number of blocks in the file
-  ``pulling`` - number of blocks currently being downloaded
-  ``copiedFromOrigin`` - number of blocks copied from the file we are
   about to replace
-  ``reused`` - number of blocks reused from a previous temporary file
-  ``copiedFromElsewhere`` - number of blocks copied from other files or
   potentially other folders
-  ``pulled`` - number of blocks actually downloaded so far
-  ``bytesTotal`` - approximate total file size
-  ``bytesDone`` - approximate number of bytes already handled (already
   reused, copied or pulled)

Where block size is 128KB.

Files/folders appearing in the event data imply that the download has
been started for that file/folder, where disappearing implies that the
downloads have been finished or failed for that file/folder. There is
always a last event emitted with no data, which implies all downloads
have finished/failed.
