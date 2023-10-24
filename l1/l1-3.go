package main

import (
	"sync"
)

func main() {
	a := [5]int{2, 4, 6, 8, 10}
	sum := variant1(a)
	sum = variant2(a)
	sum = variant3(a)
	sum = variant4(a)
	sum = sum
}

func variant1(a [5]int) int {
	sum := 0
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	for i := 0; i < len(a); i++ { //Можно и for _, v := range a { передавая в рутину v
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			res := x * x
			mu.Lock()
			sum += res
			mu.Unlock()
		}(a[i])
	}
	wg.Wait()
	return sum
}

func variant2(a [5]int) int {
	sum := 0
	wg := new(sync.WaitGroup)
	channel := make(chan int, 5)
	for i := 0; i < len(a); i++ { //Можно и for _, v := range a {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			channel <- x * x
		}(a[i])
	}
	wg.Wait()
	close(channel)
	for v := range channel {
		sum += v
	}
	return sum
}

func variant3(a [5]int) int { //Можно и for i, v := range a {
	sum := 0
	wg := new(sync.WaitGroup)
	out := [5]int{}
	for i := 0; i < len(a); i++ {
		wg.Add(1)
		go func(x int, i int) {
			defer wg.Done()
			out[i] = x * x
		}(a[i], i)
	}
	wg.Wait()
	for _, v := range out {
		sum += v
	}
	return sum
}

func variant4(a [5]int) int {
	sum := concAdder()
	wg := new(sync.WaitGroup)
	for i := 0; i < len(a); i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			sum(x * x)
		}(a[i])
	}
	wg.Wait()
	return sum(0)
}

func concAdder() func(int) int {
	mu := new(sync.Mutex)
	sum := 0
	return func(x int) int {
		mu.Lock()
		sum += x
		mu.Unlock()
		return sum
	}
}
