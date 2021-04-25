// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package demo_api

import (
	"fmt"
	"gin-demo/pkg/config"
	demo1 "gin-demo/pkg/service/demo_api/v1/demo1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartGin(cfg *config.Config, r *gin.Engine) {
	r.Use()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, I'm gin-demo!")
	})
	demo1.Routes(r)
	r.Run(fmt.Sprintf(":%s", cfg.Api.Port))
}
