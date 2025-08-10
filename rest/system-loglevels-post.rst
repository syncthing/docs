POST /rest/system/loglevels
===========================

.. versionadded:: 2.0.0

Changes the log level for specified facilities. Post an object with the log
facilites as keys and desired log level (``DEBUG``, ``INFO``, ``WARN``, or
``ERROR``) as the values.

.. code-block:: bash

    $ curl -H X-API-Key:abc123 -d '{"beacon":"DEBUG","discovery":"WARN"}' 'http://localhost:8384/rest/system/loglevels'
