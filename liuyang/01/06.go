package main

import (
	"fmt"
	"sync"
	"time"
)

var done bool = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait()
	}

	fmt.Println(name, " starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	fmt.Println("start writing")
	done = true
	fmt.Println(name, "wakes all")
	c.Broadcast()
}

func main() {
	c := sync.NewCond(&sync.Mutex{})
	go read("reader1", c)
	go read("reader2", c)
	go read("reader3", c)
	write("writer", c)
	time.Sleep(3 * time.Second)
}
