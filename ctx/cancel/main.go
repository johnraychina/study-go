package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	cancelContext()
}

func cancelContext() {
	ctx := context.Background()
	// 一个子任务cancel 其他子任务，本质是大家select 监听同一个channel: subCtx.Done()
	subCtx, cancel := context.WithCancel(ctx)
	for i := 0; i < 10; i++ {
		localI := i
		go func(innerI int) {
			//if innerI == 5 {
			//	cancel()
			//}

			for {
				select {
				case <-subCtx.Done():
					fmt.Println(innerI, " cancelled")
					return
				default:
					fmt.Println(innerI, "working ", time.Now())
					time.Sleep(100 * time.Millisecond)
				}
			}
		}(localI)
	}

	// 超时则取消
	<-time.NewTimer(5 * time.Second).C
	cancel()
	<-time.NewTimer(5 * time.Second).C
	fmt.Println("end")
}
