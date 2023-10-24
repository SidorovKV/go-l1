package main

func main() {
	nums := []int{-78, -19, -3, 10, 14, 16, 18, 21, 58, 108}
	target := 10
	index, ok := binarySearch(nums, target)
	if ok {
		println(index)
	} else {
		println("Not Found")
	}

	target = 11
	index, ok = binarySearch(nums, target)
	if ok {
		println(index)
	} else {
		println("Not Found")
	}
}

func binarySearch(nums []int, target int) (uint, bool) {
	var low, high uint = 0, uint(len(nums) - 1)
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid, true
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return 0, false
}
