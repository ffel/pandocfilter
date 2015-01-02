package w2

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

func (t *Tree) Value(key string, value interface{}) (bool, interface{}) {
	list, isList := value.([]interface{})

	if isList {
		fmt.Fprintf(t.buff, "%s+ %q: list:\n", t.indent(), key)

		t.update(len(list), true)

		return true, nil
	}

	// a cstring is a special type of map
	isStr, tstring, cstring := CString(value)

	if isStr {
		fmt.Fprintf(t.buff, "%s+ %q: string %q: %v\n", t.indent(), key, tstring, cstring)
		t.update(0, false)

		// prevent walker from going into this tc object
		return false, value
	}

	// check for other type of map
	set, isSet := value.(map[string]interface{})

	if isSet {
		fmt.Fprintf(t.buff, "%s+ %q: map:\n", t.indent(), key)

		t.update(len(set), true)

		return true, nil
	}

	// value is not identifies as something special
	fmt.Fprintf(t.buff, "%s+ %q: value: %v\n", t.indent(), key, value)

	t.update(0, false)

	return true, value
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
