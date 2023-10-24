package main

import (
	"time"
)

func main() {
	var secondsToWork time.Duration = 2

	channel := make(chan int)
	ws, rs := make(chan bool), make(chan bool)
	//timeout := time.After(secondsToWork * time.Second)

	go writer(ws, channel, time.After(secondsToWork*time.Second))
	go reader(rs, channel, time.After(secondsToWork*time.Second))
	<-ws
	close(ws)
	<-rs
	close(rs)
}

func writer(ws chan bool, c chan int, stop <-chan time.Time) {
	counter := 0
	for {
		select {
		case <-stop:
			ws <- true
			close(c)
			return
		default:
			c <- counter
			counter++
		}
	}
}

func reader(rs chan bool, c <-chan int, stop <-chan time.Time) {
	for {
		select {
		case <-stop:
			rs <- true
			return
		case <-c:
			//Считано)))
		default:
			continue
		}
	}
}
