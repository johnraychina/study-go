package main

import "fmt"

// 实现 embedded interface 与 普通实现，有什么区别？

type Splittable interface {
	Split() string
	Do() string
	Join() string
}

type Work struct {
	//Splittable
}

func (p Work) Split() string {
	return "1,2,3"
}
func (p Work) Do() string {
	return "do"
}
func (p Work) Join() string {
	return "6"
}

func main() {
	// embedded interface: 如果只实现部分接口方法
	// 如果只实现部分接口方法，调用没实现的方法时，会在运行时报错：panic: runtime error: invalid memory address or nil pointer dereference
	// todo 为什么编译器不直接报错呢？
	w := &Work{}
	a := w.Split()
	b := w.Do()
	c := w.Join()
	fmt.Println(a, b, c)
}
