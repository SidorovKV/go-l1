package main

import "fmt"

func main() {
	a := [...]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	out := make(map[int][]float64)
	for _, v := range a {
		key := int(v/10) * 10
		slice, ok := out[key]
		if !ok {
			slice = []float64{v}
			out[key] = slice
		} else {
			slice = append(slice, v)
			out[key] = slice
		}
	}
	fmt.Println(out)
}
