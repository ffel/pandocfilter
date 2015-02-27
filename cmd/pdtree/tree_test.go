package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/ffel/pandocfilter"
)

/*
meta is formatted json produced by pandoc based upon the following

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

func ExampleTree() {
	json := decode(meta)

	f := NewTree()

	_ = pandocfilter.Walk(f, "", json)

	fmt.Println(f.String())

	// Output:
	// + "": list:
	//     + "0": map:
	//         + "unMeta": map:
	//             + "author": map:
	//                 + "c": list:
	//                     + "0": map:
	//                         + "c": list:
	//                             + "0" - Str: "ffel"
	//                         + "t": value: MetaInlines
	//                 + "t": value: MetaList
	//             + "date": map:
	//                 + "c": list:
	//                     + "0" - Str: "december"
	//                     + "1" - Space
	//                     + "2" - Str: "2014"
	//                 + "t": value: MetaInlines
	//             + "title": map:
	//                 + "c": list:
	//                     + "0" - Str: "fix"
	//                     + "1" - Space
	//                     + "2" - Str: "meta"
	//                     + "3" - Space
	//                     + "4" - Str: "header"
	//                 + "t": value: MetaInlines
	//     + "1": list:
	//         + "0": map:
	//             + "c": list:
	//                 + "0": value: 1
	//                 + "1": list:
	//                     + "0": value: hello
	//                     + "1": list:
	//                     + "2": list:
	//                 + "2": list:
	//                     + "0" - Str: "Hello"
	//             + "t": value: Header
	//         + "1": map:
	//             + "c": list:
	//                 + "0" - Str: "Ordinary"
	//                 + "1" - Space
	//                 + "2" - Str: "text"
	//             + "t": value: Para
}
