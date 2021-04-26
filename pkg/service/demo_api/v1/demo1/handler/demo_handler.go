// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package handler

import (
	"fmt"
	"gin-demo/pkg/logger"
	"gin-demo/pkg/models/project"
	rc "gin-demo/pkg/service/demo_api/v1/demo1/resource_controller"
	. "gin-demo/pkg/util/api"
	"gin-demo/pkg/util/paramsdefault"
	validate "gin-demo/pkg/util/validation"
	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {

	inputs := &project.APIGetProjectsInput{}

	//params check
	if err := paramsdefault.Set(inputs); err != nil {
		logger.Errorf("get default value of parameter error, %v", err)
		ReturnError(c, InvalidDefaultValue, ErrMsg(ErrMsgInvalidDefault))
		return
	}

	_, err := validate.Bind(&inputs, c)
	if err != nil {
		logger.Errorf("validate bind error, %v", err)
		ReturnError(c, InvalidRequestFormat, fmt.Errorf("%v", err))
		return
	}

	//err = validate.Validate.Struct(bind)
	//if err != nil {
	//	if validatorErr, ok := err.(*validator.InvalidValidationError); ok {
	//		logger.Errorf("input validate error, %v", validatorErr)
	//		ReturnError(c, InvalidRequestFormat, fmt.Errorf("%v", validatorErr))
	//		return
	//	}
	//	for _, err := range err.(validator.ValidationErrors) {
	//		logger.Errorf("input validator error, %v", err.Translate(validate.Trans))
	//		ReturnError(c, InvalidRequestFormat, fmt.Errorf(err.Translate(validate.Trans)))
	//		return
	//	}
	//}
	total, outputs, err := rc.GetProjects(*inputs)
	if err != nil {
		logger.Errorf("get projects from mysql error, %v", err)
		ReturnError(c, InternalError, ErrMsg(ErrMsgInternalErr))
		return
	}

	ReturnItems(c, outputs, total)
	return
}
