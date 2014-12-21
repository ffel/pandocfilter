Readme
======

Pandoc [filters](http://johnmacfarlane.net/pandoc/scripting.html) are a
convenient way to extend Pandoc.

The golang filter in this package dumps the received json data in a
compact way. It may be convenient to
[reformat](http://jsonformatter.curiousconcept.com/)[^1].

Unfortunately, go conversion with <https://mholt.github.io/json-to-go/>
is a bit disappointing so far.

[^1]: [This](http://jsonformat.com/) tends to produce a bit more compact
    results
