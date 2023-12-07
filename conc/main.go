package main

import (
	"fmt"
	"github.com/sourcegraph/conc"
	"sync"
)

var count int

func main() {
	var m sync.Mutex
	wg := conc.NewWaitGroup()
	for i := 0; i < 10; i++ {
		wg.Go(func() {
			if count == 6 {
				panic("six error")
			}

			m.Lock()
			defer m.Unlock()
			count++
		})
	}
	wg.WaitAndRecover()

	fmt.Println(count)
}
