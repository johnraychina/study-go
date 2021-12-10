package main

import "time"

func main() {
	t := time.Now()
	t2 := t.UTC()
	println(t.Unix())
	println(t.String())
	println(t2.Unix())
	println(t2.String())

	zero := time.Unix(0, 0)
	println(zero.Unix())
	println(zero.String())

	bizDate := time.Unix(1638115200, 0)
	println(bizDate.String())
}
