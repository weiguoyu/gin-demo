package db

import (
	"context"
	"github.com/gocraft/dbr/v2"
)

type ConnTX struct {
	*dbr.Tx
	ctx        context.Context
	InsertHook InsertHook
	UpdateHook UpdateHook
	DeleteHook DeleteHook
}

func (tx *ConnTX) Select(columns ...string) *SelectQuery {
	return &SelectQuery{tx.Tx.Select(columns...), tx.ctx, 0, false, "", ""}
}

func (tx *ConnTX) SelectBySql(query string, value ...interface{}) *SelectQuery {
	return &SelectQuery{tx.Tx.SelectBySql(query, value...), tx.ctx, 0, true, "", ""}
}

func (tx *ConnTX) SelectAll(columns ...string) *SelectQuery {
	return &SelectQuery{tx.Tx.Select("*"), tx.ctx, 0, false, "", ""}
}

func (tx *ConnTX) InsertInto(table string) *InsertQuery {
	return &InsertQuery{tx.Tx.InsertInto(table), tx.ctx, tx.InsertHook}
}

func (tx *ConnTX) DeleteFrom(table string) *DeleteQuery {
	return &DeleteQuery{tx.Tx.DeleteFrom(table), tx.ctx, tx.DeleteHook, table, false}
}

func (tx *ConnTX) Update(table string) *UpdateQuery {
	return &UpdateQuery{tx.Tx.Update(table), tx.ctx, tx.UpdateHook, table, false}
}
