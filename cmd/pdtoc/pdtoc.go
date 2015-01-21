// pandoc filter which prints a tree of section header and
// label to stderr

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ffel/pandocfilter"
)

func main() {
	pandocfilter.Run(pdtoc{})
}

type pdtoc struct{}

func (p pdtoc) Value(key string, value interface{}) (bool, interface{}) {

	ok, t, c := pandocfilter.IsTypeContents(value)

	if ok && t == "Header" {
		p.printHeader(c)
	}

	return true, value
}

func (p pdtoc) printHeader(json interface{}) {
	level, err := pandocfilter.GetNumber(json, "0")

	if err != nil {
		return
	}

	ref, err := pandocfilter.GetString(json, "1", "0")

	if err != nil {
		return
	}

	label, err := pandocfilter.GetObject(json, "2")

	if err != nil {
		return
	}

	col := &collector{}

	pandocfilter.Walk(col, "", label)

	fmt.Fprintf(os.Stderr, "%s- %s(#%s)\n",
		strings.Repeat("  ", int(level-1)), col.value, ref)
}

// collector walks the header c and collects the Str
type collector struct {
	value string
}

func (coll *collector) Value(key string, value interface{}) (bool, interface{}) {
	ok, t, c := pandocfilter.IsTypeContents(value)

	if ok && t == "Str" {
		coll.value += c.(string) + " "
	}

	return true, value
}
