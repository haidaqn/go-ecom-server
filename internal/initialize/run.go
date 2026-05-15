package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count_total",
		Help: "Total number of ping requests.",
	},
)

func ping(c *gin.Context) {
	pingCounter.Inc() //1 2 3
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func Run() {
	LoadConfig()
	InitLogger()
	InitMySql()
	InitRedis()
	InitKafka()

	r := InitRouter()

	prometheus.MustRegister(pingCounter)

	r.GET("/ping/200", ping)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run(":8888")
}
