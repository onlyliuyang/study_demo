package main

import "fmt"

func main() {
	multiply := func(values []int, multiplier int) []int {
		multipledValues := make([]int, len(values))
		for i, v := range values {
			multipledValues[i] = v * multiplier
		}
		return multipledValues
	}

	add := func(values []int, additive int) []int {
		addValues := make([]int, len(values))
		for i, v := range values {
			addValues[i] = v + additive
		}
		return addValues
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range multiply(add(multiply(ints, 2), 1), 2) {
		fmt.Println(v)
	}
}
