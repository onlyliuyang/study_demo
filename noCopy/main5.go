package main

import (
	"fmt"
	"time"
)

func main()  {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCount := 0
	loop:
		for {
			select {
			case <- done:
				break loop
			default:

			}

			//模拟工作行为
			workCount++
			time.Sleep(1 * time.Second)
		}
	fmt.Println("Achived %v cycles of work before signaled to stop.\n", workCount)
}
