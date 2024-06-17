package main

import (
	"fmt"
)

func main() {
	bitLen := uint16(8)

	maskSequence := int64(1<<bitLen - 1)
	fmt.Printf("%b \n", bitLen-1)
	fmt.Printf("%b \n", 1<<bitLen-1)
	fmt.Printf("%b \n", maskSequence)

	fmt.Printf("%x \n", -1)

}
