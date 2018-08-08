GET /rest/db/status
===================

Returns information about the current status of a folder.

Parameters: ``folder``, the ID of a folder.

.. code-block:: bash

    {
      "globalBytes": 0,
      "globalDeleted": 0,
      "globalDirectories": 0,
      "globalFiles": 0,
      "globalSymlinks": 0,
      "ignorePatterns": false,
      "inSyncBytes": 0,
      "inSyncFiles": 0,
      "invalid": "",
      "localBytes": 0,
      "localDeleted": 0,
      "localDirectories": 0,
      "localFiles": 0,
      "localSymlinks": 0,
      "needBytes": 0,
      "needDeletes": 0,
      "needDirectories": 0,
      "needFiles": 0,
      "needSymlinks": 0,
      "pullErrors": 0,
      "receiveOnlyChangedBytes": 0,
      "receiveOnlyChangedDeletes": 0,
      "receiveOnlyChangedDirectories": 0,
      "receiveOnlyChangedFiles": 0,
      "receiveOnlyChangedSymlinks": 0,
      "sequence": 0,
      "state": "idle",
      "stateChanged": "2018-08-08T07:04:57.301064781+02:00",
      "version": 0
    }

The various fields have the following meaning:

global*:
  Data in the cluster latest version.

inSync*:
  Data that is locally the same as the cluster latest version.

local*:
  Data that is locally present, regardless of whether it's the same or different version as the cluster latest version.

need*:
  Data that is needed to become up to date with the cluster latest version (i.e., data that is out of sync).

receiveOnlyChanged*:
  Data that has changed locally in a receive only folder, and thus not been sent to the cluster.

invalid:
  Deprecated, always empty.

pullErrors:
  The number of files that failed to sync during the last sync operations.

sequence:
  The current folder sequence number.

state:
  The current folder state.

stateChanged:
  When the folder state last changed.

version:
  Deprecated, equivalent to the sequence number.

.. note::
  This is an expensive call, increasing CPU and RAM usage on the device. Use sparingly.
