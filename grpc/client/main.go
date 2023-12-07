package main

import (
	"context"
	"fmt"
	pt "github.com/protocol"
	"google.golang.org/grpc"
)

const (
	host = "127.0.0.1:18887"
)

func main()  {
	//客户端连接信息
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接服务器失败", err)
		return
	}
	defer conn.Close()

	//获得grpc句柄
	c := pt.NewHelloServerClient(conn)

	//远程调用
	rl, err := c.SayHello(context.Background(), &pt.HelloRequest{Name:"Scott"})
	if err != nil {
		fmt.Println("cloud not get Hello server...")
		return
	}

	fmt.Println("HelloServer resp: ", rl.Message)

	//远程调用GetHelloMsg接口
	r2, err := c.GetHelloMsg(context.Background(), &pt.HelloRequest{Name:"Scott"})
	if err != nil {
		fmt.Println("could not get hello msg...")
		return
	}
	fmt.Println("HelloServer resp: ", r2.Msg)
}
