package main

import (
	"github.com/ffel/pandocfilter"
	"io/ioutil"
	"log"
)

func main() {
	pandocfilter.Run(includes{})
}

type includes struct{}

// basic code block:
//
// {"c":[["",[],[]],"the code"],"t":"CodeBlock"}
//
// with include:
//
// {"c":[["",["go"],[["include","code.go"]]],"the code"],"t":"CodeBlock"}
//
func (inc includes) Value(key string, value interface{}) (bool, interface{}) {
	ok, t, c := pandocfilter.IsTypeContents(value)
	if ok && t == "CodeBlock" {

		// unwrap CodeBlock json structure
		var array = c.([]interface{})
		var meta = array[0].([]interface{})
		var namevals = meta[2].([]interface{})

		// we will put new content at this address
		var cb = &array[1]

		// now find and process includes (last wins)
		for _, nv := range namevals {
			v := nv.([]interface{})

			if v[0] == "include" {
				*cb = loadCodeBlock(v[1].(string))
			}
		}
		return true, nil
	}
	return true, value
}

func loadCodeBlock(name string) string {
	data, err := ioutil.ReadFile(name)
	checkFatal(err)
	return string(data)
}

func checkFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
