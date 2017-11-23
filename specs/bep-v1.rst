.. _bep-v1:

Block Exchange Protocol v1
==========================

Introduction and Definitions
----------------------------

The Block Exchange Protocol (BEP) is used between two or more *devices* thus
forming a *cluster*. Each device has one or more *folders* of files
described by the *local model*, containing metadata and block hashes. The
local model is sent to the other devices in the cluster. The union of all
files in the local models, with files selected for highest change version,
forms the *global model*. Each device strives to get its folders in sync
with the global model by requesting missing or outdated blocks from the
other devices in the cluster.

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
    |      Reliable Transport     |
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

In this document, in diagrams and text, "bit 0" refers to the *most
significant* bit of a word; "bit 15" is thus the least significant bit of a
16 bit word (int16) and "bit 31" is the least significant bit of a 32 bit
word (int32). Non protocol buffer integers are always represented in network
byte order (i.e., big endian) and are signed unless stated otherwise, but
when describing message lengths negative values do not make sense and the
most significant bit MUST be zero.

The protocol buffer schemas in this document are in ``proto3`` syntax. This
means, among other things, that all fields are optional and will assume
their default value when missing. This does not nececessarily mean that a
message is *valid* with all fields empty - for example, an index entry for a
file that does not have a name is not useful and MAY be rejected by the
implementation. However the folder label is for human consumption only so an
empty label should be accepted - the implementation will have to choose some
way to represent the folder, perhaps by using the ID in it's place or
automatically generating a label.

Pre-authentication messages
---------------------------

AFTER establishing a connection, but BEFORE performing any authentication,
devices MUST exchange Hello messages.

Hello messages are used to carry additional information about the peer,
which might be of interest to the user even if the peer is not permitted to
communicate due to failing authentication. Note that the certificate based
authentication may be considered part of the TLS handshake that precedes the
Hello message exchange, but even in the case that a connection is rejected a
Hello message must be sent before the connection is terminated.

Hello messages MUST be prefixed with an int32 containing the magic number
**0x2EA7D90B**, followed by an int16 representing the size of the message,
followed by the contents of the Hello message itself.

.. code-block:: none

     0                   1
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |             Magic             |
    |           (32 bits)           |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |             Length            |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                               /
    \             Hello             \
    /                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

The Hello message itself is in protocol buffer format with the following schema:

.. code-block:: proto

    message Hello {
        string device_name    = 1;
        string client_name    = 2;
        string client_version = 3;
    }

Fields (Hello message)
^^^^^^^^^^^^^^^^^^^^^^

The **device_name** is a human readable (configured or auto detected) device
name or host name, for the remote device.

The **client_name** and **client_version** identifies the implementation. The
values SHOULD  be simple strings identifying the implementation name, as a
user would expect to see it, and the version string in the same manner. An
example client name is "syncthing" and an example client version is "v0.7.2".
The client version field SHOULD follow the patterns laid out in the `Semantic
Versioning <http://semver.org/>`__ standard.

Immediately after exchanging Hello messages, the connection MUST be dropped
if the remote device does not pass authentication.

Post-authentication Messages
----------------------------

Every message post authentication is made up of several parts:

- A header length word
- A **Header**
- A message length word
- A **Message**

.. code-block:: none

     0                   1
     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |         Header Length         |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                               /
    \            Header             \
    /                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    |         Message Length        |
    |           (32 bits)           |
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
    /                               /
    \            Message            \
    /                               /
    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

The header length word is 16 bits. It indicates the length of the following
**Header** message. The Header is in protocol buffer format. The Header
describes the type and compression status of the following message.

The message is preceded by the 32 bit message length word and is one of the
concrete BEP messages described below, identified by the **type** field of
the Header.

As always, the length words are in network byte order (big endian).

.. code-block:: proto

    message Header {
        MessageType        type        = 1;
        MessageCompression compression = 2;
    }

    enum MessageType {
        CLUSTER_CONFIG    = 0;
        INDEX             = 1;
        INDEX_UPDATE      = 2;
        REQUEST           = 3;
        RESPONSE          = 4;
        DOWNLOAD_PROGRESS = 5;
        PING              = 6;
        CLOSE             = 7;
    }

    enum MessageCompression {
        NONE = 0;
        LZ4  = 1;
    }

When the **compression** field is **NONE**, the message is directly in
protocol buffer format.

