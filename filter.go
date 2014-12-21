// based upon (among others) <https://gobyexample.com/line-filters>
// as an alternative, probably <http://godoc.org/io#PipeReader> could be used
//
// cat filter.go | go run filter.go
//
// pandoc readme.md -o readme.html --filter ./pandocfilter
//
// see http://nathanleclaire.com/blog/2014/07/19/demystifying-golangs-io-dot-reader-and-io-dot-writer-interfaces/
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// the default scanner reads line after line
	scanner := bufio.NewScanner(os.Stdin)

	// the original example does not use stdout explicitly, but it is possible
	stdout := bufio.NewWriter(os.Stdout)
	defer func() { stdout.Flush() }()

	// dump intermediate result
	d, err := os.Create("dump.json")
	defer d.Close()

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	dump := bufio.NewWriter(d)
	defer func() { dump.Flush() }()

	// use multi as a T that writes to stdout and dump
	// multi is an io.Writer and cannot be flushed
	multi := io.MultiWriter(stdout, dump)

	proc := empty{}

	for scanner.Scan() {
		fmt.Fprintln(multi, proc.process(scanner.Text()))
	}

	// Check for errors during `Scan`. End of file is
	// expected and not reported by `Scan` as an error.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

type processor interface {
	process(in string) string
}

type uc struct {
}

func (u uc) process(in string) string {
	return strings.ToUpper(in)
}

type empty struct {
}

func (e empty) process(in string) string {
	return in
}
