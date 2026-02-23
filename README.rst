Syncthing Docs
==============

This repo is the source behind https://docs.syncthing.net/.

Editing
-------

To edit the documentation you need a GitHub account. Once you have created one
and logged in, you can edit any page by navigating to the corresponding file and
clicking the edit (pen) icon. This will create a so called "fork" and a "pull
request", which will be approved by one of the existing documentation team
members. Once you have made a contribution or two, you can be added to the
documentation team and perform edits without requiring approval.

In the long run, learning to use Git_ and running Sphinx_ on your computer is
beneficial.

First steps to run it locally::

  git clone https://github.com/syncthing/docs.git
  cd docs
  pip install -r requirements.txt
  make html
  # open _build/html/index.html

You can also use our Docker image to build the documentation, which is the
same thing the build server does in the end:

  ./docker-build.sh html

Structure
---------

The documentation is divided into an index page (``index.rst``) and various
subsections. The sections are:

- Introductory information in ``intro``.
- Information for users in ``users``.
- Information for developers in ``dev``.

The documentation uses the `rst format`_. For a starting point check out the
`reStructuredText Primer`_.

.. _Git: https://www.git-scm.com/
.. _Sphinx: https://www.sphinx-doc.org/
.. _`rst format`: https://docutils.sourceforge.io/docs/ref/rst/restructuredtext.html
.. _`reStructuredText Primer`: https://www.sphinx-doc.org/en/master/usage/restructuredtext/basics.html

License
=======

All documentation and protocol specifications are licensed under the `Creative
Commons Attribution 4.0 International License
<https://creativecommons.org/licenses/by/4.0/>`__.
