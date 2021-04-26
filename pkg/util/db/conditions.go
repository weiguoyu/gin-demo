package db

import (
	"gin-demo/pkg/constants"
	"github.com/gocraft/dbr/v2"
)

func handlerCondition(tag, column, tableAlias string, value interface{}) dbr.Builder {
	switch tag {
	case constants.LTE:
		return Lte(getAliasColumn(tableAlias, column), value)
	case constants.GTE:
		return Gte(getAliasColumn(tableAlias, column), value)
	case constants.ALL:
		return Eq(getAliasColumn(tableAlias, column), value)
	}
	return nil
}
