package main

import (
	"fmt"
	"github.com/bruceshao/lockfree"
	"sync"
	"sync/atomic"
	"time"
)

var (
	goSize   = 10000
	sizePero = 10000
	total    = goSize * sizePero
)

type eventHandler[T uint64] struct {
	signal   chan struct{}
	gcounter uint64
	now      time.Time
}

func (h *eventHandler[T]) OnEvent(v uint64) {
	cur := atomic.AddUint64(&h.gcounter, 1)
	if cur == uint64(total) {
		fmt.Printf("evnetHandler has been consumed already, read count: %v, time cose: %v\n", total, time.Since(h.now))
		close(h.signal)
		return
	}

	if cur%10000000 == 0 {
		fmt.Printf("evnet handler consume :%v\n", cur)
	}
}

func (h *eventHandler[T]) wait() {
	<-h.signal
}

func main() {
	//lockfree计时
	now := time.Now()

	//创建事件处理器
	handler := &eventHandler[uint64]{
		signal: make(chan struct{}, 0),
		now:    now,
	}

	//创建消费端串行处理的lockfree
	lf := lockfree.NewLockfree[uint64](1024*1024, handler, lockfree.NewSleepBlockStrategy(time.Millisecond))

	//启动lockfree
	if err := lf.Start(); err != nil {
		panic(err)
	}

	//获取生产者对象
	producer := lf.Producer()

	//并发写入
	var wg sync.WaitGroup
	wg.Add(goSize)
	for i := 0; i < goSize; i++ {
		go func(i int) {
			for j := 0; j < sizePero; j++ {
				err := producer.Write(uint64(i*sizePero + j + 1))
				if err != nil {
					panic(err)
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Printf("producer has been writed, write count:%v, time cost:%v\n", total, time.Since(now).String())
	handler.wait()

	//关闭LockFree
	lf.Close()
}
