package main

import (
	"fmt"
	"time"
)

func goRoutineD(ch chan int, i int) {
	time.Sleep(time.Second * 3)
	ch <- i
}

func goRoutineE(chs chan string, i string) {
	time.Sleep(time.Second * 3)
	chs <- i
}

func main() {
	ch := make(chan int, 5)
	chs := make(chan string, 5)

	go goRoutineD(ch, 5)
	go goRoutineE(chs, "ok")

	//time.Sleep(5 * time.Second)

	select {
	case msg := <-chs:
		fmt.Println("receive the data ", msg)
	case msg := <-ch:
		fmt.Println("receive the data ", msg)
		//default:
		//	fmt.Println("no data received")
		//	time.Sleep(time.Second * 1)
	}
}
