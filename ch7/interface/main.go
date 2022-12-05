package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var i int
	var w io.Writer

	var buf *bytes.Buffer
	w = os.Stdout
	w = new(bytes.Buffer)
	fmt.Printf("%T, %T, %T\n", i, w, buf)
}
