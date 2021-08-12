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

- `syncthingtray <https://github.com/Martchus/syncthingtray>`__

Android
~~~~~~~

- `syncthing-android <https://github.com/syncthing/syncthing-android>`_

  A wrapper app for the Syncthing binary.

- `Syncthing-Fork <https://github.com/catfriend1/syncthing-android>`_

  An alternative wrapper app for the Syncthing binary with extended
  functionality.

.. _contrib-windows:

Windows
~~~~~~~

- `SyncTrayzor <https://github.com/canton7/SyncTrayzor>`_

  Windows host for Syncthing.  Installer, auto-start, built-in browser, tray
  icon, and more.

macOS
~~~~~

- `syncthing-macos <https://github.com/syncthing/syncthing-macos>`_

  syncthing-macos is a native macOS Syncthing tray application bundle.
  It hosts and wraps Syncthing, making it behave more like a native macOS application and less like a command-line utility with a web browser interface.

Linux
~~~~~

- `Syncthing Icon <https://extensions.gnome.org/extension/989/syncthing-icon/>`_

  A GNOME Shell extension displaying a Syncthing status icon in the top bar.

- `Syncthing Indicator <https://extensions.gnome.org/extension/1070/syncthing-indicator/>`_

  A GNOME Shell indicator for starting, monitoring and controlling the Syncthing daemon using SystemD.

- `syncthing-quick-status <https://github.com/serl/syncthing-quick-status>`_

  Small bash application with minimal dependencies, for a simple colorful representation of the current status.

Packages and Bundlings
----------------------

Cross-platform
~~~~~~~~~~~~~~

- Webi: `syncthing <https://webinstall.dev/syncthing>`__

  Mac, Linux: ::

    $ curl -sS https://webinstall.dev/syncthing | bash

  Windows 10 ::

    > curl.exe -A MS https://webinstall.dev/syncthing | powershell

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

- Official Community Repository: `syncthing <https://www.archlinux.org/packages/?name=syncthing>`__

- Arch User Repository: `syncthing-discosrv <https://aur.archlinux.org/packages/syncthing-discosrv>`__ and `syncthingtray <https://aur.archlinux.org/packages/syncthingtray>`__

Docker
~~~~~~

- `docker-syncthing <https://docs.linuxserver.io/images/docker-syncthing>`_ from `LinuxServer <https://www.linuxserver.io>`_

- Dockerfiles for `Syncthing <https://github.com/firecat53/dockerfiles/tree/main/syncthing>`_ and `Syncthing Discovery Server <https://github.com/firecat53/dockerfiles/tree/main/syncthing_discovery>`_.
  Latest binary releases used for both.

- `docker-syncthing <https://github.com/joeybaker/docker-syncthing>`_
  A fully baked docker container that allows custom config and will keep your
  settings and data past docker image restarts.

- `syncthing-docker-scratch <https://github.com/djtm/syncthing-docker-scratch>`_
  Builds docker containers from scratch base and/or runs the containers in
  docker or rkt.

- `rpi-syncthing <https://github.com/funkyfuture/docker-rpi-syncthing>`_
  Configurable image for the Raspberry Pi.

- `syncthing for hassio <https://github.com/bestlibre/hassio-addons/tree/master/syncthing>`_
  A docker based addon for `hassio <https://www.home-assistant.io/hassio/>`_

Gentoo
~~~~~~

Official net-p2p package: `syncthing <https://packages.gentoo.org/packages/net-p2p/syncthing>`__

FreeBSD
~~~~~~~

FreshPorts: `syncthing <https://www.freshports.org/net/syncthing>`__

macOS
~~~~~

MacPorts: `syncthing <https://ports.macports.org/port/syncthing/summary>`__ ::

    $ sudo port install syncthing

OpenBSD
~~~~~~~

Official ports: `syncthing <https://cvsweb.openbsd.org/cgi-bin/cvsweb/ports/net/syncthing>`__ and `QSyncthingTray <https://cvsweb.openbsd.org/cgi-bin/cvsweb/ports/net/qsyncthingtray>`__

OpenSUSE
~~~~~~~~

Official packages: `syncthing <https://software.opensuse.org/package/syncthing>`__ and `syncthingtray <https://software.opensuse.org/package/syncthingtray>`__

Synology NAS (DSM)
~~~~~~~~~~~~~~~~~~

