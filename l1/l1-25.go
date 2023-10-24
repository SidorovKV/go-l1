package main

import (
	"fmt"
	"syscall"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Start", t.Second(), t.Nanosecond()/1000000)
	customSleep(500)
	t = time.Now()
	fmt.Println("End", t.Second(), t.Nanosecond()/1000000)

	t = time.Now()
	fmt.Println("Start", t.Second(), t.Nanosecond()/1000000)
	customSleep2(500)
	t = time.Now()
	fmt.Println("End", t.Second(), t.Nanosecond()/1000000)

	t = time.Now()
	fmt.Println("Start", t.Second(), t.Nanosecond()/1000000)
	customSleep3(500)
	t = time.Now()
	fmt.Println("End", t.Second(), t.Nanosecond()/1000000)
}

// Возможно не работает на маках
func customSleep(milliseconds int64) {
	if milliseconds <= 0 {
		return
	}
	seconds := milliseconds / 1000
	d := milliseconds % 1000
	nanos := d * 1000000
	syscall.Nanosleep(&syscall.Timespec{Sec: seconds, Nsec: nanos}, nil)
}

func customSleep2(milliseconds int64) {
	t := time.Now().Add(time.Millisecond * time.Duration(milliseconds))
	for time.Now().Before(t) {
	}
}

func customSleep3(milliseconds int64) {
	<-time.After(time.Millisecond * time.Duration(milliseconds))
}
