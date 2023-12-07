package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []int{0, 1, 2, 3, 4}
	b := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}

	res := reflect.TypeOf(a).Kind() == reflect.TypeOf(b).Kind()
	fmt.Println(res)

	fmt.Println(reflect.DeepEqual(a, b))
}
