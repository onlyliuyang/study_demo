package main

import (
	"fmt"
	"github.com/testProject/liuyang/autoLoadConfig/autoLoadConfig"
	"time"
)

//https://cloud.tencent.com/developer/article/1165960

func initConfig(file string) {
	//打开配置文件
	conf, err := autoLoadConfig.NewConfig(file)
	if err != nil {
		fmt.Printf("read config file err: %v\n", err)
		return
	}

	//添加观察者
	conf.AddObserver(autoLoadConfig.AppConfigManager)

	//第一次读取配置文件
	var appConfig autoLoadConfig.AppConfig
	appConfig.Hostname, err = conf.GetString("hostname")
	if err != nil {
		fmt.Printf("get hostname err: %v\n", err)
		return
	}
	fmt.Println("Hostname: ", appConfig.Hostname)

	appConfig.KafkaPort, err = conf.GetInt("kafkaPort")
	if err != nil {
		fmt.Printf("get kafka port err:%v\n", err)
		return
	}
	fmt.Println("Kafka port :", appConfig.KafkaPort)

	//把取到的值存储到atomic.Value
	autoLoadConfig.AppConfigManager.Config.Store(&appConfig)
	fmt.Println("first load success.")
}

func run() {
	for {
		appConfig := autoLoadConfig.AppConfigManager.Config.Load().(*autoLoadConfig.AppConfig)
		fmt.Println("Hostname:", appConfig.Hostname)
		fmt.Println("kafkaPort:", appConfig.KafkaPort)
		fmt.Printf("%v\n", "--------------------")
		time.Sleep(5 * time.Second)
	}
}

func main() {
	configFile := "./autoLoadConfig/conf/app.cfg"
	initConfig(configFile)
	run()
}
