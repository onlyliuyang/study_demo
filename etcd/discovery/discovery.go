package discovery

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"strings"
	"sync"
)

type EtcdDiscoveryConfig struct {
	Client     *clientv3.Client
	Prefix     string
	Key        string
	Value      string
	TTLSeconds int
	Callbacks  DiscoveryCallbacks
}

type DiscoveryCallbacks struct {
	OnStartDiscovering func(services []Service)
	OnServiceChanged   func(service []Service, event DiscoveryEvent)
	OnStopDiscovering  func()
}

type EtcdDiscovery struct {
	EtcdDiscoveryConfig
	myKey        string
	session      *concurrency.Session
	watchContext context.Context
	watchCancel  context.CancelFunc
	services     map[string]string
	mu           sync.RWMutex
}

type Service struct {
	Path string
	Name string
	Val  string
}

type DiscoveryEvent struct {
	Type string
	Service
}

const (
	PutEvent    = "PUT"
	DeleteEvent = "DELETE"
)

func New(config EtcdDiscoveryConfig) (*EtcdDiscovery, error) {
	session, err := concurrency.NewSession(config.Client, concurrency.WithTTL(config.TTLSeconds))
	if err != nil {
		return nil, err
	}

	config.Prefix = strings.TrimSuffix(config.Prefix, "/") + "/"
	return &EtcdDiscovery{
		EtcdDiscoveryConfig: config,
		myKey:               config.Prefix + config.Key,
		session:             session,
		watchContext:        nil,
		watchCancel:         nil,
		services:            make(map[string]string),
		mu:                  sync.RWMutex{},
	}, nil
}

//注册
func (d *EtcdDiscovery) Register(ctx context.Context) error {
	lease := d.session.Lease()
	_, err := d.Client.Put(ctx, d.myKey, d.Value, clientv3.WithLease(lease))
	return err
}

//取消注册
func (d *EtcdDiscovery) UnRegister(ctx context.Context) error {
	_, err := d.Client.Delete(ctx, d.myKey)
	return err
}

//关闭
func (d *EtcdDiscovery) Close() error {
	if d.watchCancel != nil {
		d.watchCancel()
	}
	return d.session.Close()
}

//监听
func (d *EtcdDiscovery) Watch(ctx context.Context) error {
	d.watchContext, d.watchCancel = context.WithCancel(ctx)
	resp, err := d.Client.Get(d.watchContext, d.Prefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}

	services := make(map[string]string)
	for _, kv := range resp.Kvs {
		services[string(kv.Key)] = string(kv.Value)
	}
	d.setServices(services)

	if d.Callbacks.OnStartDiscovering != nil {
		d.Callbacks.OnStartDiscovering(d.ListServices())
	}

	defer func() {
		if d.Callbacks.OnStopDiscovering != nil {
			d.Callbacks.OnStopDiscovering()
		}
	}()

	defer d.watchCancel()

	ch := d.Client.Watch(d.watchContext, d.Prefix, clientv3.WithPrefix())
	for {
		select {
		case <-d.watchContext.Done():
			return nil
		case wr, ok := <-ch:
			if !ok {
				return fmt.Errorf("watch closed")
			}

			if wr.Err() != nil {
				return wr.Err()
			}

			for _, ev := range wr.Events {
				key, val := string(ev.Kv.Key), string(ev.Kv.Value)
				switch ev.Type {
				case mvccpb.PUT:
					d.addService(key, val)
				case mvccpb.DELETE:
					d.delService(key)
				}

				if d.Callbacks.OnServiceChanged != nil {
					event := DiscoveryEvent{
						Type:    mvccpb.Event_EventType_name[int32(ev.Type)],
						Service: d.serviceFromKv(key, val),
					}
					d.Callbacks.OnServiceChanged(d.ListServices(), event)
				}
			}
		}
	}
}

func (d *EtcdDiscovery) setServices(services map[string]string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.services = services
}

func (d *EtcdDiscovery) addService(key, val string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.services[key] = val
}

func (d *EtcdDiscovery) delService(key string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	delete(d.services, key)
}

func (d *EtcdDiscovery) serviceFromKv(k, v string) Service {
	return Service{
		Path: k,
		Name: strings.TrimSuffix(k, d.Prefix),
		Val:  v,
	}
}

func (d *EtcdDiscovery) ListServices() []Service {
	d.mu.RLock()
	defer d.mu.RUnlock()

	iterms := make([]Service, 0, len(d.services))
	for k, v := range d.services {
		iterms = append(iterms, d.serviceFromKv(k, v))
	}
	return iterms
}
