.. _bep-v1:

Block Exchange Protocol v1
==========================

Introduction and Definitions
----------------------------

BEP is used between two or more *devices* thus forming a *cluster*. Each
device has one or more *folders* of files described by the *local
model*, containing metadata and block hashes. The local model is sent to
the other devices in the cluster. The union of all files in the local
models, with files selected for highest change version, forms the
*global model*. Each device strives to get its folders in sync with the
global model by requesting missing or outdated blocks from the other
devices in the cluster.

File data is described and transferred in units of *blocks*, each being
128 KiB (131072 bytes) in size.

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT",
"SHOULD", "SHOULD NOT", "RECOMMENDED", "MAY", and "OPTIONAL" in this
document are to be interpreted as described in RFC 2119.

Transport and Authentication
----------------------------

BEP is deployed as the highest level in a protocol stack, with the lower
level protocols providing encryption and authentication.

::

    +-----------------------------+
    |   Block Exchange Protocol   |
    |-----------------------------|
    | Encryption & Auth (TLS 1.2) |
    |-----------------------------|
    |             TCP             |
    |-----------------------------|
    v             ...             v

The encryption and authentication layer SHALL use TLS 1.2 or a higher
revision. A strong cipher suite SHALL be used, with "strong cipher
suite" being defined as being without known weaknesses and providing
Perfect Forward Secrecy (PFS). Examples of strong cipher suites are
given at the end of this document. This is not to be taken as an
exhaustive list of allowed cipher suites but represents best practices
at the time of writing.

The exact nature of the authentication is up to the application, however
it SHALL be based on the TLS certificate presented at the start of the
connection. Possibilities include certificates signed by a common
trusted CA, preshared certificates, preshared certificate fingerprints
or certificate pinning combined with some out of band first
verification. The reference implementation uses preshared certificate
fingerprints (SHA-256) referred to as "Device IDs".

There is no required order or synchronization among BEP messages except
as noted per message type - any message type may be sent at any time and
the sender need not await a response to one message before sending
another.

The underlying transport protocol MUST guarantee reliable packet delivery.

Pre-authentication messages
---------------------------

AFTER establishing a connection, but BEFORE performing any authentication,
*devices* MUST exchange Hello messages.

Hello messages are used to carry additional information about the peer, which
might be of interest to the user even if the peer is not permitted to
communicate due to failing authentication.

Hello messages MUST be prefixed with a magic number **0x9F79BC40**
represented in network byte order (BE), followed by 4 bytes representing the
size of the message in network byte order (BE), followed by the content of
the Hello message itself. The size of the contents of Hello message MUST be
less or equal to 1024 bytes.

::

    Prefix Structure:

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                             Magic                             |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                            Length                             |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                    Content of HelloMessage                    \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

    HelloMessage Structure:

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \              Device Name (length + padded data)               \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \              Client Name (length + padded data)               \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \             Client Version (length + padded data)             \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


Fields (HelloMessage)
^^^^^^^^^^^^^^^^^^^^^

The **Device Name** is a human readable (configured or auto detected) device
name or host name, for the remote device.

The **Client Name** and **Client Version** identifies the implementation. The
values SHOULD  be simple strings identifying the implementation name, as a
user would expect to see it, and the version string in the same manner. An
example Client Name is "syncthing" and an example Client Version is "v0.7.2".
The Client Version field SHOULD follow the patterns laid out in the `Semantic
Versioning <http://semver.org/>`__ standard.

XDR
^^^

::

    struct HelloMessage {
        string DeviceName<64>;
        string ClientName<64>;
        string ClientVersion<64>;
    };

Immediately after exchanging Hello messages, the connection should be
dropped if device does not pass authentication.

Post-authentication Messages
----------------------------

Every message starts with one 32 bit word indicating the message version, type
and ID, followed by the length of the message. The header is in network byte
order, i.e. big endian. In this document, in diagrams and text, "bit 0" refers
to the *most significant* bit of a word; "bit 31" is thus the least
significant bit of a 32 bit word.

::

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |  Ver  |       Message ID      |      Type     |   Reserved  |C|
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                            Length                             |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

