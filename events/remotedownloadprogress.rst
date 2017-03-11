Remote Download Progress
------------------------

This event is emitted when a :ref:`download-progress` message is
received. It returns a map ``data`` of filenames with a count of
downloaded blocks. The files in questions are currently being
downloaded on the remote ``device`` and belong to ``folder``.

.. code-block:: json

    {
       "time" : "2017-03-07T00:11:37.65838955+01:00",
       "globalID" : 170,
       "data" : {
          "state" : {
             "tahr64-6.0.5.iso" : 1784
          },
          "device" : "F4HSJVO-CP2C3IL-YLQYLSU-XTYODAG-PPU4LGV-PH3MU4N-G6K56DV-IPN47A",
          "folder" : "Dokumente"
       },
       "type" : "RemoteDownloadProgress",
       "id" : 163
    }
