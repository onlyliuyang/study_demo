package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int32 = 0

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//atomic.AddInt32(&counter, 1)
			for {
				val := atomic.LoadInt32(&counter)
				newVal := val + 1
				ok := atomic.CompareAndSwapInt32(&counter, val, newVal)
				if ok {
					return
				}
			}

			//atomic.StoreInt32(&counter, 210)
		}()
	}
	wg.Wait()
	fmt.Println(counter)

}
