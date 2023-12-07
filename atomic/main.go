package main

import (
	"fmt"
	"net/url"
	"path"
	"sync"
	"sync/atomic"
)

func atomicAdd() {
	var data1, data2 int64
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&data1, 1)
			data2++
		}()
	}
	wg.Wait()
	fmt.Println("data1:", data1, " data2:", data2)
}

func main() {
	//atomicAdd()
	addrUrl := "https://yuwencdn.kaishustory.com/kstory/pangu/image/a700a018-2e64-4ae7-8314-c11dc84f76c4_info_w=750_h=460_s=50655.jpg"
	r, _ := url.Parse(addrUrl)
	fmt.Println(r.Query())
	fmt.Println(path.Base(addrUrl))
}
