package main

import (
	"context"
	"fmt"
	"github.com/testProject/etcd/discovery"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"sync"
	"time"
)

func main()  {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:            []string{"127.0.0.1:2379"},
		AutoSyncInterval:     0,
		DialTimeout:          time.Second * 5,
		DialKeepAliveTime:    0,
		DialKeepAliveTimeout: 0,
		MaxCallSendMsgSize:   0,
		MaxCallRecvMsgSize:   0,
		TLS:                  nil,
		Username:             "",
		Password:             "",
		RejectOldCluster:     false,
		DialOptions:          nil,
		LogConfig:            nil,
		Context:              nil,
		PermitWithoutStream:  false,
	})

	if err != nil {
		fmt.Println("connect fail, err: ", err)
		return
	}

	defer client.Close()

	worker := func(i int, run bool) {
		id := fmt.Sprintf("worker-%d", i)
		val := fmt.Sprintf("127.0.0.%d", i)

		sd, err := discovery.New(discovery.EtcdDiscoveryConfig{
			Client:     client,
			Prefix:     "/services",
			Key:        id,
			Value:      val,
			TTLSeconds: 2,
			Callbacks:  discovery.DiscoveryCallbacks{
				OnStartDiscovering: func(services []discovery.Service) {
					log.Printf("[%s], onstarted, services: %v", id, services)
				},
				OnServiceChanged: func(service []discovery.Service, event discovery.DiscoveryEvent) {
					log.Printf("[%s], onchanged, services: %v, event: %v", id, service, event)
				},
				OnStopDiscovering: func() {
					log.Printf("[%s], onstoped", id)
				},
			},
		})

		if err != nil {
			log.Fatalf("failed to create service etcdiscovery: %v", err)
		}
		defer sd.Close()

		if !run {
			if err = sd.UnRegister(context.Background()); err != nil {
				log.Fatalf("failed to unregister service [%s]: %v", id, err)
			}
			return
		}

		if err := sd.Register(context.Background()); err != nil {
			log.Fatalf("failed to register service [%s]: %v", id, err)
		}

		if err := sd.Watch(context.Background()); err != nil {
			log.Fatalf("failed to watch service [%s]: %v", id, err)
		}
	}

	wg := sync.WaitGroup{}
	for i:=0; i<10; i++ {
		id := i
		wg.Add(1)
		go func(id int) {
			worker(id, true)
			defer wg.Done()
		}(id)
	}

	go func() {
		time.Sleep(2 * time.Second)
		worker(3, true)
	}()

	go func() {
		time.Sleep(5 * time.Second)
		worker(4, false)
	}()

	go func() {
		time.Sleep(10 * time.Second)
		worker(5, false)
	}()
	wg.Wait()
}
