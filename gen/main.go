package main

import (
	"fmt"
	"johnraychina/study-go/gen/gen"
)

func main() {
	generateUint32Example()
	//var x int = 10
	//fmt.Println(x)
	//
	//var y int64 = int64(x)
	//fmt.Println(y)
	//
	//
	//
	//intContainer := &Container{}
	//intContainer.Put(7)
	//intContainer.Put(42)
	//val := intContainer.Get()
	//fmt.Printf("%T, %v \n", val, val)
	//
	//tp := reflect.TypeOf(val)
	//fmt.Printf("val:%v type:%v \n", val, tp)
	//
	//intVal := val.(int)
	//
	//fmt.Println(intVal)

}

// go:generate ./gen.sh ./template/container.temp.go gen uint32 container
func generateUint32Example() {
	var u uint32 = 42
	c := gen.NewUint32Container()
	c.Put(u)
	v := c.Get()
	fmt.Printf("generateExample: %d (%T)\n", v, v)
}

// go:generate ./gen.sh ./template/container.temp.go gen string container
func generateStringExample() {
	var s string = "Hello"
	c := gen.NewStringContainer()
	c.Put(s)
	v := c.Get()
	fmt.Printf("generateExample: %s (%T)\n", v, v)
}

//Container is a generic container, accepting anything.
type Container []interface{}

//Put adds an element to the container.
func (c *Container) Put(elem interface{}) {
	*c = append(*c, elem)
}

//Get gets an element from the container.
func (c *Container) Get() interface{} {
	elem := (*c)[0]
	*c = (*c)[1:]
	return elem
}
