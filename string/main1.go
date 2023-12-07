package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
)

func main() {
	urls := []string{
		"https://www.baidu.com",
		"https://www.sina.com.cn",
		"https://www.sohu.com",
	}
	num := 3
	for i := 0; i < num; i++ {
		fmt.Println(urls[i], i)
		resp, _ := http.Get(urls[i])
		_, _ = ioutil.ReadAll(resp.Body)
	}
	fmt.Printf("此时groutine个数= %d\n", runtime.NumGoroutine())
}
