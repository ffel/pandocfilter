package main

import (
	"regexp"

	"github.com/ffel/pandocfilter"
)

func main() {
	pandocfilter.Run(frac{})
}

var patt *regexp.Regexp

func init() {
	// http://godoc.org/regexp#example-Regexp-ReplaceAllString
	patt = regexp.MustCompile(`(\S+)\s*/\s*(\S+)`)
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

	return WrapTMath(t, patt.ReplaceAllString(c, `\frac{$1}{$2}`))
}

// should go to types.go

// TMath expects value to be a tc with t == "Math", it returns
// inner type "DisplayMath|InnerMath" and content math
func TMath(value interface{}) (string, string) {
	// dive into the data structure, outside in
	// the tree dump of dptree is of great help
	s1 := value.(map[string]interface{})
	s2 := s1["c"].([]interface{})
	s3 := s2[0].(map[string]interface{})
	t := s3["t"].(string)
	c := s2[1].(string)

	return t, c
}

func WrapTMath(typeMath, math string) interface{} {
	// inside out
	s1 := make(map[string]interface{})
	s1["c"] = make([]interface{}, 0)
	s1["t"] = typeMath
	s2 := make([]interface{}, 2)
	s2[0] = s1
	s2[1] = math
	s3 := make(map[string]interface{})
	s3["c"] = s2
	s3["t"] = "Math"

	return s3
}
