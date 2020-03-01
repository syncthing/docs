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
``.stignore`` can be created containing file patterns to ignore. The
``.stignore`` file must be placed in the root of the folder. The
``.stignore`` file itself will never be synced to other devices, although it can
``#include`` files that *are* synchronized between devices. All patterns are
relative to the folder root.
The contents of the ``.stignore`` file must be UTF-8 encoded.

.. note::

    Note that ignored files can block removal of an otherwise empty directory.
    See below for the (?d) prefix to allow deletion of ignored files.

Patterns
--------

The ``.stignore`` file contains a list of file or path patterns. The
*first* pattern that matches will decide the fate of a given file.

Regular file names match themselves, i.e. the pattern foo matches the files foo,
subdir/foo as well as any directory named foo. Spaces are treated as regular characters.

+--------------+-----------+----------+-----------------------------------------------------------------------------------+
| Character    | Where     | Affects  | Action                                                                            |
+==============+===========+==========+===================================================================================+
| ``/``        | Beginning | Folders  | Match in the root directory only. ``/foo`` matches foo but not ``subdir/foo``.    |
|              |           |          |                                                                                   |
+--------------+-----------+----------+-----------------------------------------------------------------------------------+
| ``!``        | Beginning | All      | Ignore all other patterns and force sync of matching files. Matching directory    |
|              |           |          | trees will be synced entirely regardless of other ignore patterns.                |
|              |           |          |                                                                                   |
+--------------+-----------+----------+-----------------------------------------------------------------------------------+
| ``?d``       | Beginning | Files    | Files with names matching the pattern will be removed if they prevent directory   |
|              |           |          | deletion. Examples: ``?.DS_Store``, ``?.Recycle``                                 |
|              |           |          |                                                                                   |
+--------------+-----------+----------+-----------------------------------------------------------------------------------+
| ``?i``       | Beginning | All      | Enable case-insensitiv e pattern matching. The (?i) prefix can be combined with   |
|              |           |          | other patterns, for example the pattern ``(?i)!picture* .png`` indicates that     |
|              |           |          | Picture1.PNG should be synchronized. On Mac OS and Windows, patterns are always   |
|              |           |          | case-insensitive.                                                                 |
|              |           |          |                                                                                   |
+--------------+-----------+----------+-----------------------------------------------------------------------------------+
| ``//``       | Beginning | N/A      | Comments. Contents of the line will have no effect.                               |
|              |           |          |                                                                                   |
+--------------+-----------+----------+-----------------------------------------------------------------------------------+
| ``#include`` | Beginning | N/A      | Load sync patterns from the referenced file. Including a file more than once,     |
|              |           |          | or a file that does not exist, results in an error. Patterns are always           |
|              |           |          | relative to root - including files in a subdirectory will not affect that         |
|              |           |          | directory. Example: ``#include more-patterns.txt``.                               |
|              |           |          |                                                                                   |
+--------------+-----------+----------+-----------------------------------------------------------------------------------+
| ``*``        | Anywhere  | All      | Match zero or more characters in a file/directory name, but not directory         |
|              |           |          | separators. ``te*st`` matches ``test``, ``subdir/telerest`` but not               |
|              |           |          | ``tele/rest``.                                                                    |
|              |           |          |                                                                                   |
+--------------+-----------+----------+-----------------------------------------------------------------------------------+
| ``**``       | Anywhere  | All      | Match every character including directory separators. ``te**st`` matches          |
|              |           |          | ``test``, ``subdir/telerest`` and ``tele/sub/dir/rest``.                          |
|              |           |          |                                                                                   |
+--------------+-----------+----------+-----------------------------------------------------------------------------------+
| ``?``        | Anywhere  | All      | Match a single character that is not the directory separator.                     |
|              |           |          |                                                                                   |
+--------------+-----------+----------+-----------------------------------------------------------------------------------+
| ``[]``       | Anywhere  | All      | Acts like regex character range ``[a-z]`` matches all alphabet characters.        |
|              |           |          |                                                                                   |
+--------------+-----------+----------+-----------------------------------------------------------------------------------+


Notes
-----
-  Prefixes can be specified in any order (e.g. "(?d)(?i)"), but cannot be in a
   single pair of parentheses (not ":strike:`(?di)`").
-  Include patterns (that begin with ``!``) cause Syncthing to traverse and
   :ref:`watch <scanning>` the entire directory tree regardless of other
   ignore patterns.
-  Windows does not support escaping ``\[foo - bar\]``.

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
