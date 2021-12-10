package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	fmt.Println(time.Now().UTC().Unix())

	// Mutex
	mutex := sync.Mutex{}

	mutex.Lock()

	// RWMutex
	rw := sync.RWMutex{}
	rw.RLock()
	rw.RUnlock()
	rw.Lock()
	rw.Unlock()

	// WaitGroup
	group := sync.WaitGroup{}
	go func() {
		// producer
		group.Add(1)
	}()
	go func() {
		// worker
		group.Done()
	}()
	// main procedure
	group.Wait()

	// Once
	once := sync.Once{}
	once.Do(func() {
		println("init")
	})

	// Cond condition: 基于锁 + 等待链表
	// 按条件加锁等待，适用于条件长期得不到满足的场景
	// 由于有wait，所以会比自旋更省cpu
	c := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		go listen(c)
	}

	time.Sleep(1 * time.Second)

	go broadcast(c)

	// 订阅系统中断，结束程序
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	// ErrGroup

	// singleflight.Group
	// 在资源的获取非常昂贵时（例如：访问缓存、数据库），就很适合使用 golang/sync/singleflight.Group 优化服务。
	singleflight.Group
}

type service struct {
	requestGroup singleflight.Group
}

func broadcast(c *sync.Cond) {
	c.L.Lock()
	atomic.StoreInt64(&status, 1)
	c.Broadcast()
	c.L.Unlock()
}

var status int64

func listen(c *sync.Cond) {
	c.L.Lock()
	for atomic.LoadInt64(&status) != 1 {
		c.Wait() // wait atomically unlocks c.L
	}
	fmt.Println("listen")
	c.L.Unlock()
}