For BEP v1 the **Version** field is set to zero. Future versions with
incompatible message formats will increment the Version field. A message
with an unknown version is a protocol error and MUST result in the
connection being terminated. A client supporting multiple versions MAY
retry with a different protocol version upon disconnection.

The **Message ID** is set to a unique value for each transmitted Request
message. In Response messages it is set to the Message ID of the corresponding
Request message. The uniqueness requirement implies that no more than 4096
request messages may be outstanding at any given moment. For message types
that do not have a corresponding response (Cluster Configuration, Index, etc.)
the Message ID field is irrelevant and SHOULD be set to zero.

The **Type** field indicates the type of data following the message header
and is one of the integers defined below. A message of an unknown type
is a protocol error and MUST result in the connection being terminated.

The **Compression** bit "C" indicates the compression used for the message.

For C=0:

-  The Length field contains the length, in bytes, of the uncompressed
   message data.

-  The message is not compressed.

For C=1:

-  The Length field contains the length, in bytes, of the compressed
   message data plus a four byte uncompressed length field.

-  The compressed message data is preceded by a 32 bit field denoting
   the length of the uncompressed message.

-  The message data is compressed using the LZ4 format and algorithm
   described in http://www.lz4.org/.

All data within the message (post decompression, if compression is in
use) MUST be in XDR (RFC 1014) encoding. All fields shorter than 32 bits
and all variable length data MUST be padded to a multiple of 32 bits.
The actual data types in use by BEP, in XDR naming convention, are the
following:

(unsigned) int:
    (unsigned) 32 bit integer

(unsigned) hyper:
    (unsigned) 64 bit integer

opaque<>
    variable length opaque data

string<>
    variable length string

The transmitted length of string and opaque data is the length of actual
data, excluding any added padding. The encoding of opaque<> and string<>
are identical, the distinction being solely one of interpretation.
Opaque data should not be interpreted but can be compared bytewise to
other opaque data. All strings MUST use the Unicode UTF-8 encoding,
normalization form C.

Cluster Config (Type = 0)
^^^^^^^^^^^^^^^^^^^^^^^^^

.. Documentation note: the structure of a message section is always:
   1. A short description of the message
   2. ASCII art overview of the message formats
   3. Description of the fields in the message.
   4. XDR syntax field descriptions.

This informational message provides information about the cluster
configuration as it pertains to the current connection. A Cluster Config
message MUST be the first message sent on a BEP connection. Additional
Cluster Config messages MUST NOT be sent after the initial exchange.

Graphical Representation
~~~~~~~~~~~~~~~~~~~~~~~~

::

    ClusterConfigMessage Structure:

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                       Number of Folders                       |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                Zero or more Folder Structures                 \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                       Number of Options                       |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                Zero or more Option Structures                 \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

    Folder Structure:

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                         Length of ID                          |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                     ID (variable length)                      \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                 Label (length + padded data)                  \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                       Number of Devices                       |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                Zero or more Device Structures                 \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                             Flags                             |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                       Number of Options                       |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                Zero or more Option Structures                 \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

    Device Structure:

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                         Length of ID                          |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                     ID (variable length)                      \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                        Length of Name                         |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                    Name (variable length)                     \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                      Number of Addresses                      |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                       Length of Address                       |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                   Address (variable length)                   \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                          Compression                          |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                      Length of Cert Name                      |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                  Cert Name (variable length)                  \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                                                               |
    +                  Max Local Version (64 bits)                  +
    |                                                               |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                             Flags                             |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                       Number of Options                       |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                Zero or more Option Structures                 \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

    Option Structure:

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                         Length of Key                         |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                     Key (variable length)                     \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                        Length of Value                        |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                    Value (variable length)                    \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

Fields (ClusterConfigMessage)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

.. Documentation note: the first time a field is mentioned it is put in **bold
   text**. We use the Space Separated names in running text and ASCII art
   diagrams, and CamelCase in the XDR syntax block at the end.

The **Folders** field contains the list of folders that will be synchronized
over the current connection.

