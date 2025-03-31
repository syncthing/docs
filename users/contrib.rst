.. _contributions:

Community Contributions
=======================

This page lists integrations, addons and packagings of Syncthing created by
the community. Like all documentation pages, it's wiki editable so please do
edit and add your own.

GUI Wrappers
------------

.. _contrib-all:

Cross-platform
~~~~~~~~~~~~~~

- `Syncthing Tray <https://martchus.github.io/syncthingtray>`__

Android
~~~~~~~

- `Syncthing-Fork <https://github.com/catfriend1/syncthing-android>`_

  An Android app for the Syncthing binary with extended functionality.

.. _contrib-windows:

Windows
~~~~~~~

We currently don't have any actively updated Windows-specific GUI wrappers here.

.. seealso:: :ref:`Cross-platform GUI Wrappers <contrib-all>`.

macOS
~~~~~

- `syncthing-macos <https://github.com/syncthing/syncthing-macos>`_

  syncthing-macos is a native macOS Syncthing tray application bundle.
  It hosts and wraps Syncthing, making it behave more like a native macOS application and less like a command-line utility with a web browser interface.

Linux
~~~~~

- `Syncthing Icon <https://extensions.gnome.org/extension/989/syncthing-icon/>`_

  GNOME Shell extension: A Syncthing status icon in the top bar.

- `Syncthing Indicator <https://extensions.gnome.org/extension/1070/syncthing-indicator/>`_

  GNOME Shell extension: A Syncthing indicator for starting, monitoring and controlling the Syncthing daemon using systemd.

- `Syncthing Toggle <https://extensions.gnome.org/extension/7180/syncthing-toggle/>`_

  GNOME Shell extension: A quick setting for turning on and off the Syncthing systemd service and opening the Web GUI.

- `SyncThingy <https://github.com/zocker-160/SyncThingy>`_

  Simple tray indicator written in C++ targeted at Flatpak users.

- `syncthing-quick-status <https://github.com/serl/syncthing-quick-status>`_

  Small bash application with minimal dependencies, for a simple colorful representation of the current status.
  
- `steamdeck-decky-syncthing  <https://github.com/theCapypara/steamdeck-decky-syncthing>`_

  A Steam Deck (Decky Loader) plugin for controlling Syncthing from the Steam Big Picture / Steam Deck UI.


Command Line Tools
------------------

- `STC <https://github.com/tenox7/stc>`_

  Syncthing Cli - a simple command line tool for getting status and performing basic operations from
  the shell / terminal without need of a web browser.

- `syncthing-graph <https://gitlab.com/andrea-trentini/syncthing-graph>`_

  Very simple graph (dot format) generator for Syncthing ``config.xml``.

- `syncthing-map <https://github.com/wsw70/syncthing-map>`_

  A cross-platform utility to map Syncthing devices and shared folders.
  Generates a visual representation of the relationships between several devices
  and their respective folders, including special folder types (send-only,
  receive-only).  Requires each device's XML configuration file as input.


Packages and Bundlings
----------------------

Cross-platform
~~~~~~~~~~~~~~

- Webi: `syncthing <https://webinstall.dev/syncthing>`__

  Mac, Linux: ::

    $ curl -sS https://webinstall.dev/syncthing | bash

  Windows 10 (build 1803) or later ::

    > curl.exe -A MS https://webinstall.dev/syncthing | powershell

.. _contrib-packages-windows:

Windows
~~~~~~~

- `Syncthing Windows Setup <https://github.com/Bill-Stewart/SyncthingWindowsSetup>`_

  A lightweight yet full-featured Windows installer built using Inno Setup.  Supports both
  admin and regular user installation, auto-start, firewall integration as well as silent
  installation.


Debian / Ubuntu
~~~~~~~~~~~~~~~


- Official packages: https://apt.syncthing.net/

