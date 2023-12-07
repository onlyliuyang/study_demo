package main

import "fmt"

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	result := 1
	for i := 1; i <= num; i++ {
		result += i
	}
	return result
}

func zeroes(n int) int {
	count := 0
	for n > 0 {
		count += n / 5
		n /= 5
	}
	return count
}

func main() {
	//n := factorial(5)
	//fmt.Println(n)
	fmt.Println(zeroes(15))
}
