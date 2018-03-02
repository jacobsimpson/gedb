package gedb

import (
	"database/sql/driver"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImplementsInterface(*testing.T) {
	var _ driver.Driver = &gedb{}
	// Just compiling is enough to pass this test.
}

func TestReadHeader(t *testing.T) {
	assert := assert.New(t)
	db := &gedb{}
	header, err := db.readHeader(strings.NewReader("AAAA - 0.1.2\nthis is some more"))
	assert.NoError(err)
	assert.Equal(header.Name, "AAAA")
	assert.Equal(header.MajorVersion, 0)
	assert.Equal(header.MinorVersion, 1)
	assert.Equal(header.PatchVersion, 2)
}
