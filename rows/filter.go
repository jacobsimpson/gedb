package rows

import "database/sql/driver"

type Criteria interface {
	apply(row []driver.Value) bool
}

type FieldEqualsValue struct {
	FieldName   string
	Value       interface{}
	fieldNumber int
}

func (c *FieldEqualsValue) apply(row []driver.Value) bool {
	return row[c.fieldNumber] == c.Value
}

func NewFilter(input driver.Rows, criteria []Criteria) driver.Rows {
	columns := map[string]int{}
	for i, c := range input.Columns() {
		columns[c] = i
	}
	for _, c := range criteria {
		if f, ok := c.(*FieldEqualsValue); ok {
			f.fieldNumber = columns[f.FieldName]
		}
	}
	return &filter{
		input:    input,
		criteria: criteria,
	}
}

type filter struct {
	input    driver.Rows
	criteria []Criteria
}

// Columns returns the names of the columns. The number of
// columns of the result is inferred from the length of the
// slice. If a particular column name isn't known, an empty
// string should be returned for that entry.
func (r *filter) Columns() []string {
	return r.input.Columns()
}

// Close closes the rows iterator.
func (r *filter) Close() error {
	return r.input.Close()
}

// Next is called to populate the next row of data into
// the provided slice. The provided slice will be the same
// size as the Columns() are wide.
//
// Next should return io.EOF when there are no more rows.
func (r *filter) Next(dest []driver.Value) error {
	row := make([]driver.Value, len(dest))
	for {
		err := r.input.Next(row)
		if err != nil {
			return err
		}
		if r.apply(row) {
			for i, v := range row {
				dest[i] = v
			}
			return nil
		}
	}
}

func (r *filter) apply(row []driver.Value) bool {
	for _, c := range r.criteria {
		if !c.apply(row) {
			return false
		}
	}
	return true
}
