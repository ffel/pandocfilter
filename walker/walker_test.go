package walker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

const data string = `[ { "unMeta" : {  } },
  [ { "c" : [ 1,
          [ "hallo",
            [  ],
            [  ]
          ],
          [ { "c" : "Hallo",
              "t" : "Str"
            } ]
        ],
      "t" : "Header"
    },
    { "c" : [ { "c" : "Hallo",
            "t" : "Str"
          },
          { "c" : [  ],
            "t" : "Space"
          },
          { "c" : "Wereld!",
            "t" : "Str"
          }
        ],
      "t" : "Para"
    }
  ]
]
`

func ExampleWalk() {
	dec := json.NewDecoder(strings.NewReader(data))

	var j interface{}
	if err := dec.Decode(&j); err != nil {
		log.Fatal(err)
	}

	// http://golang.org/pkg/bytes/#NewBuffer
	d := dumper{w: &bytes.Buffer{}}

	Walk(d, j, "root", 0)

	fmt.Println(d.w.String())

	// Output:
	// + list "root"
	//     + map "0"
	//         + map "unMeta"
	//     + list "1"
	//         + list "Header"
	//             + number "0" 1
	//             + list "1"
	//                 + string "0" "hallo"
	//                 + list "1"
	//                 + list "2"
	//             + list "2"
	//                 + string "Str" "Hallo"
	//         + list "Para"
	//             + string "Str" "Hallo"
	//             + list "Space"
	//             + string "Str" "Wereld!"
}

// it should be possible to sort of override an embedded method
// http://golang.org/doc/effective_go.html#embedding

// need to test possibilities to change content

func ExampleEncode() {
	// NewDecoder and NewEncoder are used for these make sense
	// in the stdin stdout filter approach
	dec := json.NewDecoder(strings.NewReader(data))

	var j interface{}
	if err := dec.Decode(&j); err != nil {
		log.Fatal(err)
	}

	buff := &bytes.Buffer{}
	enc := json.NewEncoder(buff)

	if err := enc.Encode(&j); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buff.String())

	// output:
	// [{"unMeta":{}},[{"c":[1,["hallo",[],[]],[{"c":"Hallo","t":"Str"}]],"t":"Header"},{"c":[{"c":"Hallo","t":"Str"},{"c":[],"t":"Space"},{"c":"Wereld!","t":"Str"}],"t":"Para"}]]
}

// func Test_change(t *testing.T) {
// 	t.Error("error ...")
// }

func ExampleReplace() {
	dec := json.NewDecoder(strings.NewReader(data))

	var j interface{}
	if err := dec.Decode(&j); err != nil {
		log.Fatal(err)
	}

	// todo: fix the situation with two buffers, is confusing

	d := replacer{}
	d.dumper.w = &bytes.Buffer{}

	Walk(d, j, "root", 0)

	buff := &bytes.Buffer{}
	enc := json.NewEncoder(buff)

	if err := enc.Encode(&j); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buff.String())

	// Output:
	// [{"unMeta":{}},[{"c":[1,["hallo",[],[]],[{"c":"Hallo","t":"Str"}]],"t":"Header"},{"c":[{"c":"Hallo","t":"Str"},{"c":[],"t":"Space"},{"c":"Europe!","t":"Str"}],"t":"Para"}]]
}

type replacer struct {
	dumper
}

/*
something terribly wrong here ...

1.  difficult or impossible to pass the pointer to the
    original string value

2.  iteratinng over collection items most likely passes
    a copy of the original string: no sense to pass a referece
    to a copy
*/

func (r replacer) String(key, value string, level int) {

	if key == "Str" && value == "Wereld!" {
		value = "Europe!"
	}

	r.dumper.String(key, value, level)
}

func ExampleDummy() {
	dec := json.NewDecoder(strings.NewReader(data))

	var j interface{}
	if err := dec.Decode(&j); err != nil {
		log.Fatal(err)
	}

	dummy(&j)

	fmt.Printf("%#v\n", j)

	// buff := &bytes.Buffer{}
	// enc := json.NewEncoder(buff)

	// if err := enc.Encode(&j); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(buff.String())

	// output:
	// test changes in json
}

// try to do a walk based upon pointers
// if this does not work either, my last option is to build a
// completely new json output object based upon the input object
//
// any clues in here (yes, make val an explicit interface{} type, but
// does not fix the issue) ?
// http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
func dummy(json *interface{}) {
	switch elem := (*json).(type) {
	case []interface{}:
		for _, v := range (*json).([]interface{}) {
			dummy(&v)
		}
	case map[string]interface{}:
		for _, v := range (*json).(map[string]interface{}) {
			dummy(&v)
		}
	case string:
		// val := "boo"
		// *json = val
		fmt.Println("change", *json)
		var val interface{}
		val = "boo"
		json = &val
	default:
		fmt.Printf("?? %T - %v\n", elem, elem)
	}
}
