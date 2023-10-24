package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count uint
	mu    *sync.Mutex
}

func main() {
	wg := new(sync.WaitGroup)
	c := Counter{
		count: 0,
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
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}
