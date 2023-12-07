package main

import (
	"fmt"
	"sync"
	"time"
)

type Container struct {
	sync.Mutex
	counters map[string]int
}

func (c *Container) incr(name string) {
	c.Lock()
	defer c.Unlock()

	c.counters[name]++
}

func main()  {
	c := Container{
		Mutex:    sync.Mutex{},
		counters: map[string]int{"a":0, "b":0},
	}

	doIncrement := func(name string, n int) {
		for i:=0; i<n; i++ {
			c.incr(name)
		}
	}

	go doIncrement("a", 100)
	go doIncrement("b", 200)

	time.Sleep(2 * time.Second)
	fmt.Println(c.counters)
}
