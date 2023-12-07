package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	m       sync.RWMutex
)

func incr() {
	m.Lock()
	defer m.Unlock()
	counter++
}

func decr() {
	m.Lock()
	defer m.Unlock()
	counter--
}

func get() int {
	m.RLock()
	defer m.RUnlock()
	return counter
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()
	}

	for i := 0; i < 400; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decr()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
