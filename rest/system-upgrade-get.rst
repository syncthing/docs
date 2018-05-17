GET /rest/system/upgrade
========================

Checks for a possible upgrade and returns an object describing the
newest version and upgrade possibility.

.. code-block:: json

    {
      "latest": "v0.14.47",
      "majorNewer": false,
      "newer": true,
      "running": "v0.14.46"
    }
