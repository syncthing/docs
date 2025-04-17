LoginAttempt
------------

When authentication is enabled for the GUI, this event is emitted on every
login attempt. If either the username or password are incorrect, ``success``
is false and in any case the given username is returned.

.. code-block:: json

    {
       "id" : 187,
       "time" : "2017-03-07T00:19:24.420386143+01:00",
       "data" : {
          "remoteAddress" : "127.0.0.1",
          "username" : "somename",
          "success" : false
       },
       "type" : "LoginAttempt",
       "globalID" : 195
    }

If the ``X-Forwared-For`` header is present and the connecting host is
either on ``localhost`` or on the same LAN, it will be treated as a reverse
proxy. In this case, the ``remoteAddress`` field is filled with the leftmost
IP address from the header, and the additional ``proxy`` field retains the
original IP of the connecting host.

.. code-block:: json

    {
       "id" : 187,
       "time" : "2017-03-07T00:19:24.420386143+01:00",
       "data" : {
          "proxy" : "127.0.0.1",
          "remoteAddress" : "192.168.178.10",
          "username" : "somename",
          "success" : false
       },
       "type" : "LoginAttempt",
       "globalID" : 195
    }
