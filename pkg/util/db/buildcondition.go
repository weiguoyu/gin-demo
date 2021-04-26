package db

import (
	"gin-demo/pkg/constants"
	"gin-demo/pkg/logger"
	"gin-demo/pkg/models/dbschema"
	"gin-demo/pkg/util/modelutil"
	"gin-demo/pkg/util/stringutil"
	"github.com/fatih/structs"
	"github.com/gocraft/dbr/v2"
	"strings"
)

func getSearchWordColumn(model interface{}) []string {
	var result []string
	pkgPath := modelutil.GetModelPkgPath(model)
	modelInfos, ok := dbschema.SCHEMA_COLLECTION[pkgPath]
	if !ok {
		modelInfos = modelutil.ParseModelWithTag(model)
		dbschema.SCHEMA_COLLECTION[pkgPath] = modelInfos
	}
	fuzzys := modelInfos[constants.FUZZY]
	for _, fuzzy := range fuzzys {
		name := dbr.NameMapping(fuzzy)
		result = append(result, name)
	}
	return result
}

func getSearchFilterFromModel(model interface{}, value interface{}, tableAlias string, exclude ...string) dbr.Builder {
	if v, ok := value.(string); ok {
		var ops []dbr.Builder
		for _, column := range getSearchWordColumn(model) {
			if stringutil.Contains(exclude, column) {
				continue
			}
			column = getAliasColumn(tableAlias, column)
			// if column suffix is _id, must exact match
			if strings.HasSuffix(column, "_id") {
				ops = append(ops, Eq(column, v))
			} else {
				ops = append(ops, Like(column, v))
			}
		}
		if len(ops) == 0 {
			return nil
		}
		return Or(ops...)
	} else if value != nil {
		logger.Warnf("search_word [%+v] is not string", value)
	}
	return nil
}

func BuildFilterConditionsWithTagAndAlias(req Request, model interface{}, tableAlias string, exclude ...string) dbr.Builder {
	return buildFilterConditionsWithReq(false, req, model, tableAlias, exclude...)
}

func BuildFilterConditionsWithTag(req Request, model interface{}, exclude ...string) dbr.Builder {
	return buildFilterConditionsWithReq(false, req, model, "", exclude...)
}

func getAliasColumn(tableAlias, column string) string {
	if tableAlias != "" {
		return tableAlias + "." + column
	}
	return column
}

func buildFilterConditionsWithReq(withPrefix bool, req Request, model interface{}, tableAlias string, exclude ...string) dbr.Builder {
	var conditions []dbr.Builder
	appendCond := func(condition dbr.Builder) []dbr.Builder {
		if condition != nil {
			return append(conditions, condition)
		}
		return conditions
	}
	pkgPath := modelutil.GetModelPkgPath(model)
	modelInfos, ok := dbschema.SCHEMA_COLLECTION[pkgPath]
	if !ok {
		modelInfos = modelutil.ParseModelWithTag(model)
		dbschema.SCHEMA_COLLECTION[pkgPath] = modelInfos
	}
	var conditionTags []string
	for _, tag := range constants.TAG_SET {
		if tag.TagType == "condition" {
			conditionTags = append(conditionTags, tag.TagName)
		}
	}
	all := modelInfos[constants.ALL]
	for _, ct := range conditionTags {
		all = stringutil.Diff(all, modelInfos[ct])
	}
	for _, field := range structs.Fields(req) {
		if !field.IsExported() {
			continue
		}
		filedName := field.Name()
		column := dbr.NameMapping(field.Name())
		if stringutil.Contains(exclude, column) {
			continue
		}
		value := getReqValue(field.Value())
		if value == nil {
			continue
		}
		if column == constants.SearchWordColumnName {
			condition := getSearchFilterFromModel(model, value, tableAlias, exclude...)
			conditions = appendCond(condition)
		} else if column == constants.StartTimeColumnName {
			conditions = append(conditions, Gte(getAliasColumn(tableAlias, column), value))
		} else if column == constants.EndTimeColumnName {
			conditions = append(conditions, Lt(getAliasColumn(tableAlias, column), value))
		} else {
			if stringutil.Contains(all, filedName) {
				condition := handlerCondition(constants.ALL, column, tableAlias, value)
				conditions = appendCond(condition)
			}
			for _, ct := range conditionTags {
				if !stringutil.Contains(modelInfos[ct], filedName) {
					continue
				}
				condition := handlerCondition(ct, column, tableAlias, value)
				conditions = appendCond(condition)
			}
		}
	}
	if len(conditions) == 0 {
		return nil
	}
	return And(conditions...)
}
