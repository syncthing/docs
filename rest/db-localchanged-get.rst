GET /rest/db/localchanged
=========================

.. versionadded:: 0.14.55

Takes one mandatory parameter, ``folder``, and returns the list of files which
were changed locally in a receive-only folder.  Thus they differ from the global
state and could be reverted by pulling from remote devices again, see
:doc:`db-revert-post`.

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
