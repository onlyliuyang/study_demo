package main

import (
	"fmt"
	"sync"
)

var smap sync.Map

func Set(key, value interface{})  {
	smap.Store(key, value)
}

func Get(key interface{}) interface{} {
	value, exists := smap.Load(key)
	if exists {
		return value
	}
	return nil
}

func Delete(key interface{})  {
	smap.Delete(key)
}

func MapRange(funcs func(key, value interface{}) bool)  {
	smap.Range(funcs)
}

func main()  {
	groutines := 10
	ch := make(chan struct{}, groutines)

	for i:=0; i<groutines; i++ {
		go func(i int) {
			defer func() {
				ch <- struct{}{}
			}()
			
			Set(fmt.Sprintf("key_%d", i), i)
		}(i)
	}
	
	for i:=0; i<groutines; i++ {
		<-ch
	}
	
	MapRange(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}