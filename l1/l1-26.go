package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(checkUnique("abcd"))
	fmt.Println(checkUnique("abCdefAaf"))
	fmt.Println(checkUnique("aA"))
	fmt.Println(checkUnique("aabcd"))

}

func checkUnique(s string) bool {
	s = strings.ToLower(s)
	m := make(map[rune]bool)
	for _, c := range s {
		if m[c] {
			return false
		}
		m[c] = true
	}
	return true
}
