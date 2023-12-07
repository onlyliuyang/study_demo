package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	count := 10
	sum := 1000000
	var wg sync.WaitGroup

	gChan := make(chan struct{}, count)
	for i := 0; i < sum; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			gChan <- struct{}{}
			time.Sleep(1 * time.Second)
			fmt.Println(i)
			<-gChan //执行完毕
		}(i)
	}
	wg.Wait()
}
