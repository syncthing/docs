.. role:: strike

.. _ignoring-files:

Ignoring Files
==============

Synopsis
--------

::

    .stignore

Description
-----------

If some files should not be synchronized to (or from) other devices, a file called
``.stignore`` can be created containing file patterns to ignore.  The ``.stignore``
file must be placed in the root of the synced folder (files in other locations are
not applied).  The ``.stignore`` file itself will never be synced to other devices,
although it can ``#include`` files that *are* synchronized between devices.  All
patterns are relative to the synced folder root.  The contents of the ``.stignore``
file must be UTF-8 encoded.

.. note::

    Note that ignored files can block removal of an otherwise empty directory.
    See below for the (?d) prefix to allow deletion of ignored files.

Patterns
--------

The ``.stignore`` file contains a list of file or path patterns. The
*first* pattern that matches will decide the fate of a given file.

-  Regular file names match themselves, i.e. the pattern ``foo`` matches
   the files ``foo``, ``subdir/foo`` as well as any directory named
   ``foo``. Spaces are treated as regular characters, except for leading
   and trailing spaces, which are automatically trimmed.

-  **Asterisk** (``*``) matches zero or more characters in a filename, but does not
   match the directory separator. ``te*ne`` matches ``telephone``,
   ``subdir/telephone`` but not ``tele/phone``.

-  **Double asterisk** (``**``) matches as above, but also directory separators.
   ``te**ne`` matches ``telephone``, ``subdir/telephone`` and
   ``tele/sub/dir/phone``.

-  **Question mark** (``?``) matches a single character that is not the directory
   separator. ``te??st`` matches ``tebest`` but not ``teb/st`` or
   ``test``.

-  **Square brackets** (``[]``) denote a character range: ``[a-z]`` matches
   any lower case character.

-  **Curly brackets** (``{}``) denote a set of comma separated alternatives:
   ``{banana,pineapple}`` matches either ``banana`` or ``pineapple``.

-  **Backslash** (``\``) "escapes" a special character so that it loses its
   special meaning. For example, ``\{banana\}`` matches ``{banana}`` exactly
   and does not denote a set of alternatives as above.

   .. note::

      Escaped characters are not supported on Windows, where ``\`` is the
      path separator. If you still need to match files that have square or
      curly brackets in their names, one possible workaround is to replace
      them with ``?``, which will then match any character. For example,
      you can type ``?banana?`` to match both ``[banana]`` and
      ``{banana}``, and so on.

-  A pattern beginning with ``/`` matches in the root of the synced folder only.
   ``/foo`` matches ``foo`` but not ``subdir/foo``.

-  A pattern beginning with ``#include`` results in loading patterns
   from the named file. It is an error for a file to not exist or be
   included more than once. Note that while this can be used to include
   patterns from a file in a subdirectory, the patterns themselves are
   still relative to the synced folder *root*. Example:
   ``#include more-patterns.txt``.

   Any ``#include`` directives inside a file loaded by ``#include`` require paths
   specified relative to the directory containing the loaded file, rather than the
   synchronised root directory.

-  A pattern beginning with a ``!`` prefix negates the pattern: matching files
   are *included* (that is, *not* ignored). This can be used to override
   more general patterns that follow.

   .. note::

      Negated patterns that can match items below the folder root will cause
      Syncthing to traverse otherwise ignored directories. If the
      :ref:`watcher <scanning>` is enabled, those directories will also be
      watched. Directories ignored before the first negated pattern can
      however be safely skipped, since the first matching pattern wins. For
      example::

         /foo
         /bar
         !baz
         *

      The directories ``foo`` and ``bar`` will be entirely ignored. However any
      other directories present must be scanned entirely to find any items
      named `baz`, despite the fact that they will be ignored due to the
      ``*``. As a special case, top-level rooted patterns (e.g. ``!/foo``) do
      not cause this behaviour::

         !/baz
         *

      In this case, only the directory ``baz`` will be scanned, since
      everything else is ignored by the ``*`` pattern.

-  A pattern beginning with a ``(?i)`` prefix enables case-insensitive pattern
   matching. ``(?i)test`` matches ``test``, ``TEST`` and ``tEsT``. The
   ``(?i)`` prefix can be combined with other patterns, for example the
   pattern ``(?i)!picture*.png`` indicates that ``Picture1.PNG`` should
   be synchronized. On Mac OS and Windows, patterns are always case-insensitive.

-  A pattern beginning with a ``(?d)`` prefix enables removal of these files if
   they are preventing directory deletion. This prefix should be used by any OS
   generated files which you are happy to be removed.

   .. note::

      Prefixes can be specified in any order (e.g. ``(?d)(?i)``), but cannot
      be combined in a single pair of parentheses like :strike:`(?di)`.

-  A line beginning with ``//`` is a comment and has no effect. The same double
   slashes in any other place are interpreted literally, e.g. trying to do
   ``file // comment`` will make Syncthing look for a file called ``file // comment``.

Example
-------

Given a directory layout starting at the synced folder's root::

    .DS_Store
    .stignore
    foo
    foofoo
    bar/
        baz
        quux
        quuz
    bar2/
        baz
        frobble
    My Pictures/
        Img15.PNG

and an ``.stignore`` file with the contents::

    (?d).DS_Store
    !frobble
    !quuz
    foo
    *2
    qu*
    (?i)my pictures

all files and directories called "foo", ending in a "2" or starting with
"qu" will be ignored. The end result becomes::

    .DS_Store     # ignored, will be deleted if gets in the way of parent directory removal
    foo           # ignored, matches "foo"
    foofoo        # synced, does not match "foo" but would match "foo*" or "*foo"
    bar/          # synced
        baz       # synced
        quux      # ignored, matches "qu*"
        quuz      # synced, matches "qu*" but is excluded by the preceding "!quuz"
    bar2/         # synced, despite matching "*2" due to child frobble
        baz       # ignored, due to parent being ignored
        frobble   # synced, due to "!frobble"
    My Pictures/  # ignored, matched case insensitive "(?i)my pictures" pattern
        Img15.PNG # ignored, due to parent being ignored

.. note::
  Please note that directory patterns ending with a slash
  ``some/directory/`` matches the content of the directory, but not the
  directory itself. If you want the pattern to match the directory and its
  content, make sure it does not have a ``/`` at the end of the pattern.

.. versionadded:: 1.19.0

   Default patterns can be configured which will take effect when automatically
   accepting a folder from a remote device.  The GUI suggests same the patterns
   when adding a folder manually.  In either case, the ``.stignore`` file is
   created with these defaults if none is present yet.