When the compression field is **LZ4**, the message consists of a 32 bit
integer describing the uncompressed message length followed by a single LZ4
block. After decompressing the LZ4 block it should be interpreted as a
protocol buffer message just as in the uncompressed case.

Message Subtypes
----------------

Cluster Config
^^^^^^^^^^^^^^

.. Documentation note: the structure of a message section is always:
   1. A short description of the message
   2. Protocol buffer schema of the message
   3. Description of the fields in the message.

This informational message provides information about the cluster
configuration as it pertains to the current connection. A Cluster Config
message MUST be the first post authentication message sent on a BEP
connection. Additional Cluster Config messages MUST NOT be sent after the
initial exchange.

Protocol Buffer Schema
~~~~~~~~~~~~~~~~~~~~~~

.. code-block:: proto

    message ClusterConfig {
        repeated Folder folders = 1;
    }

    message Folder {
        string id                   = 1;
        string label                = 2;
        bool   read_only            = 3;
        bool   ignore_permissions   = 4;
        bool   ignore_delete        = 5;
        bool   disable_temp_indexes = 6;

        repeated Device devices = 16;
    }

    message Device {
        bytes           id                         = 1;
        string          name                       = 2;
        repeated string addresses                  = 3;
        Compression     compression                = 4;
        string          cert_name                  = 5;
        int64           max_sequence               = 6;
        bool            introducer                 = 7;
        uint64          index_id                   = 8;
        bool            skip_introduction_removals = 9;
    }

    enum Compression {
        METADATA = 0;
        NEVER    = 1;
        ALWAYS   = 2;
    }

Fields (Cluster Config Message)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

.. Documentation note: the first time a field is mentioned it is put in
   **bold text**. We use the space separated names in running text and
   snake_case in the protocol buffer schema.

The **folders** field contains the list of folders that will be synchronized
over the current connection.

Fields (Folder Message)
~~~~~~~~~~~~~~~~~~~~~~~

The **id** field contains the folder ID, which is the unique identifier of
the folder.

The **label** field contains the folder label, the human readable name of
the folder.

The **read only** field is set for folders that the device will accept no
updates from the network for.

The **ignore permissions** field is set for folders that the device will not
accept or announce file permissions for.

The **ignore delete** field is set for folders that the device will ignore
deletes for.

The **disable temp indexes** field is set for folders that will not dispatch
and do not wish to receive progress updates about partially downloaded files
via Download Progress messages.

The **devices** field is a list of devices participating in sharing this
folder.

Fields (Device Message)
~~~~~~~~~~~~~~~~~~~~~~~

The device **id** field is a 32 byte number that uniquely identifies the
device. For instance, the reference implementation uses the SHA-256 of the
device X.509 certificate.

The **name** field is a human readable name assigned to the described device
by the sending device. It MAY be empty and it need not be unique.

The list of **addresses** is that used by the sending device to connect to
the described device.

The **compression** field indicates the compression mode in use for this
device and folder. The following values are valid:

:0: Compress metadata. This enables compression of metadata messages such as Index.
:1: Compression disabled. No compression is used on any message.
:2: Compress always. Metadata messages as well as Response messages are compressed.

The **cert name** field indicates the expected certificate name for this
device. It is commonly blank, indicating to use the implementation default.

The **max sequence** field contains the highest sequence number of the files
in the index. See :ref:`deltaidx` for the usage of this field.

The **introducer** field is set for devices that are trusted as cluster
introducers.

The **index id** field contains the unique identifier for the current set of
index data. See :ref:`deltaidx` for the usage of this field.

The **skip introduction removals** field signifies if the remote device has
opted to ignore introduction removals for the given device. This setting is
copied across as we are being introduced to a new device.

Index and Index Update
^^^^^^^^^^^^^^^^^^^^^^

The Index and Index Update messages define the contents of the senders
folder. An Index message represents the full contents of the folder and
thus supersedes any previous index. An Index Update amends an existing
index with new information, not affecting any entries not included in
the message. An Index Update MAY NOT be sent unless preceded by an
Index, unless a non-zero Max Sequence has been announced for the
given folder by the peer device.

The Index and Index Update messages are currently identical in format,
although this is not guaranteed to be the case in the future.

Protocol Buffer Schema
~~~~~~~~~~~~~~~~~~~~~~

