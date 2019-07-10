POST /rest/system/config
========================

Post the full contents of the configuration, in the same format as returned by
the corresponding GET request. When posting the configuration succeeds,
the posted configuration is immediately applied, except for changes that require a restart. Query
:ref:`rest-config-insync` to check if a restart is required.

This endpoint is the main point to control Syncthing, even if the change only
concerns a very small part of the config: The usual workflow is to get the
config, modify the needed parts and post it again.
