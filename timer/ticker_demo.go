package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	var statesCount = make(map[int32]int32)
	statesCount[1]++
	statesCount[1]++
	println(statesCount[1])

	var maxEndTime time.Time
	print(maxEndTime.IsZero())
	//print(time.Millisecond)

	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//}()
	//
	//
	////ctx, cancelFunc := context.WithCancel(context.Background())
	//ctx, cancelFunc := context.WithTimeout(context.Background(), 100 * 1000 * time.Millisecond)
	//go StartTickerWithTimer2(ctx, 1)
	//go StartTickerWithTimer2(ctx, 2)
	//
	//time.Sleep(100 * 1000 * time.Millisecond)
	//cancelFunc()
}

func StartTickerWithTimer2(ctx context.Context, id int) {

	go func(c context.Context) {
		ticker := time.NewTicker(2000 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Ticker ", id, " end")
			case t := <-ticker.C:
				fmt.Println("Ticker ", id, ", at ", t)
			}
		}

	}(ctx)

}
