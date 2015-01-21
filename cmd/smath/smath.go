package main

import (
	"regexp"

	"github.com/ffel/pandocfilter"
)

func main() {
	pandocfilter.Run(frac{})
}

var fracPatt *regexp.Regexp
var dotPatt *regexp.Regexp

func init() {
	// http://godoc.org/regexp#example-Regexp-ReplaceAllString
	fracPatt = regexp.MustCompile(`(\S+)\s*/\s*(\S+)`)
	dotPatt = regexp.MustCompile(`\.`)
}

type frac struct{}

func (f frac) Value(key string, value interface{}) (bool, interface{}) {

	ok, t, _ := pandocfilter.IsTypeContents(value)

	if ok && t == "Math" {
		return false, f.resolve(value)
	}

	return true, value
}

func (f frac) resolve(value interface{}) interface{} {
	t, c := TMath(value)

	c = fracPatt.ReplaceAllString(c, `\frac{$1}{$2}`)
	c = dotPatt.ReplaceAllString(c, `\cdot{}`)

	return WrapTMath(t, c)
}

// should go to types.go

// TMath expects value to be a tc with t == "Math", it returns
// inner type "DisplayMath|InnerMath" and content math
func TMath(value interface{}) (string, string) {
	// dive into the data structure, outside in
	// the tree dump of dptree is of great help
	//
	// have a look at
	// https://github.com/mitchellh/mapstructure
	// this could work if we're able to define the math part
	// in a proper go struct.  This appears to be difficult
	// or even impossible as the math data structure uses interface{}
	// collections to its full extent:
	// the "c" list has a map and a string value
	// the latter map has a list and a string value
	// maybe we could help things a little by adding a
	//
	// GetString("c", "0", "c")
	//
	// all the info is essentially there "c" and "t" refer to map fields
	// every thing else is to be converted to an int which refers to a
	// slice
	//
	// to convert back to an interface{} which is understood by pandoc
	// is completely else. The scenario here is that we change a
	// value in math field.  It must be possible to use a copy of this
	// math field to do a
	//
	// PutString("c", "0", "c", "newvalue")
	//
	// The other scenario is that one element is replaced by something
	// completely else, for instance, a code block for an image
	s1 := value.(map[string]interface{})
	s2 := s1["c"].([]interface{})
	s3 := s2[0].(map[string]interface{})
	t := s3["t"].(string)
	c := s2[1].(string)

	return t, c
}

// type jmap map[string]interface{}
// type jslice []interface{}

func WrapTMath(typeMath, math string) interface{} {
	// explicit struct is possible, even simpler if we use
	// jmap and jslice aliases for map[string]interface{} and
	// []interface{}
	m := map[string]interface{}{
		"c": []interface{}{
			map[string]interface{}{
				"c": []interface{}{},
				"t": typeMath,
			},
			math,
		},
		"t": "Math",
	}

	return m
}
