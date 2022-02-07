GET /rest/db/remoteneed
=======================

.. versionadded:: 0.14.43

Takes the mandatory parameters ``folder`` and ``device``, and returns the list
of files which are needed by that remote device in order for it to become in
sync with the shared folder.

The results can be paginated using the :ref:`common pagination parameters
<rest-pagination>`.

.. code-block:: json

    {
      "files": [
        {
          "flags": "0755",
          "sequence": 6,
          "modified": "2015-04-20T23:06:12+09:00",
          "name": "ls",
          "size": 34640,
          "version": [
            "5157751870738175669:1"
          ]
        }
      ],
      "page": 1,
      "perpage": 100
    }

.. note:: This is an expensive call, increasing CPU and RAM usage on the device.
          Use sparingly.
