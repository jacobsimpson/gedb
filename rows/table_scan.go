package rows

import (
	"database/sql/driver"
	"fmt"

	"github.com/jacobsimpson/gedb/storage"
)

func NewTableScan(store storage.Store, tableMetadata *storage.TableMetadata) driver.Rows {
	scanner := store.Scan(tableMetadata)
	return &tableScan{
		scanner: scanner,
	}
}

type tableScan struct {
	scanner storage.RowScanner
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
	if err := r.scanner.Scan(); err != nil {
		return err
	}
	row := r.scanner.Row()
	for i, v := range row.Data() {
		dest[i] = v
	}
	return nil
}
