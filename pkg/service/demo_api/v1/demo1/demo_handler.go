// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package demo1

import (
	"gin-demo/pkg/logger"
	. "gin-demo/pkg/util/api"
	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {

	output := "Hello World!"

	logger.Infof("get projects success")

	ReturnSuccess(c, output)
	return
}
