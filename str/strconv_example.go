package main

import "strconv"

func main() {
	i, err := strconv.ParseInt("0001", 10, 32)
	println(i)
	if err != nil {
		print(err)
	}
}
