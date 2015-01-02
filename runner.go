package pandocfilter

import (
	"encoding/json"
	"log"
	"os"
)

// Run reads json from stdin, runs filter and writes to stdout
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