The **Options** field is a list of options that apply to the current
connection. The options are used in an implementation specific manner. The
options list is conceptually a map of keys to values, although it is
transmitted in the form of a list of key and value pairs, both of string type.
Key ID:s are implementation specific. An implementation MUST ignore unknown
keys. An implementation MAY impose limits on the length keys and values. The
options list may be used to inform devices of relevant local configuration
options such as rate limiting or make recommendations about request
parallelism, device priorities, etc. An empty options list is valid for
devices not having any such information to share. Devices MAY NOT make any
assumptions about peers acting in a specific manner as a result of sent
options.


Fields (Folder Structure)
~~~~~~~~~~~~~~~~~~~~~~~~~

The **ID** field contains the folder ID, as a human readable string.

The **Label** field contains the folder label, as human readable name for the folder.

The **Devices** field is list of devices participating in sharing this folder.

The **Flags** field contains flags that affect the behavior of the folder. The
folder Flags field contains the following single bit flags:

::

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                           Reserved                    |T|D|P|R|
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

:Bit 31 ("R", Read Only):
    is set for folders that the device will accept no updates from the network
    for.

:Bit 30 ("P", Ignore Permissions):
    is set for folders that the device will not accept or announce file
    permissions for.

:Bit 29 ("D", Ignore Deletes):
    is set for folders that the device will ignore deletes for.

:Bit 28 ("T", Disable Temporary Indexes):
    is set for folders that will not dispatch and do not wish to receive
    progress updates about partially downloaded files via DownloadProgress
	messages.

The **Options** field contains a list of options that apply to the folder.

Fields (Device Structure)
~~~~~~~~~~~~~~~~~~~~~~~~~

The device **ID** field is a 32 byte number that uniquely identifies the
device. For instance, the reference implementation uses the SHA-256 of the
device X.509 certificate.

The **Name** field is a human readable name assigned to the described device
by the sending device. It MAY be empty and it need not be unique.

The list of **Addressess** is that used by the sending device to connect to
the described device.

The **Compression** field indicates the compression mode in use for this
device and folder. The following values are valid:

:0: Compress metadata. This enables compression of metadata messages such as Index.
:1: Compression disabled. No compression is used on any message.
:2: Compress always. Metadata messages as well as Response messages are compressed.

The **Cert Name** field indicates the expected certificate name for this
device. It is commonly blank, indicating to use the implementation default.

The **Max Local Version** field contains the highest local file
version number of the files already known to be in the index sent by
this device. If nothing is known about the index of a given device, this
field MUST be set to zero. When receiving a Cluster Config message with
a non-zero Max Local Version for the local device ID, a device MAY elect
to send an Index Update message containing only files with higher local
version numbers in place of the initial Index message.

The **Flags** field indicates the sharing mode of the folder and other device
& folder specific settings. See the discussion on Sharing Modes. The Device
Flags field contains the following single bit flags:

::

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |          Reserved         |Pri|          Reserved       |I|R|T|
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

:Bit 31 ("T", Trusted): is set for devices that participate in trusted
   mode.

:Bit 30 ("R", Read Only): is set for devices that participate in read
   only mode.

:Bit 29 ("I", Introducer): is set for devices that are trusted as
   cluster introducers.

:Bits 16 through 28: are reserved and MUST be set to zero.

:Bits 14-15 ("Pri", Priority): indicate the device's upload priority for this
   folder. Possible values are:

   :00: The default. Normal priority.

   :01: High priority. Other devices SHOULD favour requesting files
      from this device over devices with normal or low priority.

   :10: Low priority. Other devices SHOULD avoid requesting files from
      this device when they are available from other devices.

   :11: Sharing disabled. Other devices SHOULD NOT request files from
      this device.

:Bits 0 through 14: are reserved and MUST be set to zero.

Exactly one of the T and R bits MUST be set.

The **Options** field contains a list of options that apply to the device.

XDR
~~~

::

    struct ClusterConfigMessage {
        Folder Folders<1000000>;
        Option Options<64>;
    };

    struct Folder {
        string ID<256>;
        string Label<256>;
        Device Devices<1000000>;
        unsigned int Flags;
        Option Options<64>;
    };

    struct Device {
        opaque ID<32>;
        string Name<64>;
        string Addresses<64>;
        unsigned int Compression;
        string CertName<64>;
        hyper MaxLocalVersion;
        unsigned int Flags;
        Option Options<64>;
    };

    struct Option {
        string Key<64>;
        string Value<1024>;
    };

