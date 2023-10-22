.. _syncthing:
.. role:: strike

Syncthing
=========

Synopsis
--------

::

    syncthing [serve]
              [--audit] [--auditfile=<file|-|-->] [--browser-only] [--device-id]
              [--generate=<dir>] [--gui-address=<address>] [--gui-apikey=<key>]
              [--home=<dir> | --config=<dir> --data=<dir>]
              [--logfile=<filename>] [--logflags=<flags>]
              [--log-max-old-files=<num>] [--log-max-size=<num>]
              [--no-browser] [--no-console] [--no-restart] [--paths] [--paused]
              [--no-default-folder] [--skip-port-probing]
              [--reset-database] [--reset-deltas] [--unpaused] [--allow-newer-config]
              [--upgrade] [--no-upgrade] [--upgrade-check] [--upgrade-to=<url>]
              [--verbose] [--version] [--help] [--debug-*]

    syncthing generate
              [--home=<dir> | --config=<dir>]
              [--gui-user=<username>] [--gui-password=<password|->]
              [--no-default-folder] [--skip-port-probing] [--no-console]
              [--help]

    syncthing decrypt (--to=<dir> | --verify-only)
              [--password=<pw>] [--folder-id=<id>] [--token-path=<file>]
              [--continue] [--verbose] [--version] [--help]
              <path>

    syncthing cli
              [--home=<dir> | --config=<dir> --data=<dir>]
              [--gui-address=<address>] [--gui-apikey=<key>]
              [--help]
              <command> [command options...] [arguments...]

Description
-----------

Syncthing lets you synchronize your files bidirectionally across multiple
devices. This means the creation, modification or deletion of files on one
machine will automatically be replicated to your other devices. We believe your
data is your data alone and you deserve to choose where it is stored. Therefore
Syncthing does not upload your data to the cloud but exchanges your data across
your machines as soon as they are online at the same time.

The ``syncthing`` core application is a command-line program which usually runs
in the background and handles the synchronization. It provides a built-in, HTML
and JavaScript based user interface to be controlled from a web browser. This
frontend communicates with the core application through some HTTP APIs, which
other apps like graphical system integration helpers can use as well, for
greatest flexibility. A link to reach the GUI and API is printed among the first
few log messages.

Options
-------

.. cmdoption:: --allow-newer-config

    Try loading a config file written by a newer program version, instead of
    failing immediately.

.. cmdoption:: --audit

    Write events to timestamped file ``audit-YYYYMMDD-HHMMSS.log``.

.. cmdoption:: --auditfile=<file|-|-->

    Use specified file or stream (``"-"`` for stdout, ``"--"`` for stderr) for
    audit events, rather than the timestamped default file name.

.. cmdoption:: --browser-only

   Open the web UI in a browser for an already running Syncthing instance.

.. cmdoption:: --device-id

   Print device ID to command line.

.. cmdoption:: --generate=<dir>

    Generate key and config in specified dir, then exit.

.. cmdoption:: --gui-address=<address>

    Override GUI listen address. Set this to an address (``0.0.0.0:8384``)
    or a URL (``http://0.0.0.0:8384``). Supported schemes are ``http`` for
    plain HTTP, ``https`` for HTTP over TLS, ``unix`` for plain Unix sockets
    or ``unixs`` for TLS over Unix sockets. A Unix socket could look like this:
    ``unix:///run/syncthing/syncthing.socket`` (notice the three slashes: two
    as part of the URL structure, one to specify an absolute path).

.. cmdoption:: --gui-apikey=<string>

    Override the API key needed to access the GUI / REST API.

.. cmdoption:: --gui-password=<password|->

    Specify new GUI authentication password, to update the config file.  Read
    from the standard input stream if only a single dash (``-``) is given.  A
    plaintext password is hashed before writing to the config file, but an
    already bcrypt-hashed input is stored verbatim.  As a special case, giving
    the existing password hash as password will leave it untouched.

.. cmdoption:: --gui-user=<username>

    Specify new GUI authentication user name, to update the config file.

.. cmdoption:: --help, -h

    Show help text about command line usage.  Context-sensitive depending on the
    given subcommand.

.. cmdoption:: --home=<dir>

    Set common configuration and data directory. The default configuration
    directory is ``$XDG_STATE_HOME/syncthing`` or
    ``$HOME/.local/state/syncthing`` (Unix-like),
    ``$HOME/Library/Application Support/Syncthing`` (Mac) and
    ``%LOCALAPPDATA%\Syncthing`` (Windows).

.. cmdoption:: --config=<dir>

    Set configuration directory. Alternative to ``--home`` and must be used
    together with ``--data``.

.. cmdoption:: --data=<dir>

    Set data (e.g. database) directory. Alternative to ``--home`` and must be used
    together with ``--config``.

.. cmdoption:: --logfile=<filename>

    Set destination filename for logging (use ``"-"`` for stdout, which is the
    default option).

.. cmdoption:: --logflags=<flags>

    Select information in log line prefix. The ``--logflags`` value is a sum of
    the following:

    -  1: Date
    -  2: Time
    -  4: Microsecond time
    -  8: Long filename
    - 16: Short filename

    To prefix each log line with date and time, set ``--logflags=3`` (1 + 2 from
    above). The value 0 is used to disable all of the above. The default is to
    show time only (2).

