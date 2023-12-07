package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"time"
)

var P *ants.PoolWithFunc

func init()  {
	P, _ = ants.NewPoolWithFunc(10000, func(i interface{}) {
		err := HandleJob(i)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
}

//HandleJob 异步注册任务
func HandleJobs(c time.Time)  {
	copyC := c
	go func() {
		err := P.Invoke(copyC)
		if err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("异步任务执行成功！", copyC)
}

//HandleJob 真正要执行的任务
func HandleJob(c interface{}) (err error) {
	t := c.(time.Time)
	time.Sleep(time.Second * 3)
	fmt.Println("执行任务时间：", t.Format("2006-01-02 15:04:05"))
	return
}

func main()  {
	defer P.Release()

	var wg sync.WaitGroup
	ch := make(chan struct{}, 200)
	for i:=0; i<10000; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			HandleJobs(time.Now())
			<-ch
		}(i)
	}
	wg.Wait()
}