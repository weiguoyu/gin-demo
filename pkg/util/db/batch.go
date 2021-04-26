package db

type InsertBatchFactory interface {
	New(conn *Conn) InsertBatch
	NewWithTx(tx *ConnTX) InsertBatch
}

type insertBatchFactory struct {
	table   string
	columns []string
}

func NewInsertBatchFactory(table string, columns ...string) InsertBatchFactory {
	return &insertBatchFactory{table: table, columns: columns}
}

func (f *insertBatchFactory) New(conn *Conn) InsertBatch {
	return newInsertBatch(conn, f.table, f.columns...)
}

func (f *insertBatchFactory) NewWithTx(tx *ConnTX) InsertBatch {
	return newInsertBatchWithTx(tx, f.table, f.columns...)
}

type InsertBatch interface {
	Insert(structVal interface{}) InsertBatch
	Exec() error
}

type insertBatch struct {
	count int
	query *InsertQuery
}

func newInsertBatch(conn *Conn, table string, columns ...string) InsertBatch {
	return &insertBatch{query: conn.InsertInto(table).Columns(columns...)}
}

func newInsertBatchWithTx(tx *ConnTX, table string, columns ...string) InsertBatch {
	return &insertBatch{query: tx.InsertInto(table).Columns(columns...)}
}

func (b *insertBatch) Insert(structVal interface{}) InsertBatch {
	if structVal != nil {
		b.query.Record(structVal)
		b.count++
	}
	return b
}

func (b *insertBatch) Exec() error {
	if b.count == 0 {
		return nil
	}
	_, err := b.query.Exec()
	return err
}
