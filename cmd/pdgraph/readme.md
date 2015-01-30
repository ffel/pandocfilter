Pandoc Graphs
=============

Links in Pandoc {#sec1}
---------------

1.  Section headers have an inplicit or explicit label in pandoc

2.  These labels enable internal links for example to the [next
    section](#graph-output) possible.

Graph Output
------------

Based upon the data described in the [previous section](#sec1) we can
extract *graphs*, and dump these in [trivial graph
format](http://en.wikipedia.org/wiki/Trivial_Graph_Format) or `tgf` for
short.

This readme should produce the following graph:

``` {.tgf}
sec1 Links in Pandoc
graph-output Graph Output
#
sec1 graph-output next section
graph-output sec1 previous section
```

Test
----

The script can be tested

    go build
    pandoc readme.md -o readme.json --filter ./pdgraph 2> graph.tgf

File `graph.tgf` is a plain text file and can be opened with
[yed](http://www.yworks.com/en/products/yfiles/yed/).
