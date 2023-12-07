package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("new rand stream exit")
			defer close(randStream)

			for {
				select {
				case <-done:
					return
				case randStream <- rand.Int():

				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("3 random ints")
	for i := 0; i < 3; i++ {
		<-randStream
	}
	close(done)
	time.Sleep(1 * time.Second)
}
