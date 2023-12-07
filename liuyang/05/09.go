package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	for val := range readNum(ctx) {
		fmt.Println(val)
	}
}

func readNum(ctx context.Context) <-chan int {
	dataChan := make(chan int)
	var i int = 0
	go func() {
		defer close(dataChan)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Channel is Closed")
				return
			case dataChan <- i:
				i++
			}
		}
	}()
	return dataChan
}
