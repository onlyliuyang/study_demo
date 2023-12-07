package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	ch := make(chan os.Signal)
	signal.Notify(ch)

	fmt.Println("start...")
	select {
	case s := <-ch:
		fmt.Println("end...", s)
	}
}
