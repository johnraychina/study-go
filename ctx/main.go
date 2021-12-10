package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// text context
	//type emptyCtx int
	//e1 := new(emptyCtx)
	//e2 := new(emptyCtx)
	//if e1 != e2 {
	//	fmt.Println("e1 != e2")
	//}
	//
	//b1 := context.Background()
	//b2 := context.Background()
	//
	//if b1 == b2 {
	//	fmt.Println("b1==b2")
	//}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)
	//go handleIgnoreContext(ctx)
	//cancel()

	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
	time.Sleep(2 * time.Second)

}

func handleIgnoreContext(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("tick")
	}
}
func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
