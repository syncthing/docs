Login Attempt
-------------

When authentication is enabled for the GUI, this event is emitted on every
login attempt. If either the username or password are incorrect, ``success``
is false and in any case the given username is returned.

.. code-block:: json

    {
       "id" : 187,
       "time" : "2017-03-07T00:19:24.420386143+01:00",
       "data" : {
          "username" : "somename",
          "success" : false
       },
       "type" : "LoginAttempt",
       "globalID" : 195
    }
