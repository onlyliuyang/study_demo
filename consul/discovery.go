package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/testProject/consul/discovery"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main()  {
	consulAddr := flag.String("consul.addr", "127.0.0.1", "consul.address")
	consulPort := flag.Int("consul.port", 8500, "consul.port")
	serviceName := flag.String("service.name", "instance", "service name")

	flag.Parse()

	ctx, _ := context.WithCancel(context.Background())
	client := discovery.NewDiscoveryClient(*consulAddr, *consulPort)

	c := make(chan os.Signal, 1)
	go func() {
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	}()

	//定时刷新
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <- ticker.C:
			instance, err := client.DiscoveryServices(discovery.ServiceInfo{
				Ctx:            ctx,
				ServiceName:    *serviceName,
				InstanceId:     "",
				HealthCheckUrl: "",
				InstanceHost:   "",
				InstancePort:   0,
				Meta:           nil,
				Weights:        nil,
			})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("get instance num", len(instance))
			for _, v := range instance {
				fmt.Printf("instance ID:%s, address:%s, port:%d\n", v.ID, v.Address, v.Port)
			}
		case <-c:
			log.Println("discovery service exit.")
			os.Exit(0)
		}
	}
}
