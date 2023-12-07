package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var createNum int32

func createBuffer() interface{} {
	atomic.AddInt32(&createNum, 1)
	buffer := make([]byte, 1024)
	return buffer
}

func main() {
	bufferPool := sync.Pool{
		New: createBuffer,
	}

	workerPool := 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(workerPool)

	for i := 0; i < workerPool; i++ {
		go func() {
			defer wg.Done()
			buffer := bufferPool.Get()
			_ = buffer.([]byte)
			defer bufferPool.Put(buffer)
		}()
	}
	wg.Wait()

	fmt.Printf("%d buffer objects were created.\n", createNum)
}
