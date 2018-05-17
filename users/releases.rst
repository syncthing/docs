.. _release-channels:

Release Channels
================

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

.. [#] Auto upgrades are not enabled in builds delivered via APT or Snap.
.. [#] Yes, there is intentionally no difference here.

Life Cycle
----------

Every new feature and bugfix begins its life in the development branch,
``master``. Once a month the current ``master`` becomes a *release
candidate*. This version is identified by "-rc" in it's name, for example
"0.14.35-rc.1".

Those running the candidate channel will update to this release candidate.
For the next three weeks it is tested "in the wild". Any new, serious issues
that are discovered are fixed, and new release candidates "0.14.35-rc.2" etc
are created as needed. These release candidates do not include any new
features or non-essential bugfixes added to ``master`` in the meantime.

Stable releases are given version numbers without any suffix - "0.14.35".
Unless any serious issues were discovered, this release is exactly identical
to the "-rc.1" release candidate three weeks prior.

The cycle then restarts one week later with a new release candidate based on
the current ``master`` branch.

Releases are timed so that stable releases happen on the first Tuesday of
the month. The next candidate release is one week later, on the second
Tuesday of the month.

How to Choose
-------------

Built-in / GitHub
~~~~~~~~~~~~~~~~~

For releases obtained from Syncthing.net or GitHub, with built-in upgrade
functionality, the choice is made in the "Settings" dialog. Set the
"Automatic upgrade" drop down to either "Stable releases only" or "Stable
releases and release candidates".

APT (Debian)
~~~~~~~~~~~~

The choice between stable and candidate is done in the APT source
configuration. Please see `our APT instructions
<https://apt.syncthing.net/>`__.

Snap
~~~~

The ``snap`` tool can be told to install the candidate channel, but defaults
to the stable channel. See the Snap documentation for detail.

Some Other Distribution Channel
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

If you are getting packages from your Linux distribution, NAS vendor, etc.,
then you should be getting the *stable* channel. If you get a release
candidate you should complain to your distributor or vendor and refer them
to this page.

FAQ
---

Which bugfixes trigger a new release candidate?
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Those that fix a regression since the last release. Lets say the current
release is 0.14.35. We release 0.14.36-rc.1 and discover a new problem that
is not present in 0.14.35. This gets fixed and we release a new 0.14.36-rc.2
candidate. However, if we discover and fix a problem that's been present
since 0.14.20, this fix will instead be incorporated in the next regular
cycle.

What's the difference between the latest candidate and the following stable release?
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Nothing. If we release 0.14.36-rc.1 and no serious problems are discovered
during the next twelve days, this is the exact software that will become
0.14.36 for general consumption. Since the version number is different it
requires a rebuild and the release signatures / hashes are different. If you
are on the candidate channel, your Syncthing will "upgrade" from
0.14.36-rc.1 to 0.14.36 when we make the release. This is normal.
