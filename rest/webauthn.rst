.. _rest-webauthn-registration:

POST /rest/webauthn/register-start
----------------------------------

.. versionadded:: TODO 1.28.0

``POST .../register-start`` begins a WebAuthn registration ceremony
and ``POST .../register-finish`` finshes it,
adding the newly created credential to a list of pending credentials.
Pending credentials may be persisted
by including them in a request to ``POST /rest/webauthn/state``.

``POST .../register-start`` takes no parameters and returns a JSON object
suitable as the parameter to ``navigator.credentials.create()``
after base64url decoding binary values - i.e.,
it contains a single attribute ``publicKey`` with the shape of
`PublicKeyCredentialCreationOptionsJSON
<https://w3c.github.io/webauthn/#dictdef-publickeycredentialcreationoptionsjson>`_.

Example response:

.. code-block:: json

    {
      "publicKey": {
        "rp": {
          "name": "Syncthing @ DEVICENAME",
          "id": "localhost"
        },
        "user": {
          "name": "asdf",
          "displayName": "asdf",
          "id": "4_KyxKWr6x2KvB3GGHYLkjmn1M6xTip5ITZQUgaUzJW5e023M0j4NBOkgR-4aQarM7RRCv7TGkmOD53kQBPhLQ"
        },
        "challenge": "VopAfwRL52Jc1E_H0yi-kEmb59s4IfJ1UN2zSjY_5CA",
        "pubKeyCredParams": [
          { "type": "public-key", "alg": -7 },
          { "type": "public-key", "alg": -35 },
          { "type": "public-key", "alg": -36 },
          { "type": "public-key", "alg": -257 },
          { "type": "public-key", "alg": -258 },
          { "type": "public-key", "alg": -259 },
          { "type": "public-key", "alg": -37 },
          { "type": "public-key", "alg": -38 },
          { "type": "public-key", "alg": -39 },
          { "type": "public-key", "alg": -8 }
        ],
        "timeout": 300000,
        "authenticatorSelection": {
          "requireResidentKey": false,
          "userVerification": "preferred"
        }
      }
    }


POST /rest/webauthn/register-finish
-----------------------------------

.. versionadded:: TODO 1.28.0

``POST .../register-start`` begins a WebAuthn registration ceremony
and ``POST .../register-finish`` finshes it,
adding the newly created credential to a list of pending credentials.
Pending credentials may be persisted
by including them in a request to ``POST /rest/webauthn/state``.

``POST .../register-finish`` takes a request body with the shape of
`RegistrationResponseJSON
<https://w3c.github.io/webauthn/#dictdef-registrationresponsejson>`_ - i.e.,
a ``PublicKeyCredential`` object
(the result of calling ``navigator.credentials.create()``)
with base64url encoded binary values.
It returns a JSON representation of the pending registered credential,
which can be added as an element of the ``credentials`` array
in the JSON body of a `/rest/webauthn/state <rest-webauthn-state_>`_ request.


.. note::
   WebAuthn credentials are "config-like"
   and are managed in the "GUI" section of the settings GUI.
   In order to follow the convention of changes being pending
   until the user presses the "save" button,
   this API call does not yet permanently save the returned credential.
   To permanently save the credential and activate it as a login option,
   the returned object must be saved by appending it to the ``credentials`` list
   in the JSON body of a `/rest/webauthn/state <rest-webauthn-state_>`_ request.


Example request:

.. code-block:: json

    {
      "type": "public-key",
      "id": "VxT1FCv2nrNwCTGmOnNDoUAY3p6RJyvBzF7y-dsD5Ll73Mve76m9okIX7C5cDf2elKxtBRRmcnMUuVnPk3TUuA",
      "rawId": "VxT1FCv2nrNwCTGmOnNDoUAY3p6RJyvBzF7y-dsD5Ll73Mve76m9okIX7C5cDf2elKxtBRRmcnMUuVnPk3TUuA",
      "authenticatorAttachment": "cross-platform",
      "response": {
        "clientDataJSON": "eyJjaGFsbGVuZ2UiOiJWb3BBZndSTDUySmMxRV9IMHlpLWtFbWI1OXM0SWZKMVVOMnpTallfNUNBIiwib3JpZ2luIjoiaHR0cHM6Ly9sb2NhbGhvc3Q6ODM4NCIsInR5cGUiOiJ3ZWJhdXRobi5jcmVhdGUifQ",
        "attestationObject": "o2NmbXRkbm9uZWdhdHRTdG10oGhhdXRoRGF0YVjESZYN5YgOjGh0NBcPZHZgW4_krrmihjLHmVzzuoMdl2NFAAAABAAAAAAAAAAAAAAAAAAAAAAAQFcU9RQr9p6zcAkxpjpzQ6FAGN6ekScrwcxe8vnbA-S5e9zL3u-pvaJCF-wuXA39npSsbQUUZnJzFLlZz5N01LilAQIDJiABIVgg1ZEbVe7_o93_XuuRl98qhHa-cmsJrpL_Rw5wrpEqgqIiWCCpp0NlSL-xBR9lDc5Th5Y1WsGLs0vS5jgjxh_kS1D_0Q",
        "transports": ["nfc", "usb"]
      },
      "clientExtensionResults": {}
    }

