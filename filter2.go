// a much simpler piece of code
//
// pandoc readme.md -o readme.html --filter ./pandocfilter
//
// -- pandocfilter$ pandoc readme.md -o readme.html --filter ./filter2
// 2014/12/22 21:58:34 json: cannot unmarshal array into Go value of type map[string]interface {}
// pandoc: not enough input
//
// it does not run the code recursively

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// go run filter2.go < dump.json > filter2.out

func main() {
	dec := json.NewDecoder(os.Stdin)
	var json interface{}
	if err := dec.Decode(&json); err != nil {
		log.Println(err)
		return
	}

	process(json, "")
	// enc := json.NewEncoder(os.Stdout)
	// for {
	// 	var v map[string]interface{}
	// 	if err := dec.Decode(&v); err != nil {
	// 		log.Println(err)
	// 		return
	// 	}
	// 	// for k := range v {
	// 	// 	if k != "Name" {
	// 	// 		delete(v, k)
	// 	// 	}
	// 	// }
	// 	if err := enc.Encode(&v); err != nil {
	// 		log.Println(err)
	// 	}
	// }
}

func process(json interface{}, prefix string) {
	// for k, v := range m {
	// 	switch vv := v.(type) {
	// 	case string:
	// 		fmt.Println(k, "is string", vv)
	// 	case int:
	// 		fmt.Println(k, "is int", vv)
	// 	case []interface{}:
	// 		fmt.Println(k, "is an array:")
	// 		for i, u := range vv {
	// 			fmt.Println(i, u)
	// 		}
	// 	default:
	// 		fmt.Println(k, "is of a type I don't know how to handle")
	// 	}
	// }
	switch elem := json.(type) {
	case string:
		fmt.Printf("%s* string: %v\n", prefix, json)
	case int:
		fmt.Printf("%s* int: %v\n", prefix, json)
	case []interface{}:
		fmt.Printf("%s* slice:\n", prefix)
		// use type assertion to tell compiler that iterating is possible
		for _, v := range json.([]interface{}) {
			process(v, prefix+"  ")
		}
	case map[string]interface{}:
		fmt.Printf("%s* map:\n", prefix)
		// use type assertion to tell compiler that iterating is possible
		for _, v := range json.(map[string]interface{}) {
			process(v, prefix+"  ")
		}
	default:
		fmt.Printf("don't know how to handle %v\n", elem)
	}
}
