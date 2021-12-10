package main

import (
	"fmt"
)

// 练习 9.4: 创建一个流水线程序，支持用channel连接任意数量的goroutine，在跑爆内存之前，可以创建多少流水线阶段？一个变量通过整个流水线需要用多久？（这个练习题翻译不是很确定）
func main() {
	in := make(chan int64)
	out := make(chan int64)

	in <- 1
	for i := 0; i < 2; i++ {
		go makePipeline(in, out)
		in = make(chan int64)
		out = make(chan int64)
	}

	value := <-out
	fmt.Println(value)
}

func makePipeline(in chan int64, out chan int64) {
	value := <-in
	out <- (value + 1)
}
