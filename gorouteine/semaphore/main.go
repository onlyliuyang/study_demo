package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"runtime"
	"sync"
	"time"
)

const (
	//同时运行的goroutine上限
	Limit = 10
	//信号量权重
	Weight = 1
)

func main() {
	start := time.Now()
	names := []int{1, 2, 3, 4, 5}

	for i := 0; i < 1000; i++ {
		names = append(names, i)
	}

	sem := semaphore.NewWeighted(Limit)
	var wg sync.WaitGroup

	for _, name := range names {
		wg.Add(1)
		go func(name int) {
			defer wg.Done()
			err := sem.Acquire(context.Background(), Weight)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			time.Sleep(1 * time.Second)
			fmt.Printf("Items is %v, This NumGroutine is %v \n", name, runtime.NumGoroutine())
			sem.Release(Weight)
		}(name)
	}
	wg.Wait()

	times := time.Since(start)
	fmt.Println(times)
}