.. code-block:: proto

    message Index {
        string            folder = 1;
        repeated FileInfo files  = 2;
    }

    message IndexUpdate {
        string            folder = 1;
        repeated FileInfo files  = 2;
    }

    message FileInfo {
        string       name           = 1;
        FileInfoType type           = 2;
        int64        size           = 3;
        uint32       permissions    = 4;
        int64        modified_s     = 5;
        int32        modified_ns    = 11;
        uint64       modified_by    = 12;
        bool         deleted        = 6;
        bool         invalid        = 7;
        bool         no_permissions = 8;
        Vector       version        = 9;
        int64        sequence      = 10;

        repeated BlockInfo Blocks         = 16;
        string             symlink_target = 17;
    }

    enum FileInfoType {
        FILE              = 0;
        DIRECTORY         = 1;
        SYMLINK_FILE      = 2 [deprecated = true];
        SYMLINK_DIRECTORY = 3 [deprecated = true];
        SYMLINK           = 4;
    }

    message BlockInfo {
        int64 offset = 1;
        int32 size   = 2;
        bytes hash   = 3;
    }

    message Vector {
        repeated Counter counters = 1;
    }

    message Counter {
        uint64 id    = 1;
        uint64 value = 2;
    }

Fields (Index Message)
~~~~~~~~~~~~~~~~~~~~~~

The **folder** field identifies the folder that the index message pertains to.

The **files** field is a list of files making up the index information.

Fields (FileInfo Message)
~~~~~~~~~~~~~~~~~~~~~~~~~

The **name** is the file name path relative to the folder root. Like all
strings in BEP, the Name is always in UTF-8 NFC regardless of operating
system or file system specific conventions. The name field uses the slash
character ("/") as path separator, regardless of the implementation's
operating system conventions. The combination of folder and name uniquely
identifies each file in a cluster.

The **type** field contains the type of the described item. The type is one
of **file (0)**, **directory (1)**, or **symlink (4)**.

The **size** field contains the size of the file, in bytes. For directories
and symlinks the size is zero.

The **permissions** field holds the common Unix permission bits. An
implementation MAY ignore or interpret these as is suitable on the host
operating system.

The **modified_s** time is expressed as the number of seconds since the Unix
Epoch (1970-01-01 00:00:00 UTC). The **modified_ns** field holds the
nanosecond part of the modification time.

The **modified_by** field holds the short id of the client that last made
any modification to the file whether add, change or delete.  This will be
overwritten every time a change is made to the file by the last client to do
so and so does not hold history.

The **deleted** field is set when the file has been deleted. The block list
SHALL be of length zero and the modification time indicates the time of
deletion or, if the time of deletion is not reliably determinable, the last
known modification time.

The **invalid** field is set when the file is invalid and unavailable for
synchronization. A peer MAY set this bit to indicate that it can temporarily
not serve data for the file.

The **no permissions** field is set when there is no permission information
for the file. This is the case when it originates on a file system which
does not support permissions. Changes to only permission bits SHOULD be
disregarded on files with this bit set. The permissions bits MUST be set to
the octal value 0666.

The **version** field is a version vector describing the updates performed
to a file by all members in the cluster. Each counter in the version vector
is an ID-Value tuple. The ID is the first 64 bits of the device ID. The
Value is a simple incrementing counter, starting at zero. The combination of
Folder, Name and Version uniquely identifies the contents of a file at a
given point in time.

The **sequence** field is the value of a device local monotonic clock at the
time of last local database update to a file. The clock ticks on every local
database update, thus forming a sequence number over database updates.

The **blocks** list contains the size and hash for each block in the file.
Each block represents a 128 KiB slice of the file, except for the last block
which may represent a smaller amount of data. The block list is empty for
directories and symlinks.

The **symlink_target** field contains the symlink target, for entries of
symlink type. It is empty for all other entry types.

Request
^^^^^^^

The Request message expresses the desire to receive a data block
corresponding to a part of a certain file in the peer's folder.

Protocol Buffer Schema
~~~~~~~~~~~~~~~~~~~~~~

.. code-block:: proto

    message Request {
        int32  id             = 1;
        string folder         = 2;
        string name           = 3;
        int64  offset         = 4;
        int32  size           = 5;
        bytes  hash           = 6;
        bool   from_temporary = 7;
    }

