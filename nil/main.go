package main

import "fmt"

func main() {
	//var f1 Fish = nil //Cannot use 'nil' as the type Fish
	//fmt.Println(f1)

	var f2 *Fish = nil
	NilInterface(f2)

}

type ObjectI interface {
	Id() int
	Name() string
}
type Fish struct {
	ObjectI
	id   int
	name string
}

func NilInterface(val ObjectI) {
	if val == nil {
		fmt.Println("nil val")
	} else {
		// 这里会panic
		fmt.Println(val.Name())
	}
}

func NilSlice() {
}

func EmptySlice() {
}
