package main

import (
	"fmt"
	"time"
)

// channel 调用close后，会变成一个类似于slice的东西，有可能返回零值
// channel 应该配合go使用，直接在main执行路径上使用容易导致deadlock
// channel 没close，无数据时，从中取数据，程序会被阻塞，陷入等待
func main() {
	fetchEmptyOpenChannel()      // 直接报错 deadlock
	fetchEmptyClosedChannel()    // 直接取会返回零值
	selectEmptyClosedChannel()   // 返回零值
	selectEmptyOpenChannel()     // 没close，慢分支等数据，快分支会走default
	goSelectEmptyClosedChannel() // close, 返回零值
	goSelectEmptyOpenChannel()   // 没close，慢分支等数据，快分支会走default
	time.Sleep(1 * time.Second)
}

func fetchEmptyOpenChannel() {
	ch := make(chan int, 1)
	ch <- 1
	//close(ch)

	x1 := <-ch
	fmt.Println(x1)

	x2 := <-ch
	fmt.Println(x2)
}

func fetchEmptyClosedChannel() {
	ch := make(chan int, 1)
	ch <- 1
	close(ch)

	x1 := <-ch
	fmt.Println(x1)

	x2 := <-ch
	fmt.Println(x2)
}

//  select + empty closed channel，
func selectEmptyClosedChannel() {
	ch := make(chan int, 1)
	close(ch)
	//go func() {
	select {
	case i := <-ch:
		fmt.Println("i=", i)
	default:
		fmt.Println("nothing")
	}
	//}()
}

// select + empty open channel 不会返回零值
func selectEmptyOpenChannel() {
	ch := make(chan int, 1)
	//go func() {
	select {
	case i := <-ch:
		fmt.Println("i=", i)
	default:
		fmt.Println("nothing")
	}
	//}()
}

// go select + empty closed channel，
func goSelectEmptyClosedChannel() {
	ch := make(chan int, 1)
	close(ch)
	go func() {
		select {
		case i := <-ch:
			fmt.Println("i=", i)
		default:
			fmt.Println("nothing")
		}
	}()
}

// go select + empty open channel 不会返回零值
func goSelectEmptyOpenChannel() {
	ch := make(chan int, 1)
	go func() {
		select {
		case i := <-ch:
			fmt.Println("i=", i)
		default:
			fmt.Println("nothing")
		}
	}()
}
