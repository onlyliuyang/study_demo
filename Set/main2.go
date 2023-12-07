package main

import "fmt"

func main() {
	var num int = 10
	ch := make(chan struct{}, 10)

	for i := 0; i < num; i++ {
		go func(i int) {
			defer func() {
				ch <- struct{}{}
			}()

			for {
				fmt.Printf("goroutine %d\n", i)
			}
		}(i)
	}

	for i := 0; i < num; i++ {
		<-ch
	}
}
