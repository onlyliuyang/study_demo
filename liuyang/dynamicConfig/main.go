package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/testProject/liuyang/dynamicConfig/config"
	"net/http"
)

func main() {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		fmt.Println("Current redis host is: ", config.GlobalConfig.Get("service.redis.host"))
		fmt.Println(config.GlobalConfig.GetInt("service.redis.port"))
		context.JSON(http.StatusOK, gin.H{
			"message": "You are welcome",
			"config":  config.GlobalConfig,
		})
	})
	r.Run(":9292")
}
