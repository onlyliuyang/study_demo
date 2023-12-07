package discovery

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//https://blog.csdn.net/skh2015java/article/details/111321515

//服务实例结构体
type InstanceInfo struct {
	ID                string            `json:"id"`
	Service           string            `json:"service,omitempty"`
	Name              string            `json:"name"`
	Tags              []string          `json:"tags"`
	Address           string            `json:"address"`
	Port              int               `json:"port"`
	Meta              map[string]string `json:"meta,omitempty"`
	EnableTagOverride bool              `json:"enable_tag_override"`
	Check             Check             `json:"check,omitempty"`
	Weights           Weights           `json:"weights,omitempty"`
	CurWeight         int               `json:"cur_weight,omitempty"`
}

//健康检查相关配置
type Check struct {
	//DeregisterCritcalServiceAfter string   `json:"deregister_critcal_service_after"`
	Args                          []string `json:"args,omitempty"`
	HTTP                          string   `json:"http"`
	Interval                      string   `json:"interval,omitempty"`
	TTL                           string   `json:"ttl,omitempty"`
}

//权重
type Weights struct {
	Passing int `json:"passing"`
	Warning int `json:"warning"`
}

//服务发现的客户端
type DiscoveryClient struct {
	host string
	port int
}

//服务信息
type ServiceInfo struct {
	Ctx            context.Context
	ServiceName    string
	InstanceId     string
	HealthCheckUrl string
	InstanceHost   string
	InstancePort   int
	Meta           map[string]string
	Weights        *Weights
}

func NewDiscoveryClient(host string, port int) *DiscoveryClient {
	return &DiscoveryClient{
		host: host,
		port: port,
	}
}

//服务注册
func (c *DiscoveryClient) Register(serviceInfo ServiceInfo) error {
	instanceInfo := &InstanceInfo{
		ID:                serviceInfo.InstanceId,
		Name:              serviceInfo.ServiceName,
		Tags:              nil,
		Address:           serviceInfo.InstanceHost,
		Port:              serviceInfo.InstancePort,
		Meta:              serviceInfo.Meta,
		EnableTagOverride: false,
		Check: Check{
			//DeregisterCritcalServiceAfter: "30s",
			Args:                          nil,
			HTTP:                          fmt.Sprintf("http://%s:%d%s", serviceInfo.InstanceHost, serviceInfo.InstancePort, serviceInfo.HealthCheckUrl),
			Interval:                      "15s",
			TTL:                           "",
		},
	}

	if serviceInfo.Weights != nil {
		instanceInfo.Weights = *serviceInfo.Weights
	} else {
		instanceInfo.Weights = Weights{
			Passing: 10,
			Warning: 1,
		}
	}

	byteData, err := json.Marshal(instanceInfo)
	if err != nil {
		log.Printf("json format err: %s", err)
		return err
	}

	url := fmt.Sprintf("http://%s:%d%s", c.host, c.port, "/v1/agent/service/register")
	req, err := http.NewRequest("PUT", url, bytes.NewReader(byteData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	client.Timeout = time.Second * 2
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Register service err: %s", err)
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("register service http request errCode : %v", resp.StatusCode)
		return fmt.Errorf("register service http request errCode : %v", resp.StatusCode)
	}

	log.Println("register service success")
	return nil
}

//服务注销
func (c *DiscoveryClient) Deregister(serviceInfo ServiceInfo) error {
	url := fmt.Sprintf("http://%s:%d%s/%s", c.host, c.port, "/v1/agent/service/deregister", serviceInfo.InstanceId)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		log.Printf("req format err: %s", err)
		return err
	}

	client := http.Client{}
	client.Timeout = time.Second * 2
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Register service err: %s", err)
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("deregister service http request errCode : %v", resp.StatusCode)
		return fmt.Errorf("deregister service http request errCode : %v", resp.StatusCode)
	}

	log.Println("deregister service success")
	return nil
}

//服务发现
func (c *DiscoveryClient) DiscoveryServices(serviceInfo ServiceInfo) ([]*InstanceInfo, error) {
	url := fmt.Sprintf("http://%s:%d%s/%s", c.host, c.port, "/v1/health/service", serviceInfo.ServiceName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("req format err: %s", err)
		return nil, err
	}

	client := http.Client{}
	client.Timeout = time.Second * 2
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("discover service err : %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("discover service http request errCode : %v", resp.StatusCode)
		return nil, fmt.Errorf("discover service http request errCode : %v", resp.StatusCode)
	}

	var serviceList []struct {
		Service InstanceInfo `json:"service"`
	}
	err = json.NewDecoder(resp.Body).Decode(&serviceList)
	if err != nil {
		log.Printf("format service info err : %s", err)
		return nil, err
	}

	instances := make([]*InstanceInfo, len(serviceList))
	for i := 0; i < len(serviceList); i++ {
		instances[i] = &serviceList[i].Service
	}
	return instances, nil
}
