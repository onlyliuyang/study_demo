package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-tick.C:
			fmt.Println("aaaaa")
			tick.Reset(5 * time.Second)
		}
	}
}
