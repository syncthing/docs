.. _event-api:

Event API
=========

Description
-----------

Syncthing provides a simple long polling interface for exposing events from the
core utility towards a GUI.

To receive events, perform a HTTP GET of ``/rest/events`` or
``/rest/events/disk``. The latter returns only :ref:`local-change-detected` and
:ref:`remote-change-detected` events, the former all other events unless filtered.

To filter the event list, in effect creating a specific subscription for
only the desired event types, add a parameter
``events=EventTypeA,EventTypeB,...`` where the event types are any from the
list below.

The optional parameter ``since=<lastSeenID>`` sets the ID of the last event
you've already seen. Syncthing returns a JSON encoded array of event objects,
starting at the event just after the one with this last seen ID. The default
value is 0, which returns all events. There is a limit to the number of events
buffered, so if the rate of events is high or the time between polling calls is
long some events might be missed. This can be detected by noting a discontinuity
in the event IDs.

If no new events are produced since ``<lastSeenID>``, the HTTP call blocks and
waits for new events to happen before returning. By default it times out after
60 seconds returning an empty array. The time out duration can be customized
with the optional parameter ``timeout=seconds``.

To receive only a limited number of events, add the ``limit=n`` parameter with a
suitable value for ``n`` and only the *last* ``n`` events will be returned. This
can be used to catch up with the latest event ID after a disconnection for
example: ``/rest/events?since=0&limit=1``.

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

Event Types
-----------

.. toctree::
    :maxdepth: 2
    :glob:

    ../events/*
