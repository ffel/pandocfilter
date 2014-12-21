// https://gobyexample.com/line-filters
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	// Wrapping the unbuffered `os.Stdin` with a buffered
	// scanner gives us a convenient `Scan` method that
	// advances the scanner to the next token; which is
	// the next line in the default scanner.
	scanner := bufio.NewScanner(os.Stdin)

	// the original example does not use stdout explicitly, but it is possible
	stdout := bufio.NewWriter(os.Stdout)

	defer func() { stdout.Flush() }()

	// as an alternative, probably <http://godoc.org/io#PipeReader> could be used

	// this may be very interesting
	// http://nathanleclaire.com/blog/2014/07/19/demystifying-golangs-io-dot-reader-and-io-dot-writer-interfaces/

	// try to see intermediate result
	d, _ := os.Create("dump.json")

	defer d.Close()

	dump := bufio.NewWriter(d)

	defer func() { dump.Flush() }()

	multi := io.MultiWriter(stdout, dump)

	for scanner.Scan() {
		// `Text` returns the current token, here the next line,
		// from the input.
		ucl := strings.ToUpper(scanner.Text())

		// Write out the uppercased line.
		//fmt.Println(ucl)

		fmt.Fprintln(multi, ucl)
	}

	// Check for errors during `Scan`. End of file is
	// expected and not reported by `Scan` as an error.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// func ExampleWriter() {
// 	w := bufio.NewWriter(os.Stdout)
// 	fmt.Fprint(w, "Hello, ")
// 	fmt.Fprint(w, "world!")
// 	w.Flush() // Don't forget to flush!
// 	// Output: Hello, world!
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"

// 	"github.com/gorilla/websocket"
// )

// func main() {

// }

// // https://gobyexample.com/line-filters

// // http://golang-examples.tumblr.com/post/41863665491/read-a-line-from-stdin
// func code() {
// 	bio := bufio.NewReader(os.Stdin)
// 	line, hasMoreInLine, err := bio.ReadLine()
// }

// // http://golang.org/src/bufio/example_test.go

// // http://talks.golang.org/2012/chat.slide#39

// type socket struct {
// 	io.Reader
// 	io.Writer
// 	done chan bool
// }

// var chain = NewChain(2) // 2-word prefixes

// func socketHandler(ws *websocket.Conn) {
// 	r, w := io.Pipe()
// 	go func() {
// 		_, err := io.Copy(io.MultiWriter(w, chain), ws)
// 		w.CloseWithError(err)
// 	}()
// 	s := socket{r, ws, make(chan bool)}
// 	go match(s)
// 	<-s.done
// }
