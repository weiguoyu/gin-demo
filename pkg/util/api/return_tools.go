// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package api

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// Normal Ret code.
var NormalRetCode = 0

// Api normal response body struct.
type SuccessRespJson struct {
	Data    interface{} `json:"data,required"`
	RetCode int         `json:"ret_code,required"`
}

// Api normal response body struct.
type SuccessItemsRespJson struct {
	Data       interface{} `json:"data,required"`
	RetCode    int         `json:"ret_code,required"`
	TotalCount uint32      `json:"total_count"`
}

// Api error response body struct.
type ErrRespJson struct {
	RetCode int         `json:"ret_code,required"`
	Msg     interface{} `json:"message,omitempty"`
}

// Api normal return: http status code always keep 200, ret code defined by parameter.
func ReturnSuccess(c *gin.Context, data interface{}) {
	var body = SuccessRespJson{
		Data:    data,
		RetCode: NormalRetCode,
	}
	c.JSON(StatusOk, body)
}

// Api normal return: http status code always keep 200, Ret code defined by parameter.
// Unlike ReturnSuccess, the Data field should be a list, TotalCount field is required.
func ReturnItems(c *gin.Context, dataItems interface{}, totalCount uint32) {
	var body = SuccessItemsRespJson{
		Data:       dataItems,
		TotalCount: totalCount,
		RetCode:    NormalRetCode,
	}
	c.JSON(StatusOk, body)
}

// Api error return: ret code defined by parameter, Msg filed is required.
// Default http status code is 500, can by defined in args[0].
func ReturnError(c *gin.Context, retCode int, errMessage error, args ...int) {
	var body = ErrRespJson{
		Msg:     errMessage.Error(),
		RetCode: retCode,
	}
	httpStatusCode := StatusInternalServerError
	if len(args) >= 1 {
		httpStatusCode = args[0]
	}
	c.JSON(httpStatusCode, body)
}

func Error(err string) error {
	return errors.New(err)
}
