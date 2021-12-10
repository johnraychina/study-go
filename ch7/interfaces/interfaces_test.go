package interfaces

import (
	"bytes"
	"io"
	"os"
	"time"
)

func main() {
	var w io.Writer
	w = os.Stdout         // OK: *os.File has Write method
	w = new(bytes.Buffer) // OK: *bytes.Buffer has Write method
	//w = time.Second         // compile error: time.Duration lacks Write method

	var rwc io.ReadWriteCloser
	rwc = os.Stdout // OK: *os.File has Read, Write, Close methods
	//rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method
}
