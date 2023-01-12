.. _releases:

Versions & Releases
===================

.. _semver:

Major, Minor, or Patch
----------------------

Since the 1.0.0 release, Syncthing uses a `semver
<https://semver.org/>`__-like [1]_ three part version number, **x.y.z** where *x*
is the major version, *y* is the minor version, and *z* is the patch
version. We decide the version number for a new release based on the
following criteria:

- Is the new version protocol incompatible with the previous one, so that
  they cannot connect to each other or otherwise can't sync files for some
  reason? That's a new *major* version. (This hasn't happened yet.)

- Are there changes in the REST API so that integrations or wrappers
  need changes, or did the database schema or configuration change so that a
  downgrade might be problematic? That's a new *minor* version.

- If there are no specific concerns as above, it's a new *patch* version.

Release Channels
----------------

There are two different release channels that can be selected. The *stable*
channel is the more stable one, while *candidate* releases are closer to
development. Candidate releases get promoted to stable after a certain
period of testing.

There are a few trade-offs between the two:

=========================  =========================  ======================
\                                   Stable                   Candidate
=========================  =========================  ======================
**Stability**              More Stable                More Experimental
**Features & Fixes**       One month behind           Latest
**Auto Upgrades**          Optional                   Mandatory [#]_
**Anon. Usage Reporting**  Optional                   Mandatory
**Support**                Fully supported            Fully supported [#]_
=========================  =========================  ======================

Run the candidate channel if you are technically savvy and enjoy new
features. Run the stable channel if you want to minimize the amount of
surprises you might run into.

.. [#] Auto upgrades are not enabled in builds delivered via APT.
.. [#] Yes, there is intentionally no difference here.

Schedule
~~~~~~~~

Barring blocking issues, stable versions are released *on the first Tuesday
of the month*. A new candidate releases is made *on the second Tuesday of the
month*.

How to Choose
~~~~~~~~~~~~~

Built-in / GitHub
^^^^^^^^^^^^^^^^^

For releases obtained from Syncthing.net or GitHub, with built-in upgrade
functionality, the choice is made in the "Settings" dialog. Set the
"Automatic upgrade" drop down to either "Stable releases only" or "Stable
releases and release candidates".

APT (Debian)
^^^^^^^^^^^^

The choice between stable and candidate is done in the APT source
configuration. Please see `our APT instructions
<https://apt.syncthing.net/>`__.

Docker
^^^^^^

Docker images are pushed to several tags. By pulling a specific tag you can
select the release channel you want:

- ``syncthing/syncthing:latest`` points to the latest stable release, e.g. ``v1.9.0``
- ``syncthing/syncthing:rc`` points to the latest release candidate, e.g. ``v1.10.0-rc.3``
- ``syncthing/syncthing:nightly`` points to the latest nightly development build, e.g. ``v1.10.0-rc.3.dev.1.gd2e32957``

There are also tags for the major and minor versions, for example
``syncthing/syncthing:1.9`` for the latest stable 1.9 patch release and
``syncthing/syncthing:1`` for the latest stable release with major version
one.

Some Other Distribution Channel
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

If you are getting packages from your Linux distribution, NAS vendor, etc.,
then you should be getting the *stable* channel. If you get a release
candidate you should complain to your distributor or vendor and refer them
to this page.

Nightly Builds
--------------

It's also possible to run the nightly development builds. These are not
releases in any sense of the word, they are simply builds of the current
``main`` branch of the repository. These builds are not supported and may
contain functionality that is changed or removed before the actual release.

To use nightly releases, in Advanced Settings -> Options,

- Change "Releases URL" to ``https://upgrades.syncthing.net/nightly.json``
- Enable "Upgrade To Pre Releases"

Restart Syncthing and it will upgrade to the latest nightly. New nightlies
are published at midnight UTC.

FAQ
~~~

What's the relationship between candidate and release exactly?
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Every new feature and bugfix begins its life in the development branch,
``main``. Once a month the current ``main`` becomes a *release
candidate*. This version is identified by "-rc" in its name, for example
``v1.5.0-rc.1``.

Those running the candidate channel will update to this release candidate.
For the next three weeks it is tested "in the wild". Any new, serious issues
that are discovered are fixed, and new release candidates ``v1.5.0-rc.2`` etc
are created as needed. These release candidates do not include any new
features or non-essential bugfixes added to ``main`` in the meantime.

Stable releases are given version numbers without any suffix - ``v1.5.0``.
Unless any serious issues were discovered, this release is exactly identical
to the "-rc.1" release candidate three weeks prior.

The cycle then restarts one week later with a new release candidate based on
the current ``main`` branch.

Which bugfixes trigger a new release candidate?
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Those that fix a regression since the last release. Lets say the current
release is ``v1.5.0``. We release ``v1.5.1-rc.1`` and discover a new problem that
is not present in ``v1.5.0``. This gets fixed and we release a new ``v1.5.1-rc.2``
candidate. However, if we discover and fix a problem that's been present
since ``v1.4.0``, this fix will instead be incorporated in the next regular
cycle.

What's the difference between the latest candidate and the following stable release?
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Nothing. If we release ``v1.5.1-rc.1`` and no serious problems are discovered
during the next twelve days, this is the exact software that will become
``v1.5.1`` for general consumption. Since the version number is different it
requires a rebuild and the release signatures / hashes are different. If you
are on the candidate channel, your Syncthing will "upgrade" from
``v1.5.1-rc.1`` to ``v1.5.1`` when we make the release. This is normal.

---

.. [1] SemVer-*like* because semver is absolutist about what constitutes an
       API change, in a way that isn't super helpful to the average user of a
       program like Syncthing.

Historical Releases
-------------------

This table lists the historically released versions of Syncthing, which Go
version they were built with, and which date they were released.

.. csv-table:: Syncthing Releases
   :file: releases.csv
   :header-rows: 1
   :align: left
