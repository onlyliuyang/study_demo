package main

import (
	"sync/atomic"
	"time"
)

func loadConfig() map[string]string {
	//从线程或数据库中读取配置，然后放在map中并返回
	return make(map[string]string)
}

func requests() chan int {
	//将从外界中接收到的请求放入chan中
	return make(chan int)
}

func main() {
	//config 变量用来存放服务的配置信息
	var config atomic.Value
	config.Store(loadConfig())

	go func() {
		for {
			//每10秒钟定时的拉取最新的配置信息，并且更新到config变量中
			time.Sleep(10 * time.Second)
			config.Store(loadConfig())
		}
	}()

	//创建工作线程，每个工作线程都会根据它所读取到的最新的配置来处理请求
	for i := 0; i < 100; i++ {
		go func() {
			for r := range requests() {
				// 对应于取值操作 c := config
				// 由于load()返回的是一个interface类型，所以先强制转一下
				c := config.Load().(map[string]string)
				_, _ = r, c
			}
		}()
	}
}
