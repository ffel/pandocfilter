Readme
======

Pandoc [filters](http://johnmacfarlane.net/pandoc/scripting.html) are a
convenient way to extend Pandoc.

The golang filter in this package dumps the received json data in a
compact way. It may be convenient to
[reformat](http://jsonformatter.curiousconcept.com/)[^1].

Unfortunately, go conversion with <https://mholt.github.io/json-to-go/>
is a bit disappointing so far.

The produced json is [not
documented](https://groups.google.com/forum/#!topic/pandoc-discuss/GhzVXBLEvng),
so we have to work it out for ourselves.

Json processing in golang is described in the web log [json and
go](http://blog.golang.org/json-and-go).

My motivation for writing a pandoc filter is that I want to have a "dsl"
for math. I'd like to support shorthands for

$$\frac{a\cdot b}{c}\cdot\frac{e+f}{d}$$

It would be nice to say:

``` {.txt}
a.b/c . e+f/d
```

[^1]: [This](http://jsonformat.com/) tends to produce a bit more compact
    results
