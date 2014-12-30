package pandocfilter

import (
	"bytes"
	"fmt"
	"strings"
)

// NewTree initiates a Tree object and returns its pointer
func NewTree() *Tree {
	return &Tree{make([]int, 0, 10), &bytes.Buffer{}}
}

// Tree prints a tree view of pandoc json, it also duplicates
// the pandoc json as does Duplicator
type Tree struct {
	levels []int
	buff   *bytes.Buffer
}

func (t *Tree) List(key string, json []interface{}) (bool, interface{}) {
	fmt.Fprintf(t.buff, "%s+ list %q\n", t.indent(), key)

	t.update(len(json), true)

	return true, json
}

func (t *Tree) Map(key string, json map[string]interface{}) (bool, interface{}) {
	fmt.Fprintf(t.buff, "%s+ map %q\n", t.indent(), key)

	t.update(len(json), true)

	return true, json
}

func (t *Tree) Text(key string, value string) interface{} {
	fmt.Fprintf(t.buff, "%s+ text %q: %q\n", t.indent(), key, value)

	t.update(0, false)

	return value
}

func (t *Tree) Number(key string, value float64) interface{} {
	fmt.Fprintf(t.buff, "%s+ number %q: %v\n", t.indent(), key, value)

	t.update(0, false)

	return value
}

func (t *Tree) Bool(key string, value bool) interface{} {
	fmt.Fprintf(t.buff, "%s+ list %q: %t\n", t.indent(), key, value)

	t.update(0, false)

	return value
}

func (t *Tree) String() string {
	return t.buff.String()
}

// update updates the indent level
func (t *Tree) update(l int, add bool) {

	if add {
		t.levels = append(t.levels, l)
	}

	for i := len(t.levels) - 1; i >= 0; i-- {
		if t.levels[i] == 0 {
			// pop last element
			t.levels = t.levels[:len(t.levels)-1]
		} else {
			t.levels[i] = t.levels[i] - 1
			break
		}
	}
}

func (t *Tree) indent() string {
	return strings.Repeat("    ", len(t.levels))
}
