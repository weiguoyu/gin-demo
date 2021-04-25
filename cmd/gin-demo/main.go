// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package gin_demo

import (
	"flag"
	"gin-demo/pkg/config"
	"gin-demo/pkg/logger"
	"gin-demo/pkg/service/demo_api"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//init config
	cfg := flag.String("c", "/etc/gin-demo/config.yaml", "configuration file")
	flag.Parse()
	config.ParseConfig(*cfg)
	c := config.ReadConf()

	//init logger
	logger.Setup("git-demo")

	// init router
	gin.SetMode(c.Common.Gin.Mode)
	routes := gin.New()
	go demo_api.StartGin(c, routes)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		os.Exit(0)
	}()
	select {}

}
