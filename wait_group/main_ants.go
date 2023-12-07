package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"time"
)

func main() {
	pool, err := ants.NewPool(5)
	if err != nil {
		panic(err)
	}
	defer pool.Release()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		err = pool.Submit(func() {
			fmt.Printf("Task %d is done\n", i)
			time.Sleep(1 * time.Second)
			wg.Done()
		})
		if err != nil {
			panic(err)
		}
	}
	wg.Wait()
}
