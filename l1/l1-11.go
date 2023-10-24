package main

import "fmt"

type void struct{}

func main() {
	fmt.Println(intersection([]int{234, 10, 654, 20, -45, 30, 35}, []int{654, -45, 10, 999, 20, 40}))
}

func intersection(nums1 []int, nums2 []int) []int {
	t := void{}
	m := make(map[int]void)
	res := make([]int, 0)
	for _, v := range nums1 {
		m[v] = t
	}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			res = append(res, v)
		}
	}
	return res
}
