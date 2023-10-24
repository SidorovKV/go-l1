package main

import "fmt"

func main() {
	a := [100]int{}
	for i := range a {
		a[i] = i + 1
	}
	channel1, channel2 := make(chan int, len(a)), make(chan int, len(a))
	quit := make(chan struct{})

	go func(a [100]int) {
		for _, v := range a {
			channel1 <- v
		}
		close(channel1)
	}(a)

	go func() {
		for v := range channel1 {
			channel2 <- v * v
		}
		close(channel2)
	}()

	go func() {
		for v := range channel2 {
			fmt.Println(v)
		}
		quit <- struct{}{}
	}()

	<-quit
	fmt.Println("end")
}
