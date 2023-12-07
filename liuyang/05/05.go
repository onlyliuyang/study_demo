package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() {
		defer close(done)
	})

	internal := time.Tick(1 * time.Second)
	for {
		select {
		case <-done:
			return
		case ttt := <-internal:
			fmt.Println(ttt)
		}
	}
}
