GET /rest/db/file
=================

Returns most data available about a given file, including version and
availability. Takes ``folder`` and ``file`` parameters.

.. code-block:: json

    {
      "availability": [
        "I6KAH76-66SLLLB-5PFXSOA-UFJCDZC-YAOMLEK-CP2GB32-BV5RQST-3PSROAU"
      ],
      "global": {
        "flags": "0644",
        "sequence": 3,
        "modified": "2015-04-20T22:20:45+09:00",
        "name": "util.go",
        "numBlocks": 1,
        "size": 9642,
        "version": [
          "5407294127585413568:1"
        ]
      },
      "local": {
        "flags": "0644",
        "sequence": 4,
        "modified": "2015-04-20T22:20:45+09:00",
        "name": "util.go",
        "numBlocks": 1,
        "size": 9642,
        "version": [
          "5407294127585413568:1"
        ]
      }
    }

Fields
------


The *name* is the file name path relative to the folder root (UTF-8 NFC). The name field uses the slash character ("/") as path separator, regardless of the operating system conventions. The combination of Folder and Name uniquely identifies each file in a cluster.

The *version* field is a version vector describing the updates performed to file by all members in the cluster. Each counter in the version vector is an ID-Value tuple. The ID is used the first 64 bits of the device ID. The Value is a simple incrementing counter, starting at zero. The combination of Folder, Name and Version uniquely identifies the contents of a file at a given point in time.

The field *numBlocks* is the number of blocks the file has been divided in. Each block have one entry in the version vector. 

The *localVersion* field is the value of a device local monotonic clock at the time of last local database update to a file. The clock ticks on every local database update.

The *flags* field is made up of the following single bit flags::

    0                   1                   2                   3
    0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |              Reserved       |U|S|P|I|D|   Unix Perm. & Mode   |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+ 

The lower 12 bits hold the common Unix permission and mode bits. An implementation MAY ignore or interpret these as is suitable on the host operating system.

Bit 19 ("D") is set when the file has been deleted. The block list SHALL be of length zero and the modification time indicates the time of deletion or, if the time of deletion is not reliably determinable, the last known modification time.

Bit 18 ("I") is set when the file is invalid and unavailable for synchronization. A peer MAY set this bit to indicate that it can temporarily not serve data for the file.

Bit 17 ("P") is set when there is no permission information for the file. This is the case when it originates on a non-permission- supporting file system. Changes to only permission bits SHOULD be disregarded on files with this bit set. The permissions bits MUST be set to the octal value 0666.

Bit 16 ("S") is set when the file is a symbolic link. The block list SHALL be of one or more blocks since the target of the symlink is stored within the blocks of the file.

Bit 15 ("U") is set when the symbolic links target does not exist. On systems where symbolic links have types, this bit being means that the default file symlink SHALL be used. If this bit is unset bit 19 will decide the type of symlink to be created.

Bit 0 through 14 are reserved for future use and SHALL be set to zero.

The *modified* is the modification time in ISO date format.

In the rare occasion that a file is simultaneously and independently modified by two devices in the same cluster and thus end up on the same *version* number after modification, the *modified* field is used as a tie breaker (higher being better), followed by the hash values of the file blocks (lower being better).

The field *size* is the file size in bytes. 

