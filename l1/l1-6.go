package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithDeadline(context.Background(), time.Now().Add(200*time.Millisecond))
	ctx3, cancel3 := context.WithTimeout(context.Background(), 200*time.Millisecond)
	stop := make(chan struct{})
	defer cancel2()
	defer cancel3()

	go one(ctx)

	go two(time.After(100 * time.Millisecond))

	go three(time.NewTimer(200 * time.Millisecond))

	go four(time.NewTicker(100 * time.Millisecond))

	go five(ctx2)

	go six(ctx3)

	go seven(stop)

	time.Sleep(300 * time.Millisecond)
	cancel()
	stop <- struct{}{}
	close(stop)

	time.Sleep(1 * time.Second)
	fmt.Println("main done")
}

func one(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("one done")
			return
		default:
			continue
		}
	}
}

func two(stop <-chan time.Time) {
	for {
		select {
		case <-stop:
			fmt.Println("two done")
			return
		default:
			continue
		}
	}
}

func three(stop *time.Timer) {
	for {
		select {
		case <-stop.C:
			fmt.Println("three done")
			return
		default:
			continue
		}
	}
}

func four(stop *time.Ticker) {
	for {
		select {
		case <-stop.C:
			fmt.Println("four done")
			stop.Stop()
			return
		default:
			continue
		}
	}
}

func five(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("five done")
			return
		default:
			continue
		}
	}
}

func six(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("six done")
			return
		default:
			continue
		}
	}
}

func seven(stop <-chan struct{}) {
	for {
		select {
		case <-stop:
			fmt.Println("seven done")
			return
		default:
			continue
		}
	}
}
