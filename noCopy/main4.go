package main

import (
	"fmt"
	"time"
)

func main()  {
	var c chan int
	c = make(chan int, 1)

	go func() {
		c <- 100
	}()

	for {
		select {
		case <-c:
			fmt.Println("读到数据")
		case <-time.After(2 * time.Second):
			fmt.Println("time out")

		}
	}

}
