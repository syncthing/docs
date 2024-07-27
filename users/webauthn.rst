WebAuthn (Passkey) Authentication
=================================

Syncthing can be configured to allow GUI authentication using `WebAuthn <https://webauthn.guide>`_ (passkeys)
as an alternative to a password.

To enable WebAuthn, the GUI must use HTTPS (see config :ref:`gui.tls <gui-tls>`)
and must be served at exactly the address ``https://localhost:8384``,
unless configured otherwise as described in :ref:`webauthn-custom-gui-address`.


.. _webauthn-custom-gui-address:

Customizing the GUI address
---------------------------

The GUI address can be customized via the advanced GUI settings
:stconf:opt:`gui.webauthnRpId` and :stconf:opt:`gui.webauthnOrigin`.

If you access the GUI at some other address than ``https://localhost:<port>``,
you'll need to set the ``webauthnRpId`` setting to the domain name or a parent domain name of that address
and ``webauthnOrigin`` to the full address including scheme and port (except the default port), but not path.
For example, if you serve the GUI at the address ``https://syncthing.mydomain.org:8443/syncthing/gui``,
set ``webauthnRpId`` to one of ``mydomain.org`` or ``syncthing.mydomain.org``
and set ``webauthnOrigin`` to ``https://syncthing.mydomain.org:8443``.

WebAuthn authentication will be enabled if you have at least one `credential` registered.
A credential is a public-private key pair stored either on an external security key,
or a `platform credential` stored on your computer or phone.
Some platforms might sync platform credentials between devices signed into the same cloud account.

Use the settings GUI to register a new credential.

.. note::
   We use the term "passkey" more inclusively here than usual.
   For technical reasons, the term "passkey" usually means
   a credential that consumes storage space on the authenticator device.
   Some external security keys have limited storage capacity
   and therefore also have a limited capacity for storing passkeys.

   However, because Syncthing has only a single user account per installation,
   we can enable the same use cases as passkeys
   but with credentials that do not need to consume storage space.
   We therefore sometimes refer to WebAuthn credentials in Syncthing as "passkeys",
   because they enable most of the same UI flows as passkeys,
   even though they do not consume storage space on external security keys like passkeys usually do.