Example response:

.. code-block:: json

    {
      "id": "VxT1FCv2nrNwCTGmOnNDoUAY3p6RJyvBzF7y-dsD5Ll73Mve76m9okIX7C5cDf2elKxtBRRmcnMUuVnPk3TUuA==",
      "rpId": "localhost",
      "nickname": "",
      "publicKeyCose": "pQECAyYgASFYINWRG1Xu_6Pd_17rkZffKoR2vnJrCa6S_0cOcK6RKoKiIlggqadDZUi_sQUfZQ3OU4eWNVrBi7NL0uY4I8Yf5EtQ_9E=",
      "signCount": 4,
      "transports": ["nfc", "usb"],
      "requireUv": false,
      "createTime": "2024-07-21T15:24:01+02:00",
      "lastUseTime": "2024-07-21T15:24:01+02:00"
    }


.. _rest-webauthn-state:

GET /rest/webauthn/state
------------------------

Returns the state of currently registered WebAuthn credentials.
The credential data model is described `below <rest-webauthn-credential_>`_.

.. versionadded:: TODO 1.28.0

Example ``GET`` response:

.. code-block:: json

    {
      "credentials": [
        {
          "id": "cTVm-CWdvbMOX7v4QdUxJgPZ5TWpFuliLDWNcI9chOw02DBJcZjmvHDOwpGEwxS6Lk6H8eikYbystBghaJuq-g==",
          "rpId": "localhost",
          "nickname": "Security key",
          "publicKeyCose": "pQECAyYgASFYIC9CP0p82dtJiRKYfUGSYeVaccOsNAmYgIz-EAl1GzbyIlggtcbhDVA8bUpjK_GH3QpGL9i_y9GfoTM1pg0jyEBf88M=",
          "signCount": 644,
          "transports": ["nfc", "usb"],
          "requireUv": true,
          "createTime": "2024-07-13T13:58:07Z",
          "lastUseTime": "2024-07-21T12:55:43Z"
        },
        {
          "id": "4gvuaMwVUnv6a0cNRzUm4hkbeTgVsf7HUBbgXBoSB9A57AagRbZvWCUMaBjroYhnWBubRq_29uo4CGFtfWwpdg==",
          "rpId": "localhost",
          "nickname": "",
          "publicKeyCose": "pQECAyYgASFYIGxYkCaHAkelm7Mu5JGtaQFdcAPqlWlhOFuGah4eom7KIlggtvPzU9tMFtxElKqr3zXO2YZAlIKAbUOvbTA93tx39Rc=",
          "signCount": 115,
          "transports": ["nfc", "usb"],
          "requireUv": false,
          "createTime": "2024-07-13T14:07:20Z",
          "lastUseTime": "2024-07-13T15:36:44Z"
        }
      ]
    }


POST /rest/webauthn/state
-------------------------

.. versionadded:: TODO 1.28.0

Updates the WebAuthn state to match the body,
except the following rules are applied to the new ``credentials`` value:

- Each item must have an ``id`` that already exists in the currently stored ``credentials`` value,
  or in the list of pending credentials stored by `/rest/webauthn/register-finish <rest-webauthn-registration_>`_.
  Items with any other ``id`` are ignored.
- For each already existing item, all attributes except ``nickname`` and ``requireUv`` are ignored.

The credential data model is described in the `next section <rest-webauthn-credential_>`_.

Assuming that ``id: "VxT1FCv2..."``
was previously returned from `/rest/webauthn/register-finish <rest-webauthn-registration_>`_
as in the example above,
this example request would change the nickname of "Security key" to "My security key",
delete the credential with ``id: "4gvuaMwV..."``
and persist the pending credential with the nickname "New security key":