Fields
~~~~~~

The **id** is the request identifier. It will be matched in the
corresponding **Response** message. Each outstanding request must have a
unique ID.

The **folder** and **name** fields are as documented for the Index message.
The **offset** and **size** fields specify the region of the file to be
transferred. This SHOULD equate to exactly one block as seen in an Index
message.

The *hash* field MAY be set to the expected hash value of the block. If set,
the other device SHOULD ensure that the transmitted block matches the
requested hash. The other device MAY reuse a block from a different file and
offset having the same size and hash, if one exists.

The **from temporary** field is set to indicate that the read should be
performed from the temporary file (converting name to it's temporary form)
and falling back to the non temporary file if any error occurs. Knowledge of
contents of temporary files comes from DownloadProgress messages.

Response
^^^^^^^^

The Response message is sent in response to a Request message.

Protocol Buffer Schema
~~~~~~~~~~~~~~~~~~~~~~

.. code-block:: proto

    message Response {
        int32     id   = 1;
        bytes     data = 2;
        ErrorCode code = 3;
    }

    enum ErrorCode {
        NO_ERROR     = 0;
        GENERIC      = 1;
        NO_SUCH_FILE = 2;
        INVALID_FILE = 3;
    }

Fields
~~~~~~

The **id** field is the request identifier. It must match the ID of the
**Request** that is being responded to.

The **data** field contains either the requested data block or is empty if
the requested block is not available.

The **code** field contains an error code describing the reason a Request
could not be fulfilled, in the case where zero length data was returned. The
following values are defined:

:0: No Error (data should be present)

:1: Generic Error

:2: No Such File (the requested file does not exist, or the offset is
   outside the acceptable range for the file)

:3: Invalid (file exists but has invalid bit set or is otherwise
   unavailable)

DownloadProgress
^^^^^^^^^^^^^^^^

The DownloadProgress message is used to notify remote devices about partial
availability of files. By default, these messages are sent every 5 seconds,
and only in the cases where progress or state changes have been detected.
Each DownloadProgress message is addressed to a specific folder and MUST
contain zero or more FileDownloadProgressUpdate messages.

Protocol Buffer Schema
~~~~~~~~~~~~~~~~~~~~~~

.. code-block:: proto

    message DownloadProgress {
        string                              folder  = 1;
        repeated FileDownloadProgressUpdate updates = 2;
    }

    message FileDownloadProgressUpdate {
        FileDownloadProgressUpdateType update_type   = 1;
        string                         name          = 2;
        Vector                         version       = 3;
        repeated int32                 block_indexes = 4;
    }

    enum FileDownloadProgressUpdateType {
        APPEND = 0;
        FORGET = 1;
    }

Fields (DownloadProgress Message)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

The **folder** field represents the ID of the folder for which the update is
being provided.

The **updates** field is a list of progress update messages.

Fields (FileDownloadProgressUpdate Message)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

The **update type** indicates whether the update is of type **append (0)**
(new blocks are available) or **forget (1)** (the file transfer has
completed or failed).

The **name** field defines the file name from the global index for which
this update is being sent.

The **version** message defines the version of the file for which this
update is being sent.

The **block indexes** field is a list of positive integers, where each
integer represents the index of the block in the FileInfo message Blocks
array that has become available for download.

For example an integer with value 3 represents that the data defined in the
fourth BlockInfo message of the FileInfo message of that file is now
available. Please note that matching should be done on **name** AND
**version**. Furthermore, each update received is incremental, for example
the initial update message might contain indexes 0, 1, 2, an update 5
seconds later might contain indexes 3, 4, 5 which should be appended to the
original list, which implies that blocks 0-5 are currently available.

Block indexes MAY be added in any order. An implementation MUST NOT assume
that block indexes are added in any specific order.

The **forget** field being set implies that previously advertised file is no
longer available, therefore the list of block indexes should be truncated.

Messages with the **forget** field set MUST NOT have any block indexes.

Any update message which is being sent for a different **version** of the
same file name must be preceded with an update message for the old version
of that file with the **forget** field set.

As a safeguard on the receiving side, the value of **version** changing
between update messages implies that the file has changed and that any
indexes previously advertised are no longer available. The list of available
block indexes MUST be replaced (rather than appended) with the indexes
specified in this message.

Ping
^^^^

The Ping message is used to determine that a connection is alive, and to
keep connections alive through state tracking network elements such as
firewalls and NAT gateways. A Ping message is sent every 90 seconds, if no
other message has been sent in the preceding 90 seconds.

Protocol Buffer Schema
~~~~~~~~~~~~~~~~~~~~~~

.. code-block:: proto

    message Ping {
    }


Close
^^^^^

The Close message MAY be sent to indicate that the connection will be torn
down due to an error condition. A Close message MUST NOT be followed by
further messages.

Protocol Buffer Schema
~~~~~~~~~~~~~~~~~~~~~~

.. code-block:: proto

    message Close {
        string reason = 1;
    }

Fields
~~~~~~

The **reason** field contains a human readable description of the error
condition.

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

Send Only
^^^^^^^^^

In send-only mode, a device does not apply any updates from the cluster, but
publishes changes of its local folder to the cluster as usual. The local
folder can be seen as a "master copy" that is never affected by the actions
of other cluster devices.

::

    +------------+     Updates      /---------\
    |            |  ----------->   /           \
    |   Device   |                 |  Cluster  |
    |            |                 \           /
    +------------+                  \---------/

.. _deltaidx:

Delta Index Exchange
--------------------

Index data must be exchanged whenever two devices connect so that one knows
the files available on the other. In the most basic case this happens by way
of sending an ``Index`` message followed by one or more ``Index Update``
messages. Any previous index data known for a remote device is removed and
replaced with the new index data received in an ``Index`` message, while the
contents of an ``Index Update`` message is simply added to the existing
index data.

For situations with large indexes or frequent reconnects this can be quite
inefficient. A mechanism can then be used to retain index data between
connections and only transmit any changes since that data on connection
start. This is called "delta indexes". To enable this mechanism the
**sequence** and **index ID** fields are used.

Sequence:
    Each index item (i.e., file, directory or symlink) has a sequence number
    field. It contains the value of a counter at the time the index item was
    updated. The counter increments by one for each change. That is, as files
    are scanned and added to the index they get assigned sequence numbers
    1, 2, 3 and so on. The next file to be changed or detected gets sequence
    number 4, and future updates continue in the same fashion.

Index ID:
    Each folder has an Index ID. This is a 64 bit random identifier set at
    index creation time.

Given the above, we know that the tuple {index ID, maximum sequence number}
uniquely identifies a point in time of a given index. Any further changes
will increase the sequence number of some item, and thus the maximum
sequence number for the index itself. Should the index be reset or removed
(i.e., the sequence number reset to zero), a new index ID must be generated.

By letting a device know the {index ID, maximum sequence number} we have for
their index data, that device can arrange to only transmit ``Index Update``
messages for items with a higher sequence number. This is the delta index
mechanism.

The index ID and maximum sequence number known for each device is
transmitted in the ``Cluster Config`` message at connection start.

For this mechanism to be reliable it is essential that outgoing index
information is ordered by increasing sequence number. Devices announcing a
non-zero index ID in the ``Cluster Config`` message MUST send all index data
ordered by increasing sequence number. Devices not intending to participate
in delta index exchange MUST send a zero index ID or, equivalently, not send
the ``index_id`` attribute at all.

Message Limits
--------------

An implementation MAY impose reasonable limits on the length of messages and
message fields to aid robustness in the face of corruption or broken
implementations. An implementation should strive to keep messages short
and to the point, favouring more and smaller messages over fewer and larger.
For example, favour a smaller Index message followed by one or more Index
Update messages rather than sending a very large Index message.

The Syncthing implementation imposes a hard limit of 500,000,000 bytes on
all messages. Attempting to send or receive a larger message will result in
a connection close. This size was chosen to accommodate Index messages
containing a large block list. It's intended that the limit may be further
reduced in a future protocol update supporting variable block sizes (and
thus shorter block lists for large files).

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
messages and then Index records. The Index records are received and both
peers recompute their knowledge of the data in the cluster. In this example,
peer A has four missing or outdated blocks. At 5 through 8 peer A sends
requests for these blocks. The requests are received by peer B, who
retrieves the data from the folder and transmits Response records (9 through
12). Device A updates their folder contents and transmits an Index Update
message (13). Both peers enter idle state after 13. At some later time 14,
the ping timer on device B expires and a Ping message is sent. The same
process occurs for device A at 15.

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

