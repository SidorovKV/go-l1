package main

import (
	"fmt"
)

func main() {
	c := make(chan interface{})
	typeIdentifier(6)
	typeIdentifier(true)
	typeIdentifier("hello")
	typeIdentifier(c)
	close(c)
}

func typeIdentifier(in interface{}) {
	switch in.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan interface{}:
		fmt.Println("channel")
	default:
		fmt.Println("unknown")
	}
}
