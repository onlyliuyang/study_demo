package main

import "fmt"

func main() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case v1 := <-c1:
			c1Count++
			fmt.Println(v1)
		case v2 := <-c2:
			c2Count++
			fmt.Println(v2)
		}
	}
	fmt.Println(c1Count, c2Count)
}
