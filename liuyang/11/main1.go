package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func letter(ch chan string) {
	defer wg.Done()

	for i := 0; i < 26; i++ {
		ch <- fmt.Sprintf("%c", 'A'+i)
	}
	close(ch)
}

func number(ch chan int) {
	defer wg.Done()
	for i := 1; i < 29; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	leChar := make(chan string, 26)
	intChar := make(chan int, 28)
	wg.Add(3)
	go letter(leChar)
	go number(intChar)
	go func() {
		defer wg.Done()
		for i := range intChar {
			fmt.Printf("%d%d%s%s", i, <-intChar, <-leChar, <-leChar)
		}
	}()
	wg.Wait()
}