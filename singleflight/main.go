package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"sync/atomic"
	"time"
)

var count int32

func main() {
	var wg sync.WaitGroup
	now := time.Now()
	sg := &singleflight.Group{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			//Getcontent(1)
			SingleGetcontent(sg, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("耗时: %s", time.Since(now))
}

func Getcontent(id int) (string, error) {
	atomic.AddInt32(&count, 1)
	time.Sleep(time.Duration(count) * time.Millisecond)
	return fmt.Sprintf("获取第%d个内容", id), nil
}

func SingleGetcontent(sg *singleflight.Group, id int) (string, error) {
	v, err, _ := sg.Do(fmt.Sprintf("%d", id), func() (interface{}, error) {
		return Getcontent(id)
	})
	//fmt.Println(ok)
	return v.(string), err
}
