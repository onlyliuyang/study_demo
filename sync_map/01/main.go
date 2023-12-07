package main

import (
	"fmt"
	"sync"
)

func main() {
	var mp sync.Map

	mp.Store("name", "liuyang")
	value, res := mp.Load("name")
	fmt.Println(value, res)

	val, res := mp.LoadOrStore("name2", "zhangsan")
	fmt.Println(val, res)
}
