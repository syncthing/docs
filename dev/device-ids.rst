.. _device-ids:

Understanding Device IDs
========================

Description
-----------

Every device is identified by a device ID. The device ID is used for address
resolution, authentication and authorization. The term "device ID" could
interchangeably have been "key ID" since the device ID is a direct property of
the public key in use.

Keys
----

To understand device IDs we need to look at the underlying mechanisms. At first
startup, Syncthing will create a public/private keypair.

Currently this is a 3072 bit RSA key. The keys are saved in the form of the
private key (``key.pem``) and a self signed certificate (``cert.pem``). The self
signing part doesn't actually add any security or functionality as far as
Syncthing is concerned but it enables the use of the keys in a standard TLS
exchange.

The typical certificate will look something like this, inspected with
``openssl x509``::

    Certificate:
        Data:
            Version: 3 (0x2)
            Serial Number: 0 (0x0)
            Signature Algorithm: sha1WithRSAEncryption
            Issuer: CN=syncthing
            Validity
                Not Before: Mar 30 21:10:52 2014 GMT
                Not After : Dec 31 23:59:59 2049 GMT
            Subject: CN=syncthing
            Subject Public Key Info:
                Public Key Algorithm: rsaEncryption
                RSA Public Key: (3072 bit)
                    Modulus (3072 bit):
                        00:da:83:8a:c0:95:af:0a:42:af:43:74:65:29:f2:
                        30:e3:b9:12:d2:6b:70:93:da:0b:7b:8a:1e:e5:79:
                        ...
                        99:09:4c:a9:7b:ba:4a:6a:8b:3b:e6:e7:c7:2c:00:
                        90:aa:bc:ad:94:e7:80:95:d2:1b
                    Exponent: 65537 (0x10001)
            X509v3 extensions:
                X509v3 Key Usage: critical
                    Digital Signature, Key Encipherment
                X509v3 Extended Key Usage:
                    TLS Web Server Authentication, TLS Web Client Authentication
                X509v3 Basic Constraints: critical
                    CA:FALSE
        Signature Algorithm: sha1WithRSAEncryption
            68:72:43:8b:83:61:09:68:f0:ef:f0:43:b7:30:a6:73:1e:a8:
            d9:24:6c:2d:b4:bc:c9:e8:3e:0b:1e:3c:cc:7a:b2:c8:f1:1d:
            ...
            88:7e:e2:61:aa:4c:02:e3:64:b0:da:70:3a:cd:1c:3d:86:db:
            df:54:b9:4e:be:1b

We can see here that the certificate is little more than a container for the
public key; the serial number is zero and the Issuer and Subject are both
"syncthing" where a qualified name might otherwise be expected.

An advanced user could replace the ``key.pem`` and ``cert.pem`` files with a
keypair generated directly by the ``openssl`` utility or other mechanism.

Device IDs
----------

To form a device ID the SHA-256 hash of the certificate data in DER form is
calculated. This means the hash covers all information under the
``Certificate:`` section above.

The hashing results in a 256 bit hash which we encode using base32. Base32
encodes five bits per character so we need 256 / 5 = 51.2 characters to encode
the device ID. This becomes 52 characters in practice, but 52 characters of
base32 would decode to 260 bits which is not a whole number of bytes. The
base32 encoding adds padding to 280 bits (the next multiple of both 5 and 8
bits) so the resulting ID looks something like::

    MFZWI3DBONSGYYLTMRWGC43ENRQXGZDMMFZWI3DBONSGYYLTMRWA====

The padding (``====``) is stripped away, the device ID split into four
groups, and `check
digits <https://forum.syncthing.net/t/v0-9-0-new-node-id-format/478>`__
are added for each group. For presentation purposes the device ID is
grouped with dashes, resulting in the final value::

    MFZWI3D-BONSGYC-YLTMRWG-C43ENR5-QXGZDMM-FZWI3DP-BONSGYY-LTMRWAD

Connection Establishment
~~~~~~~~~~~~~~~~~~~~~~~~

Now we know what device IDs are, here's how they are used in Syncthing. When
you add a device ID to the configuration, Syncthing will attempt to
connect to that device. The first thing we need to do is figure out the IP and
port to connect to. There are three possibilities here:

-  The IP and port can be set statically in the configuration. The IP
   can equally well be a host name, so if you have a static IP or a
   dynamic DNS setup this might be a good option.

-  Using local discovery, if enabled. Every Syncthing instance on a LAN
   periodically broadcasts information about itself (device ID, address,
   port number). If we've seen one of these broadcasts for a given
   device ID that's where we try to connect.

