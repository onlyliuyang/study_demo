package main

import "fmt"

func bubbleSort(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func quickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}

	middle := arr[0]
	left, right := make([]int, 0), make([]int, 0)
	for i := 1; i < length; i++ {
		if arr[i] > middle {
			right = append(right, arr[i])
		} else {
			left = append(left, arr[i])
		}
	}
	middleS := []int{middle}
	left = quickSort(left)
	right = quickSort(right)
	arr = append(append(left, middleS...), right...)
	return arr
}

func main() {
	arr := []int{23, 5, 1, 4, 68, 9, 4, 5, 65, 2}
	//bubbleSort(arr)
	arr = quickSort(arr)
	fmt.Println(arr)
}
