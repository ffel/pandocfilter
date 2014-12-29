package pandocfilter

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

func walk(json interface{}) interface{} {
	switch /*elem := */ json.(type) {
	case []interface{}:
		result := make([]interface{}, 0, len(json.([]interface{})))
		for _, v := range json.([]interface{}) {
			result = append(result, walk(v))
		}
		return result
	case map[string]interface{}:
		typekey, tok := json.(map[string]interface{})["t"]
		contents, cok := json.(map[string]interface{})["c"]
		if tok && cok {
			// todo fix lost of c t object
			// maybe a variant of walk, walkct, that wraps the result in a
			return walkct(typekey.(string), contents)
		} else {
			result := make(map[string]interface{})
			for k, v := range json.(map[string]interface{}) {
				result[k] = walk(v)
			}
			return result
		}
	case string:
		if json.(string) == "Wereld!" {
			return "Europe!"
		} else {
			return json.(string)
		}
	default:
		return json
	}
}

func walkct(key string, json interface{}) interface{} {
	m := make(map[string]interface{})
	m["t"] = key
	m["c"] = walk(json)

	return m
}

func ExampleClone() {
	dec := json.NewDecoder(strings.NewReader(data))

	var j interface{}
	if err := dec.Decode(&j); err != nil {
		log.Fatal(err)
	}

	cl := walk(j)

	buff := &bytes.Buffer{}
	enc := json.NewEncoder(buff)

	if err := enc.Encode(&cl); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buff.String())

	// Output:
	// [{"unMeta":{}},[{"c":[1,["hallo",[],[]],[{"c":"Hallo","t":"Str"}]],"t":"Header"},{"c":[{"c":"Hallo","t":"Str"},{"c":[],"t":"Space"},{"c":"Europe!","t":"Str"}],"t":"Para"}]]
}
