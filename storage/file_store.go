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

func (s *fileStore) Scan(tableId TableId) TableScanner {
	if tableId == 1 {
		return &tableScanner{tableId: tableId, nextRow: -1}
	}
	return nil
}

type tableScanner struct {
	tableId TableId
	nextRow int
}

func (t *tableScanner) Scan() error {
	t.nextRow++
	if t.nextRow >= len(static) {
		return io.EOF
	}
	return nil
}

func (t *tableScanner) Row() Row {
	return &row{
		data: []driver.Value{static[t.nextRow].id, static[t.nextRow].name},
	}
}
