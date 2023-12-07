package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu          sync.Mutex
	CounterType int
	Name        string
	count       uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Decr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func main() {
	var c Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				c.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(c.count)
}
