.. _event-api:

Event API
=========

Description
-----------

Syncthing provides a simple long polling interface for exposing events from the
core utility towards a GUI. To receive events, see :ref:`events-get`.

Event Structure
---------------

Each event is represented by an object similar to the following::

    {
        "id": 2,
        "globalID": 3,
        "type": "DeviceConnected",
        "time": "2014-07-13T21:04:33.687836696+02:00",
        "data": {
            "addr": "172.16.32.25:22000",
            "id": "NFGKEKE-7Z6RTH7-I3PRZXS-DEJF3UJ-FRWJBFO-VBBTDND-4SGNGVZ-QUQHJAG"
        }
    }

The top level keys ``id``, ``globalID``, ``time``, ``type`` and ``data`` are always present,
though ``data`` may be ``null``.

id
    A unique ID for this event on the events API. It always increases by 1: the first
    event generated has id ``1``, the next has id ``2`` etc. If this increases by
    more than 1, then one or more events have been skipped by the events API.
globalID
    A global ID for this event, across the events API, the audit log, and any other
    sources. It may increase by more than 1, but it will always be greater
    than or equal to the id.
time
    The time the event was generated.
type
    Indicates the type of (i.e. reason for) the event and is one of the event
    types below.
data
    An object containing optional extra information; the exact structure is
    determined by the event type.

.. _event-types:

Event Types
-----------

.. toctree::
    :maxdepth: 2
    :glob:

    ../events/*
