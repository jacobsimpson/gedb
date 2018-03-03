package rows

import (
	"database/sql/driver"
	"fmt"
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

func NewTableScan(tableName string) driver.Rows {
	return &tableScan{
		tableName: tableName,
		nextRow:   0,
	}
}

type tableScan struct {
	tableName string
	nextRow   int
}

// Columns returns the names of the columns. The number of
// columns of the result is inferred from the length of the
// slice. If a particular column name isn't known, an empty
// string should be returned for that entry.
func (r *tableScan) Columns() []string {
	return []string{
		"id",
		"table_name",
	}
}

// Close closes the rows iterator.
func (r *tableScan) Close() error {
	return fmt.Errorf("gedbRows.Close -- not implemented")
}

// Next is called to populate the next row of data into
// the provided slice. The provided slice will be the same
// size as the Columns() are wide.
//
// Next should return io.EOF when there are no more rows.
func (r *tableScan) Next(dest []driver.Value) error {
	if r.nextRow >= len(static) {
		// Return io.EOF indicates there are no more rows.
		return io.EOF
	}
	dest[0] = static[r.nextRow].id
	dest[1] = static[r.nextRow].name
	r.nextRow++
	return nil
}
