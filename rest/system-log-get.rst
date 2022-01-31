GET /rest/system/log
====================

.. versionadded:: 0.12.0

Returns the list of recent log entries.  The optional ``since`` parameter limits
the results to message newer than the given timestamp in :rfc:`3339` format.

.. code-block:: json

    {
      "messages": [
        {
          "when": "2014-09-18T12:59:26.549953186+02:00",
          "message": "This is a log entry"
        }
      ]
    }

GET /rest/system/log.txt
========================

Returns the same information, formatted as a text log instead of a JSON object.
