package typeassert

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestTypeAssert(t *testing.T) {

	var w io.Writer
	w = os.Stdout
	rw := w.(io.Writer)
	fmt.Print(rw)
}
