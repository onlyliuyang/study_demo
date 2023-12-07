package main

import (
	"go.etcd.io/etcd/clientv3"
	"context"
	"fmt"
	"time"
)

const HOSTS = ":2379"

func main()  {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:            []string{"127.0.0.1:2379"},
		DialTimeout:          5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect to etcd failed, err: ", err)
		return
	}

	fmt.Println("connect to etcd success")
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "name", "liuyang")
	cancel()

	if err != nil {
		fmt.Println("put to etcd failed, err: ", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "name")

	cancel()

	if err != nil {
		fmt.Println("put to etcd failed, err: ", err)
		return
	}

	for _, v := range resp.Kvs {
		fmt.Println(v)
	}

	//监听一下
	fmt.Println("Listen: ")
	ch := cli.Watch(context.Background(), "name")
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Println("Type:%v Key:%v value:%v\n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
		}
	}
}