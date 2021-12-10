package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	// close before iteration
	//close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

	// for-range wait for a close event on the queue
	// but close(queue) wait for-range
	// which caused a deadlock!
	// fatal error: all goroutines are asleep - deadlock!
	close(queue)
}
