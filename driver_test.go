package gedb

import (
	"database/sql"
	"database/sql/driver"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImplementsDriver(t *testing.T) {
	var _ driver.Driver = &gedbDriver{}
}

func TestOpenDatabase(t *testing.T) {
	assert := assert.New(t)

	tmpDB := filepath.Join(os.TempDir(), "TestOpenDatabase.gedb")
	db, err := sql.Open("gedb", tmpDB)
	assert.NoError(err)
	defer os.Remove(tmpDB)
	assert.NotNil(db)

	db.Query("SELECT * FROM users")
	//assert.NoError(err)
	//assert.NotNil(rows)
}
