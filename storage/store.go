package storage

import "database/sql/driver"

type Store interface {
	Get(rowid RowId) Row
	Scan(tableid TableId) TableScanner
}

type TableScanner interface {
	// Advances the scanner to the next row. Will return io.EOF if there are no
	// more rows available.
	Scan() error
	// Returns the current row. Scan must be called before Row will return any
	// useful data.
	Row() Row
}

type TableId int64

type RowId int64
type Row interface {
	Data() []driver.Value
}

type row struct {
	data []driver.Value
}

func (r *row) Data() []driver.Value {
	return r.data
}
