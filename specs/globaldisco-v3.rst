.. _globaldisco-v3:

Global Discovery v3
===================

Announcements
-------------

A device should announce itself at startup. It does this by an HTTPS POST to
the announce server URL. Standard discovery currently requires the path to be
"/v2/", yet this can be up to the discovery server. The POST has a JSON payload
listing connection addresses (if any)::

	{
		addresses: ["tcp://192.0.2.45:22000", "tcp://:22202", "relay://192.0.2.99:22028"],
	}

It's OK for the "addresses" field to be either the empty list (``[]``),
``null``, or missing entirely. An announcement with the field missing
or empty is however not useful...

Any empty or unspecified IP addresses (i.e. addresses like ``tcp://:22000``,
``tcp://0.0.0.0:22000``, ``tcp://[::]:22000``) are interpreted as referring to
the source IP address of the announcement.

The device ID of the announcing device is not part of the announcement.
Instead, the server requires that the client perform certificate
authentication. The device ID is deduced from the presented certificate.

The server response is empty, with code ``204`` (No Content) on success. If no
certificate was presented, status ``403`` (Forbidden) is returned. If the
posted data doesn't conform to the expected format, ``400`` (Bad Request) is
returned.

In successful responses, the server may return a ``Reannounce-After`` header
containing the number of seconds after which the client should perform a new
announcement.

In error responses, the server may return a ``Retry-After`` header containing
the number of seconds after which the client should retry.

Performing announcements significantly more often than indicated by the
``Reannounce-After`` or ``Retry-After`` headers may result in the client being
throttled. In such cases the server may respond with status code ``429`` (Too
Many Requests).

Queries
-------

Queries are performed as HTTPS GET requests to the announce server URL. The
requested device ID is passed as the query parameter "device", in canonical
string form, i.e. ``https://discovery.syncthing.net/?device=ABC12345-....``

Successful responses will have status code ``200`` (OK) and carry a JSON payload
of the same format as the announcement above. The response will not contain
empty or unspecified addresses.

If the "device" query parameter is missing or malformed, the status code 400
(Bad Request) is returned.

If the device ID is of a valid format but not found in the registry, 404 (Not
Found) is returned.

If the client has exceeded a rate limit, the server may respond with 429 (Too
Many Requests).

Authentication
--------------

Global discovery is spoken over HTTPS and is protected against attackers in
the same manner as other HTTPS traffic. However, there are a few Syncthing
specific considerations on top of this. As mentioned above, for
announcements the client must provide a certificate to prove ownership of
the announced device ID.

In addition, Syncthing has a mechanism to verify the identity of the
discovery server.  While this would normally be accomplished by using a CA
signed certificate, Syncthing often runs in environments with outdated or
simply nonexistent root CA bundles. Instead, Syncthing can verify the
discovery server certificate fingerprint using the device ID mechanism. This
is certificate pinning and conveyed in the Syncthing configuration as a
synthetic "id" parameter on the discovery server URL:
``https://discovery.syncthing.net/?id=...``. The "id" parameter is not, in
fact, sent to the discovery server - it's used by Syncthing itself to know
which certificate to expect on the server side.

The public discovery network uses this authentication mechanism instead of
CA signed certificates.

The discovery server prints its certificate ID in this manner on startup.
