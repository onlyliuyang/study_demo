package main

import (
	"fmt"
	"github.com/bruceshao/lockfree/lockfree"
	"sync"
	"sync/atomic"
	"time"
)

type longEventHandler[T uint64] struct {
}

func (h *longEventHandler[T]) OnEvent(v uint64) {
	fmt.Printf("value = %v\n", v)
}

func main() {
	var (
		goSize    = 10
		sizePerGo = 10
		counter   = uint64(0)
	)

	//创建事件处理器
	eh := &longEventHandler[uint64]{}

	//创建消费端串行处理的Disruptor
	disruptor := lockfree.NewSerialDisruptor[uint64](1024*1024, eh, &lockfree.SchedWaitStrategy{})

	//启动disruptor
	if err := disruptor.Start(); err != nil {
		panic(err)
	}

	//获取生产对象
	producer := disruptor.Producer()
	var wg sync.WaitGroup
	wg.Add(goSize)
	for i := 0; i < goSize; i++ {
		go func() {
			for j := 0; j < sizePerGo; j++ {
				x := atomic.AddUint64(&counter, 1)
				err := producer.Write(x)
				if err != nil {
					panic(err)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("-------- write complete -------")

	time.Sleep(time.Second * 2)
	disruptor.Close()
}
