Noauth Endpoints
================

These endpoints do not require being authenticated or an API key to be used.
This way third-party services and devices can make calls like a health-check
without needing to share your API key or credentials.

GET /rest/noauth/health
-----------------------

Returns a ``{"status": "OK"}`` object.

.. code-block:: json

    {
      "status": "OK"
    }