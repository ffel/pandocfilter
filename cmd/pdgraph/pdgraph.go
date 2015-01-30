// pandoc filter that removes block quotes from the output

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ffel/pandocfilter"
)

func main() {
	g := &graph{make([]string, 0, 100), make([]string, 0, 100), ""}

	pandocfilter.Run(g)

	fmt.Fprint(os.Stderr, g.tgf())
}

type graph struct {
	nodes, edges []string
	current      string
}

func (r *graph) Value(key string, value interface{}) (bool, interface{}) {
	ok, t, c := pandocfilter.IsTypeContents(value)

	if ok && t == "Header" {
		ref, lab := r.header(c)
		// fmt.Fprintf(os.Stderr, "h %s %s\n", ref, lab)
		r.current = ref
		r.nodes = append(r.nodes, ref+" "+lab)
	} else if ok && t == "Link" {
		ref, lab := r.link(c)

		if strings.IndexRune(ref, '#') == 0 {
			// fmt.Fprintf(os.Stderr, "l %s %s\n", strings.TrimPrefix(ref, "#"), lab)
			r.edges = append(r.edges,
				r.current+" "+strings.TrimPrefix(ref, "#")+" "+lab)
		}
	}

	return true, value
}

func (r *graph) tgf() string {
	txt := ""
	for _, n := range r.nodes {
		txt += n + "\n"
	}
	txt += "#\n"
	for _, e := range r.edges {
		txt += e + "\n"
	}

	return txt
}

func (r *graph) header(json interface{}) (string, string) {
	ref, err := pandocfilter.GetString(json, "1", "0")

	if err != nil {
		return "-", "-"
	}

	label, err := pandocfilter.GetObject(json, "2")

	if err != nil {
		return "-", "-"
	}

	col := &collector{}

	pandocfilter.Walk(col, "", label)

	return ref, col.value
}

func (r *graph) link(json interface{}) (string, string) {
	ref, err := pandocfilter.GetString(json, "1", "0")

	if err != nil {
		return "-", "-"
	}

	label, err := pandocfilter.GetObject(json, "0")

	if err != nil {
		return "-", "-"
	}

	col := &collector{}

	pandocfilter.Walk(col, "", label)

	return ref, col.value
}

// collector walks the header c and collects the Str
// copied from pdtoc
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
