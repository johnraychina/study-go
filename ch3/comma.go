package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(comma("1234567"))

	//intToArray()

	x, err := strconv.Atoi("123") // x is an int
	if err != nil {
		fmt.Println(err)
	}
	y, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
	fmt.Println(x, y)

}

func intToArray() {
	x := 123
	y := fmt.Sprintf("test%d", x)
	fmt.Println(y, strconv.Itoa(x)) // "123 123"
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	//n : prefix length that to be inserted
	for n := len(s); n > 3; n -= 3 {
		s = s[:n-3] + "," + s[n-3:]
	}
	return s
}
