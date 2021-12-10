package main

import "sync"

func main() {
	list := []string{"Alice", "Bob", "Cathy"}
	var group sync.WaitGroup
	//group.Add(len(list))
	for _, p := range list {
		go func(string) {
			defer group.Done()
			println(p)
		}(p)
		group.Wait() // blocked: fatal error: all goroutines are asleep - deadlock!
	}
	// group.Wait() should be here
}
