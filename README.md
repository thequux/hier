Hier: Hosting In Every Repo
===========================

Hier adds fossil-style project management integrated into a git repo,
so that it is transfered with every clone, push, and pull.

Usage
=====

Simply run `hier ui` in a git checkout! Hier can also serve from a
bare repo by passing the name of the repo as a third argument.

If you wish, for some reason, to prohibit Hier from serving a
particular repo (possibly while you are developing Hier, in case of an
errant libgit2 API call, or to protect a private repo from being
served by a multiplexing script), you can set `hier.verboten` to true
in the repo's git config.
