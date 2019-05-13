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

If some files should not be synchronized to other devices, a file called
``.stignore`` can be created containing file patterns to ignore. The
``.stignore`` file must be placed in the root of the folder. The
``.stignore`` file itself will never be synced to other devices, although it can
``#include`` files that *are* synchronized between devices. All patterns are
relative to the folder root.

.. note::

    Note that ignored files can block removal of an otherwise empty directory.
    See below for the (?d) prefix to allow deletion of ignored files.

Patterns
--------

The ``.stignore`` file contains a list of file or path patterns. The
*first* pattern that matches will decide the fate of a given file.

-  Regular file names match themselves, i.e. the pattern ``foo`` matches
   the files ``foo``, ``subdir/foo`` as well as any directory named
   ``foo``. Spaces are treated as regular characters.

-  Asterisk matches zero or more characters in a filename, but does not
   match the directory separator. ``te*st`` matches ``test``,
   ``subdir/telerest`` but not ``tele/rest``.

-  Double asterisk matches as above, but also directory separators.
   ``te**st`` matches ``test``, ``subdir/telerest`` and
   ``tele/sub/dir/rest``.

-  Question mark matches a single character that is not the directory
   separator. ``te??st`` matches ``tebest`` but not ``teb/st`` or
   ``test``.

-  Characters enclosed in square brackets ``[]`` are interpreted as a character range ``[a-z]``. Before using this syntax you should have a basic understanding of regular expression character classes.

-  A pattern beginning with ``/`` matches in the current directory only.
   ``/foo`` matches ``foo`` but not ``subdir/foo``.

-  A pattern beginning with ``#include`` results in loading patterns
   from the named file. It is an error for a file to not exist or be
   included more than once. Note that while this can be used to include
   patterns from a file in a subdirectory, the patterns themselves are
   still relative to the folder *root*. Example:
   ``#include more-patterns.txt``.

-  A pattern beginning with a ``!`` prefix negates the pattern: matching files
   are *included* (that is, *not* ignored). This can be used to override
   more general patterns that follow.

-  A pattern beginning with a ``(?i)`` prefix enables case-insensitive pattern
   matching. ``(?i)test`` matches ``test``, ``TEST`` and ``tEsT``. The
   ``(?i)`` prefix can be combined with other patterns, for example the
   pattern ``(?i)!picture*.png`` indicates that ``Picture1.PNG`` should
   be synchronized. On Mac OS and Windows, patterns are always case-insensitive.

-  A pattern beginning with a ``(?d)`` prefix enables removal of these files if
   they are preventing directory deletion. This prefix should be used by any OS
   generated files which you are happy to be removed.

-  A line beginning with ``//`` is a comment and has no effect.

-  Windows does not support escaping ``\[foo - bar\]``.

.. note::

   Prefixes can be specified in any order (e.g. "(?d)(?i)"), but cannot be in a
   single pair of parentheses (not ":strike:`(?di)`").

Example
-------

Given a directory layout::

    .DS_Store
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
    bar2/         # ignored, matched "*2"
        baz       # ignored, due to parent being ignored
        frobble   # ignored, due to parent being ignored; "!frobble" doesn't help
    My Pictures/  # ignored, matched case insensitive "(?i)my pictures" pattern
        Img15.PNG # ignored, due to parent being ignored

.. note::
  Please note that directory patterns ending with a slash
  ``some/directory/`` matches the content of the directory, but not the
  directory itself. If you want the pattern to match the directory and its
  content, make sure it does not have a ``/`` at the end of the pattern.
