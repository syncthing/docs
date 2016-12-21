.. _introducer:

Introducer Configuration
========================

The introducer feature allows a degree of automation for adding unconnected devices to a cluster. When two devices connect they exchange a list of mutually shared folders and the devices connected to those shares. If a device has a remote device configured as an introducer it will then act on this information, adding unknown devices listed by the introducer and connecting them to the currently configured folders as per the introducer's settings. This is done on a per-folder basis and only devices which are in common with mutual shares are included. No additional shares are imported, nor are devices which are not in common between the introducer and device's shared folder list.

The introduction process involves the autoconfiguration of device IDs, labels and configured address settings, but no other device-specific settings. For each offered device autoconfiguration is only applied once and is done so when a device connects to an introducer; a restart, after configuring a remote device to introduce, will force this. Once autoconfigured, device-specific settings will currently not receive any updates from an introducer.

If an introducer adds or removes any devices or shares, or changes device-share settings, however, this change will be reflected to devices upon their next connection. Similarly, if an introduced device is no longer present on an introducer, or no longer shares any mutual folders with the device, it will be automatically removed when devices in the cluster next connect to the introducer.

Note that devices which are introduced cannot be removed so long as the introducer device is still marked as such, and if they are unshared from a folder they will be re-added.

Introducer status is transferable; that is, an introducers' introducer will become your introducer as well.
