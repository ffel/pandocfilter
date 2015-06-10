package main

import (
	"encoding/json"
	"github.com/ffel/pandocfilter"
	"os"
	"reflect"
	"testing"
)

func TestJson(t *testing.T) {
	// decode input
	in_file := openFile("data/input.json")
	defer closeFile(in_file)

	var pandoc interface{}
	dec := json.NewDecoder(in_file)
	checkFatal(dec.Decode(&pandoc))

	// apply filter
	output := pandocfilter.Walk(includes{}, "", pandoc)

	// check
	x_file := openFile("data/expected.json")
	defer closeFile(x_file)

	var expected interface{}
	dec = json.NewDecoder(x_file)
	checkFatal(dec.Decode(&expected))

	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected: %v\nbut got: %v\n", expected, output)
	}
}

func openFile(name string) (file *os.File) {
	file, err := os.Open(name)
	checkFatal(err)
	return file
}

func closeFile(file *os.File) {
	err := file.Close()
	checkFatal(err)
}
