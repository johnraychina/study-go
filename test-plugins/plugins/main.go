package main

import "fmt"

var V int

func F() {
	// go build -buildmode=plugin -o plugin.so main.go
	fmt.Printf("插件执行第%d次 \n ", V)
}
