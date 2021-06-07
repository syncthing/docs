ListenAddressesChanged
----------------------

This event is emitted when a :ref:`listen address <listen-addresses>` changes.

.. code-block:: json

    {
       "type" : "ListenAddressesChanged",
       "id" : 70,
       "time" : "2017-03-06T15:01:24.88340663+01:00",
       "globalID" : 70,
       "data" : {
          "address" : {
             "Fragment" : "",
             "RawQuery" : "",
             "Scheme" : "dynamic+https",
             "Path" : "/endpoint",
             "RawPath" : "",
             "User" : null,
             "ForceQuery" : false,
             "Host" : "relays.syncthing.net",
             "Opaque" : ""
          },
          "wan" : [
             {
                "ForceQuery" : false,
                "User" : null,
                "Host" : "31.15.66.212:443",
                "Opaque" : "",
                "Path" : "/",
                "RawPath" : "",
                "RawQuery" : "id=F4HSJVO-CP2C3IL-YLQYLSU-XTYODAG-PPU4LGV-PH3MU4N-G6K56DV-IPN47A&pingInterval=1m0s&networkTimeout=2m0s&sessionLimitBps=0&globalLimitBps=0&statusAddr=:22070&providedBy=",
                "Scheme" : "relay",
                "Fragment" : ""
             }
          ],
          "lan" : [
             {
                "RawQuery" : "id=F4HSJVO-CP2C3IL-YLQYLSU-XTYODAG-PPU4LGV-PH3MU4N-G6K56DV-IPN47A&pingInterval=1m0s&networkTimeout=2m0s&sessionLimitBps=0&globalLimitBps=0&statusAddr=:22070&providedBy=",
                "Scheme" : "relay",
                "Fragment" : "",
                "RawPath" : "",
                "Path" : "/",
                "Host" : "31.15.66.212:443",
                "Opaque" : "",
                "ForceQuery" : false,
                "User" : null
             }
          ]
       }
    }
