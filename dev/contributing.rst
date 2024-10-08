.. _contribution-guidelines:

Contribution Guidelines
=======================

Getting Started
---------------

So you want to contribute? Great! Here's a short checklist with the most
important points:

- Don't worry. You are not expected to get everything right on the first
  attempt, we'll guide you through it.

- Make sure there is an `issue
  <https://github.com/syncthing/syncthing/issues>`__ that describes the
  change you want to do. If the thing you want to do does not have an issue
  yet, please file one before starting work on it.

- Fork the repository and make your changes in a new branch. If you already
  have push access to the Syncthing repository, do *not* create a new branch
  there. We do all changes as pull requests from personal forks.

Authorship
----------

All code authors are listed in the AUTHORS file. When your first pull request
is accepted, the maintainer will add your details to the AUTHORS file, the
NICKS file and the list of authors in the GUI. Commits must be made with the
same name and email as listed in the AUTHORS file. To accomplish this, ensure
that your git configuration is set correctly prior to making your first
commit::

    $ git config --global user.name "Jane Doe"
    $ git config --global user.email janedoe@example.com

You must be reachable on the given email address. If you do not wish to use
your real name for whatever reason, using a nickname or pseudonym is perfectly
acceptable.

Team Membership
---------------

Contributing useful changes via pull requests will at some point get you
invited to the ``contributors`` team -- typically, after having contributed
five or more nontrivial changes during the last year. This team gives you
push access to most repositories, subject to the guidelines below.

The ``maintainers`` team is for long standing contributors with the added
responsibility of reviewing major changes.

Code Review
-----------

Commits will generally fall into one of the three categories below, with
different requirements on review.

Trivial:
  A small change or refactor that is obviously correct. These may be pushed
  without review by any member of the ``maintainers`` team. Examples:
  `removing dead code <https://github.com/syncthing/syncthing/commits/main>`__,
  :commit:`updating values <2aa028facb7ccbe48e85c8c444501cc3fb38ef24>`.

Minor:
  A new feature, bugfix or refactoring that may need extra eyes on it to weed
  out mistakes, but is architecturally simple or at least uncontroversial.
  Minor changes must go through a pull request and can be merged on approval
  by any other developer on the ``contributors`` or ``maintainers`` team.
  Examples: `adding caching <https://github.com/syncthing/syncthing/pull/2432/files>`__,
  `fixing a small bug <https://github.com/syncthing/syncthing/pull/2406/files>`__.

Major:
  A complex new feature or bugfix, a large refactoring, or a change to the
  underlying architecture of things. A major change must be reviewed by a
  member of the ``maintainers`` team.

Infrastructure:
  Changes to the build system, release process, or other infrastructure
  components. Iteration may sometimes happen on branches in the main repo,
  to test interactions with GitHub Actions, etc. These should be reviewed by
  a member of the ``maintainers`` team.

Coding Style
------------

General
~~~~~~~

- All text files use Unix line endings. The git settings already present in
  the repository attempt to enforce this.

- When making changes, follow the brace and parenthesis style of the
  surrounding code.

Go Specific
~~~~~~~~~~~

- Follow the conventions laid out in `Effective
  Go <https://go.dev/doc/effective_go>`__ as much as makes
  sense. The review guidelines in `Go Code Review Comments
  <https://github.com/golang/go/wiki/CodeReviewComments>`__ should generally
  be followed.

- Each commit should be ``go fmt`` clean.

- Imports are grouped per ``goimports`` standard; that is, standard
  library first, then third party libraries after a blank line.

Commits
-------

- Commit messages (and pull request titles) should follow the [conventional
  commits](https://www.conventionalcommits.org/en/v1.0.0/) specification and
  be in lower case.

- We use a scope description in the commit message subject. This is the
  component of Syncthing that the commit affects. For example, ``gui``,
  ``protocol``, ``scanner``, ``upnp``, etc -- typically, the part after
  ``lib/`` or ``cmd/`` in the package path. If the commit doesn't affect a
  specific component, such as for changes to the build system or
  documentation, the scope should be omitted. The same goes for changes
  that affect many components which would be cumbersome to list.

- Commits that resolve an existing issue must include the issue number
  as ``(fixes #123)`` at the end of the commit message subject. A correctly
  formatted commit message looks like this::

    feat(dialer): add env var to disable proxy fallback (fixes #3006)

- If the commit message subject doesn't say it all, one or more paragraphs of
  describing text should be added to the commit message. This should explain
  why the change is made and what it accomplishes.

- When drafting a pull request, please feel free to add commits with
  corrections and merge from ``main`` when necessary. This provides a clear time
  line with changes and simplifies review. Do not, in general, rebase your
  commits, as this makes review harder.

- Pull requests are merged to ``main`` using squash merge. The "stream of
  consciousness" set of commits described in the previous point will be reduced
  to a single commit at merge time. The pull request title and description will
  be used as the commit message.

Tests
-----

Yes please, do add tests when adding features or fixing bugs. Also, when a
pull request is filed a number of automatic tests are run on the code. This
includes:

- That the code actually builds and the test suite passes.

- That the code is correctly formatted (``go fmt``).

- That the commits are based on a reasonably recent ``main``.

- That the author is listed in AUTHORS.

- That the output from ``go lint`` and ``go vet`` is clean. (This checks for a
  number of potential problems the compiler doesn't catch.)

Branches
--------

- ``main`` is the main branch containing good code that will end up
  in the next release. You should base your work on it. It won't ever
  be rebased or force-pushed to.

- ``vx.y`` branches exist to make patch releases on otherwise obsolete
  minor releases. Should only contain fixes cherry picked from ``main``.
  Don't base any work on them.

- ``infrastructure`` is a specific branch from which builds for the
  infrastructure components (usage reporting server, crash reporting server,
  relay pool server, etc) are sometimes made. It may be ahead of ``main``.
  Do not base any work on it.

- Other branches are probably topic branches and may be subject to rebasing.
  Don't base any work on them unless you specifically know otherwise.
  Generally, avoid creating branches on the main repo, preferring instead to
  have topic branches on your own fork.

Tags
----

All releases are tagged semver style as ``vx.y.z``. The maintainer doing the
release signs the tag using their GPG key.

Licensing
---------

All contributions are made under the same MPLv2 license as the rest of the
project, except documentation, user interface text and translation strings
which are licensed under the Creative Commons Attribution 4.0 International
License. You retain the copyright to code you have written.
