package api

import "fmt"

func ErrMsg(msg map[string]string, args ...interface{}) error {
	data := fmt.Errorf(msg[DefaultLang], args...)
	return data
}

var ErrMsgInvalidParameter = map[string]string{
	En:   "invalid parameter %s",
	ZhCn: "非法的参数：%s",
}

var ErrMsgInvalidDefault = map[string]string{
	En:   "invalid parameter default value",
	ZhCn: "非法的参数默认值",
}

var ErrMsgInternalErr = map[string]string{
	En:   "internal error",
	ZhCn: "内部错误",
}
