The available subcommands are grouped into several nested hierarchies and some
parts dynamically generated from the running Syncthing instance.  On every
level, the ``--help`` option lists the available properties, actions and
commands for the user to discover interactively.  The top-level groups are:

config
    Access the live configuration in a running instance over the REST API to
    retrieve (get) or update (set) values in a fine-grained way.  The hierarchy
    is based on the same structure as used in the JSON / XML representations.

show
    Show system properties and status of a running instance.  The output is
    passed on directly from the REST API response and therefore requires parsing
    JSON format.

operations
    Control the overall program operation such as restarting or handling
    upgrades, as well as triggering some actions on a per-folder basis.  Can
    also configure the default ignore patterns from a plain text ignore file.

errors
    Examine pending error conditions that need attention from the user, or
    acknowledge (clear) them.

debug
    Various tools to aid in diagnosing problems or collection information for
    bug reports.  Some of these commands access the database directly and can
    therefore only work when Syncthing is not running.

``-`` (a single dash)
    Reads subsequent commands from the standard input stream, without needing to
    call the ``syncthing cli`` command over and over.  Exits on any invalid
    command or when EOF (end-of-file) is received.
