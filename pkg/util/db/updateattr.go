package db

import (
	"gin-demo/pkg/util/reflectutil"
	"gin-demo/pkg/util/stringutil"
	"gin-demo/pkg/util/structutil"
	"github.com/fatih/structs"
	"reflect"
	"time"
)

func BuildUpdateAttributes(req Request, columns ...string) map[string]interface{} {
	attributes := make(map[string]interface{})
	for _, field := range structs.Fields(req) {
		if !field.IsExported() {
			continue
		}
		column := structutil.GetFieldName(field)
		f := field.Value()
		v := reflect.ValueOf(f)
		if !stringutil.Contains(columns, column) {
			continue
		}
		if !reflectutil.ValueIsNil(v) {
			switch v := f.(type) {
			case string, bool, int32, uint32, time.Time:
				attributes[column] = v
			default:
				attributes[column] = v
			}
		}
	}
	return attributes
}
