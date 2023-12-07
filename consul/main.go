package main

import (
	"github.com/hashicorp/consul/api"
	"log"
	"net"
)

func main()  {
	//使用默认配置创建consul客户端
	consulClient, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}

	//注册服务
	consulClient.Agent().ServiceRegister(&api.AgentServiceRegistration{
		Kind:              "",
		ID:                "MyService",
		Name:              "MyService",
		Tags:              nil,
		Port:              5050,
		Address:           "127.0.0.1",
		SocketPath:        "",
		TaggedAddresses:   nil,
		EnableTagOverride: false,
		Meta:              nil,
		Weights:           nil,
		Check:             &api.AgentServiceCheck{
			CheckID:                        "MyService",
			Name:                           "",
			Args:                           nil,
			DockerContainerID:              "",
			Shell:                          "",
			Interval:                       "10s",
			Timeout:                        "1s",
			TTL:                            "",
			HTTP:                           "",
			Header:                         nil,
			Method:                         "",
			Body:                           "",
			TCP:                            "127.0.0.1:5050",
			UDP:                            "",
			Status:                         "",
			Notes:                          "",
			TLSServerName:                  "",
			TLSSkipVerify:                  false,
			GRPC:                           "",
			GRPCUseTLS:                     false,
			H2PING:                         "",
			H2PingUseTLS:                   false,
			AliasNode:                      "",
			AliasService:                   "",
			SuccessBeforePassing:           0,
			FailuresBeforeWarning:          0,
			FailuresBeforeCritical:         0,
			DeregisterCriticalServiceAfter: "",
		},
		Checks:            nil,
		Proxy:             nil,
		Connect:           nil,
		Namespace:         "",
		Partition:         "",
	})

	//运行完注销服务
	defer consulClient.Agent().ServiceDeregister("MyService")
	ler, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ler.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			log.Printf("IP: %s connected", conn.RemoteAddr().String())
		}()
	}
}