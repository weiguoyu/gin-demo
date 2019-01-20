package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/huyujie/gin-demo/logger"
	"net/http"
)

func GetHello(c *gin.Context) {
	logger.Info("roger-info")
	name := "roger-infof"
	logger.Infof("err:%s", name)
	c.String(http.StatusOK, "Hello World!")
}
