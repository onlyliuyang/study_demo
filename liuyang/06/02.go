package main

import (
	"fmt"
)

func main() {
	//fmt.Println(testMain())
	fmt.Println(testMain2())
}

func testMain() int {
	var i int = 0
	//var wg sync.WaitGroup

	defer func() {
		i++
		fmt.Println("defer1 ", i)
	}()

	defer func() {
		i++
		fmt.Println("defer2 ", i)
	}()
	return i
}

func testMain2() (i int) {
	defer func() {
		i++
		fmt.Println("defer1 ", i)
	}()

	defer func() {
		i++
		fmt.Println("defer2 ", i)
	}()
	return i
}
