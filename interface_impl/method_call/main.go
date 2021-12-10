package main

// call with pointer vs call with object

type Cat struct {
	Name string
}

type Duck interface {
	Quack()
}

func (c Cat) Quack() { // 使用结构体实现接口
	println(c.Name)
}

//func (c *Cat) Quack() {  // 使用结构体指针实现接口
//	println(c.Name)
//}

func main() {
	var s Duck = Cat{Name: "meow"} // 使用结构体初始化变量
	println(&s)
	s.Quack()

	//var p Duck = &Cat{} // 使用结构体指针初始化变量
	//println(p)
	//p.Quack()
}
