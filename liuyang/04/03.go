package main

import (
	"fmt"
	"math/rand"
)

func main() {
	repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
		fmt.Println(fn(), " ", "123")
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():

				}
			}
		}()
		return valueStream
	}

	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:

					}
				}
			}
		}()
		return valueStream
	}
	//
	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case value := <-valueStream:
					//fmt.Println(value, " 123")
					takeStream <- value
					//case takeStream <- valueStream:

				}
			}
		}()
		return takeStream
	}

	toString := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case stringStream <- v.(string):
				}
			}
		}()
		return stringStream
	}

	done := make(chan interface{})
	defer close(done)

	//for num := range repeat(done, 1) {
	//	fmt.Println(num)
	//}
	//time.Sleep(2 * time.Second)

	randInt := func() interface{} { return rand.Int() }

	var message string
	for token := range toString(done, take(done, repeat(done, "I", "am"), 5)) {
		message += token
	}
	fmt.Println("message: ", message)

	for num := range take(done, repeatFn(done, randInt), 10) {
		fmt.Println(num)
		//fmt.Printf("%v\n", num)
	}
}
