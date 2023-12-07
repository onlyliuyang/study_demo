package autoLoadConfig

import (
	"fmt"
	"sync/atomic"
)

type AppConfig struct {
	Hostname  string
	Port      int
	KafkaHost string
	KafkaPort int
}

type AppConfigMgr struct {
	Config atomic.Value
}

var AppConfigManager = &AppConfigMgr{}

func (a *AppConfigMgr) Callback(config *Config) {
	appConfig := &AppConfig{}
	hostname, err := config.GetString("hostname")
	if err != nil {
		fmt.Printf("get hostname err:%v\n", err)
		return
	}

	appConfig.Hostname = hostname

	kafkaPort, err := config.GetInt("kafkaPort")
	if err != nil {
		fmt.Printf("get kafka port err:%v\n", err)
		return
	}
	appConfig.KafkaPort = kafkaPort

	AppConfigManager.Config.Store(appConfig)
}
