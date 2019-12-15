Untrusted (Encrypted) Devices
=============================

It is possible to set a password on a folder when it's shared with another
device. Data sent will be encrypted by this password, and data received will
be decrypted by the same password.

As an example, lets assume a *trusted* device ``T1``, maybe your laptop. You
have sensitive documents here but they are in cleartext from Syncthing's
point of view (perhaps protected by full disk encryption). There is also an
*untrusted* device ``U1``, perhaps a cloud server, where we want to sync
data but in unreadable form.

We set a folder password on ``T1`` when sharing the folder with ``U1``. Data
on disk on ``T1`` is not affected, but data sent to ``U1`` becomes encrypted
-- garbage, if you don't know the password.

.. graphviz::
    :align: center

    digraph g {
        rankdir=LR;
        "T1" [label="T1\n(Clear text)", style=filled, color="/accent3/1"];
        "U1" [label="U1\n(Encrypted)", style=filled, color="/accent3/2"];

        T1 -> U1 [label="Encrypted by T1"];
        U1 -> T1 [label="Decrypted by T1"];
    }

From this setup it's also possible to add further trusted devices, say
``T2``, and have these sync the data from the untrusted device ``U1``
without being in contact with ``T1``. Using the *same folder password* on
``T2`` makes the existing data on ``U2`` intelligeble and the plaintext data
becomes available.

.. graphviz::
    :align: center

    digraph g {
        rankdir=LR;
        "T1" [style=filled, color="/accent3/1"];
        "U1" [style=filled, color="/accent3/2"];
        "T2" [style=filled, color="/accent3/1"];

        T1 -> U1 [label="Encrypted by T1"];
        U1 -> T1 [label="Decrypted by T1"];
        U1 -> T2 [label="Decrypted by T2"];
        T2 -> U1 [label="Encrypted by T2"];
    }

Configuration
-------------

GUI
~~~

TBD

config.xml
~~~~~~~~~~

This is the configuration on a trusted device. Here the folder ``default``
is shared with three devices. The device ``373HSRP`` is a traditional
trusted peer. The other two devices ``CJBIJBJ`` and ``I6KAH76`` are both
untrusted and will get encrypted folder data, using different passwords.

.. code-block:: text

    <folder id="default" ...>
        <device id="373HSRP-..."></device>
        <device id="CJBIJBJ-..." encryptionPassword="foo"></device>
        <device id="I6KAH76-..." encryptionPassword="bar"></device>
    </folder>

There is no specific configuration required on the untrusted devices; they
will simply accept the encrypted data as is.
