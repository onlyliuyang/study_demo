package main

import (
	"fmt"
	"math/rand"
)

func main() {
	newRandStream := func() <-chan int {
		randStrem := make(chan int)
		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStrem)

			for {
				randStrem <- rand.Int()
			}
		}()
		return randStrem
	}

	randStrem := newRandStream()
	fmt.Println("3 random ints")

	for i := 0; i < 30; i++ {
		fmt.Printf("%d: %d\n", i, <-randStrem)
	}
}