.. cmdoption:: --log-max-old-files=<num>

    Number of old files to keep (zero to keep only current).  Applies only when
    log rotation is enabled through ``--log-max-size``.

.. cmdoption:: --log-max-size=<num>

    Maximum size of any log file (zero to disable log rotation).

.. cmdoption:: --no-browser

    Do not start a browser.

.. cmdoption:: --no-console

    Hide the console window. (On Windows only)

.. cmdoption:: --no-default-folder

    Don't create a default folder when generating an initial configuration /
    starting for the first time.

.. cmdoption:: --no-restart

    Do not restart Syncthing when it exits. The monitor process will still run
    to handle crashes and writing to logfiles (if configured to).

.. cmdoption:: --no-upgrade

    Disable automatic upgrades.  Equivalent to the ``STNOUPGRADE`` environment
    variable, see below.

.. cmdoption:: --paths

    Print the paths used for configuration, keys, database, GUI overrides,
    default sync folder and the log file.

.. cmdoption:: --paused

    Start with all devices and folders paused.

.. cmdoption:: --reset-database

    Reset the database, forcing a full rescan and resync. Create `.stfolder`
    folders in each sync folder if they do not already exist. **Caution**:
    Ensure that all sync folders which are mountpoints are already mounted.
    Inconsistent versions may result if the mountpoint is later mounted and
    contains older versions.

.. cmdoption:: --reset-deltas

    Reset delta index IDs, forcing a full index exchange.

.. cmdoption:: --skip-port-probing

    Don't try to find unused random ports for the GUI and listen address when
    generating an initial configuration / starting for the first time.

.. cmdoption:: --unpaused

    Start with all devices and folders unpaused.

.. cmdoption:: --upgrade

    Perform upgrade.

.. cmdoption:: --upgrade-check

    Check for available upgrade.

.. cmdoption:: --upgrade-to=<url>

    Force upgrade directly from specified URL.

.. cmdoption:: --verbose

    Print verbose log output.

.. cmdoption:: --version

    Show version.

.. cmdoption:: --to=<dir>

    Destination directory where files should be stored after decryption.

.. cmdoption:: --verify-only

    Don't write decrypted files to disk (but verify plaintext hashes).

.. cmdoption:: --password=<pw>

    Folder password for decryption / verification.  Can be passed through the
    ``FOLDER_PASSWORD`` environment variable instead to avoid recording in a
    shell's history buffer or sniffing from the running processes list.

.. cmdoption:: --folder-id=<id>

    Folder ID of the encrypted folder, if it cannot be determined automatically.

.. cmdoption:: --token-path=<file>

    Path to the token file within the folder (used to determine folder ID).

.. cmdoption:: --continue

    Continue processing next file in case of error, instead of aborting.

Exit Codes
----------

0
    Success / Shutdown
1
    Error
2
    Upgrade not available
3
    Restarting
4
    Upgrading

Exit codes over 125 are usually returned by the shell/binary loader/default
signal handler. Exit codes over 128+N on Unix usually represent the signal which
caused the process to exit. For example, ``128 + 9 (SIGKILL) = 137``.

Subcommands
-----------

The command line syntax actually supports different modes of operation through
several subcommands, specified as the first argument.  If omitted, the default
``serve`` is assumed.

The initial setup of a device ID and default configuration can be called
explicitly with the ``generate`` subcommand.  It can also update the configured
GUI authentication credentials, without going through the REST API.  An existing
device certificate is left untouched.  If the configuration file already exists,
it is validated and updated to the latest configuration schema, including adding
default values for any new options.

The ``decrypt`` subcommand is used in conjunction with untrusted (encrypted)
devices, see the relevant section on :ref:`decryption <untrusted-decrypt>` for
details.  It does not depend on Syncthing to be running, but works on offline
data.

To work with the REST API for debugging or automating things in Syncthing, the
``cli`` subcommand provides easy access to individual features.  It basically
saves the hassle of handling HTTP connections and API authentication.

.. include:: ../includes/cli-commands.rst

Proxies
-------

Syncthing can use a SOCKS, HTTP, or HTTPS proxy to talk to the outside
world. The proxy is used for outgoing connections only - it is not possible
to accept incoming connections through the proxy. The proxy is configured
through the environment variable ``all_proxy``. Somewhat unusually, this
variable must be named in lower case - it is not ":strike:`ALL_PROXY`". For
example::

    $ export all_proxy=socks://192.0.2.42:8081

Development Settings
--------------------

The following environment variables modify Syncthing's behavior in ways that
are mostly useful for developers. Use with care.
If you start Syncthing from within service managers like systemd or supervisor,
path expansion may not be supported.

.. include:: ../includes/env-vars.rst

See Also
--------

:manpage:`syncthing-config(5)`, :manpage:`syncthing-stignore(5)`,
:manpage:`syncthing-device-ids(7)`, :manpage:`syncthing-security(7)`,
:manpage:`syncthing-networking(7)`, :manpage:`syncthing-versioning(7)`,
:manpage:`syncthing-faq(7)`
