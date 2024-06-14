Creating a Release
==================

Prerequisites
-------------

- About ten minutes to half an hour of free time.

- The main branch in a clean and buildable state, full of commits you are proud of and know the users will love. This is of course the default state at any given time.

- A normal computer (real or virtual) that has a command line and can run bash scripts. Macs and Linux boxes are good choices here. If you know what you're doing I'm sure it's entirely possible to do it on Windows as well - but then you're on your own. In a pinch you can use ``secure.syncthing.net`` as it has all the required tools installed, although you'll need to add your git config, keys etc.

- The source repo, Git, Go, and GPG. It doesn't much matter which version of Go, we are not going to be building the release artifacts with it. If you can build the project you are good to go. If you can't build it, please don't attempt to make a release of it.

- Push access to the repo, being able to bypass the pull request requirements. This means ``admin`` or ``maintainers`` group.

- An account on ``secure.syncthing.net`` with sudo access.

- A TeamCity account on ``build.syncthing.net`` with deploy access.

Release Procedure
-----------------

The procedure differs slightly depending on whether we're doing a release candidate or a stable release. Candidate releases require work to prepare the changelog, which will just be reused for the stable release. The stable release on the other hand requires a slightly different release process and is announced more widely.

Release Candidates - Write a Change Log
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Most of the change log is machine generated from the closed issues. We do however need to make sure that issues belong to the correct milestone, have the correct labels, and that the issue subject makes sense as a line in the change log. To our help we have the purpose written tool `grt <https://github.com/syncthing/github-release-tool>`__. The grt tool requires your GitHub token to manage milestones and issues; you set the environment variable ``GITHUB_TOKEN`` while you are working on the release (but hopefully not all the time - programs can and do steal environment data).

To ensure that all closed issues are tagged with the milestone for the release you are doing, use the following command. First, find the commit hash or tag of the last commit on the *previous* release - changes since this point is what we are going to consider part of this release. If there haven't been any special releases or branching you can simply use the previous release as the starting point.

.. code-block:: bash

    $ grt milestone v0.14.50 --from=v0.14.49

Visit the milestone in your browser and double check the issue subjects and labels. Remember that only closed issues (not pull requests) will appear in the change log. If there are specific things to note about this release, such as changed APIs or config formats, briefly describe these changes in the notes field. You can preview the change log using grt:

.. code-block:: bash

    $ grt changelog v0.14.50
    Bugfixes:

     - #5063: panic: cannot start already running folder
     - #5073: lib/logger: tests fail due to compilation error with go 1.11

In principle the output should be a complete, valid release note for the release in question. Pipe it to a text file. If you're preparing a release candidate you should specify the full candidate tag (e.g. v0.14.50-rc.1) to the command above to get the correct title on the changelog.

Add further notes or commentary to taste, if required.

Stable Release - Write a Change Log
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Use the change log from the corresponding release candidate, just change the version in the first line and the first sentence that describes the release status.

Prepare the Release Branch
~~~~~~~~~~~~~~~~~~~~~~~~~~

Releases come from the ``release`` branch. If you are making a new candidate release you will want to fast forward ``release`` to point at current ``main`` ``HEAD``. If you are making a stable release from the latest RC the ``release`` branch is already in the right place.

.. code-block:: bash

    $ git checkout release
    $ git merge --ff-only main

If there's been some funky business with the ``release`` branch and it can't be fast forwarded to ``main``, 1) the previous release manager screwed up, 2) don't do a merge, just reset the branch to the right place.

Don't push the branch yet, we want to create the tag first.

Create and Push the Tag
~~~~~~~~~~~~~~~~~~~~~~~

Make sure you push the tag before the release branch, as the latter
currently triggers the release job on the CI (and requires the tag already being in place).

.. code-block:: bash

    $ git tag -a -s -F ~/changelog.txt v0.14.50
    $ git push --tags
    $ git push origin release:release

The changelog file is the one you prepared previously.

You will need your PGP key at hand for this step. It should be your personal PGP key, whatever you would normally use. If you don't have one you'll need to create one for the purpose. Keep it around, keep it secure, upload the public part to a key server.

If your remote spec is nondefault, tailor the push command to suit. We deliberately pushed the tags before the release branch, because the builder may start building the release branch immediately and needs to see the right tags at that point.

Build the Packages
~~~~~~~~~~~~~~~~~~

If you are building a release candidate and fast forwarded the ``release`` branch the build server will already have started building it. If not, jump in on the build server and trigger the Release/Syncthing job, for the ``release`` branch, while checking the options to rebuild all dependencies in the chain. We need the rebuild for those binaries to pick up the new tag.

Once the build succeeds, log in on ``secure.syncthing.net``. If something failed in the build it's hopefully "just" a flaky test - redo the build.

Create the GitHub release
~~~~~~~~~~~~~~~~~~~~~~~~~

From this point on we will work on ``secure.s.n``, as the ``release`` user.

.. code-block:: bash

    jb@laptop$ ssh secure.syncthing.net
    jb@secure$ sudo su - release
    release@secure$

We will use grt to create the release with the appropriate change log, and possibly close the milestone. If we are doing a candidate release we need to specify the tag including the candidate suffix:

.. code-block:: bash

    $ grt release v0.14.50-rc.1

This will create a v0.14.50-rc.1 release, with the "pre-release" bit set, and leave the v0.14.50 milestone open. For a stable release:

.. code-block:: bash

    $ grt release v0.14.50

The milestone will be closed.

Sign and upload the archives
~~~~~~~~~~~~~~~~~~~~~~~~~~~~

At this point the build should have completed and the artifacts should have been uploaded to ``secure.s.n``. If the build number was 1234 and the version v0.14.50 the files will be in ``/home/incoming/build-1234-v0.14.50``. Run the following scripts. None of them should fail, barring connectivity issues - so if they do, you get to fix whatever it is without any guidance from me. Sorry.

.. code-block:: bash

    $ sign-upload-debian /home/incoming/build-1234-v0.14.50

Publishes the Debian archives to apt.syncthing.net.

.. code-block:: bash

    $ upload-release /home/incoming/build-1234-v0.14.50

Publishes the regular release archives to GitHub.

Stable Releases - Trigger update of the website
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

The main website needs to be redeployed to reflect the new release version on the download page.

.. code-block:: bash

    $ ./deploy-web

Stable Releases - Create a post on the forum
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

In the Announce/Releases category. Use the tag message as the template, make the header a link to the release, make the issue numbers to be links to the corresponding issues. You can use ``grt changelog v0.14.50 --md`` to get the change log with issue links in proper Markdown.

Stable Releases - Optionally, tweet it
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

If you have the Twitter account and the release isn't a cake-in-your-face screwup fix that you'd rather no one ever heard about and want to just silently roll out to everyone during the night.

Stable Releases - Create new milestone
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Create a milestone for the next release, with the due date set to the first
Tuesday of the next month.
Which version number to bump depends on what kind of changes are already in main (see :ref:`semver`).
This might change in the time until the first candidate is released.

Merge Release Into Main
~~~~~~~~~~~~~~~~~~~~~~~

If this was a non-first candidate release with cherry picked commits on it, merge ``release`` back into ``main`` and push ``main``.
