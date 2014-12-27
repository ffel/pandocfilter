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
