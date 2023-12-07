package main

import (
	"fmt"
	"sync"
)

var m sync.Map

func main()  {
	//新增
	m.Store(1, "one")
	m.Store(2, "two")

	//key 不存在情况
	v, ok := m.LoadOrStore(3, "three")
	fmt.Println(v, ok)

	//key存在情况
	v, ok = m.LoadOrStore(1, "thisOne")
	fmt.Println(v, ok)

	//Load
	v, ok = m.Load(1)
	if ok {
		fmt.Println(v)
	}

	//Range
	f := func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	}

	m.Range(f)

	m.Delete(1)
	fmt.Println(m.Load(1))
}
