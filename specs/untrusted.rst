Untrusted Device Encryption
===========================

Threat Model / Primary Goals
----------------------------

An "untrusted" device can participate in a Syncthing cluster with the
following assumptions and limitations;

The untrusted device should *not* be able to observe:

- File data
- File or directory names, symlink names, symlink targets
- File modification time, permissions, or modification history (version vectors)

The untrusted device *will* be able to observe:

- Which other devices are paired with it

- File sizes (although files grow slightly due to block
  overhead, and some files are padded up to an even kilobyte, file sizes
  can be determined at least to the closest kilobyte)

- When and which parts of files are changed by the other devices

- Which identical blocks are reused within any given file

The last two points (identifying changed and reused blocks) are required by
the syncing mechanism, in order to avoid transferring all unchanged file
data when a file block changes. Blocks and block hashes are encrypted with a
per-file key so correlation is not possible between files -- just within any
given file.

In addition the untrusted device *must not* be able to modify, remove or
introduce data by itself without detection.

Primitives Used
---------------

The user input to the system is the *folder ID*, which is a short string
identifying a given folder between devices, and the *password*. From this we
generate a *folder key* using ``scrypt`` (32 bytes)::

    folderKey = Scrypt(password, "syncthing" + folderID)

The string "syncthing" with the folder ID concatenated make up the salt. The
folder key is used to encrypt file names using AES-SIV::

    encryptedFilename = AES-SIV(filename, folderKey)

To make the encrypted file name usable again as a file name, we encode it
using base32 and add slashes at strategic places.

From the folder key and the plaintext file name we derive the *file key* by
xor:ing the folder key with the SHA256 of the plaintext file name::

    fileKey = folderKey ^ SHA256(filename)

This file key is used for all other encryption, specifically file block
hashes and data blocks. In file metadata, block hashes are encrypted using
AES-SIV with the file key::

    encryptedBlockHash = AES-SIV(blockHash, fileKey)

Data blocks are encrypted using XChaCha20-Poly1305 with random nonces and
appended to the nonce itself::

    encryptedBlock = nonce + XChaCha20-Poly1305.Seal(blockData, fileKey)

The original file metadata descriptor is encrypted in the same manner and
attached to the encrypted-file metadata.

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

Implementation Details
----------------------

Metadata Encryption
~~~~~~~~~~~~~~~~~~~

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
