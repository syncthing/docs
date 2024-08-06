WebAuthn (Passkey) Authentication
=================================

Syncthing can be configured to allow GUI authentication using `WebAuthn <https://webauthn.guide>`_ (passkeys)
as an alternative to a password.
WebAuthn offers a passwordless login experience that some users may find preferable.

To enable WebAuthn, the GUI must use HTTPS (see config :ref:`gui.tls <gui-tls>`)
and must be served at exactly the address ``https://localhost:8384``,
unless configured otherwise as described in :ref:`webauthn-custom-gui-address`.

WebAuthn authentication will be enabled if you have at least one `credential`, also called a `passkey`, registered.
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


.. _webauthn-require2fa:

The 2FA setting
---------------

Each credential (passkey) has a checkbox setting labeled "2FA" in the GUI.
When checked, Syncthing will enforce that this credential uses
`two-factor authentication (2FA) <https://en.wikipedia.org/wiki/Multi-factor_authentication>`_.
The technical name for this is `User Verification (UV) <https://www.w3.org/TR/webauthn/#user-verification>`_.

For example:

- If the credential is stored on a smartphone,
  the phone may prompt for screen unlock to authenticate you to the phone before unlocking the passkey.
  This could be a PIN, swipe pattern, fingerprint or face recognition, according to the phone's settings.

  Smartphones typically always require 2FA,
  so this setting may not make a noticeable difference for smartphone-based credentials.

- An external security key may prompt for a PIN configured on the security key,
  or use a built-in fingerprint reader.
  With the 2FA setting disabled, you would only need to plug in the security key
  and usually press a button on it,
  but would need no additional factor beyond possessing the security key.

  Some older models of security keys do not support 2FA.

.. note::

  No biometrics, PIN or other data is sent to the server -
  Syncthing does **not** collect or store biometric information.
  Instead, the second factor is only verified locally by your authenticator
  (for example, a USB security key or a smartphone) before unlocking the passkey for login.

If you have some credentials with 2FA enabled and some with 2FA disabled,
you might get prompted for 2FA even when using a credential that doesn't require it.
This is because Syncthing doesn't know beforehand which credential you're going to use,
so it needs to pessimistically request 2FA in case it is required for the credential you choose.


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