Index (Type = 1) and Index Update (Type = 6)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

The Index and Index Update messages define the contents of the senders
folder. An Index message represents the full contents of the folder and
thus supersedes any previous index. An Index Update amends an existing
index with new information, not affecting any entries not included in
the message. An Index Update MAY NOT be sent unless preceded by an
Index, unless a non-zero Max Local Version has been announced for the
given folder by the peer device.

Graphical Representation
~~~~~~~~~~~~~~~~~~~~~~~~

::

    IndexMessage Structure:

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                       Length of Folder                        |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                   Folder (variable length)                    \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                        Number of Files                        |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \               Zero or more FileInfo Structures                \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                             Flags                             |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                       Number of Options                       |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                Zero or more Option Structures                 \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

    FileInfo Structure:

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                        Length of Name                         |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                    Name (variable length)                     \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                             Flags                             |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                                                               |
    +                      Modified (64 bits)                       +
    |                                                               |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                   Version (variable length)                   \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                                                               |
    +                    Local Version (64 bits)                    +
    |                                                               |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                       Number of Blocks                        |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \               Zero or more BlockInfo Structures               \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

    Vector Structure:

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                      Number of Counters                       |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                Zero or more Counter Structures                \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

    Counter Structure:

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                                                               |
    +                          ID (64 bits)                         +
    |                                                               |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                                                               |
    +                        Value (64 bits)                        +
    |                                                               |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


    BlockInfo Structure:

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                             Size                              |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                        Length of Hash                         |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                    Hash (variable length)                     \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

Fields (Index Message)
~~~~~~~~~~~~~~~~~~~~~~

The **Folder** field identifies the folder that the index message pertains to.

**Files**

The **Flags** field is reserved for future use and MUST currently be set to
zero.

The **Options** list is implementation defined and as described in the
ClusterConfig message section.

Fields (FileInfo Structure)
~~~~~~~~~~~~~~~~~~~~~~~~~~~

The **Name** is the file name path relative to the folder root. Like all
strings in BEP, the Name is always in UTF-8 NFC regardless of operating
system or file system specific conventions. The Name field uses the
slash character ("/") as path separator, regardless of the
implementation's operating system conventions. The combination of Folder
and Name uniquely identifies each file in a cluster.

The **Flags** field is made up of the following single bit flags:

::

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |            Reserved       |U|S|P|D|I|R|   Unix Perm. & Mode   |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

:The lower 12 bits: hold the common Unix permission and mode bits. An
   implementation MAY ignore or interpret these as is suitable on the
   host operating system.

:Bit 19 ("R"): is set when the file has been deleted. The block list
   SHALL be of length zero and the modification time indicates the time
   of deletion or, if the time of deletion is not reliably determinable,
   the last known modification time.

:Bit 18 ("I"): is set when the file is invalid and unavailable for
   synchronization. A peer MAY set this bit to indicate that it can
   temporarily not serve data for the file.

:Bit 17 ("D"): is set when the item represents a directory. The block
   list SHALL be of length zero.

:Bit 16 ("P"): is set when there is no permission information for the
   file. This is the case when it originates on a file system which
   does not support permissions. Changes to only permission bits SHOULD
   be disregarded on files with this bit set. The permissions bits MUST
   be set to the octal value 0666.

:Bit 15 ("S"): is set when the file is a symbolic link. The block list
   SHALL be of one or more blocks since the target of the symlink is
   stored within the blocks of the file.

:Bit 14 ("U"): is set when the symbolic links target does not exist. On
   systems where symbolic links have types, this bit being means that
   the default file symlink SHALL be used. If this bit is unset bit 19
   will decide the type of symlink to be created.

:Bit 0 through 13: are reserved for future use and SHALL be set to
   zero.

The **Modified** time is expressed as the number of seconds since the Unix
Epoch (1970-01-01 00:00:00 UTC).

The **Version** field is a version vector describing the updates performed
to a file by all members in the cluster. Each counter in the version
vector is an ID-Value tuple. The ID is used the first 64 bits of the
device ID. The Value is a simple incrementing counter, starting at zero.
The combination of Folder, Name and Version uniquely identifies the
contents of a file at a given point in time.