- Debian packages: `syncthing <https://packages.debian.org/search?keywords=syncthing>`__, `syncthing-discosrv <https://packages.debian.org/search?keywords=syncthing-discosrv>`__ and `syncthing-relaysrv <https://packages.debian.org/search?keywords=syncthing-relaysrv>`_


Fedora / CentOS
~~~~~~~~~~~~~~~

For Fedora Syncthing is now in the official repo : https://src.fedoraproject.org/rpms/syncthing

Unofficial `RPM repo of Syncthing <https://copr.fedorainfracloud.org/coprs/daftaupe/syncthing/>`_ (`sources <https://gitlab.com/daftaupe/syncthing-rpm>`_)

ArchLinux
~~~~~~~~~

- Official Community Repository: `syncthing <https://archlinux.org/packages/?name=syncthing>`__

- Arch User Repository: `syncthingtray <https://aur.archlinux.org/packages/syncthingtray>`__

Docker
~~~~~~

- `docker-syncthing <https://docs.linuxserver.io/images/docker-syncthing>`_ from `LinuxServer <https://www.linuxserver.io>`__

- Dockerfiles for `Syncthing <https://github.com/firecat53/dockerfiles/tree/main/syncthing>`_ and `Syncthing Discovery Server <https://github.com/firecat53/dockerfiles/tree/main/syncthing_discovery>`_.
  Latest binary releases used for both.

- `docker-syncthing <https://github.com/joeybaker/docker-syncthing>`__
  A fully baked docker container that allows custom config and will keep your
  settings and data past docker image restarts.

- `syncthing-docker-scratch <https://github.com/djtm/syncthing-docker-scratch>`_
  Builds docker containers from scratch base and/or runs the containers in
  docker or rkt.

- `rpi-syncthing <https://github.com/funkyfuture/docker-rpi-syncthing>`_
  Configurable image for the Raspberry Pi.

- `Syncthing for Home Assistant OS <https://github.com/Poeschl/Hassio-Addons/tree/master/syncthing>`_
  A docker based addon for `Home Assistant Operating System <https://www.home-assistant.io/installation/#compare-installation-methods>`_

Gentoo
~~~~~~

Official net-p2p package: `syncthing <https://packages.gentoo.org/packages/net-p2p/syncthing>`__

FreeBSD
~~~~~~~

FreshPorts: `syncthing <https://www.freshports.org/net/syncthing>`__

macOS
~~~~~

MacPorts: `syncthing <https://ports.macports.org/port/syncthing/>`__ ::

    $ sudo port install syncthing

OpenBSD
~~~~~~~

Official ports: `syncthing <https://cvsweb.openbsd.org/cgi-bin/cvsweb/ports/net/syncthing>`__ and `QSyncthingTray <https://cvsweb.openbsd.org/cgi-bin/cvsweb/ports/net/qsyncthingtray>`__

OpenSUSE
~~~~~~~~

Official packages: `syncthing <https://software.opensuse.org/package/syncthing>`__ and `qsyncthingtray <https://software.opensuse.org/package/qsyncthingtray>`__

Synology NAS (DSM)
~~~~~~~~~~~~~~~~~~

- Synocommunity: add ``http://packages.synocommunity.com/`` to the Package
  Center in DSM or view the `browsable repository
  <https://synocommunity.com/packages>`__. Numerous CPU architectures are
  supported. SPK's may be older versions, however you can execute a Syncthing
  version upgrade via the web GUI after installation.

QNAP NAS (QTS)
~~~~~~~~~~~~~~

`Syncthing QPKG <https://www.myqnap.org/product/syncthing/>`__ (Qnap
Package) available for ALL models x86, x86\_64, Arm (all including new models).
Syncthing running as root <https://www.myqnap.org/product/syncthing-run-as-root/>

RockStor
~~~~~~~~

`Docker container <https://rockstor.com/docs/docker-based-rock-ons/syncthing.html>`_ and `registry entry <https://github.com/rockstor/rockon-registry/blob/master/syncthing.json>`_

Cloudron
~~~~~~~~

