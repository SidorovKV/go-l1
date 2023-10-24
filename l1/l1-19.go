package main

import "fmt"

func main() {
	fmt.Println(reverseString("абырвалг"))
	fmt.Println(reverseString("123456789"))
}

func reverseString(s string) string {
	runes := []rune(s)
	out := make([]rune, len(s))
	for i, j := 0, len(runes)-1; i < len(runes); i, j = i+1, j-1 {
		out[i] = runes[j]
	}
	return string(out)
}
