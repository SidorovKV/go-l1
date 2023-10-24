package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := []int{5, 3, 9, 1, 6, 8, 7, 2, 4}

	quicksortInts(a)
	fmt.Println(a)
	fmt.Println(isSortedInts(a))

	b := [15]int{}
	for v := range b {
		b[v] = rand.Intn(200) - 100
	}
	fmt.Println(b)
	quicksortInts(b[:])
	fmt.Println(b)
	fmt.Println(isSortedInts(b[:]))

	f := [15]float64{}
	for v := range f {
		f[v] = rand.Float64()*100 - 50
	}
	fmt.Printf("%.2f\n", f)
	quicksortFloats64(f[:])
	fmt.Printf("%.2f\n", f)
	fmt.Println(isSortedFloats64(f[:]))
}

// Можно сделать обобщённую функцию через женерик, но нужен импорт "cmp"
// func quicksort[T cmp.Ordered](nums []T) {...}
func quicksortInts(nums []int) {
	p := nums[len(nums)/2]
	l, r := 0, len(nums)-1
	for l < r {
		for nums[l] < p {
			l++
		}
		for nums[r] > p {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
		}
		if nums[l] == nums[r] {
			l++
		}
	}
	left := nums[:r]
	right := nums[r:]
	if len(left) > 1 {
		quicksortInts(left)
	}
	if len(right) > 1 {
		quicksortInts(right)
	}
}

func quicksortFloats64(nums []float64) {
	p := nums[len(nums)/2]
	l, r := 0, len(nums)-1
	for l < r {
		for nums[l] < p {
			l++
		}
		for nums[r] > p {
			r--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
		if nums[l] == nums[r] {
			l++
		}
	}
	left := nums[:r]
	right := nums[r:]
	if len(left) > 1 {
		quicksortFloats64(left)
	}
	if len(right) > 1 {
		quicksortFloats64(right)
	}
}

func isSortedInts(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			return false
		}
	}
	return true
}

func isSortedFloats64(nums []float64) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			return false
		}
	}
	return true
}
