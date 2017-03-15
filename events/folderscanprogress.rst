Folder Scan Progress
--------------------

Emitted in regular intervals (folder setting ProgressIntervalS, 2s by default)
during scans giving the amount of bytes already scanned and to be scanned in
total , as well as the current scanning rates in bytes per second.

.. code-block:: json

    {
       "data" : {
          "total" : 1,
          "rate" : 0,
          "current" : 0,
          "folder" : "bd7q3-zskm5"
       },
       "globalID" : 29,
       "type" : "FolderScanProgress",
       "time" : "2017-03-06T15:00:58.072004209+01:00",
       "id" : 29
    }
