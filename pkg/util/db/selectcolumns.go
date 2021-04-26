package db

import (
	"gin-demo/pkg/constants"
	"gin-demo/pkg/models/dbschema"
	"gin-demo/pkg/util/modelutil"
	"gin-demo/pkg/util/stringutil"
	"github.com/gocraft/dbr/v2"
)

func getColumnsFromStructWithTag(s interface{}, tableAlias string) []string {
	var result []string
	pkgPath := modelutil.GetModelPkgPath(s)
	modelInfos, ok := dbschema.SCHEMA_COLLECTION[pkgPath]
	if !ok {
		modelInfos = modelutil.ParseModelWithTag(s)
		dbschema.SCHEMA_COLLECTION[pkgPath] = modelInfos
	}
	all := modelInfos[constants.ALL]
	exclude := modelInfos[constants.SELECTEXCLUDE]
	columns := stringutil.Diff(all, exclude)
	for _, column := range columns {
		name := dbr.NameMapping(column)
		prefix := ""
		if tableAlias != "" {
			prefix = tableAlias + "."
		}
		result = append(result, prefix+"`"+name+"`")
	}
	return result
}

func GetColumnsFromStructWithTag(s interface{}) []string {
	return getColumnsFromStructWithTag(s, "")
}

func GetColumnsFromStructWithTagAndAlias(s interface{}, tableAlias string) []string {
	return getColumnsFromStructWithTag(s, tableAlias)
}
