GET /rest/folder/errors
=======================

.. versionadded:: 0.14.53

Takes one mandatory parameter, ``folder``, and returns the list of errors
encountered during scanning or pulling.

The results can be paginated using the :ref:`common pagination parameters
<rest-pagination>`.

.. code-block:: json

    {
      "folder": "nnhic-sxuae",
      "errors": [
        {
          "path": "noperm.txt",
          "error": "hashing: open /path/to/folder/noperm.txt: permission denied"
        }
      ],
      "page": 1,
      "perpage": 100
    }
