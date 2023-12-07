package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan<- interface{}
	//c = make(chan interface{})

	//close(c)

	select {
	case c <- "aa":
		fmt.Println("from c")
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
