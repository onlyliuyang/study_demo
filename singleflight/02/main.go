package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync/atomic"
	"time"
)

type Result string

func find(ctx context.Context, query string) (Result, error) {
	return Result(fmt.Sprintf("result for %q", query)), nil
}

func main() {
	var sg singleflight.Group
	const n = 50
	waited := int32(n)
	done := make(chan struct{})
	key := "http://www.baidu.com"

	for i := 0; i < n; i++ {
		go func(j int) {
			result := sg.DoChan(key, func() (interface{}, error) {
				ret, err := find(context.Background(), key)
				return ret, err
			})

			select {
			case a := <-result:
				fmt.Println(a)
			case <-time.After(time.Second * 5):
				sg.Forget(key)
			}

			if atomic.AddInt32(&waited, -1) == 0 {
				close(done)
			}

			//fmt.Printf("index: %d, val: %v, shared: %v\n", j, v, shared)
		}(i)
	}

	select {
	case <-done:
	case <-time.After(time.Second):
		fmt.Println("Do hangs")
	}
}
