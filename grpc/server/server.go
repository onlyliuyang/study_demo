package main

import (
	"context"
	pt "github.com/protocol"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	host = "127.0.0.1:18887"
)

//对象要和proto内定义的服务一样
type HelloServer struct {
	pt.UnimplementedHelloServerServer
}

func (s *HelloServer) SayHello(ctx context.Context, in *pt.HelloRequest) (*pt.HelloReplay, error) {
	return &pt.HelloReplay{Message:"Hello" + in.GetName()}, nil
}

func (s *HelloServer) GetHelloMsg(ctx context.Context, in *pt.HelloRequest) (*pt.HelloMessage, error) {
	return &pt.HelloMessage{Msg:"this is from server!"}, nil
}

func main()  {
	ln, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}

	//创建一个grpc句柄
	srv := grpc.NewServer()
	//将server结构体注册到grpc服务中
	pt.RegisterHelloServerServer(srv, &HelloServer{})
	//监听grpc
	err = srv.Serve(ln)
	log.Println("Grpc Server Listen...")
	if err != nil {
		log.Fatal("网络启动异常")
	}
}