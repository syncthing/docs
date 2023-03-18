RemoteChangeDetected
--------------------

Generated upon scan whenever a file is locally updated due to a remote change.
Files that are updated locally produce a :doc:`localchangedetected` event.

.. note:: This event is not included in :doc:`/rest/events-get` endpoint without
   a mask specified, but needs to be selected explicitly.

.. code-block:: json

   {
      "time" : "2017-03-06T23:58:21.844739891+01:00",
      "globalID" : 123,
      "data" : {
         "type" : "file",
         "action" : "deleted",
         "folder": "Dokumente",
         "folderID" : "Dokumente",
         "path" : "testfile",
         "label" : "Dokumente",
         "modifiedBy" : "BPDFDTU"
      },
      "type" : "RemoteChangeDetected",
      "id" : 2
   }

.. deprecated:: v1.1.2
  The ``folderID`` field is a legacy name kept only for compatibility.  Use the
  ``folder`` field with identical content instead.
