package main

import (
	"strings"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]rune(v)[:100])
}

func main() {
	someFunc()
}

func createHugeString(l int) string {
	out := &strings.Builder{}
	for i := 0; i < l; i++ {
		out.WriteString("a")
	}
	return out.String()
}
