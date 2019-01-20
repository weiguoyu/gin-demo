package routers

import (
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/huyujie/gin-demo/logger"
	"github.com/huyujie/gin-demo/routers/api/v1"
	"net/http"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//api v1
	apiv1 := r.Group("/api/v1")
	apiv1.Use(ginzap.Ginzap(logger.Logger, time.RFC3339, false))
	{
		apiv1.GET("/hello", v1.GetHello)
	}

	return r
}
