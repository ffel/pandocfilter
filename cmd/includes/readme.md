CodeBlock includes filter
=========================

It parses CodeBlock includes written like this in the Markdown source:

    ~~~ {.go include="data/code.go" }
    any text here will be replaced...
    ~~~

so that the filtered output contains:

~~~
package test

import (
    "testing"
)

func TestHello(t *testing.T) {
    t.Fail()
}
~~~

You can try this by running:

~~~
pandoc --filter includes.exe -t html test.md -o test.html
~~~
