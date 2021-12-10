package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)

	fmt.Printf("%T \n", 3)

	v := reflect.ValueOf(3)
	fmt.Println(v.String())
	x := v.Interface()
	fmt.Printf("%d\n", x)
}
