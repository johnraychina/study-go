package main

import "time"

func main() {
	// send data to a close channel

	ch := make(chan int, 1)
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			<-ticker.C
			ch <- time.Now().Second()
		}
	}()

	go func() {
		for {
			println(<-ch)
		}
	}()

	time.Sleep(time.Second * 5)
	// channel should be closed by sender
	// or-else the sender will be at risk of panic on closed channel
	close(ch)
	println("channel closed")
	time.Sleep(time.Second * 5)
}
