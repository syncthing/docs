allowedNetworks
===============

``allowedNetworks`` is an advanced device setting that affects connection
handling. The default is for this setting to be empty, indicating that there
is no restriction on the allowed networks for a device.

By setting this to a comma separated list of networks, connections to the
given device will be limited to those networks. The networks refer to the
address of the *remote* device, not the network that the local device is
present on.

Given a value of `192.168.0.0/16, 172.16.0.0/12` Syncthing will:

 - Allow connections from the device from addresses in the specified
   networks.

 - Reject connections from the device from addresses outside the specified
   networks.

 - Attempt connections to addresses in the specified networks (manually
   configured or discovered).

 - Not attempt connections to addresses outside the specified networks,
   regardless of whether manually configured or automatically discovered.
