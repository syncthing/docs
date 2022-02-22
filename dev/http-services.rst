HTTP Utility Services API
=========================

These services are available via HTTP on the GUI listen address, but in contrast
to the REST API do not require any form of authentication or an API key.


GET /meta.js
------------

Returns a JavaScript snippet defining a ``metadata`` variable with information
about the serving Syncthing instance.  It is meant to be consumed directly from
an HTML ``<script>`` tag, such as for generating the needed CSRF token header
name.

.. code-block:: bash

    $ curl http://localhost:8384/meta.js
    var metadata = {"deviceID":"S7UKX27-GI7ZTXS-GC6RKUA-7AJGZ44-C6NAYEB-HSKTJQK-KJHU2NO-CWV7EQW"};


GET /qr/
--------

Encodes the content given in the ``text`` URL parameter into a QR code and
returns it as a black-and-white PNG image.  Can be used to encode device IDs for
quick exchange in a machine-readable optical format.

.. code-block:: bash

    $ curl http://localhost:8384/qr/?text=Hello%2C%20world%21 | display
