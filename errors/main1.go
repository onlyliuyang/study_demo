package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3}
	x := a[:1]
	fmt.Println(len(x), cap(x))
	y := a[2:]
	fmt.Println(len(y), cap(y))
	x = append(x, y...)
	fmt.Println(a)
	x = append(x, y...)
	fmt.Println(a)
	fmt.Println(a, x)
}
