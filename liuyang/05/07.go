package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan interface{})
	go func() {
		time.Sleep(time.Second * 20)
		close(done)
	}()

	for {
		select {
		case <-done:
			fmt.Println("done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
			time.Sleep(4 * time.Second)
		}
	}
}
