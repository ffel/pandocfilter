package pandocfilter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// this is similar content as passed by pandoc to the filter
// pandoc hello.md -o hello.html --filter ./filter
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

// decode accepts data as a reader to convert into a data object
// much the same way as used in the filter implementation
func decode(in string) interface{} {
	dec := json.NewDecoder(strings.NewReader(in))

	var json interface{}

	if err := dec.Decode(&json); err != nil {
		log.Fatal(err)
	}

	return json
}

// encode accepts a data object and convert it to json.
// it uses a writer internally
// much the same way as used in the filter implementation
func encode(js interface{}) string {
	buff := &bytes.Buffer{}
	enc := json.NewEncoder(buff)

	if err := enc.Encode(&js); err != nil {
		log.Fatal(err)
	}

	return buff.String()
}

func ExampleDuplicator() {
	json := decode(data)

	f := Duplicator{}

	object := Walk(f, "", json)

	fmt.Println(encode(object))

	// Output:
	// [{"unMeta":{}},[{"c":[1,["hallo",[],[]],[{"c":"Hallo","t":"Str"}]],"t":"Header"},{"c":[{"c":"Hallo","t":"Str"},{"c":[],"t":"Space"},{"c":"Wereld!","t":"Str"}],"t":"Para"}]]
}
