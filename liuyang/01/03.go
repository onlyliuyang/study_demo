package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	go func() {
		defer func() {
			close(done)
		}()
		time.Sleep(5 * time.Second)
	}()

	workCount := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}

		//模拟工作行为
		workCount++
		time.Sleep(1 * time.Second)
	}

	fmt.Println(workCount)
}
