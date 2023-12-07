package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	//channel 长度
	poolCount = 1
	//复用的goroutine数量
	gorouteineCount = 5
)

func main() {
	start := time.Now()
	jobsChan := make(chan int, poolCount)

	//workers
	var wg sync.WaitGroup
	for i := 0; i < gorouteineCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range jobsChan {
				time.Sleep(1 * time.Second)
				fmt.Printf("Items is %v, This NumGoroutine is %v\n", item, runtime.NumGoroutine())
			}
		}()
	}

	//senders
	for i := 0; i < 100000; i++ {
		jobsChan <- i
	}

	//关闭channel
	close(jobsChan)
	wg.Wait()

	times := time.Since(start)
	fmt.Println(times)
}
