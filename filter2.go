// json readers and writers are provided for in std lib ...

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

	process(json, "", "")
	// enc := json.NewEncoder(os.Stdout)

	// next steps:
	//
	// -  parse based upon "t" values into proper structs
	// -  write to json again
	//
	// wanneer je nette structs gebruikt, kan je in veel gevallen
	// gebruik maken van de nul waarde van slices om weer naar
	// json te gaan:  een space heeft een lege slice van strings, en
	// zo zijn er meer!
}

func process(json interface{}, indent, key string) {
	switch elem := json.(type) {
	case string, float64, bool:
		fmt.Printf("%s%s %T: %v\n", indent, key, elem, json)
	case []interface{}:
		fmt.Printf("%s%s slice:\n", indent, key)
		// use type assertion to tell compiler that iterating is possible
		for i, v := range json.([]interface{}) {
			process(v, indent+"  ", fmt.Sprintf("%v:", i))
		}
	case map[string]interface{}:
		fmt.Printf("%s%s map:\n", indent, key)
		// use type assertion to tell compiler that iterating is possible
		for k, v := range json.(map[string]interface{}) {
			process(v, indent+"  ", k+":")
		}
	default:
		fmt.Printf("don't know how to handle %T %v\n", elem, elem)
	}
}
