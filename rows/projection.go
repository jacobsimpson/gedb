package rows

import (
	"database/sql/driver"
)

func NewProjection(input driver.Rows, fields ...string) driver.Rows {
	requested := map[string]bool{}
	for _, f := range fields {
		requested[f] = true
	}
	columns := []string{}
	columnNumbers := []int{}
	for i, c := range input.Columns() {
		if _, ok := requested[c]; ok {
			columns = append(columns, c)
			columnNumbers = append(columnNumbers, i)
		}
	}
	return &projection{
		input:         input,
		columns:       columns,
		columnNumbers: columnNumbers,
	}
}

type projection struct {
	input         driver.Rows
	columns       []string
	columnNumbers []int
}

// Columns returns the names of the columns. The number of
// columns of the result is inferred from the length of the
// slice. If a particular column name isn't known, an empty
// string should be returned for that entry.
func (r *projection) Columns() []string {
	return r.columns
}

// Close closes the rows iterator.
func (r *projection) Close() error {
	return r.input.Close()
}

// Next is called to populate the next row of data into
// the provided slice. The provided slice will be the same
// size as the Columns() are wide.
//
// Next should return io.EOF when there are no more rows.
func (r *projection) Next(dest []driver.Value) error {
	row := make([]driver.Value, len(r.input.Columns()))
	if err := r.input.Next(row); err != nil {
		return err
	}

	for i, _ := range dest {
		dest[i] = row[r.columnNumbers[i]]
	}
	return nil
}
