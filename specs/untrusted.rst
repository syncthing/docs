.. _untrusted:

.. warning::
   This feature is currently in public testing and not yet recommended to be
   used for important data. Related UI controls are hidden by a feature flag -
   check the release notes for information on how to test it.

Untrusted Device Encryption
===========================

Threat Model / Primary Goals
----------------------------

An "untrusted" device can participate in a Syncthing cluster with the
following assumptions and limitations;

The untrusted device can *not* observe:

- File data

- File or directory names, symlink names, symlink targets

- File modification time, permissions, or modification history (version
  vectors)

The untrusted device *will* be able to observe:

- File sizes [#sizes]_

- Which parts of files are changed by the other devices and when

The last point is required by the syncing mechanism, in order to avoid
transferring all unchanged file data when a file block changes. Blocks and
block hashes are encrypted with a per-file key and depends on the block
offset, so correlation is not possible between blocks at different offsets
or different files.

In addition the untrusted device *must not* be able to modify, remove or
introduce data by itself without detection.

Primitives Used
---------------

The user input to the system is the *folder ID*, which is a short string
identifying a given folder between devices, and the *password*. From this we
generate a *folder key* (32 bytes) using ``scrypt``::

    folderKey = scrypt.Key(password, "syncthing" + folderID)

The string "syncthing" with the folder ID concatenated make up the salt. The
folder key is used to encrypt file names using AES-SIV without nonce::

    encryptedFilename = AES-SIV(filename, folderKey)

Given the key length of 32 bytes the algorithm in use will be AES-128
("AES-SIV-256"). To make the encrypted file name usable again as a file
name, we encode it using base32 and add slashes at strategic places.

From the folder key and the plaintext file name we derive the *file key* by
HKDF of the folder key and the plaintext file name::

    fileKey = HKDF(SHA256, folderKey+filename, salt: "syncthing", info: nil)

This file key is used for all other encryption, specifically file block
hashes and data blocks. In file metadata, block hashes are encrypted using
AES-SIV with the file key::

    encryptedBlockHash = AES-SIV(blockHash, fileKey)

Data blocks are encrypted using XChaCha20-Poly1305 with random nonces and
appended to the nonce itself::

    encryptedBlock = nonce + XChaCha20-Poly1305.Seal(block, nonce, fileKey)

The original file metadata descriptor is encrypted in the same manner and
attached to the encrypted-file metadata.

Devices sharing a folder need to use the same password.
To ensure that a *password token* in the form of an arbitrary, but commonly
known string encrypted using AES-SIV with the folder key is sent in the
:ref:`cluster-config`::

    passwordToken = AES-SIV("syncthing" + folderID, folderKey)

Thus an encrypted device can verify all its connected devices use the same
password comparing the encrypted token, without knowing the password itself.

.. note::

    In Syncthing a file is made up of a number of equal size data blocks,
    followed by a usually shorter last data block. The full size data blocks
    are at minimum 128 KiB, ranging up to 16 MiB in multiples of two. The
    last data block can in principle be as small as one byte. For untrusted
    folders the size of the last data block is padded up to a kilobyte if it
    was shorter to begin with. The untrusted device isn't allowed to request
    less than a kilobyte of data.

    I don't actually know if this block padding serves a purpose. It was
    added to address a worry that something might break or leak if an
    attacker is allowed to repeatedly request single-byte data blocks of
    their choosing. If there is nothing to worry about here we can remove
    the padding. //jb

.. note::

    While a well behaved implementation is expected to request data blocks
    precisely as announced in the file metadata there is no enforcement of
    this. This means that an attacker on the untrusted side can repeatedly
    request arbitrary ranges of a file and receive the encrypted result.
    With the restriction above, the minimum block size that can be requested
    in 1024 bytes.


Implementation Details
----------------------

Metadata Encryption
~~~~~~~~~~~~~~~~~~~

The Syncthing protocol is essentially two-phase:

- A device sends file metadata (a ``FileInfo`` structure) for a new or changed file

- The other side determines which blocks it needs to construct the new file, and requests these blocks

For untrusted devices a fake FileInfo is constructed, with an encrypted
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

The fake FileInfo encrypts and adjusts a couple of attributes:

- The name is encrypted (with the folder key), base32 encoded, and slashes
  are inserted after the first and third characters, and then every 200
  characters.

- The file size is adjusted for the per block overhead, and rounded up so that
  the last block is a multiple of 1024 bytes.

- The block size is adjusted for block overhead.

Other file attributes are set to static values, for example the modification
time is set to UNIX epoch time 1234567890 and permissions are set to 0644.

The block list is encrypted and adjusted:

- The offset and size are adjusted to account for block overhead

- The hash is encrypted using AES-SIV (with the file key)

The resulting encrypted hash can't be used for data verification by the
untrusted device, but it can be used as a form of "token" referring to a
given data block for reuse purposes.

Finally, the whole original FileInfo (in protobuf form) is encrypted using
XChaCha20-Poly1305 with the file key and attached to the fake FileInfo. This
is retained on the untrusted side and passed along to trusted devices, where
it will be used in place of the fake FileInfo.

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

Incoming Metadata
~~~~~~~~~~~~~~~~~

File metadata sent from the untrusted device is always decrypted. This means
the original FileInfo is discarded and the attached encrypted FileInfo is
decrypted and used instead. If the FileInfo does not decrypt it's considered
a protocol error and the connection is dropped. This means only file
metadata created by a trusted device is accepted.

Data Encryption
~~~~~~~~~~~~~~~

When an untrusted device makes a request for a data block, the trusted
device:

1. decrypts the given filename,
2. reads the corresponding plaintext data block,
3. pads the block with random data if the read returned less than 1024 bytes,
4. encrypts it using the file encryption key and a random nonce, and
5. responds with the result.

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
            label = "nonce (24 B) | <h> ciphertext (variable) | tag (16 B)"
            shape = "record"
        ]
        u:h -> e:h [ label = "XChaCha20-Poly1305" ]
    }

This is repeated for all required blocks. At the end, the untrusted device
appends the fake FileInfo (which includes the original, encrypted, FileInfo)
to the file. This serves no purpose during normal operations, but enables
offline decryption of an encrypted folder without database access and, in
principle, scanning an encrypted folder to populate the database should it
be lost or corrupted.

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

Incoming Data
~~~~~~~~~~~~~

Making a request to an untrusted device is mostly the reverse of the above.
The file name is encrypted and the block offset and size adjusted. The
resulting data is decrypted and thereby also authenticated, meaning it must
have originated on a trusted device. Contrary to the usual case we cannot
simply make arbitrary range requests -- only the precise blocks that were
encrypted to begin with will decrypt properly.

---

.. [#sizes] Although files grow slightly due to block
    overhead, and some files are padded up to an even kilobyte, file sizes
    can be determined at least to the closest kilobyte.
