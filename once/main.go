package main

import (
	"fmt"
	"sync"
)

type singleLeton map[string]string

var (
	once sync.Once
	instance singleLeton
)

func New() singleLeton  {
	once.Do(func() {
		instance = make(map[string]string)
	})
	return instance
}

func main()  {
	s := New()
	s["test1"] = "test1"

	s1 := New()
	s1["test2"] = "test2"

	fmt.Println(s)
	fmt.Println(s1)
}
