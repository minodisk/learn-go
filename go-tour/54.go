package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var r Reader
	var w Writer

	// os.Stdout implements Writer
	r = os.Stdin
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}
