package main

import (
	"fmt"
	"time"
)

func main() {
	myTimer := time.NewTimer(time.Second * 1)
	var i int = 0
	for {
		select {
		case <-myTimer.C:
			i++
			fmt.Println("count ", i)
			myTimer.Reset(time.Second * 1)
		}
	}

	myTimer.Stop()
}
