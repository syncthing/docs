GET /rest/db/completion
=======================

Returns the completion percentage (0 to 100) for a given device and
folder. Takes ``device`` and ``folder`` parameters.

.. code-block:: json

    {
      "completion": 100,
      "globalBytes": 156793013575,
      "needBytes": 0,
      "needDeletes": 0
    }

.. note::
  This is an expensive call, increasing CPU and RAM usage on the device. Use sparingly.
