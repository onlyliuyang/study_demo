package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	letterChan := make(chan rune, 1)
	numberChan := make(chan int, 1)

	go func() {
		defer wg.Done()
		for i := 1; i <= 26; i++ {
			numberChan <- i
			letter := <-letterChan
			fmt.Printf("%d%c", i, letter)
		}
	}()

	go func() {
		defer wg.Done()
		for letter := 'B'; letter <= 'Z'; letter++ {
			<-numberChan
			letterChan <- letter
		}
	}()

	letterChan <- 'A' // 启动第一个协程

	wg.Wait()
}
