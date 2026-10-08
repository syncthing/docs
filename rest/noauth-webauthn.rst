POST /rest/noauth/auth/webauthn-start
=====================================

.. versionadded:: TODO 1.28.0

Begins a WebAuthn authentication ceremony.

No parameters. The response is a JSON object
suitable as the parameter to ``navigator.credentials.get()``
after base64url decoding binary values - i.e.,
it contains a single attribute ``publicKey`` with the shape of
`PublicKeyCredentialRequestOptionsJSON
<https://w3c.github.io/webauthn/#dictdef-publickeycredentialrequestoptionsjson>`_.

Example response:

.. code-block:: json

    {
      "publicKey": {
        "challenge": "tlGvyFeTIOEPWVJLWZuiRCBEl2dVnC0ZvWt4Epmk-rk",
        "timeout": 120000,
        "rpId": "localhost",
        "allowCredentials": [
          {
            "type": "public-key",
            "id": "XW6tWsMNphd3rbESk4n9HEtd-h2MUdkHWQV6k2vuAzz8F9UoDTAVj3D-DWF_0z6q4R03mRJbtUPMDdNVr2Km-A",
            "transports": ["usb"]
          }
        ],
        "userVerification": "discouraged"
      }
    }


POST /rest/noauth/auth/webauthn-finish
======================================

.. versionadded:: TODO 1.28.0

Finishes a WebAuthn authentication ceremony, logging the user into the GUI if successful.

The request body is a JSON object containing two attributes:
a required ``credential`` attribute with the shape of
`AuthenticationResponseJSON
<https://w3c.github.io/webauthn/#dictdef-authenticationresponsejson>`_ - i.e.,
a ``PublicKeyCredential`` object
(the result of calling ``navigator.credentials.get()``)
with base64url encoded binary values,
and an optional Boolean ``stayLoggedIn`` attribute.
If ``stayLoggedIn`` is ``false`` or absent, the returned session cookie will expire with the browser session,
if ``true`` the cookie will persist for a time after the browser session ends.

The response on success is status 204 (No content) with no response body
and a ``Set-Cookie`` header containing the session cookie.

Example request:

.. code-block:: json

    {
        "credential": {
            "type": "public-key",
            "id": "XW6tWsMNphd3rbESk4n9HEtd-h2MUdkHWQV6k2vuAzz8F9UoDTAVj3D-DWF_0z6q4R03mRJbtUPMDdNVr2Km-A",
            "rawId": "XW6tWsMNphd3rbESk4n9HEtd-h2MUdkHWQV6k2vuAzz8F9UoDTAVj3D-DWF_0z6q4R03mRJbtUPMDdNVr2Km-A",
            "authenticatorAttachment": "cross-platform",
            "response": {
                "clientDataJSON": "eyJjaGFsbGVuZ2UiOiJ0bEd2eUZlVElPRVBXVkpMV1p1aVJDQkVsMmRWbkMwWnZXdDRFcG1rLXJrIiwib3JpZ2luIjoiaHR0cHM6Ly9sb2NhbGhvc3Q6ODM4NCIsInR5cGUiOiJ3ZWJhdXRobi5nZXQifQ",
                "authenticatorData": "SZYN5YgOjGh0NBcPZHZgW4_krrmihjLHmVzzuoMdl2MBAAABPA",
                "signature": "MEUCIQDjTizDIioXQFPrMih8UaAAo9R6sdYCMedrBxpSeYkd2wIgIMI-5h_CHJHa04EFN4HPsFO4nLCW8XR3iu5cRu5X4-w",
                "userHandle": null
            },
            "clientExtensionResults": {}
        },
        "stayLoggedIn": true
    }

Example response headers (other headers omitted):

.. code-block::

    HTTP/1.1 204 No Content
    Set-Cookie: sessionid-STLYOU4=banm5zZNHRAzJXmHUwWLjZmoJ9p4huCGscVSxnbXjgSR6CLuES3vQr2u5uX3Zt43; Path=/; Max-Age=604800; Secure
