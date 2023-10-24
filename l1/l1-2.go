package main

import (
	"fmt"
	"sync"
)

func main() {
	a := [5]int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)
	for i := 0; i < len(a); i++ { //Можно и for _, v := range a { передавая в рутину v
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Println(x * x)
		}(a[i])
	}
	// Ожидание завершения горутин через WaitGroup
	wg.Wait()

	//----
	fmt.Println("-----")
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
		fmt.Println(v)
	}

	//----Если важен порядок
	fmt.Println("-----")
	out := [5]int{}
	for i := 0; i < len(a); i++ { //Можно и for i, v := range a {
		wg.Add(1)
		go func(x int, i int) {
			defer wg.Done()
			out[i] = x * x
		}(a[i], i)
	}
	wg.Wait()
	for _, v := range out {
		fmt.Println(v)
	}

	//---- Синхронизируемся с мейном, но теряем конкурентность, если я правильно понимаю
	fmt.Println("-----")
	f := func(x int) <-chan struct{} {
		done := make(chan struct{})
		go func(x int) {
			fmt.Println(x * x)
			done <- struct{}{}
		}(x)
		return done
	}
	for _, v := range a {
		<-f(v)
	}
}
