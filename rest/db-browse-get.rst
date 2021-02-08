GET /rest/db/browse
===================

Returns the directory tree of the global model. Directories are always
JSON objects (map/dictionary), and files are always arrays of
modification time and size. The first integer is the files modification
time, and the second integer is the file size.

The call takes one mandatory ``folder`` parameter and two optional
parameters. Optional parameter ``levels`` defines how deep within the
tree we want to dwell down (0 based, defaults to unlimited depth)
Optional parameter ``prefix`` defines a prefix within the tree where to
start building the structure.

.. code-block:: bash

    $ curl -s http://localhost:8384/rest/db/browse?folder=j663y-3ct3e&prefix=DCIM&levels=2
    [
        {
            "modTime" : "2020-10-02T23:48:52.076996974+02:00",
            "name" : "100ANDRO",
            "size" : 128,
            "type" : "FILE_INFO_TYPE_DIRECTORY"
        },
        {
            "children" : [
                {
                    "modTime" : "2020-12-16T23:31:34.5009668+01:00",
                    "name" : "IMG_20201114_124821.jpg",
                    "size" : 10682189,
                    "type" : "FILE_INFO_TYPE_FILE"
                },
                {
                    "modTime" : "2020-12-16T23:31:35.0106367+01:00",
                    "name" : "IMG_20201213_122451.jpg",
                    "size" : 7936351,
                    "type" : "FILE_INFO_TYPE_FILE"
                },
                {
                    "modTime" : "2020-12-13T12:25:05.017097469+01:00",
                    "name" : "IMG_20201213_122504.jpg",
                    "size" : 8406507,
                    "type" : "FILE_INFO_TYPE_FILE"
                },
                {
                    "modTime" : "2020-12-13T12:25:06.127097469+01:00",
                    "name" : "IMG_20201213_122505.jpg",
                    "size" : 8381931,
                    "type" : "FILE_INFO_TYPE_FILE"
                },
                {
                    "modTime" : "2020-12-13T12:53:29.707298401+01:00",
                    "name" : "IMG_20201213_125329.jpg",
                    "size" : 4388331,
                    "type" : "FILE_INFO_TYPE_FILE"
                },
            ],
            "modTime" : "2020-10-09T13:04:42.4410738+02:00",
            "name" : "Camera",
            "size" : 128,
            "type" : "FILE_INFO_TYPE_DIRECTORY"
        },
    ]

.. note::
  This is an expensive call, increasing CPU and RAM usage on the device. Use sparingly.
