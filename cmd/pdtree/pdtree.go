// pandoc filter which prints the pandoc tree to stderr
package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ffel/pandocfilter"
)

func main() {
	decoder := json.NewDecoder(os.Stdin)

	var pandoc interface{}

	if err := decoder.Decode(&pandoc); err != nil {
		log.Println(err)
		return
	}

	filter := NewTree()

	out := pandocfilter.Walk(filter, "", pandoc)

	// print to stderr to not interfere with json to stdout
	log.Printf("pandoc tree in go:\n%s\n", filter.String())

	encoder := json.NewEncoder(os.Stdout)

	if err := encoder.Encode(&out); err != nil {
		log.Println(err)
		return
	}
}
