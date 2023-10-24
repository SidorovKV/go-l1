package main

import (
	"fmt"
	"sync"
)

func main() {
	ourMap := make(map[int]int, 10000)
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			ourMap[i] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(ourMap))
}
