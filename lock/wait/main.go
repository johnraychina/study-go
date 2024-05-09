package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1) // WaitGroup 在调用 Wait() 之后不能再调用 Add() 方法的。
	}()
	wg.Wait()
}