- Synocommunity: add ``http://packages.synocommunity.com/`` to the Package
  Center in DSM or view the `browsable repository
  <https://synocommunity.com/packages>`__. Numerous CPU architectures are
  supported. SPK's may be older versions, however you can execute a Syncthing
  version upgrade via the web GUI after installation.

- Kastelo maintains `Synology packages <https://docs.kastelo.net/synology/>`__
  in a custom package source. Packages auto upgrade via the Synology package
  manager.

QNAP NAS (QTS)
~~~~~~~~~~~~~~

`Syncthing QPKG <https://qnapclub.eu/en/qpkg/692>`__ (Qnap
Package) available for ALL models x86, x86\_64, Arm (all including new models).

RockStor
~~~~~~~~

`Docker container <http://rockstor.com/docs/docker-based-rock-ons/syncthing.html>`_ and `registry entry <https://github.com/rockstor/rockon-registry/blob/master/syncthing.json>`_

ClearOS / WikiSuite
~~~~~~~~~~~~~~~~~~~~

Syncthing is part of `WikiSuite <http://wikisuite.org/>`_, and thus packaged for `ClearOS <http://wikisuite.org/How-to-install-Syncthing-on-ClearOS>`_.

Cloudron
~~~~~~~~

Syncthing is available as a 1-click install on `Cloudron <https://cloudron.io>`_. For those unaware,
Cloudron makes it easy to run apps on your server and keep them up-to-date and secure.

.. image:: https://cloudron.io/img/button.svg
   :target: https://cloudron.io/button.html?app=net.syncthing.cloudronapp2

There is a `demo available <https://my-demo.cloudron.me>`_ (username: cloudron password: cloudron)

The Cloudron package is developed `here <https://git.cloudron.io/cloudron/syncthing-app>`_.

WD My Cloud NAS
~~~~~~~~~~~~~~~

Packages for OS3 available on `WDCommunity <https://wdcommunity.com>`_.

Integrations
------------

REST API Bindings
~~~~~~~~~~~~~~~~~

- Python: https://github.com/blakev/python-syncthing (https://pypi.python.org/pypi/syncthing)
- PHP: https://github.com/terzinnorbert/syncthing-rest

Ports
~~~~~

- Swift: `pulse-swift <https://source.ind.ie/project/pulse-swift/tree/master>`_

  Currently still in development and is "not yet usable by any standard". Only the Block Exchange Protocol layer and the Connection layer are completed.

- Java: `a-sync <https://github.com/davide-imbriaco/a-sync>`_

  This implements the BEP, discovery and relay protocols providing a command
  line utility to access a Syncthing network, a service to proxy the relay
  protocol over http and a client library for the BEP protocol.

Configuration management
~~~~~~~~~~~~~~~~~~~~~~~~

- `puppet-syncthing <https://github.com/whefter/puppet-syncthing>`_
- `ansible-syncthing <https://github.com/le9i0nx/ansible-syncthing>`_
- Command line interface: `syncthingmanager <https://github.com/classicsc/syncthingmanager>`_

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

- `Syncthing-GTK <https://github.com/kozec/syncthing-gtk>`_
-  https://github.com/syncthing/syncthing-lite
-  https://github.com/sieren/QSyncthingTray
-  https://github.com/akissa/pysyncthing
-  https://github.com/retgoat/syncthing-ruby
-  https://github.com/sodacode/Windows-Syncthing-Installer
-  https://github.com/gutenye/syncthing-kindle
-  https://github.com/m0ppers/syncthing-bar (OSX 10.10 only)
-  https://github.com/graboluk/stiko
-  https://github.com/ALinuxNinja/salt-syncthing
-  https://www.asustor.com/apps/app_detail?id=552
-  https://susestudio.com/a/qkdvwb/syncthing
-  https://source.ind.ie/project/pulse-swift/tree/master
-  https://github.com/icaruseffect/syncthing-ubuntu-indicator
-  https://github.com/bloones/SyncThingWin
-  https://github.com/thunderbirdtr/syncthing_rpm
-  https://github.com/dapperstout/pulse-java
-  https://github.com/cebe/pulse-php-discover
-  https://github.com/sebw/bitbar-plugins
-  https://github.com/nhojb/SyncthingBar
-  https://github.com/iss0/SyncthingTray
-  https://github.com/alex2108/syncthing-tray
