package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
)

func InitMetrics(r *gin.Engine) {
	recordMetrics()
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "Myapp_processed_ops_total",
		Help: "The Total number of processed events",
	})
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}
