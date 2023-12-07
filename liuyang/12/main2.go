package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	numberChan := make(chan int)
	letterChan := make(chan rune)

	go func() {
		defer wg.Done()
		for i := 1; i <= 26; i++ {
			select {
			case numberChan <- i:
				letter := <-letterChan
				fmt.Printf("%d%c", i, letter)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for letter := 'A'; letter <= 'Z'; letter++ {
			select {
			case <-numberChan:
				letterChan <- letter
			}
		}
	}()

	wg.Wait()
}
