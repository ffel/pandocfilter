package w2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

/*
json produced by pandoc based upon the following

  ---
  author:
  - ffel
  date: december 2014
  title: fix meta header
  ...

  Hello
  =====

  Ordinary text
*/
const meta = `[ { "unMeta" : { "author" : { "c" : [ { "c" : [ { "c" : "ffel",
                        "t" : "Str"
                      } ],
                  "t" : "MetaInlines"
                } ],
            "t" : "MetaList"
          },
        "date" : { "c" : [ { "c" : "december",
                  "t" : "Str"
                },
                { "c" : [  ],
                  "t" : "Space"
                },
                { "c" : "2014",
                  "t" : "Str"
                }
              ],
            "t" : "MetaInlines"
          },
        "title" : { "c" : [ { "c" : "fix",
                  "t" : "Str"
                },
                { "c" : [  ],
                  "t" : "Space"
                },
                { "c" : "meta",
                  "t" : "Str"
                },
                { "c" : [  ],
                  "t" : "Space"
                },
                { "c" : "header",
                  "t" : "Str"
                }
              ],
            "t" : "MetaInlines"
          }
      } },
  [ { "c" : [ 1,
          [ "hello",
            [  ],
            [  ]
          ],
          [ { "c" : "Hello",
              "t" : "Str"
            } ]
        ],
      "t" : "Header"
    },
    { "c" : [ { "c" : "Ordinary",
            "t" : "Str"
          },
          { "c" : [  ],
            "t" : "Space"
          },
          { "c" : "text",
            "t" : "Str"
          }
        ],
      "t" : "Para"
    }
  ]
]`

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
	json := decode(meta)

	f := Duplicator{}

	object := Walk(f, "", json)

	fmt.Println(encode(object))

	// Output:
	// [{"unMeta":{"author":{"c":[{"c":[{"c":"ffel","t":"Str"}],"t":"MetaInlines"}],"t":"MetaList"},"date":{"c":[{"c":"december","t":"Str"},{"c":[],"t":"Space"},{"c":"2014","t":"Str"}],"t":"MetaInlines"},"title":{"c":[{"c":"fix","t":"Str"},{"c":[],"t":"Space"},{"c":"meta","t":"Str"},{"c":[],"t":"Space"},{"c":"header","t":"Str"}],"t":"MetaInlines"}}},[{"c":[1,["hello",[],[]],[{"c":"Hello","t":"Str"}]],"t":"Header"},{"c":[{"c":"Ordinary","t":"Str"},{"c":[],"t":"Space"},{"c":"text","t":"Str"}],"t":"Para"}]]
}
