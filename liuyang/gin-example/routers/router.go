package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	v1 := r.Group("/admin/")
	v1.GET("/info", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
		})
	})

	return r
}
