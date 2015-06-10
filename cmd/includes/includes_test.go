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
	in_file := ReadFile("data/input.json")
	defer CloseFile(in_file)

	dec := json.NewDecoder(in_file)
	var pandoc interface{}
	if err := dec.Decode(&pandoc); err != nil {
		panic(err)
	}

	// apply filter
	output := pandocfilter.Walk(includes{}, "", pandoc)

	// check
	x_file := ReadFile("data/expected.json")
	defer CloseFile(x_file)

	dec = json.NewDecoder(x_file)
	var expected interface{}
	if err := dec.Decode(&expected); err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected: %v\nbut got: %v\n", expected, output)
	}
}

func ReadFile(name string) (file *os.File) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return file
}

func CloseFile(file *os.File) {
	if err := file.Close(); err != nil {
		panic(err)
	}
}
