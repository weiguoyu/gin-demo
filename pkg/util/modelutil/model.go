package modelutil

import (
	"reflect"

	"gin-demo/pkg/constants"
)

func GetModelPkgPath(model interface{}) string {
	modelVal := reflect.TypeOf(model)
	if k := modelVal.Kind(); k == reflect.Ptr {
		modelVal = modelVal.Elem()
	}
	if modelVal.Kind() != reflect.Struct {
		return ""
	}
	return modelVal.PkgPath() + "/" + modelVal.Name()
}

func ParseModelWithTag(model interface{}) map[string][]string {
	m := make(map[string][]string)
	modelVal := reflect.TypeOf(model)
	if k := modelVal.Kind(); k == reflect.Ptr {
		modelVal = modelVal.Elem()
	}
	if modelVal.Kind() != reflect.Struct {
		return m
	}
	for i := 0; i < modelVal.NumField(); i++ {
		field := modelVal.Field(i)
		for _, tag := range constants.TAG_SET {
			v, ok := field.Tag.Lookup(tag.TagName)
			if !ok || v != tag.TagValue {
				continue
			}
			if _, ok := m[tag.TagName]; ok {
				m[tag.TagName] = append(m[tag.TagName], field.Name)
			} else {
				m[tag.TagName] = []string{field.Name}
			}
		}
		if _, ok := m[constants.ALL]; ok {
			m[constants.ALL] = append(m[constants.ALL], field.Name)
		} else {
			m[constants.ALL] = []string{field.Name}
		}
	}
	return m
}