.. code-block:: json

    {
        "credentials": [
            {
              "id": "cTVm-CWdvbMOX7v4QdUxJgPZ5TWpFuliLDWNcI9chOw02DBJcZjmvHDOwpGEwxS6Lk6H8eikYbystBghaJuq-g==",
              "rpId": "localhost",
              "nickname": "My security key",
              "publicKeyCose": "pQECAyYgASFYIC9CP0p82dtJiRKYfUGSYeVaccOsNAmYgIz-EAl1GzbyIlggtcbhDVA8bUpjK_GH3QpGL9i_y9GfoTM1pg0jyEBf88M=",
              "signCount": 644,
              "transports": ["nfc", "usb"],
              "requireUv": true,
              "createTime": "2024-07-13T13:58:07Z",
              "lastUseTime": "2024-07-21T12:55:43Z"
            },
            {
              "id": "VxT1FCv2nrNwCTGmOnNDoUAY3p6RJyvBzF7y-dsD5Ll73Mve76m9okIX7C5cDf2elKxtBRRmcnMUuVnPk3TUuA==",
              "rpId": "localhost",
              "nickname": "New security key",
              "publicKeyCose": "pQECAyYgASFYINWRG1Xu_6Pd_17rkZffKoR2vnJrCa6S_0cOcK6RKoKiIlggqadDZUi_sQUfZQ3OU4eWNVrBi7NL0uY4I8Yf5EtQ_9E=",
              "signCount": 4,
              "transports": ["nfc", "usb"],
              "requireUv": false,
              "createTime": "2024-07-21T15:24:01+02:00",
              "lastUseTime": "2024-07-21T15:24:01+02:00"
            }
        ]
    }


.. _rest-webauthn-credential:

The credential data model
-------------------------

Items in the ``credentials`` field of the `WebAuthn state <rest-webauthn-state_>`_ have the following attributes:

- ``id``

  The base64url-encoded `credential ID <https://www.w3.org/TR/webauthn/#credential-id>`_ of the credential.
  This is created by the authenticator and cannot be changed.

- ``rpId``

  The value of the :stconf:opt:`gui.webauthnRpId` setting in effect at the time this credential was created.
  This is set automatically and cannot be changed.

  If :stconf:opt:`gui.webauthnRpId` is changed after creating a credential,
  the credential can no longer be used unless the :stconf:opt:`gui.webauthnRpId` value is restored.
  This attribute is used in the settings GUI to highlight credentials that cannot currently be used
  and show what :stconf:opt:`gui.webauthnRpId` to restore to in order to make them usable again.

- ``publicKeyCose``

  The base64url-encoded public key of the credential, in `COSE_Key format <https://www.w3.org/TR/webauthn/#sctn-encoded-credPubKey-examples>`_.
  This is created by the authenticator and cannot be changed.

- ``signCount``

  The `signature counter <https://www.w3.org/TR/webauthn/#signature-counter>`_ of the credential.
  A decrease in the signature counter may indicate that the credential has been cloned.
  Syncthing displays a warning if this happens, but does not otherwise act on it.

- ``nickname``

  A user-chosen nickname for the credential.
  If empty or not set, the GUI will use the abbreviated credential ID (``id``) as the name of the credential.
  This can be edited in the settings GUI.

- ``requireUv``

  If set to ``true``, this credential requires `User Verification (UV) <https://www.w3.org/TR/webauthn/#user-verification>`_,
  for example a PIN or a biometric.
  This means that logging in with this credential is two-factor authentication (2FA):
  something you have (the credential private key)
  combined with something you know (a PIN) or something you are (a biometric).

  This can be enabled or disabled in the settings GUI.

  .. note::

    The PIN or biometric is not sent to the server -
    Syncthing does **not** collect or store biometric information.
    Instead, the PIN or biometric is only verified locally by your authenticator
    (for example, a USB security key or a smartphone) before unlocking the passkey for login.

- ``transports``

  A list of hints the browser may use to determine how to communicate with the authenticator
  that holds the private key for this credential -
  for example, this may be ``["nfc", "usb"]`` if the credential is stored on a USB security key
  or ``["hybrid", "internal"]`` if the credential is stored on a smartphone or laptop.

  This is set automatically and cannot not be changed.
  Changing it could make the credential unusable,
  since the browser might conclude it has no way to communicate with the authenticator
  if none of the transports listed here is available on the platform.
  If this happens, you can attempt to make the credential usable again by deleting the attribute.

- ``createTime`` and ``lastUseTime``

  Timestamps recording when this credential was created and when it was last used to log in to the GUI.
  Used only to help the user identify and distinguish credentials in the GUI;
  not used for any security decisions.