The **Local Version** field is the value of a device local monotonic clock
at the time of last local database update to a file. The clock ticks on
every local database update.

The **Blocks** list contains the size and hash for each block in the file.
Each block represents a 128 KiB slice of the file, except for the last
block which may represent a smaller amount of data.

The hash algorithm is implied by the **Hash** length. Currently, the hash
MUST be 32 bytes long and computed by SHA256.

XDR
~~~

::

    struct IndexMessage {
        string Folder<256>;
        FileInfo Files<1000000>;
        unsigned int Flags;
        Option Options<64>;
    };

    struct FileInfo {
        string Name<8192>;
        unsigned int Flags;
        hyper Modified;
        Vector Version;
        hyper LocalVersion;
        BlockInfo Blocks<10000000>;
    };

    struct Vector {
        Counter Counters<>;
    };

    struct Counter {
        unsigned hyper ID;
        unsigned hyper Value;
    };

    struct BlockInfo {
        unsigned int Size;
        opaque Hash<64>;
    };

Request (Type = 2)
^^^^^^^^^^^^^^^^^^

The Request message expresses the desire to receive a data block
corresponding to a part of a certain file in the peer's folder.

Graphical Representation
~~~~~~~~~~~~~~~~~~~~~~~~

::

    RequestMessage Structure:

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                       Length of Folder                        |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                   Folder (variable length)                    \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                        Length of Name                         |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                    Name (variable length)                     \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                                                               |
    +                       Offset (64 bits)                        +
    |                                                               |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                             Size                              |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                        Length of Hash                         |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                    Hash (variable length)                     \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                             Flags                             |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                       Number of Options                       |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                Zero or more Option Structures                 \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

Fields
~~~~~~

The Folder and Name fields are as documented for the Index message. The
Offset and Size fields specify the region of the file to be transferred.
This SHOULD equate to exactly one block as seen in an Index message.

The Hash field MAY be set to the expected hash value of the block, or
may be left empty (zero length). If set, the other device SHOULD ensure
that the transmitted block matches the requested hash. The other device
MAY reuse a block from a different file and offset having the same size
and hash, if one exists.

The **Flags** field is made up of the following single bit flags:
::

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                           Reserved                          |T|
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

