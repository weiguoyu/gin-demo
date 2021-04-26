package dbschema

import (
	"fmt"
	"strings"
)

var SCHEMA_COLLECTION = make(map[string]map[string][]string)

var TablePrifixNameMap = map[string]string{}

var TablePrimaryMap = map[string]string{}

func getPrefix(tableId string) string {
	return strings.Split(tableId, "_")[0] + "_"
}

func GetTableName(tableId string) (string, error) {
	prefix := getPrefix(tableId)
	if tableName, ok := TablePrifixNameMap[prefix]; ok {
		return tableName, nil
	} else {
		return "", fmt.Errorf("获取资源类型失败")
	}
}

var DeleteCheckerMap = map[string]map[string]string{}
