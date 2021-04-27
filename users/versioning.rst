.. _versioning:

File Versioning
===============

Syncthing supports archiving the old version of a file when it is deleted or
replaced with a newer version from the cluster. This is called "file
versioning" and uses one of the available *versioning strategies* described
below. File versioning is configured per folder, on a per-device basis, and
defaults to "no file versioning", i.e. no old copies of files are kept.

.. note::
    Versioning applies to changes received *from other devices*. That is, if
    Alice has versioning turned on and Bob changes a file, the old version
    will be archived on Alice's computer when that change is synced from
    Bob. If Alice changes a file locally on her own computer Syncthing will
    not and can not archive the old version.

Trash Can File Versioning
-------------------------

This versioning strategy emulates the common "trash can" approach. When a file
is deleted or replaced due to a change on a remote device, it is moved to
the trash can in the ``.stversions`` folder. If a file with the same name was
already in the trash can it is replaced.

A configuration option is available to clean the trash can from files older
than a specified number of days. If this is set to a positive number of days,
files will be removed when they have been in the trash can that long. Setting
this to zero prevents any files from being removed from the trash can
automatically.

Simple File Versioning
----------------------

With "Simple File Versioning" files are moved to the ``.stversions`` folder
(inside your shared folder) when replaced or deleted on a remote device. This
option also takes a value in an input titled "Keep Versions" which tells
Syncthing how many old versions of the file it should keep. For example, if
you set this value to 5, if a file is replaced 5 times on a remote device, you
will see 5 time-stamped versions on that file in the ".stversions" folder on
the other devices sharing the same folder.

Staggered File Versioning
-------------------------

