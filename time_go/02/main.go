package main

import (
	"fmt"
)

func main() {
	//var prices []int
	//prices = make([]int, 0)
	//prices = []int{7, 1, 5, 3, 6, 4}
	//fmt.Println(maxPrices(prices))

	fmt.Println(1 ^ 1)
	fmt.Println(1 ^ 2)
	fmt.Println(1 ^ 0)

}

func maxPrices(prices []int) int {

}

func singleNumber(nums []int) int {
	n := 0
	for i := 0; i < len(nums); i++ {
		n ^= nums[i]
	}
	return n
}
