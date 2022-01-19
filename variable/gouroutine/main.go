package main

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"sync"
)

func main() {

	defer func() {
		if p := recover(); p != nil {
			fmt.Println("UpdateTaskExecution 20分发失败")
			panic(p)
		}
	}()

	var batchError error
	var wg sync.WaitGroup
	ctx := context.Background()
	wg.Add(40)
	for i := 0; i < 40; i++ {
		WithGo(ctx, func(ctx context.Context) {
			defer func() {
				wg.Done()
				if p := recover(); p != nil {
					batchError = errors.New(fmt.Sprint(p))
					fmt.Println("defer 1: ", p)
				}
			}()
			panic("接口查询失败")
		})
	}

	wg.Wait()
	if batchError != nil {
		fmt.Println("Batch error ", batchError)
		panic(batchError)
	}
	fmt.Println("OK")
}

func WithGo(parentCtx context.Context, process func(ctx context.Context)) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("batch 处理失败", r)
				debug.PrintStack()
			}
		}()
		process(parentCtx)
	}()
}