With "Staggered File Versioning" files are also moved to a different folder
when replaced or deleted on a remote device (just like "Simple File
Versioning"), however, versions are automatically deleted if they are older
than the maximum age or exceed the number of files allowed in an interval.

With this versioning method it's possible to specify where the versions are
stored, with the default being the ``.stversions`` folder inside the normal
folder path. If you set a custom version path, please ensure that it's on the
same partition or filesystem as the regular folder path, as moving files there
may otherwise fail. You can use an absolute path (this is recommended) or a
relative path. Relative paths are interpreted relative to Syncthing's current
or startup directory.

The following intervals are used and they each have a maximum number of files
that will be kept for each.

1 Hour
    For the first hour, the oldest version in every 30-seconds interval is
    kept.
1 Day
    For the first day, the oldest version in every hour is kept.
30 Days
    For the first 30 days, the oldest version in every day is kept.
Until Maximum Age
    Until maximum age, the oldest version in every week is kept.
Maximum Age
    The maximum time to keep a version in days. For example, to keep replaced or
    deleted files in the ".stversions" folder for an entire year, use 365. If
    only for 10 days, use 10.
    **Note: Set to 0 to keep versions forever.**

This means that there is only one version in each interval and as files age they
will be deleted unless when the interval they are entering is empty. By keeping
the oldest versions this versioning scheme preserves the file if it is
overwritten.

For more info, check the `unit test file
<https://github.com/syncthing/syncthing/blob/main/lib/versioner/staggered_test.go#L32>`__
that shows which versions are deleted for a specific run.

External File Versioning
------------------------

This versioning method delegates the decision on what to do to an
external command (e.g. a program or a command line script). Just prior
to a file being replaced, the command will be executed. The file needs
to be removed from the folder in the process, or otherwise Syncthing
will report an error. The command can use the following templated
arguments:

..
    This to be added when actually relevant.

    %FOLDER_FILESYSTEM%
      Filesystem type for the underlying folder.

%FOLDER_PATH%
  Path to the folder

%FILE_PATH%
  Path to the file within the folder

Note that the former expands to the path of the actual Syncthing folder,
and the latter to the path inside that folder. For instance, if you use
the default ``Sync`` folder in Windows, and the full path to the file is
``C:\Users\User\Sync\Family photos\IMG_2021-03-01.jpg``, then the
``%FOLDER_PATH%`` will be ``C:\Users\User\Sync``, and the
``%FILE_PATH%`` will be ``Family photos\IMG_2021-03-01.jpg``.

Example for Unixes
~~~~~~~~~~~~~~~~~~

Let's say I want to keep the latest version of each file as they are replaced
or removed; essentially I want a "trash can"-like behavior. For this, I create
the following script and store it as ``/Users/jb/bin/onlylatest.sh`` (i.e. the
``bin`` directory in my home directory):

.. code-block:: bash

    #!/bin/sh
    set -eu

    # Where I want my versions stored
    versionspath=~/.trashcan

    # The parameters we get from Syncthing
    folderpath="$1"
    filepath="$2"

    # First ensure the dir where we need to store the file exists
    outpath=$(dirname "$versionspath/$filepath")
    mkdir -p "$outpath"
    # Then move the file there
    mv -f "$folderpath/$filepath" "$versionspath/$filepath"

I must ensure that the script has execute permissions (``chmod 755
onlylatest.sh``), then configure Syncthing with command ``/Users/jb/bin/onlylatest.sh %FOLDER_PATH% %FILE_PATH%``

Let's assume I have a folder "default" in ~/Sync, and that within that folder
there is a file ``docs/letter.txt`` that is being replaced or deleted. The
script will be called as if I ran this from the command line::

    $ /Users/jb/bin/onlylatest.sh /Users/jb/Sync docs/letter.txt

The script will then move the file in question to
``~/.trashcan/docs/letter.txt``, replacing any previous version of that letter
that may already have been there.

Examples for Windows
~~~~~~~~~~~~~~~~~~~~

Move to a given folder using the command prompt (:abbr:`CMD`)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

On Windows we can use a batch script to perform the same "trash can"-like
behavior as mentioned above. I created the following script and saved it as
``C:\Users\mfrnd\Scripts\onlylatest.bat``.

.. code-block:: batch

    @echo off

    rem Enable UTF-8 encoding to deal with multilingual folder and file names
    chcp 65001

    rem We need command extensions for md to create intermediate folders in one go
    setlocal EnableExtensions

    rem Where I want my versions stored
    set "VERSIONS_PATH=%USERPROFILE%\.trashcan"

    rem The parameters we get from Syncthing, '~' removes quotes if any
    set "FOLDER_PATH=%~1"
    set "FILE_PATH=%~2"

    rem First ensure the dir where we need to store the file exists
    for %%F in ("%VERSIONS_PATH%\%FILE_PATH%") do set "OUTPUT_PATH=%%~dpF"
    if not exist "%OUTPUT_PATH%" md "%OUTPUT_PATH%" || exit /B

    rem Finally move the file, overwrite existing file if any
    move /Y "%FOLDER_PATH%\%FILE_PATH%" "%VERSIONS_PATH%\%FILE_PATH%"

Finally, I set ``"C:\Users\mfrnd\Scripts\onlylatest.bat" "%FOLDER_PATH%"
"%FILE_PATH%"`` as the command name in Syncthing.

Move to the Recycle Bin using PowerShell
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

We can use PowerShell to send files directly to the Recycle Bin, which
mimics the behaviour of deleting them using the Windows Explorer.
Firstly, create the following script and save it in your preferred
location, e.g. ``C:\Users\User\Scripts\SendToRecycleBin.ps1``.

.. code-block:: powershell

    # PowerShell has no native method to recycle files, so we use Visual
    # Basic to perform the operation. If succeeded, we also include the
    # recycled file in the Syncthing's DEBUG output.
    Add-Type -AssemblyName Microsoft.VisualBasic
    [Microsoft.VisualBasic.FileIO.FileSystem]::DeleteFile($args,'OnlyErrorDialogs','SendToRecycleBin')
    if ($?) {
      Write-Output ("Recycled " + $args + ".")
    }

Alternatively, the script can be expanded to send only deleted files to
the Recycle Bin, and permanently delete modified ones, which makes it
more consistent with how the Explorer works.

.. code-block:: powershell

    # PowerShell has no native method to recycle files, so we use Visual
    # Basic to perform the operation.
    Add-Type -AssemblyName Microsoft.VisualBasic

    # We need to test if a Syncthing .tmp file exists. If it does, we assume
    # a modification and delete the existing file. If if does not, we assume
    # a deletion and recycle the current file. If succeeded, we also include
    # the deleted/recycled file in the Syncthing's DEBUG output.
    if (Test-Path -LiteralPath ((Split-Path -Path $args) + "\~syncthing~" + (Split-Path -Path $args -Leaf) + ".tmp")) {
      [Microsoft.VisualBasic.FileIO.FileSystem]::DeleteFile($args,'OnlyErrorDialogs','DeletePermanently')
      if ($?) {
        Write-Output ("Deleted " + $args + ".")
      }
    } else {
      [Microsoft.VisualBasic.FileIO.FileSystem]::DeleteFile($args,'OnlyErrorDialogs','SendToRecycleBin')
      if ($?) {
        Write-Output ("Recycled " + $args + ".")
      }
    }

Finally, we set the command name in Syncthing to ``powershell.exe
-ExecutionPolicy Bypass -File "C:\Users\User\Scripts\SendToRecycleBin.ps1"
"%FOLDER_PATH%\%FILE_PATH%"``.

The only caveat that you should be aware of is that if your Syncthing
folder is located on a portable storage, such as a USB stick, or if you
have the Recycle Bin disabled, then the script will end up deleting all
files permanently.
