package main

import "fmt"

func main() {
	fmt.Println(makeSet([]string{"cat", "cat", "dog", "cat", "tree"}))
}

func makeSet(in []string) []string {
	void := struct{}{}
	set := make(map[string]struct{})
	for _, v := range in {
		set[v] = void
	}
	out := make([]string, 0, len(set))
	for k := range set {
		out = append(out, k)
	}
	return out
}
