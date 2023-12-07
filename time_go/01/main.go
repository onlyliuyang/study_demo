package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int = 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			for j := 0; j < 10000; j++ {
				count++
			}
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
