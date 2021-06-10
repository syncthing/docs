Failure
-------

Generated for specific errors that will also be sent to the usage
reporting server, if enabled in the configuration.  These are usually
of special interest to the developers to troubleshoot complex errors.
The ``data`` field contains a textual error message.

.. code-block:: json

    {
      "id": 93,
      "globalID": 93,
      "type": "Failure",
      "time": "2021-06-07T21:22:03.414609034+02:00",
      "data": "index handler got paused while already paused"
    }
