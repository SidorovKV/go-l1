package main

import (
	"fmt"
	"sync"
)

func main() {
	variant7_1()
	variant7_2()
}

func variant7_1() {
	ourMap := make(map[int]int)
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			ourMap[i] = 2 * i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println(ourMap[5])
	fmt.Println(len(ourMap))
}

func variant7_2() {
	ourMap := &sync.Map{}
	wg := new(sync.WaitGroup)
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ourMap.Store(i, 2*i)
		}(i)
	}
	wg.Wait()
	fmt.Println(ourMap.Load(5))
	fmt.Println(ourMap.Load(1000))
}