-  Using global discovery, if enabled. Every Syncthing instance
   announces itself to the global discovery service (device ID and
   external port number - the internal address is not announced to the
   global server). If we don't have a static address and haven't seen
   any local announcements the global discovery server will be queried
   for an address.

Once we have an address and port a TCP connection is established and a TLS
handshake performed. As part of the handshake both devices present their
certificates. Once the handshake has completed and the peer certificate is
known, the following steps are performed:

#. Calculate the remote device ID by processing the received certificate as above.

#. Weed out a few possible misconfigurations - i.e. if the device ID is
   that of the local device or of a device we already have an active
   connection to. Drop the connection in these cases.

#. Verify the remote device ID against the configuration. If it is not a
   device ID we are expecting to talk to, drop the connection.

#. Verify the certificate ``CommonName`` against the configuration. By
   default, we expect it to be ``syncthing``, but when using custom
   certificates this can be changed.

#. If everything checks out so far, accept the connection.

An Aside About Collisions
-------------------------

The SHA-256 hash is cryptographically collision resistant. This means
that there is no way that we know of to create two different messages
with the same hash.

You can argue that of course there are collisions - there's an infinite
amount of inputs and a finite amount of outputs - so by definition there
are infinitely many messages that result in the same hash.

I'm going to quote `stack
overflow <http://stackoverflow.com/questions/4014090/is-it-safe-to-ignore-the-possibility-of-sha-collisions-in-practice>`__
here:

    The usual answer goes thus: what is the probability that a rogue
    asteroid crashes on Earth within the next second, obliterating
    civilization-as-we- know-it, and killing off a few billion people ?
    It can be argued that any unlucky event with a probability lower
    than that is not actually very important.

    If we have a "perfect" hash function with output size n, and we have
    p messages to hash (individual message length is not important),
    then probability of collision is about p2/2n+1 (this is an
    approximation which is valid for "small" p, i.e. substantially
    smaller than 2n/2). For instance, with SHA-256 (n=256) and one
    billion messages (p=10^9) then the probability is about 4.3\*10^-60.

    A mass-murderer space rock happens about once every 30 million years
    on average. This leads to a probability of such an event occurring
    in the next second to about 10^-15. That's 45 orders of magnitude
    more probable than the SHA-256 collision. Briefly stated, if you
    find SHA-256 collisions scary then your priorities are wrong.

It's also worth noting that the property of SHA-256 that we are using is not
simply collision resistance but resistance to a preimage attack, i.e. even if
you can find two messages that result in a hash collision that doesn't help you
attack Syncthing (or TLS in general). You need to create a message that hashes
to exactly the hash that my certificate already has or you won't get in.

Note also that it's not good enough to find a random blob of bits that happen to
have the same hash as my certificate. You need to create a valid DER-encoded,
signed certificate that has the same hash as mine. The difficulty of this is
staggeringly far beyond the already staggering difficulty of finding a SHA-256
collision.

Problems and Vulnerabilities
----------------------------

As far as I know, these are the issues or potential issues with the
above mechanism.

Discovery Spoofing
~~~~~~~~~~~~~~~~~~

Currently, neither the local nor global discovery mechanism is protected
by crypto. This means that any device can in theory announce itself for
any device ID and potentially receive connections for that device.

This could be a denial of service attack (we can't find the real device
for a given device ID, so can't connect to it and sync). It could also
be an intelligence gathering attack; if I spoof a given ID, I can see
which devices try to connect to it.

It could be mitigated in several ways:

-  Announcements could be signed by the device private key. This
   requires already having the public key to verify.

-  Announcements to the global announce server could be done using TLS,
   so the server calculates the device ID based on the certificate
   instead of trusting the device to tell the truth.

-  The user could statically configure IP or host name for the devices.

-  The user could run a trusted global server.

It's something we might want to look at at some point, but not a huge
problem as I see it.

Long Device IDs are Painful
~~~~~~~~~~~~~~~~~~~~~~~~~~~

It's a mouthful to read over the phone, annoying to type into an SMS or even
into a computer. And it needs to be done twice, once for each side.

This isn't a vulnerability as such, but a user experience problem. There are
various possible solutions:

-  Use shorter device IDs with verification based on the full ID ("You
   entered MFZWI3; I found and connected to a device with the ID
   MFZWI3-DBONSG-YYLTMR-WGC43E-NRQXGZ-DMMFZW-I3DBON-SGYYLT-MRWA, please
   confirm that this is correct").

-  Use shorter device IDs with an out of band authentication, a la
   Bluetooth pairing. You enter a one time PIN into Syncthing and give
   that PIN plus a short device ID to another user. On initial connect,
   both sides verify that the other knows the correct PIN before
   accepting the connection.
