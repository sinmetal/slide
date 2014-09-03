package main

import (
	"fmt"
	"os"
)

// io.Writerと同じinterface宣言
type Writer interface {
	Write(b []byte) (n int, err error)
}

func main() {
	var w Writer

	// os.Stdout implements Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}
