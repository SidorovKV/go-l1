package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	count *atomic.Uint64
	mu    *sync.Mutex
}

func main() {
	wg := new(sync.WaitGroup)
	c := Counter{
		count: &atomic.Uint64{},
		mu:    new(sync.Mutex),
	}
	for i := 0; i < 10000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	fmt.Println(c.count)
}

func (c *Counter) Inc() {
	c.count.Add(1)
}
