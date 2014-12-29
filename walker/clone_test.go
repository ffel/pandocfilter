package walker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func clone(json interface{}) interface{} {
	switch /*elem := */ json.(type) {
	case []interface{}:
		result := make([]interface{}, 0, len(json.([]interface{})))
		for _, v := range json.([]interface{}) {
			result = append(result, clone(v))
		}
		return result
	case map[string]interface{}:
		result := make(map[string]interface{})
		for k, v := range json.(map[string]interface{}) {
			result[k] = clone(v)
		}
		return result
	case string:
		if json.(string) == "Wereld!" {
			return "Europe!"
		} else {
			return json.(string)
		}
	default:
		//fmt.Printf("??? %T - %v\n", elem, elem)
		return json
	}
}

func ExampleClone() {
	dec := json.NewDecoder(strings.NewReader(data))

	var j interface{}
	if err := dec.Decode(&j); err != nil {
		log.Fatal(err)
	}

	cl := clone(j)

	// fmt.Printf("%#v\n", cl)

	buff := &bytes.Buffer{}
	enc := json.NewEncoder(buff)

	if err := enc.Encode(&cl); err != nil {
		log.Fatal(err)
	}

	fmt.Println(buff.String())

	// Output:
	// [{"unMeta":{}},[{"c":[1,["hallo",[],[]],[{"c":"Hallo","t":"Str"}]],"t":"Header"},{"c":[{"c":"Hallo","t":"Str"},{"c":[],"t":"Space"},{"c":"Europe!","t":"Str"}],"t":"Para"}]]
}
