package rows

import (
	"database/sql/driver"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockRows struct {
	columns []string
	rows    [][]driver.Value
	next    int
}

func (r *mockRows) Close() error { return nil }

func (r *mockRows) Columns() []string {
	return r.columns
}

func (r *mockRows) Next(dest []driver.Value) error {
	if r.next >= len(r.rows) {
		return io.EOF
	}
	for i, _ := range dest {
		dest[i] = r.rows[r.next][i]
	}
	return nil
}

func TestProjectionImplementsRows(t *testing.T) {
	var _ driver.Rows = &projection{}
}

func TestProjectionColumnsMatch(t *testing.T) {
	assert := assert.New(t)

	projection := NewProjection(
		&mockRows{
			columns: []string{"id", "name", "value", "type", "category", "item"},
		},
		"id",
		"value",
		"item",
	)

	assert.Equal([]string{"id", "value", "item"}, projection.Columns())
}

func TestProjectionNext(t *testing.T) {
	assert := assert.New(t)

	projection := NewProjection(
		&mockRows{
			columns: []string{"id", "name", "value", "type", "category", "item"},
			rows: [][]driver.Value{
				[]driver.Value{1, "name-1", "value-1", "type-1", "category-1", "item-1"},
			},
		},
		"id",
		"type",
		"value",
	)
	//dest []Value
	row := make([]driver.Value, 3)
	assert.NoError(projection.Next(row))
	assert.Equal([]driver.Value{1, "value-1", "type-1"}, row)
}
