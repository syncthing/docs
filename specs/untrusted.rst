Untrusted Device Encryption
===========================

Threat Model
------------

An "untrusted" device can participate in a Syncthing cluster with the
following assumptions and limitations;

The untrusted device should *not* be able to observe:

- File data
- File or directory names, symlink names, symlink targets
- File modification time, permissions

The untrusted device *will* be able to observe:

- Which other devices are paired with it
- Approximate file sizes (although files grow slightly due to block overhead, and the last block is padded up to an even kilobyte, file sizes can be determined to the closest kilobyte)
- When and which parts of files are changed by the other devices
- Which identical blocks are reused by more than one file

The last two points (identifying changed and reused blocks) are required
by the syncing mechanism, in order to avoid transferring all unchanged file
data when a file block changes. The actual block data is encrypted by
XChaCha20-Poly1305 and random nonces, but encrypted block hashes on the form
``AES-SIV(plaintext hash)`` are available to the untrusted device.

In addition the untrusted device *must not* be able to modify, remove or
introduce data by itself without detection.

Secondary Goals
~~~~~~~~~~~~~~~

Apart from fulfilling the threat model, the implementation also aims for the
following goals:

- An encrypted folder should be self contained.
  This is so that it can be copied using non-Syncthing tools and then
  either used again as an encrypted folder in Syncthing or decrypted in
  place using an appropriate tool.

- Encrypted file names should be reasonable.
  We must avoid special characters and stick to length limits appropriate for
  Windows file systems.

Encryption Keys
---------------

A password is set on the trusted device, per remote untrusted device and
folder. The actual 32-byte encryption key is generated using scrypt and the
folder ID as salt::

    scrypt.Key(password, "syncthing" + folderID, 32768, 8, 1, keySize)

This key is used both for AES-SIV where deterministic encryption is required
(file names and block hashes) and for pure data (XChaCha20-Poly1305 with
random nonces).

Metadata Encryption
-------------------

The Syncthing protocol is essentially two phase:

- A device sends file metadata for a new or changed file
- The other side determines which blocks it needs to construct the new file, and requests these blocks

For untrusted devices a fake file metadata is constructed, with an encrypted
name and block list and other metadata such as modification time and
permissions set to static values.

An original file metadata structure looks something like this:

.. graphviz::

    digraph g {
        graph [
            rankdir = "LR"
        ]
        "fileinfo" [
            label = "name | type | size | modified | ... | <b> blocks | block size"
            shape = "record"
        ]
        "blocks" [
            label = "{ <a> offset | size | hash } | { offset | size | hash } | ..."
            shape = "record"
        ]
        fileinfo:b -> blocks:a
    }

The fake metadata encrypts and adjusts a couple of attributes:

- The name is encrypted using AES-SIV, base32 encoded, and slashes are
  inserted after the first and third characters, and then every 200
  characters.

- The size is adjusted for the per block overhead, and rounded up so that
  the last block is a multiple of 1024 bytes.

- The block size is adjusted for block overhead.

Other file attributes are set to static values, for example the modification
time is set to UNIX epoch time 1234567890 and permissions are set to 0644.

The block list is encrypted and adjusted:

- The offset and size are adjusted to account for block overhead
- The hash is encrypted using AES-SIV

The resulting encrypted hash can't be used for data verification by the
untrusted device, but it can be used as a form of "token" referring to a
given data block for reuse purposes.

Finally, the whole original file metadata (in protobuf form) is encrypted
using XChaCha20-Poly1305 and attached to the fake fileinfo. This is retained
on the untrusted side and passed along to trusted devices, where it will be
used in place of the fake fileinfo.

.. graphviz::

    digraph g {
        graph [
            rankdir = "LR"
        ]
        "fileinfo" [
            label = "encrypted name | ... | adjusted size | ... | <b> encrypted blocks | adjusted block size | encrypted metadata"
            shape = "record"
        ]
        "blocks" [
            label = "{ <a> offset + n * overhead | size + overhead | encrypted hash } | { <a> offset + n * overhead | size + overhead | encrypted hash } | ..."
            shape = "record"
        ]
        fileinfo:b -> blocks:a
    }

Data Encryption
---------------

When an untrusted device makes a request for a data block, the trusted
device reads the corresponding plaintext data block, encrypts it using the
encryption key and a random nonce, and responds with the result. If the
requested block was the last block in the file and size rounding resulted in
a request for more data than was avaialble, additional random data is added
to fulfill the request.

.. graphviz::

    digraph g {
        graph [
            rankdir = "LR"
        ]
        "u" [
            label = "<h> plaintext (variable)"
            shape = "record"
        ]
        "e" [
            label = "nonce (24 B) | tag (16 B) | <h> ciphertext (variable)"
            shape = "record"
        ]
        u:h -> e:h [ label = "XChaCha20-Poly1305" ]
    }

This is repeated for all required blocks. At the end, the untrusted device
appends the fake metadata (including the correct, encrypted, metadata) to
the file. This serves no purpose during normal operations, but enables
offline decryption of an encrypted folder without database access.

.. graphviz::

    digraph g {
        graph [
            rankdir = "LR"
        ]
        "u" [
            label = "<b0> plaintext block | <b1> plaintext block | ..."
            shape = "record"
        ]
        "e" [
            label = "<b0> encrypted block | <b1> encrypted block | ... | FileInfo (variable) | len(FileInfo) (uint32)"
            shape = "record"
        ]
        u:b0 -> e:b0 [ label = "encryption" ]
        u:b1 -> e:b1
    }
