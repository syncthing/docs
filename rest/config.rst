Config Endpoints
================

.. versionadded:: 1.12.0

These endpoints facilitate access and modification of the :ref:`configuration
<config>` in a granular way. Config sent to the endpoints must be in the same
format as returned by the corresponding GET request. When posting the
configuration succeeds, the posted configuration is immediately applied, except
for changes that require a restart. Query :ref:`rest-config-insync` to check if
a restart is required.

For all endpoints supporting ``PATCH``, it takes the existing config and
unmarshals the given JSON object on top of it. This means all child objects will
replace the existing objects, not extend them. For example for
``RawListenAddresses`` in options, which is an array of strings, sending
``{RawListenAddresses: ["tcp://10.0.0.2"]}`` will replace all existing listen
addresses.

.. _rest-config:

/rest/config
------------

``GET`` returns the entire config and ``PUT`` replaces it.

.. _rest-config-insync:

/rest/config/restart-required
-----------------------------

``GET`` returns whether a restart of Syncthing is required for the current
config to take effect.

/rest/config/folders, /rest/config/devices
------------------------------------------

``GET`` returns all folders respectively devices as an array. ``PUT`` takes an array and
``POST`` a single object. In both cases if a given folder/device already exists,
it's replaced, otherwise a new one is added.

/rest/config/folders/\*id\*, /rest/config/devices/\*id\*
--------------------------------------------------------

Put the desired folder- respectively device-ID in place of \*id\*. ``GET``
returns the folder/device for the given ID, ``PUT`` replaces the entire config,
``PATCH`` replaces only the given child objects and ``DELETE`` removes the
folder/device.

/rest/config/options, /rest/config/ldap, /rest/config/gui
---------------------------------------------------------

``GET`` returns the respective object, ``PUT`` replaces the entire object and
``PATCH`` replaces only the given child objects.
