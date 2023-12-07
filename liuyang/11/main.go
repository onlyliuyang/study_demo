package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	number, letter := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Printf("%d%d", i, i+1)
				i += 2
				letter <- true
			default:
				break
			}
		}
	}()

	wait.Add(1)
	go func() {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				if i >= strings.Count(str, "")-1 {
					wait.Done()
					return
				}
				fmt.Print(str[i : i+2])
				i += 2
				number <- true
			default:
				break
			}
		}
	}()
	number <- true
	wait.Wait()
}
