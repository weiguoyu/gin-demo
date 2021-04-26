package structutil

import (
	"github.com/asaskevich/govalidator"
	"github.com/fatih/structs"
	"strings"
)

func GetColumnsFromStruct(s interface{}) []string {
	names := structs.Names(s)
	for i, name := range names {
		names[i] = govalidator.CamelCaseToUnderscore(name)
	}
	return names
}

func GetFieldName(field *structs.Field) string {
	return GetFieldNameWithTag(field, "json")
}

func GetFieldNameWithTag(field *structs.Field, tagName string) string {
	tag := field.Tag(tagName)
	t := strings.Split(tag, ",")
	if len(t) == 0 {
		return "-"
	}
	return t[0]
}
