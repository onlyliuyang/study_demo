package main

import "fmt"

func main() {
	//m := new(map[string]int)
	m := make(map[string]int)
	m1 := new(map[string]string)

	//m["name"] = 123
	m["name"] = 123
	fmt.Println(m, m1)

}
