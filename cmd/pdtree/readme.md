---
author:
- ffel
date: december 2014
title: pandoc document
...

Readme
======

This is a [markdown](http://daringfireball.net/projects/markdown/)
document.

Pandoc
------

[Pandoc](http://johnmacfarlane.net/pandoc/) is a powerfull tool that
converts markdown documents in a number of other document types.

Extension with filters
----------------------

Pandoc can be extended with
[filters](http://johnmacfarlane.net/pandoc/scripting.html).

pdtree
------

`pdtree` is a basic pandoc filter written in [Go](http://golang.org/).

All it does is that it writes a tree to `stderr` which helps to write
your own filter with
[pandocfilter](https://github.com/ffel/pandocfilter).

    pandoc readme.md -o readme.json --filter ./pdtree 2> tree.txt

The above pandoc command converts `readme.md` into `readme.json` while
meanwhile the intermediate result is processed by `./pdtree`. The
command `pdtree` should have no side effect on the final result.
Therefore, the `readme.json` should be equal to `readme2.json` in the
following command (that is without processing):

    pandoc readme.md -o readme2.json
