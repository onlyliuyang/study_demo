package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int32
	wg sync.WaitGroup
)

func main()  {
	threadNum := 5
	wg.Add(threadNum)

	for i:=0; i<threadNum; i++ {
		go incrCounter(i)
	}
	wg.Wait()
}

func incrCounter(index int)  {
	defer wg.Done()

	spinNum := 0
	for {
		old := counter
		ok := atomic.CompareAndSwapInt32(&counter, old, old-1)
		if ok {
			break
		} else {
			spinNum++
		}
	}
	fmt.Printf("thread, %d, spinnum, %d, counter, %d\n", index, spinNum, counter)
}
