package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	numOfWorkers := 10

	channel := make(chan int, numOfWorkers)
	stop := make(chan bool, numOfWorkers)
	stopped := 0

	interruptChannel := make(chan os.Signal, 2)
	signal.Notify(interruptChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		select {
		case <-interruptChannel:
			fmt.Println("received interrupt")
			cancel()
			return
		}
	}()

	for i := 0; i < numOfWorkers; i++ {
		go runWorker(i, channel, stop, ctx)
	}
	for stopped < numOfWorkers {
		select {
		case <-stop:
			stopped++
		default:
			if stopped == 0 {
				data := rand.Int()
				channel <- data
			}
		}
	}
	fmt.Println("all workers stopped")
	close(channel)
	close(stop)
	fmt.Println("done")
	os.Exit(2)
}

func runWorker(id int, c <-chan int, stop chan bool, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker", id, "stopped")
			stop <- true
			return
		case n := <-c:
			fmt.Println("worker", id, "got", n)
		}
	}
}
