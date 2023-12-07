package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"time"
)

func main()  {
	pool, _ := ants.NewPool(20)
	defer pool.Release()

	var runTimes int = 20
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		func() {
			time.Sleep(2 * time.Second)
			fmt.Println("Hello World")
		}()
		wg.Done()
	}

	for i:=0; i< runTimes; i++ {
		wg.Add(1)
		_ = pool.Submit(syncCalculateSum)
	}
	wg.Wait()

}
