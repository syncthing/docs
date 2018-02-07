Issue Management
================

Bugs, feature requests and other things we need to do are tracked as
Github issues. Issues can be of various types and in various states, and
also belong to milestones or not. This page is an attempt to document
the current practice.

Labels
------

Issues without labels are undecided - that is, we don't yet know if it's
a bug, a configuration issue, a feature request or what. Issues that are
invalid for whatever reason are closed with a short explanation of why.
Examples include "Duplicate of #123", "Discovered to be configuration
error", "Rendered moot by #123" and so on. We don't use the "invalid" or
"wontfix" labels.

bug
    The issue is a verified bug.

build
    The issue is caused by or requires changes to the build system
    (scripts or Docker image).

enhancement
    This is a new feature or an improvement of some kind, as
    opposed to a problem (bug).

frozen-due-to-age
    This is set automatically on issues when they have been closed untouched
    for a long time, together with the issue being locked against
    discussion.

good-first-issue
    These issues are good starting points for new contributors. They are
    contained in scope and size, and it should be clear what is the desired
    outcome.

not-our-bug
    Rare and temporary label used only when we want to keep an issue
    open for visibility, but the real problem is somewhere else.

protocol
    This requires a change to the protocol.

unreleased
    This is set on bugs that have never been in a released stable version -
    i.e., a bug introduced in v0.14.44-rc.1 and fixed in v0.14.44-rc.2. They
    are excluded from the release notes for v0.14.44.

Milestones
----------

Each released version gets a milestone. Issues that are resolved and will be
released as that version get added to the milestone. The release notes are
based on the issues present in the milestone.

In addition to version specific milestones there are two generic ones:

Planned
    This issue is being worked on, or will soon be worked on, by someone in
    the core team. Expect action on it within the next few days, weeks or
    months.

Unplanned (Contributions Welcome)
    This issue is not being worked on by the core team, and we don't plan on
    doing so in the foreseeable future. We still consider it a valid issue
    and welcome contributions towards resolving it.

Issues lacking a milestone are currently undecided. In practice this is
similar to Unplanned in that probably no-one is working on it, but we are
still considering it and it may end up Planned or closed instead.

Assignee
--------

Users can be assigned to issues. We don't usually do so. Sometimes
someone assigns themself to an issue to indicate "I'm working on this"
to avoid others doing so too. It's not mandatory.

Locking
-------

We don't normally lock issues (prevent further discussion on them).
There are some exceptions though;

-  "Popular" issues that attract lots of "me too" and "+1" comments.
   These are noise and annoy people with useless notifications via mail
   and in the Github interface. Once the issue is clear and it suffers
   from this symptom I may lock it.

-  Contentious bikeshedding discussions. After two sides in a discussion
   have clarified their points, there is no point arguing endlessly
   about it. As above, this may get closed.

-  Duplicates. Once an issue has been identified as a duplicate of
   another issue, it may be locked to prevent further discussion there.
   The intention is to move the discussion to the other (referenced)
   issue, while someone just doing a search and jumping on the first
   match might otherwise resurrect discussion in the duplicate.
