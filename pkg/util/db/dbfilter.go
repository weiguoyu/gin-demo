package db

import (
	"context"
	"github.com/gocraft/dbr/v2"
)

/**
说明:
1.func (conn *Conn) Select(columns ...string) *SelectQuery 该方法会走selectFilter
2.func (conn *Conn) SelectAll(columns ...string) *SelectQuery 也会走selectFilter
3.func (conn *Conn) SelectBySql(query string, value ...interface{}) *SelectQuery
该方法不会走selectFilter，如果需要数据权限控制，需自行处理
4.Select 与 SelectAll 中的Join与JoinAs需自行处理on的条件
5.insert update delete的处理与上相同
*/
func (b *SelectQuery) selectFilter(ctx context.Context, whereCond []dbr.Builder) []dbr.Builder {
	if b.hadUsedSelectFilter {
		return whereCond
	}
	b.hadUsedSelectFilter = true
	return dbFilter(ctx, b.tableName, b.tableAlias, whereCond)
}

func (b *UpdateQuery) updateFilter(ctx context.Context, whereCond []dbr.Builder) []dbr.Builder {
	if b.hadUsedUpdateFilter {
		return whereCond
	}
	b.hadUsedUpdateFilter = true
	return dbFilter(ctx, b.tableName, "", whereCond)
}

func (b *DeleteQuery) deleteFilter(ctx context.Context, whereCond []dbr.Builder) []dbr.Builder {
	if b.hadUsedDeleteFilter {
		return whereCond
	}
	b.hadUsedDeleteFilter = true
	return dbFilter(ctx, b.tableName, "", whereCond)
}

func dbFilter(ctx context.Context, tableName, tableAlias string, whereCond []dbr.Builder) []dbr.Builder {
	return whereCond
}
