package main

import (
	"context"
	"fmt"
	"math"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	doChan := make(chan interface{}, 1)
	dataChan := make(chan int)
	go func() {
		for {
			<-doChan
			select {
			case <-ctx.Done():
				fmt.Println("groutine1 done")
				return
			case v := <-dataChan:
				fmt.Println("groutine1 ", v)
			}
		}
	}()

	for i := 0; i < math.MaxInt; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("groutine2 done")
			return
		default:
			if i%2 == 0 {
				fmt.Println("groutine2 ", i)
			} else {
				doChan <- struct {
				}{}
				dataChan <- i
			}
		}
	}

	//time.Sleep(20 * time.Second)
}
