.. versionadded:: 1.2.0

    Automatic crash reporting is available from version 1.2.0.

.. _crashrep:

Automatic Crash Reporting
=========================

Automatic crash reporting sends reports of crashes or "panics" to the developers of Syncthing.
This crash report contains the version information of Syncthing (equivalent to the output of ``syncthing --version``) and a technical trace of what the various threads / routines in Syncthing were doing at the time of the crash.

The crash report does not include log data, file names, device IDs, statistics, a unique identifier, or any other personally identifiable information.
Crash reporting is enabled out of the box for most users, but can be disabled in the :ref:`advanced configuration dialog <advanced>`.

The following is an example of a crash report as sent::

    07:48:24 INFO: syncthing v1.1.4-rc.1+24-g39ecb7ad-quic3 "Erbium Earthworm" (go1.12.5 darwin-amd64) jb@kvin.kastelo.net 2019-05-21 20:36:38 UTC
    Panic at 2019-05-22T07:48:25+02:00
    panic: interface conversion: *pfilter.FilteredConn is not net.Conn: missing method Read

    goroutine 106 [running]:
    github.com/syncthing/syncthing/lib/connections.(*quicListener).Serve(0xc000158000)
            /buildagent/work/github.com/syncthing/syncthing/lib/connections/quic_listen.go:74 +0x41b
    github.com/thejerf/suture.(*Supervisor).runService.func1(0xc0001c6690, 0xc000000000, 0x54b4728, 0xc000158000)
            /Users/jb/go/pkg/mod/github.com/thejerf/suture@v3.0.2+incompatible/supervisor.go:600 +0x47
    created by github.com/thejerf/suture.(*Supervisor).runService
            /Users/jb/go/pkg/mod/github.com/thejerf/suture@v3.0.2+incompatible/supervisor.go:588 +0x5b

    goroutine 1 [runnable]:
    github.com/syncthing/syncthing/lib/auto.Assets(0xc0000fa280)
            /buildagent/work/github.com/syncthing/syncthing/lib/auto/gui.files.go:131 +0x4936
    github.com/syncthing/syncthing/lib/api.newStaticsServer(0xc0002229b5, 0x7, 0xc00003a300, 0x33, 0xc00022297c)
            /buildagent/work/github.com/syncthing/syncthing/lib/api/api_statics.go:38 +0x37
    github.com/syncthing/syncthing/lib/api.New(0x5d4a557355ee1a96, 0x95ccd59f7d44241, 0x23a9c69bae83ac87, 0x6ee52bc80a137f7b, 0x4c4bb20, 0xc000069800, 0xc00003a300, 0x33, 0x496ae2c, 0x9, ...)
            /buildagent/work/github.com/syncthing/syncthing/lib/api/api.go:108 +0xb1
    main.setupGUI(0xc000296000, 0x4c4bb20, 0xc000069800, 0x4c504e0, 0xc000148000, 0x4c27aa0, 0xc0001565f0, 0x4c27aa0, 0xc000156690, 0x4c42820, ...)
            /buildagent/work/github.com/syncthing/syncthing/cmd/syncthing/main.go:1092 +0x31f
    main.syncthingMain(0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x49651b4, 0x1, 0x0, 0x0, ...)
            /buildagent/work/github.com/syncthing/syncthing/cmd/syncthing/main.go:873 +0x1bc5


Note that the username and hostname of the machine where Syncthing was
built will be included in the crash log as part of the version string.
This information is essential for the developers to interpret the log in
context. If you compile Syncthing locally and want to prevent your build
from having such data embedded, see :ref:`versiontagging`.

For a more detailed description of the format and how the sending happens, see :ref:`crashrep-dev`.
