GET /rest/events
================

To receive events, perform a HTTP GET of ``/rest/events``.

To filter the event list, in effect creating a specific subscription for only
the desired event types, add a parameter ``events=EventTypeA,EventTypeB,...``
where the event types are any of the :ref:`event-types`.  If no filter is
specified, all events *except* :doc:`/events/localchangedetected` and
:doc:`/events/remotechangedetected` are included.

The optional parameter ``since=<lastSeenID>`` sets the ID of the last event
you've already seen. Syncthing returns a JSON encoded array of event objects,
starting at the event just after the one with this last seen ID. The default
value is 0, which returns all events. There is a limit to the number of events
buffered, so if the rate of events is high or the time between polling calls is
long some events might be missed. This can be detected by noting a discontinuity
in the event IDs.

If no new events are produced since ``<lastSeenID>``, the HTTP call blocks and
waits for new events to happen before returning. If ``<lastSeenID>`` is a
future ID, the HTTP call blocks until such ID is reached or timeouts. By
default it times out after 60 seconds returning an empty array. The time out
duration can be customized with the optional parameter ``timeout=<seconds>``.

To receive only a limited number of events, add the ``limit=<n>`` parameter with a
suitable value for ``n`` and only the *last* ``n`` events will be returned. This
can be used to catch up with the latest event ID after a disconnection for
example: ``/rest/events?since=0&limit=1``.


GET /rest/events/disk
=====================

This convenience endpoint provides the same event stream, but pre-filtered to show
only :doc:`/events/localchangedetected` and :doc:`/events/remotechangedetected`
event types.  The ``events`` parameter is not used.
