package pandocfilter

import (
	"encoding/json"
	"log"
	"os"
)

// Run reads json from stdin, runs filter and writes to stdout
// todo: this is a tight coupling between processing in and out
// and one way of processing the json, Walk
// I can imagine other ways of processing for instance one
// method that searches the entire doc and modifies other
// parts based upon the result.
func Run(filter Filter) {
	decoder := json.NewDecoder(os.Stdin)

	var pandoc interface{}

	if err := decoder.Decode(&pandoc); err != nil {
		log.Println(err)
		return
	}

	out := Walk(filter, "", pandoc)

	encoder := json.NewEncoder(os.Stdout)

	if err := encoder.Encode(&out); err != nil {
		log.Println(err)
		return
	}
}
