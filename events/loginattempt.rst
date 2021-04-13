Login Attempt
-------------

When authentication is enabled for the GUI, this event is emitted on every
login attempt. If either the username or password are incorrect, ``success``
is false and in any case the given username is returned.  The included
remote address concerns the immediate connecting host, which may not be the
origin of the request, but e.g. a reverse proxy.

.. code-block:: json

    {
       "id" : 187,
       "time" : "2017-03-07T00:19:24.420386143+01:00",
       "data" : {
          "remoteAddress" : "127.0.0.1:55530",
          "username" : "somename",
          "success" : false
       },
       "type" : "LoginAttempt",
       "globalID" : 195
    }
