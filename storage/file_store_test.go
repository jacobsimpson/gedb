package storage_test

import (
	"io"
	"testing"

	"github.com/jacobsimpson/gedb/storage"
	"github.com/stretchr/testify/assert"
)

func TestFileStore(t *testing.T) {
	assert := assert.New(t)

	s := storage.NewFileStore("name")
	assert.NotNil(s)
}
func TestFileStoreScan(t *testing.T) {
	assert := assert.New(t)

	s := storage.NewFileStore("name")
	assert.NotNil(s)

	ts := s.Scan(&storage.TableMetadata{
		TableId:   1,
		TableName: "objects",
		FirstPage: 1,
	})
	assert.NotNil(ts)

	ts.Scan()
	r := ts.Row()
	assert.NotNil(r)
	assert.Equal(2, len(r.Data()))
	assert.Equal(1, r.Data()[0])
	assert.Equal("objects", r.Data()[1])

	ts.Scan()
	r = ts.Row()
	assert.NotNil(r)
	assert.Equal(2, len(r.Data()))
	assert.Equal(2, r.Data()[0])
	assert.Equal("users", r.Data()[1])

	assert.Equal(io.EOF, ts.Scan())
}
