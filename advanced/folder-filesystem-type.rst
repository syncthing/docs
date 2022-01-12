filesystemType
==============

Syncthing has an internal abstraction for file system access with different
available implementations.  They can be configured per folder, with the
following possible values:

``basic`` (default)
    To be used if the folder is intended to store real data. Do not change
    unless you are a developer and want to test things.

``fake``
    A fake file system type to be used for testing, e.g. when you want to create
    a folder with a :abbr:`TB (Terabyte)` or more of random files that can be
    synced somewhere, or an infinitely large folder to sync files into.

    It has pseudorandom properties, i.e. data read from one fakefs can be
    written into another fakefs, read back, and it will look consistent, without
    any real data actually being stored.

    To create an empty file system, use

    .. code-block::

        <folder id="default" path="whatever" ...>
            <filesystemType>fake</filesystemType>

    You can also specify that it should be prefilled with files,

    .. code-block::

       <folder id="default" path="whatever?size=2000000" ...>
           <filesystemType>fake</filesystemType>

    which will create a file system filled with 2 :abbr:`TB (Terabyte)` of
    random data that can be scanned and synced. The prefilled data is based on a
    deterministic seed, so you can index it, restart Syncthing, and the index
    will still be correct for all the stored data.

    Check the source of `fakefs.go
    <https://github.com/syncthing/syncthing/blob/main/lib/fs/fakefs.go>`_ for
    more options and details.
