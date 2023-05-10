Introducer Configuration
========================

The introducer feature lets a device automatically add new devices. When two devices connect they exchange a list of mutually shared folders and the devices connected to those shares. In the following example:

  Local device L sets remote device R as an introducer. They share the folder "Pictures." Device R is also sharing the folder with A and B, but L only shares with R.
  
  Once L and R connect, L will add A and B automatically, as if R "introduced" A and B to L.
  
  Remote device R also shares "Videos" with device C, but not with our local L. Device C will not be added to L as it is not connected to any folders that L and R share.

To mark a remote device as an introducer, start editing it in the Web GUI and
check the option under the "Sharing" tab.

The introduction process involves the autoconfiguration of device IDs, labels and configured address settings, but no other device-specific settings. For each offered device autoconfiguration is only applied once and is done so when a device connects to an introducer; a restart, after configuring a remote device to introduce, will force this. Once autoconfigured, device-specific settings will currently not receive any updates from an introducer.

If an introducer adds or removes any devices or shares, or changes device-share settings, however, this change will be reflected to devices upon their next connection. Similarly, if an introduced device is no longer present on an introducer, or no longer shares any mutual folders with the device, it will be automatically removed when devices in the cluster next connect to the introducer.

Note that if you manually remove an introduced device or unshare it from a folder, it will be automatically re-added as long as the introducer device is still marked as such.

Introducer status is transferable; that is, an introducers' introducer will become your introducer as well.

It is not a good idea to set two devices as introducers to each other. While this will work for adding devices, removing a device may present a problem, as the two devices will be constantly "re-introducing" the removed device to each other.