Syncthing is available as a 1-click install on `Cloudron <https://www.cloudron.io>`_. For those unaware,
Cloudron makes it easy to run apps on your server and keep them up-to-date and secure.

.. image:: https://www.cloudron.io/img/button.svg
   :target: https://www.cloudron.io/button.html?app=net.syncthing.cloudronapp2

There is a `demo available <https://my.demo.cloudron.io>`_ (username: cloudron password: cloudron)

The Cloudron package is developed `here <https://git.cloudron.io/cloudron/syncthing-app>`_.

WD My Cloud NAS
~~~~~~~~~~~~~~~

Packages for OS3 available on `WDCommunity <https://wdcommunity.com>`_.

Integrations
------------

REST API Bindings
~~~~~~~~~~~~~~~~~

- Python: https://github.com/blakev/python-syncthing (https://pypi.org/project/syncthing/)
- PHP: https://github.com/terzinnorbert/syncthing-rest

Configuration management
~~~~~~~~~~~~~~~~~~~~~~~~

- `puppet-syncthing <https://github.com/whefter/puppet-syncthing>`_
- `ansible-syncthing <https://github.com/le9i0nx/ansible-syncthing>`_
- Command line interface: `syncthingmanager <https://github.com/classicsc/syncthingmanager>`_
- `syncthing-configd <https://github.com/kastelo/syncthing-configd>`_

  A daemon that automatically manages certain aspects of the Syncthing configuration, such
  as automatically accepting or removing devices / folders based on patterns.

Monitoring
~~~~~~~~~~~~~~~~~~~~~~~~

- `munin-syncthing <https://gitlab.com/daftaupe/munin-syncthing>`_

Resolving conflicts
~~~~~~~~~~~~~~~~~~~

- `syncthing-resolve-conflicts <https://github.com/dschrempf/syncthing-resolve-conflicts>`_

  A small bash script that handles synchronization conflicts in text
  files that may pop up when using Syncthing.  It is inspired by the
  `pacdiff` utility from Arch Linux.  A diff utility can be used to
  merge the files and keep them up to date.

Older, Possibly Unmaintained
----------------------------

.. note::
   These projects have not been updated in quite a while. They may still be
   usable, or they may be in disrepair. If you are the maintainer of one of
   these and you have revived the project, please update this page
   accordingly.

- `syncthing-android <https://github.com/syncthing/syncthing-android>`_ (Archived on 2024-12-03)
- `SyncTrayzor <https://github.com/canton7/SyncTrayzor>`_
- `a-sync <https://github.com/davide-imbriaco/a-sync>`_
- `syncthing-tray-gtk3 <https://github.com/abdeoliveira/syncthing-tray-gtk3>`_ (Archived as of 2023-12-29)
- `Syncthing-GTK <https://github.com/syncthing-gtk/syncthing-gtk>`_ (Fork from `Kozec <https://github.com/kozec/syncthing-gtk>`_)
-  https://github.com/syncthing/syncthing-lite
-  https://github.com/sieren/QSyncthingTray
-  https://github.com/akissa/pysyncthing
-  https://github.com/retgoat/syncthing-ruby
-  https://github.com/codabrink/Windows-Syncthing-Installer
-  https://github.com/gutenye/syncthing-kindle
-  https://github.com/m0ppers/syncthing-bar (OSX 10.10 only)
-  https://github.com/graboluk/stiko
-  https://www.asustor.com/apps/app_detail?id=552
-  https://source.small-tech.org/project/pulse-swift/tree/master
-  https://github.com/icaruseffect/syncthing-ubuntu-indicator
-  https://github.com/bloones/SyncThingWin
-  https://github.com/thunderbirdtr/syncthing_rpm
-  https://github.com/dapperstout/pulse-java
-  https://github.com/cebe/pulse-php-discover
-  https://github.com/sebw/bitbar-plugins
-  https://github.com/nhojb/SyncthingBar
-  https://github.com/jastBytes/SyncthingTray
-  https://github.com/alex2108/syncthing-tray
