package main

import (
	"bytes"
	"io"
)

const debug = false

func main() {
	//var buf *bytes.Buffer
	//if debug {
	//	buf = new(bytes.Buffer) // enable collection of output
	//}
	//f(buf) // NOTE: subtly incorrect!
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // OK
	if debug {
		// ...use buf...
	}
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
