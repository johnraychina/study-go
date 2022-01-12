package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	b := NewLeakyBucket(10, 100)
	// 统计20秒内leak请求数，看看与 20 * reqPerSecond 相差多少, 准确率在95%。
	cnt := 0
	timer := time.NewTimer(10 * time.Second)
FOR:
	for {
		select {
		case <-timer.C:
			break FOR
		default:
			requests := b.Leak(genRequest)
			Handle(requests)
			cnt += len(requests)
		}
	}
	fmt.Println(cnt)

	//input := bufio.NewScanner(os.Stdin)
	//ch := make(chan string)
	//go func() {
	//	input.Scan()
	//	ch <- input.Text()
	//}()
	//
	//
	//for {
	//	select {
	//	case <-ch:
	//		return
	//	default:
	//		requests := b.Leak(genRequest)
	//		Handle(requests)
	//	}
	//}
}

type leakyBucket struct {
	lock         *sync.Mutex
	size         int
	queue        []string
	rate         int64
	lastLeakTime time.Time
}

func NewLeakyBucket(size int, reqPerSecond int64) *leakyBucket {
	b := &leakyBucket{}
	b.lock = &sync.Mutex{}
	b.rate = reqPerSecond
	b.size = size
	b.queue = make([]string, 0, size)
	b.lastLeakTime = time.Now()
	return b
}

func (b *leakyBucket) Leak(genRequest func() []string) []string {
	b.lock.Lock()
	defer func() {
		b.lock.Unlock()
	}()

	// 计算应该拿出多少请求（当然不能超过队列中的请求）
	duration := time.Now().Sub(b.lastLeakTime)
	n := (duration.Nanoseconds() * b.rate) / 1e9
	qLen := len(b.queue)
	if n > int64(qLen) {
		n = int64(qLen)
	}
	if n > 0 {
		b.lastLeakTime = time.Now()
	}

	result := make([]string, 0, n)
	for _, s := range b.queue[0:n] {
		result = append(result, s)
	}
	b.queue = b.queue[n:]

	// 将请求放入queue中，多余的drop
	requests := genRequest()
	m := b.size - len(b.queue)
	if m > len(requests) {
		m = len(requests)
	}
	for _, r := range requests[0:m] {
		b.queue = append(b.queue, r)
	}

	fmt.Printf("%s leak: %d, put %d, drop %d, bucket cap: %d, len: %d \n", time.Now(), n, m, len(requests)-m, b.size, len(b.queue))
	return result
}

func genRequest() []string {
	i := rand.Intn(100)
	if i > 10 {
		return []string{"a", "b", "c"}
	}
	return []string{"a"}
}

func Handle(requests []string) {
	i := rand.Intn(2)
	time.Sleep(time.Millisecond * time.Duration(i))
}
