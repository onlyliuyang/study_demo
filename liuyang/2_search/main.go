package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right - left)
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{234, 53, 5, 7, 343, 4523, 6, 8}
	sort.Ints(arr)
	fmt.Println(arr)
	fmt.Println(binarySearch(arr, 7))
}
