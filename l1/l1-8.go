package main

import (
	"fmt"
)

func main() {
	fmt.Println(setBit(82, 2, true))
	fmt.Println(setBit(82, 4, false))
}

func setBit(n uint64, pos uint, b bool) uint64 {
	if b {
		n |= 1 << pos
	} else {
		n &= ^(1 << pos)
	}
	return n
}
