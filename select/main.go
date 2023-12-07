package main

import (
	"fmt"
	"time"
)

const (
	FMAT = "2006-01-02 15:04:05"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- 1
		close(ch)
	}()

	for {
		select {
		case x, ok := <-ch:
			fmt.Printf("%v, 通道读取到: x=%v, ok=%v\n", time.Now().Format(FMAT), x, ok)
			time.Sleep(500 * time.Millisecond)
			if !ok {
				ch = nil //把关闭后的ch赋值为nil, 则select读取则会阻塞
			}
			//default:
			//	fmt.Printf("%v, 没读取到信息，进入default\n", time.Now().Format(FMAT))
			//	time.Sleep(500 * time.Millisecond)
		}
	}
}
