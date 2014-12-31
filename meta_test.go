package pandocfilter

import "fmt"

// The way pandoc pack the meta header in json makes me to believe
// that it is not safe to collapse the ct objects into one element.
//
// Now we know that ct collapsing does not work for the meta header
// it is not safe to assume that it does not pose any problems
// in the ordinary text
//
// THE MAIN DISADVANTAGE IS THAT OPERATIONS ON STRINGS, NUMBERS,
// AND BOOLS DON'T GET THE CORRESPONDING TYPE PASSED
//
// het moet mogelijk zijn om reflectie op c te doen wanneer je
// op de map stuit en voor bool, string om c en t mee te geven.
func ExampleMeta() {
	json := decode(meta)

	f := NewTree()

	_ = Walk(f, "", json)

	fmt.Println(f.String())
	// output:
	// FIX
	// + list ""
	//     + map "0"
	//         + map "unMeta"
	//             + map "author"
	//                 + text "t": "MetaList"
	//                 + list "c"
	//                     + map "0"
	//                         + list "c"
	//                             + map "0"
	//                                 + text "c": "ffel"
	//                                 + text "t": "Str"
	//                         + text "t": "MetaInlines"
	//             + map "date"
	//                 + list "c"
	//                     + map "0"
	//                         + text "c": "december"
	//                         + text "t": "Str"
	//                     + map "1"
	//                         + list "c"
	//                         + text "t": "Space"
	//                     + map "2"
	//                         + text "c": "2014"
	//                         + text "t": "Str"
	//                 + text "t": "MetaInlines"
	//             + map "title"
	//                 + list "c"
	//                     + map "0"
	//                         + text "c": "fix"
	//                         + text "t": "Str"
	//                     + map "1"
	//                         + text "t": "Space"
	//                         + list "c"
	//                     + map "2"
	//                         + text "c": "meta"
	//                         + text "t": "Str"
	//                     + map "3"
	//                         + text "t": "Space"
	//                         + list "c"
	//                     + map "4"
	//                         + text "c": "header"
	//                         + text "t": "Str"
	//                 + text "t": "MetaInlines"
	//     + list "1"
	//         + map "0"
	//             + list "c"
	//                 + number "0": 1
	//                 + list "1"
	//                     + text "0": "hello"
	//                     + list "1"
	//                     + list "2"
	//                 + list "2"
	//                     + map "0"
	//                         + text "c": "Hello"
	//                         + text "t": "Str"
	//             + text "t": "Header"
	//         + map "1"
	//             + list "c"
	//                 + map "0"
	//                     + text "c": "Ordinary"
	//                     + text "t": "Str"
	//                 + map "1"
	//                     + list "c"
	//                     + text "t": "Space"
	//                 + map "2"
	//                     + text "c": "text"
	//                     + text "t": "Str"
	//             + text "t": "Para"
}

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
