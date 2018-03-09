package storage

import (
	"database/sql/driver"
	"io"
)

var static = []struct {
	id   int
	name string
}{
	{
		id:   1,
		name: "objects",
	},
	{
		id:   2,
		name: "users",
	},
}

func NewFileStore(filename string) Store {
	return &fileStore{}
}

type fileStore struct{}

func (s *fileStore) Get(rowid RowId) Row { return nil }

func (s *fileStore) Scan(tableMetadata *TableMetadata) RowScanner {
	if tableMetadata.TableId == 1 {
		return &rowScanner{tableMetadata: tableMetadata, nextRow: -1}
	}
	return nil
}

type rowScanner struct {
	tableMetadata *TableMetadata
	nextRow       int
}

func (t *rowScanner) Scan() error {
	t.nextRow++
	if t.nextRow >= len(static) {
		return io.EOF
	}
	return nil
}

func (t *rowScanner) Row() Row {
	return &row{
		data: []driver.Value{static[t.nextRow].id, static[t.nextRow].name},
	}
}
