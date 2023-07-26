REST API
========

Syncthing exposes a REST interface over HTTP on the GUI port. This is used by
the GUI (from Javascript) and can be used by other processes wishing to control
Syncthing. In most cases both the input and output data is in JSON format. The
interface is subject to change.

API Key
-------

To use the REST API an API key must be set and used. The API key can be
generated in the GUI, or set in the ``configuration/gui/apikey`` element in
the configuration file. To use an API key, set the request header
``X-API-Key`` to the API key value, or set it as a ``Bearer`` token in the
``Authorization`` header. For example, ``curl -X POST -H "X-API-Key: abc123"
http://localhost:8384/rest/...`` or ``curl -X POST -H "Authorization: Bearer
abc123" http://localhost:8384/rest/...`` can be used to invoke with ``curl``
(add ``-k`` flag when using HTTPS with a Syncthing generated or self signed
certificate).

One exception to this requirement is ``/rest/noauth``, you do not need an API
key to use those endpoints. This way third-party devices and services can do
simple calls that don't expose sensitive information without having to expose
your API key.

.. _rest-pagination:

Result Pagination
-----------------

Some `GET` endpoints take optional ``page`` and ``perpage`` arguments for
pagination.  No more than ``perpage`` (defaults to 65536 if not given) result
entries are returned in an array.  To access further entries, passing the
``page`` parameter will advance in the results by that many pages.  The actually
used parameters are always returned as attributes in the response object.

System Endpoints
----------------

.. toctree::
   :maxdepth: 1
   :glob:

   ../rest/system-*

Config Endpoints
----------------

.. toctree::
   :maxdepth: 1

   /rest/config/... <../rest/config.rst>

Cluster Endpoints
-----------------

Concerns the mesh network structure.

.. toctree::
   :maxdepth: 1
   :glob:

   ../rest/cluster-*

Folder Endpoints
----------------

Runtime state of the individual shared folders.

.. toctree::
   :maxdepth: 1
   :glob:

   ../rest/folder-*

Database Endpoints
------------------

.. toctree::
   :maxdepth: 1
   :glob:

   ../rest/db-*

Event Endpoints
---------------

.. toctree::
   :maxdepth: 1
   :glob:

   ../rest/events-*

Statistics Endpoints
--------------------

.. toctree::
   :maxdepth: 1
   :glob:

   ../rest/stats-*

Misc Services Endpoints
-----------------------

.. toctree::
   :maxdepth: 1
   :glob:

   ../rest/svc-*

Debug Endpoints
---------------

.. toctree::
   :maxdepth: 1

   /rest/debug/... <../rest/debug.rst>

Noauth Endpoints
----------------

Calls that do not require authentication.

.. toctree::
   :maxdepth: 1
   :glob:

   ../rest/noauth-*