GET /rest/db/completion
=======================

Returns the completion percentage (0 to 100) and byte / item counts. Takes
optional ``device`` and ``folder`` parameters:

- ``folder`` specifies the folder ID to calculate completion for. An empty
  or absent ``folder`` parameter means all folders as an aggregate.

- ``device`` specifies the device ID to calculate completion for. An empty
  or absent ``device`` parameter means the local device.

If a device is specified but no folder, completion is calculated for all
folders shared with that device.

Example Queries
---------------

Completion status for folder ``abcd-1234`` on device ``I6KAH76-...-3PSROAU``::

    /rest/db/completion?folder=abcd-1234&device=I6KAH76-...-3PSROAU

Aggregated completion status for device ``I6KAH76-...-3PSROAU`` (all folders shared with them)::

    /rest/db/completion?device=I6KAH76-...-3PSROAU

Completion status for folder ``abcd-1234`` on the local device::

    /rest/db/completion?folder=abcd-1234

Aggregated completion status for all folders on the local device::

    /rest/db/completion

Example Response
----------------

.. code-block:: json

    {
      "completion": 99.9937565835,
      "globalBytes": 156793013575,
      "needBytes": 9789241,
      "globalItems": 7823,
      "needItems": 412,
      "needDeletes": 0,
      "remoteState": "valid",
      "sequence": 12
    }

.. versionadded:: 1.8.0

  The ability to aggregate multiple folders by leaving out the folder ID.
  Querying data for the local device by leaving out the device ID. Returning
  the ``globalItems`` counter in the response.

.. versionadded:: 1.20.0

  Indication whether the remote device has accepted the folder (shares it with
  us) as well, and whether it is paused.  The ``remoteState`` field is
  meaningless for aggregated responses, ``unknown`` when the remote device is
  not connected.  Otherwise it can be either ``paused``, ``notSharing``, or
  ``valid`` if the remote is sharing back.
