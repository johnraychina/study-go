package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	//s := []int{7, 2, 8, -9, 4, 0}
	//
	//c := make(chan int)

	//go sum(s[:len(s)/2], c)
	//go sum(s[len(s)/2:], c)
	//
	//x, y := <-c, <-c
	//fmt.Println(x, y)
	defer func() {
		println("defer")
		if p := recover(); p != nil {
			fmt.Println("panic:", p)
		}
	}()

	in := make(chan int, 100)
	for i := 0; i < 100; i++ {
		in <- i
	}
	out := make(chan int, 10)
	for i := 0; i < 10; i++ {
		out <- <-in
	}
	// 不加close：fatal error: all goroutines are asleep - deadlock!
	close(out)
	for o := range out {
		fmt.Println(o)
	}
}
