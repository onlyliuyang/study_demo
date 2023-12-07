package main

import (
	"fmt"
	"time"
)

var realNum = make(chan int)	//设置数字值
var delta = make(chan int)	//设置的增减额

func setNumber(n int)  {
	realNum <- n
}

func changeByDelta(d int)  {
	delta <- d
}

func getNumber() int {
	return <- realNum
}

func monitor()  {
	var i int //把数值限定在方法内，groutine运行后仅在groutine内可见
	for {
		select {
			case i = <-realNum:
			case d := <- delta:
				i += d
		case realNum <- i:

		}
	}
}

func init()  {
	go monitor()
}

func main()  {
	var i int = 100
	var j int = 200

	go func() {
		setNumber(i)
	}()
	go func() {
		res := getNumber()
		fmt.Println(res)
	}()

	go func() {
		changeByDelta(j)
	}()

	time.Sleep(2 * time.Second)
}