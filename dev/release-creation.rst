Creating a Release
==================

Prerequisites
-------------

- About ten minutes to half an hour of free time.

- The master branch in a clean and buildable state, full of commits you are proud of and know the users will love. This is of course the default state at any given time.

- A normal computer (real or virtual) that has a command line and can run bash scripts. Macs and Linux boxes are good choices here. If you know what you're doing I'm sure it's entirely possible to do it on Windows as well - but then you're on your own. In a pinch you can use ``secure.syncthing.net`` as it has all the required tools installed, although you'll need to add your git config, keys etc.

- The source repo, Git, Go, and GPG. It doesn't much matter which version of Go, we are not going to be building the release artifacts with it. If you can build the project you are good to go. If you can't build it, please don't attempt to make a release of it.

- Push access to the repo, being able to bypass the pull request requirements. This means ``admin`` or ``maintainers`` group.

- An account on ``secure.syncthing.net`` with sudo access.

- An Jenkins account on ``build.syncthing.net``.

Release Procedure
=================

Update translations and documentation
-------------------------------------

.. code-block:: bash

    $ ./build.sh prerelease
    $ git add -A
    $ git commit -m "gui, man: Update docs & translations"
    $ git push

This pulls in translations from Transifex and documentation (man pages) from docs.syncthing.net. If your remote spec is nondefault, tailor the push command to suit.

Write a change log
------------------

Look at the previous ones for inspiration; they are in the tag message for each recent previous release. The first line is the version being described. Then start with a sentence describing the release type (scheduled feature and bugfix, hotfix, ...) and who should use it. Follow with the set of issues closed since last release - each being a bullet point describing the new behavior, not the problem or error as it was previously. The format is semi machine readable - bullet points are a single text line, regardless of length.

Add further notes or commentary to taste. Separate bugfixes from features if it makes sense to do so.

The change log script can assist in listing the issues, their subjects and the commit message that closed them. The format is not correct to be used as is, but it can refresh your memory enough to write the real notes.

.. code-block:: bash

    $ go run script/changeless.go

Create and push the tag

.. code-block:: bash

    $ git tag -a -s -F ~/changelog.txt v0.14.18
    $ git push --tags

The changelog file is the one you prepared previously.

You will need your PGP key at hand for this step. It should be your personal PGP key, whatever you would normally use. If you don't have one you'll need to create one for the purpose. Keep it around, keep it secure, upload the public part to a key server.

If your remote spec is nondefault, tailor the push command to suit.

Build the packages
------------------

Trigger the ‘syncthing-release’ job on the build server, giving it the newly created tag name. It'll trigger the Mac and Windows builds in parallel. Wait for them to complete and verify that nothing failed. You can create the GitHub release in the next step while waiting.


Create the GitHub release
-------------------------

Find the new tag on the GitHub release page. It'll show you the tag message before the release is created, you can copy and paste that. Tweak formatting if necessary. Fix any typos you see, but also kick yourself for not noticing before creating the tag and vow to do it better next time. Publish the release.

Sign and upload the archives
----------------------------

The release keys etc are hosted on secure.syncthing.net under he user account `release`. SSH there and become the ``release`` user.

.. code-block:: bash

    jb@laptop$ ssh secure.syncthing.net
    jb@secure$ sudo su - release
    release@secure$

This is the point where you ensure the builds you started earlier are all good. If they are not, fix that first. Run the following scripts. None of them should fail, barring connectivity issues - so if they do, you get to fix whatever it is without any guidance from me. Sorry.

.. code-block:: bash

    $ sign-upload-debian

Publishes the Debian archives to apt.syncthing.net.

.. code-block:: bash

    $ upload-snaps

Publishes the Snap packages to Ubuntu.

.. code-block:: bash

    $ sign-upload-release

Publishes the regular release archives to GitHub.

Create a post on the forum
--------------------------

In the Announce/Releases category. Use the tag message as the template, make the header a link to the release, edit the issue numbers to be links to the corresponding issues.

If you remember to, lock the previous release announcement. Or don't, as there is not much discussion on the release posts anyhow.

Optionally, tweet it
--------------------

If you have the Twitter account and the release isn't a cake-in-your-face screwup fix that you'd rather no one ever heard about and want to just silently roll out to everyone during the night.