:Bit 31 ("T", Temporary): is set to indicate that the read should be performed
    from the temporary file (converting Name to it's temporary form) and falling
    back to the non temporary file if any error occurs. Knowledge of content
	inside temporary files comes from DownloadProgress messages.

The Options list is implementation defined and as described in the
ClusterConfig message section.

XDR
~~~

::

    struct RequestMessage {
        string Folder<64>;
        string Name<8192>;
        hyper Offset;
        int Size;
        opaque Hash<64>;
        unsigned int Flags;
        Option Options<64>;
    };

Response (Type = 3)
^^^^^^^^^^^^^^^^^^^

The Response message is sent in response to a Request message.

Graphical Representation
~~~~~~~~~~~~~~~~~~~~~~~~

ResponseMessage Structure:

::

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                        Length of Data                         |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                    Data (variable length)                     \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                             Code                              |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

Fields
~~~~~~

The **Data** field contains either a full 128 KiB block, a shorter block in
the case of the last block in a file, or is empty (zero length) if the
requested block is not available.

The **Code** field contains an error code describing the reason a Request
could not be fulfilled, in the case where a zero length Data was
returned. The following values are defined:

:0: No Error (Data should be present)

:1: Generic Error

:2: No Such File (the requested file does not exist, or the offset is
   outside the acceptable range for the file)

:3: Invalid (file exists but has invalid bit set or is otherwise
   unavailable)

XDR
~~~

::

    struct ResponseMessage {
        opaque Data<>;
        int Code;
    }

DownloadProgress (Type = 8)
^^^^^^^^^^^^^^^^^^^^^^^^^^^

The DownloadProgress message is used to notify remote devices about partial
availability of files. By default, these messages are sent every 5 seconds,
and only in the cases where progress or state changes have been detected.
Each DownloadProgress message is addressed to a specific folder and MUST
contain zero or more FileDownloadProgressUpdate structures.

Graphical Representation
~~~~~~~~~~~~~~~~~~~~~~~~

::

    DownloadProgressMessage Structure:

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                 Folder (length + padded data)                 \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                       Number of Updates                       |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \      Zero or more FileDownloadProgressUpdate Structures       \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                             Flags                             |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                       Number of Options                       |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                Zero or more Option Structures                 \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

    FileDownloadProgressUpdate Structure:

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                          Update Type                          |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                  Name (length + padded data)                  \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                      Version Structure                        \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                    Number of Block Indexes                    |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    |                    Block Indexes (n items)                    |
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


Each

Fields (DownloadProgress Message)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
**Folder** represents the ID of the folder for which the update is being
provided.

The **Flags** field is reserved for future use and MUST currently be set to
zero. The **Options** field contains a list of options that apply to the update.

Fields (FileDownloadProgressUpdate Structure)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

The **Update Type** field is made up of the following single bit flags:
::

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                           Reserved                          |F|
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

:Bit 31 ("F", Forget): is set to notify that the file that was previously
    advertised is no longer available (at least as a temporary file).

The **Name** field defines the file name from the global index for which this
update is being sent.

The **Version** structure defines the version of the file for which this update
is being sent.

**Block Indexes** is a list of positive integers, where each integer represents
the index of the block in the FileInfo structure Blocks array that has become
available for download.
For example an integer with with value 3 represents that the data defined in the
fourth BlockInfo structure of the FileInfo structure of that file is now available.
Please note that matching should be done on **Name** AND **Version**.
Furthermore, each update received is incremental, for example the initial update
structure might contain indexes 0, 1, 2, an update 5 seconds later might contain
indexes 3, 4, 5 which should be appended to the original list, which implies
that blocks 0-5 are currently available.

Block indexes MAY be added in any order.
An implementation MUST NOT assume that block indexes are added in any specific
order.

**Forget** bit being set implies that the file that was previously advertised
is no longer available, therefore the list of block indexes should be truncated.

Messages with **Forget** bit set MUST NOT have any block indexes.

Any update message which is being sent for a different **Version** of the same
file name must be preceded with an update message for the old version of that
file with the **Forget** bit set.

As a safeguard on the receiving side, value of **Version** changing between
update messages implies that the file has changed, and that any indexes
previously advertised are no longer available. The list of available block
indexes MUST be replaced (rather than appended) with the indexes specified in
this message.

XDR
~~~

::

    struct DownloadProgressMessage {
        string Folder<64>;
        FileDownloadProgressUpdate Updates<1000000>;
        unsigned int Flags;
        Option Options<64>;
    }

    struct FileDownloadProgressUpdate {
        unsigned int UpdateType;
        string Name<8192>;
        Vector Version;
        int BlockIndexes<1000000>;
    }


Ping (Type = 4)
^^^^^^^^^^^^^^^

The Ping message is used to determine that a connection is alive, and to keep
connections alive through state tracking network elements such as firewalls
and NAT gateways. The Ping message has no contents. A Ping message is sent
every 90 seconds, if no other message has been sent in the preceding 90
seconds.

Close (Type = 7)
^^^^^^^^^^^^^^^^

The Close message MAY be sent to indicate that the connection will be
torn down due to an error condition. A Close message MUST NOT be
followed by further messages.

Graphical Representation
~~~~~~~~~~~~~~~~~~~~~~~~

::

    CloseMessage Structure:

     0                   1                   2                   3
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                       Length of Reason                        |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                                                               /
    \                   Reason (variable length)                    \
    /                                                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |                             Code                              |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

Fields
~~~~~~

The **Reason** field contains a human description of the error condition,
suitable for consumption by a human. The **Code** field is for a machine
readable error code. Codes are reserved for future use and MUST
currently be set to zero.

::

    struct CloseMessage {
        string Reason<1024>;
        int Code;
    }

Sharing Modes
-------------

Trusted
^^^^^^^

Trusted mode is the default sharing mode. Updates are exchanged in both
directions.

::

    +------------+     Updates      /---------\
    |            |  ----------->   /           \
    |   Device   |                 |  Cluster  |
    |            |  <-----------   \           /
    +------------+     Updates      \---------/

Read Only
^^^^^^^^^

In read only mode, a device does not apply any updates from the cluster,
but publishes changes of its local folder to the cluster as usual.
The local folder can be seen as a "master copy" that is never affected
by the actions of other cluster devices.

::

    +------------+     Updates      /---------\
    |            |  ----------->   /           \
    |   Device   |                 |  Cluster  |
    |            |                 \           /
    +------------+                  \---------/

Message Limits
--------------

An implementation MAY impose reasonable limits on the length of messages
and message fields to aid robustness in the face of corruption or broken
implementations. These limits, if imposed, SHOULD NOT be more
restrictive than the following. An implementation should strive to keep
messages short and to the point, favouring more and smaller messages
over fewer and larger. For example, favour a smaller Index message
followed by one or more Index Update messages rather than sending a very
large Index message.

=================== =================== =============
Message Type        Field               Limit
=================== =================== =============
**All Messages**
-----------------------------------------------------
|                   Total length        512 MiB

**Index and Index Update Messages**
-----------------------------------------------------
|                   Folder              64 bytes
|                   Number of Files     1.000.000
|                   Name                8192 bytes
|                   Number of Blocks    10.000.000
|                   Hash                64 bytes
|                   Number of Counters  1.000.000

**Request Messages**
-----------------------------------------------------
|                   Folder              64 bytes
|                   Name                8192 bytes

**Response Messages**
-----------------------------------------------------
|                   Data                256 KiB

**Cluster Config Message**
-----------------------------------------------------
|                   Number of Folders   1.000.000
|                   Number of Devices   1.000.000
|                   Number of Options   64
|                   Key                 64 bytes
|                   Value               1024 bytes

**Download Progress Messages**
-----------------------------------------------------
|                   Folder              64 bytes
|                   Number of Updates   1.000.000
|                   Name                8192 bytes
|                   Number of Indexes   1.000.000
=================== =================== =============

The currently defined values allow maximum file size of 1220 GiB
(10.000.000 x 128 KiB). The maximum message size covers an Index message
for the maximum file.

Example Exchange
----------------

===  =======================  ======================
 #             A                        B
===  =======================  ======================
 1   ClusterConfiguration->   <-ClusterConfiguration
 2   Index->                  <-Index
 3   IndexUpdate->            <-IndexUpdate
 4   IndexUpdate->
 5   Request->
 6   Request->
 7   Request->
 8   Request->
 9                            <-Response
10                            <-Response
11                            <-Response
12                            <-Response
13   Index Update->
...
14                            <-Ping
15   Ping->
===  =======================  ======================

The connection is established and at 1. both peers send ClusterConfiguration
messages and then Index records. The Index records are received and both peers
recompute their knowledge of the data in the cluster. In this example, peer A
has four missing or outdated blocks. At 5 through 8 peer A sends requests for
these blocks. The requests are received by peer B, who retrieves the data from
the folder and transmits Response records (9 through 12). Device A updates
their folder contents and transmits an Index Update message (13). Both peers
enter idle state after 13. At some later time 14, the ping timer on device B
expires and a Ping message is sent. The same process occurs for device A at
15.

Examples of Strong Cipher Suites
--------------------------------

======  ===========================  ==================================
ID      Name                         Description
======  ===========================  ==================================
0x009F  DHE-RSA-AES256-GCM-SHA384    TLSv1.2 DH RSA AESGCM(256) AEAD
0x006B  DHE-RSA-AES256-SHA256        TLSv1.2 DH RSA AES(256) SHA256
0xC030  ECDHE-RSA-AES256-GCM-SHA384  TLSv1.2 ECDH RSA AESGCM(256) AEAD
0xC028  ECDHE-RSA-AES256-SHA384      TLSv1.2 ECDH RSA AES(256) SHA384
0x009E  DHE-RSA-AES128-GCM-SHA256    TLSv1.2 DH RSA AESGCM(128) AEAD
0x0067  DHE-RSA-AES128-SHA256        TLSv1.2 DH RSA AES(128) SHA256
0xC02F  ECDHE-RSA-AES128-GCM-SHA256  TLSv1.2 ECDH RSA AESGCM(128) AEAD
0xC027  ECDHE-RSA-AES128-SHA256      TLSv1.2 ECDH RSA AES(128) SHA256
======  ===========================  ==================================

