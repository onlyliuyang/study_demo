package main

import "fmt"

func main() {
	data := []int{5, 23, 2, 5, 7, 8, 4, 3, 212, 5, 7, 99, 4, 3}
	QuickSort(0, len(data)-1, data)
	fmt.Println(data)
}

func QuickSort(left, right int, data []int) {
	l := left
	r := right

	pivot := data[(left+right)/2]
	for l < r {
		//从pivot左边找到大于pivot的值
		for data[l] < pivot {
			l++
		}

		//从pivot右边找到小于pivot的值
		for data[r] > pivot {
			r--
		}

		//交换位置
		data[l], data[r] = data[r], data[l]

		if l == r {
			l++
			r--
		}

		if left < r {
			QuickSort(left, r, data)
		}

		if right > l {
			QuickSort(l, right, data)
		}
	}
}
