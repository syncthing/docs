.. _crashrep-dev:

Automatic Crash Reporting for Developers
========================================

Collecting and Sending
----------------------

Syncthing runs with one *monitor process* and one *main
process*. The main process is the thing that is really "Syncthing". The
monitor process is responsible for reading the console output from the main
process, restarting it if it exits, and reporting any crashes of the main
process -- when it's allowed to do so.

When the monitor process detects a crash it creates a file
``panic-$timestamp.log`` in the config directory and attempts to upload it
to the crash reporting server -- if crash reporting is enabled. When a log
has been successfully reported it is renamed with the double file ending
``.reported.log``. Old crash logs are automatically removed after a while,
reported or not.

Report Format
-------------

A crash report is fundamentally a blob of plain UTF-8 text. It has a loose
format, documented below. A report implicitly has a "report ID" which is the
SHA-256 hash of the entire report text, in hex format.

The report consists of the following:

- One line containing the Syncthing version, exactly as reported by
  ``syncthing --version``. A leading timestamp and log level *may* be
  present but is ignored.

- Zero or more lines of plaintext data that is for human consumption only.
  The reports that Syncthing itself sends will have zero lines here, but one
  could include a report of what happened, log extracts, etc. here barring
  any privacy issues on the sender's behalf.

- A line beginning with the words ``Panic at`` followed by a timestamp in
  RFC3339 format.

- The panic backtrace as printed / formatted by the Go runtime.

Here is an example of a well formed but short report::

    07:48:24 INFO: syncthing v1.1.4 "Erbium Earthworm" (go1.12.5 darwin-amd64) jb@kvin.kastelo.net 2019-05-21 20:36:38 UTC
    Panic at 2019-05-22T07:48:25+02:00
    panic: interface conversion: *pfilter.FilteredConn is not net.Conn: missing method Read

    goroutine 106 [running]:
    github.com/syncthing/syncthing/lib/connections.(*quicListener).Serve(0xc000158000)
            /Users/jb/dev/github.com/syncthing/syncthing/lib/connections/quic_listen.go:74 +0x41b
    github.com/thejerf/suture.(*Supervisor).runService.func1(0xc0001c6690, 0xc000000000, 0x54b4728, 0xc000158000)
            /Users/jb/go/pkg/mod/github.com/thejerf/suture@v3.0.2+incompatible/supervisor.go:600 +0x47
    created by github.com/thejerf/suture.(*Supervisor).runService
            /Users/jb/go/pkg/mod/github.com/thejerf/suture@v3.0.2+incompatible/supervisor.go:588 +0x5b

Wire Protocol
-------------

To upload a crash report we need three things:

- The data comprising the report as above,
- the SHA-256 hash of the report data, making up the report ID, and
- the base URL to send the report to.

The report URL is constructed by adding the report ID to the base URL. The
default base URL of ``https://crash.syncthing.net/newcrash/`` and the report
ID ``abcd1234`` results in the URL
``https://crash.syncthing.net/newcrash/abcd1234``.

First a ``HEAD`` request is performed on the report URL. If this request
returns successfully (``200 OK``) it means the server already has the report
ID in question. We do not need to upload it.

If the HEAD request returns ``404 Not Found`` or another error we can
attempt to upload the report. This is done by a ``PUT`` request to the same
URL with the report data as the body.
