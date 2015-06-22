package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/ffel/pandocfilter"
)

// NewTree initiates a Tree object and returns its pointer
func NewTree() *Tree {
	return &Tree{&bytes.Buffer{}}
}

// Tree prints a tree view of pandoc json, it also duplicates
// the pandoc json as does Duplicator
type Tree struct {
	buff *bytes.Buffer
}

func (t *Tree) Value(level int, key string, value interface{}) (bool, interface{}) {
	_, isList := value.([]interface{})

	if isList {
		fmt.Fprintf(t.buff, "%s+ %q: list:\n", t.indent(level), key)

		return true, nil
	}

	// a cstring is a special type of map
	isTC, tval, cval := pandocfilter.IsTypeContents(value)

	// don't do anything special with known collection types
	// as the returned values are used again by pandoc

	if isTC {
		switch tval {
		case pandocfilter.Space:
			fmt.Fprintf(t.buff, "%s+ %q - %s\n", t.indent(level), key, tval)
			return false, value
		case pandocfilter.Str:
			fmt.Fprintf(t.buff, "%s+ %q - %s: %q\n", t.indent(level), key, tval, cval.(string))
			return false, value
		}
	}

	_, isSet := value.(map[string]interface{})

	if isSet {
		fmt.Fprintf(t.buff, "%s+ %q: map:\n", t.indent(level), key)

		return true, nil
	}

	// value is not identifies as something special
	fmt.Fprintf(t.buff, "%s+ %q: value: %v\n", t.indent(level), key, value)

	return true, value
}

func (t *Tree) String() string {
	return t.buff.String()
}

func (t *Tree) indent(level int) string {
	return strings.Repeat("    ", level)
}
