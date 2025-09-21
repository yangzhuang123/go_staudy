package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter struct {
	Lock  sync.Mutex
	count int
}

func main() {
	count := &Counter{count: 0}

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				count.Lock.Lock()
				count.count++
				count.Lock.Unlock()
			}
		}()
	}

	time.Sleep(time.Second * 3)

	fmt.Println(count.count)

	var counter int32
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				atomic.AddInt32(&counter, 1)
			}
		}()
	}

	time.Sleep(time.Second * 3)

	fmt.Println(counter)
}
