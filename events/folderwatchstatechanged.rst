FolderWatchStateChanged
-----------------------

The ``FolderWatchStateChanged`` event is emitted when a folder's watcher routine
encounters a new error, or when a previous error disappeared after retrying.
The event contains the ID of the affected folder and textual error messages
describing the previous (``from``) and the updated (``to``) error conditions.
If there was no error in either of these, the respective field is omitted.

.. code-block:: json

    {
      "id": 123,
      "type": "FolderWatchStateChanged",
      "time": "2022-03-14T12:34:56.890000000+01:00",
      "data": {
        "folder": "default",
        "from": "Something bad happened.",
        "to": "Something worse happened."
      }
    }
