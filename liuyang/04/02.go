package main

import "fmt"

func main() {
	generator := func(done <-chan interface{}, integres ...int) <-chan int {
		inStream := make(chan int)
		go func() {
			defer close(inStream)

			for _, i := range integres {
				select {
				case <-done:
					return
				case inStream <- i:

				}
			}
		}()
		return inStream
	}

	multiply := func(done <-chan interface{}, inStream <-chan int, multiplier int) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range inStream {
				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream
	}

	add := func(done <-chan interface{}, inStream <-chan int, additive int) <-chan int {
		addStream := make(chan int)
		go func() {
			defer close(addStream)
			for i := range inStream {
				select {
				case <-done:
					return
				case addStream <- i + additive:

				}
			}
		}()
		return addStream
	}

	done := make(chan interface{})
	defer close(done)

	inStream := generator(done, 1, 2, 3, 4, 5)
	pipline := multiply(done, add(done, multiply(done, inStream, 2), 1), 2)
	for v := range pipline {
		fmt.Println(v)
	}

}
