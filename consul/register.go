package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/testProject/consul/discovery"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main()  {
	consulAddr := flag.String("consul.addr", "127.0.0.1", "consul address")
	consulPort := flag.Int("consul.port", 8500, "consul port")

	servicePort := flag.Int("service.port", 12310, "service port")
	serviceName := flag.String("service.name", "instance", "service name")
	serviceAddr := flag.String("service.addr", "127.0.0.1", "service addr")

	flag.Parse()

	ctx, _ := context.WithCancel(context.Background())
	instanceId := *serviceName + "-" + strings.Replace(uuid.New().String(),"-","",-1)
	client := discovery.NewDiscoveryClient(*consulAddr, *consulPort)

	//将服务注册到consul
	err := client.Register(discovery.ServiceInfo{
		Ctx:            ctx,
		ServiceName:    *serviceName,
		InstanceId:     instanceId,
		HealthCheckUrl: "/health",
		InstanceHost:   *serviceAddr,
		InstancePort:   *servicePort,
		Meta:           nil,
		Weights:        nil,
	})
	if err != nil {
		log.Fatalln(err)
	}

	//开启http服务，并注册handle
	go func() {
		http.HandleFunc("/health", checkHealth)
		http.ListenAndServe(fmt.Sprintf("%s:%d", *serviceAddr, *servicePort), nil)
	}()

	//监控系统信号，等待ctrl+c系统信号通知服务关闭
	c := make(chan os.Signal, 1)
	go func() {
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	}()
	log.Printf("exit %s", <-c)
	client.Deregister(discovery.ServiceInfo{
		Ctx:            ctx,
		InstanceId:     instanceId,
	})
	log.Printf("Deregister service %s", instanceId)
}

func checkHealth(w http.ResponseWriter, _ *http.Request)  {
	io.WriteString(w, "success")
}