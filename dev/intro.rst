Syncthing Development
=====================

Controlling Syncthing from External Applications
------------------------------------------------

People all over the world have developed a number of :ref:`useful applications
<contributions>` that build around the Syncthing core, such as tray
notifications and Android support. These are made possible using two APIs:

-  A long polling interface for exposing events from
   the core utility to an external party. This :doc:`/dev/events` is useful for being
   notified of when changes to files, network connections or sync status occur.

-  A :doc:`/dev/rest` for controlling the operation of Syncthing and directly
   querying for current status.

If this covers what you need to do, there is no need to delve deeper. However,
if you would like to add functionality to Syncthing itself, or correct a bug
or two in there, please read on.


Contributing to Syncthing
-------------------------

Please see `the contribution guidelines
<https://github.com/syncthing/syncthing/blob/main/CONTRIBUTING.md>`__ in the
main repository.