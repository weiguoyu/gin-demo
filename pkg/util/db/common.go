// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package db

import (
	"context"
	"fmt"
	"gin-demo/pkg/constants"
	"gin-demo/pkg/logger"
	"gin-demo/pkg/models/dbschema"
	"gin-demo/pkg/util/reflectutil"
	"gin-demo/pkg/util/stringutil"
	"gin-demo/pkg/util/structutil"
	"github.com/fatih/structs"
	"github.com/gocraft/dbr/v2"
	"reflect"
	"strings"
)

type Request interface {
	Reset()
	String() string
	ProtoMessage()
	Descriptor() ([]byte, []int)
}

func getSearchFilter(tableName string, value interface{}, exclude ...string) dbr.Builder {
	if v, ok := value.(string); ok {
		var ops []dbr.Builder
		for _, column := range dbschema.SearchColumns[tableName] {
			if stringutil.Contains(exclude, column) {
				continue
			}
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

func GetReqValue(param interface{}) interface{} {
	return getReqValue(param)
}

func getReqValue(param interface{}) interface{} {
	switch value := param.(type) {
	case string:
		if value == "" {
			return nil
		}
		return value

	case []string:
		var values []string
		for _, v := range value {
			if v != "" {
				values = append(values, v)
			}
		}
		if len(values) == 0 {
			return nil
		}
		return values
	case []uint32:
		if len(value) == 0 {
			return nil
		}
		if len(value) == 1 {
			return value[0]
		}
		return value
	default:
		return nil
	}
}

func BuildFilterConditions(req interface{}, tableName string, exclude ...string) dbr.Builder {
	return buildFilterConditions(false, req, tableName, "", exclude...)
}

func GetDisplayColumns(displayColumns []string, wholeColumns []string) []string {
	if displayColumns == nil {
		return wholeColumns
	} else if len(displayColumns) == 0 {
		return nil
	} else {
		var newDisplayColumns []string
		for _, column := range displayColumns {
			if stringutil.Contains(wholeColumns, column) {
				newDisplayColumns = append(newDisplayColumns, column)
			}
		}
		return newDisplayColumns
	}
}

func BuildFilterConditionWithAliasPrefix(req interface{}, tableName, tableAlias string, exclude ...string) dbr.Builder {
	return buildFilterConditions(true, req, tableName, tableAlias, exclude...)
}

func BuildFilterConditionsWithPrefix(req interface{}, tableName string, exclude ...string) dbr.Builder {
	return buildFilterConditions(true, req, tableName, "", exclude...)
}

func buildFilterConditions(withPrefix bool, req interface{}, tableName, tableAlias string, exclude ...string) dbr.Builder {
	var conditions []dbr.Builder
	for _, field := range structs.Fields(req) {
		if !field.IsExported() {
			continue
		}
		column := structutil.GetFieldName(field)
		fieldValue := field.Value()

		if column == constants.SearchWordColumnName {
			if _, ok := dbschema.SearchColumns[tableName]; ok {
				value := getReqValue(fieldValue)
				condition := getSearchFilter(tableName, value, exclude...)
				if condition != nil {
					conditions = append(conditions, condition)
				}
			}
		} else if column == constants.StartTimeColumnName {
			value := getReqValue(fieldValue)
			if value != nil {
				conditions = append(conditions, Gte(column, value))
			}
		} else if column == constants.EndTimeColumnName {
			value := getReqValue(fieldValue)
			if value != nil {
				conditions = append(conditions, Lt(column, value))
			}
		} else {
			indexedColumns, ok := dbschema.IndexColumns[tableName]
			if ok && stringutil.Contains(indexedColumns, column) {
				value := getReqValue(fieldValue)
				if value != nil {
					key := column
					if withPrefix {
						if len(tableAlias) == 0 {
							key = tableName + "." + key
						} else {
							key = tableAlias + "." + key
						}
					}
					conditions = append(conditions, Eq(key, value))
				}
			} else if tGteColumns, ok := dbschema.TimeGteColumns[tableName]; ok && stringutil.Contains(tGteColumns, column) {
				if !reflectutil.ValueIsNil(reflect.ValueOf(fieldValue)) {
					conditions = append(conditions, Gte(column, getReqValue(fieldValue)))
				}
			} else if tLteColumns, ok := dbschema.TimeLteColumns[tableName]; ok && stringutil.Contains(tLteColumns, column) {
				if !reflectutil.ValueIsNil(reflect.ValueOf(fieldValue)) {
					conditions = append(conditions, Lte(column, getReqValue(fieldValue)))
				}
			}
		}
	}
	if len(conditions) == 0 {
		return nil
	}
	return And(conditions...)
}

func AddQueryJoinWithMap(query *SelectQuery, table, joinTable, primaryKey, keyField, valueField string, filterMap map[string][]string) *SelectQuery {
	var whereCondition []dbr.Builder
	for key, values := range filterMap {
		aliasTableName := fmt.Sprintf("table_label_%d", query.JoinCount)
		onCondition := fmt.Sprintf("%s.%s = %s.%s", aliasTableName, primaryKey, table, primaryKey)
		query = query.Join(dbr.I(joinTable).As(aliasTableName), onCondition)
		whereCondition = append(whereCondition, And(Eq(aliasTableName+"."+keyField, key), Eq(aliasTableName+"."+valueField, values)))
		query.JoinCount++
	}
	if len(whereCondition) > 0 {
		query = query.Where(And(whereCondition...))
	}
	return query
}

// common func to insert into table
func Insert(ctx context.Context, conn *Conn, table string, s interface{}) error {
	_, err := conn.InsertInto(table).
		Columns(dbschema.InsertColumns[table]...).
		Record(s).Exec()
	if err != nil {
		logger.Errorf("Failed to insert into %s: [%+v]", table, err)
	}
	return err
}

// common func to delete rows from table
func Delete(ctx context.Context, conn *Conn, table, idColumn string, ids []string) error {
	_, err := conn.DeleteFrom(table).
		Where(Eq(idColumn, ids)).
		Exec()

	if err != nil {
		logger.Errorf("Failed to delete rows(%+v) from table(%s): [%+v]", ids, table, err)
	}
	return err
}

func CheckAllRowsExist(ctx context.Context, conn *Conn, table, idColumn string, ids []string, otherColumns map[string][]string) (notExistRows []string, err error) {
	var idsInDB []string
	selectQuery := conn.Select(idColumn).From(table).Where(Eq(idColumn, ids))
	if otherColumns != nil {
		for column, v := range otherColumns {
			selectQuery.Where(Eq(column, v))
		}
	}
	_, err = selectQuery.Load(&idsInDB)

	if err != nil {
		logger.Errorf("Failed to get %s(%s) from %s: %+v", idColumn, ids, table, err)
		return
	}
	notExistRows = stringutil.Diff(ids, idsInDB)
	return
}

func CheckAnyRowsExistByConditions(ctx context.Context, conn *Conn, table, column string, conditions ...dbr.Builder) (bool, error) {
	query := conn.Select(column).From(table)
	for _, c := range conditions {
		query.Where(c)
	}
	query.Limit(1)

	total, err := query.Count()
	if err != nil {
		logger.Errorf("Failed to check rows exit in %s by %v: %+v", table, conditions, err)
		return false, err
	}
	return total > 0, nil
}

func CheckAnyRowsExist(ctx context.Context, conn *Conn, table, column string, columnValues []string) (bool, error) {
	selectQuery := conn.Select(column).From(table).
		Where(Eq(column, columnValues)).
		Limit(1)

	total, err := selectQuery.Count()
	if err != nil {
		logger.Errorf("Failed to get %s(%s) from %s: %+v",
			column, columnValues, table, err)
		return false, err
	}

	if total > 0 {
		return true, nil
	}

	return false, nil
}

func CheckRowExist(ctx context.Context, conn *Conn, table, column, value string) (bool, error) {
	var total, err = conn.Select(column).
		From(table).
		Where(Eq(column, value)).Count()

	if err != nil {
		logger.Errorf("Failed to check if %s(%s) exists in %s: %+v",
			column, value, table, err)
		return false, err
	}
	if total > 0 {
		return true, nil
	}
	return false, nil
}

//return: the ids not exist
func GetNotExistRows(ctx context.Context, conn *Conn, table, idColumn string, ids []string, otherColumns map[string][]string) ([]string, error) {
	selectQuery := conn.Select(idColumn).
		From(table).
		Where(Eq(idColumn, ids))

	if otherColumns != nil {
		for column, v := range otherColumns {
			selectQuery.Where(Eq(column, v))
		}
	}

	var idsInDB []string
	_, err := selectQuery.Load(&idsInDB)

	if err != nil {
		logger.Errorf("Failed to check if %s(%s) exist in %s: %+v", idColumn, ids, table, err)
		return nil, err
	}

	return stringutil.Diff(ids, idsInDB), nil
}
