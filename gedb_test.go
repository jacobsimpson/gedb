package gedb

import (
	"database/sql"
	"database/sql/driver"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImplementsDriver(t *testing.T) {
	var _ driver.Driver = &gedb{}
}

func TestOpenDatabase(t *testing.T) {
	assert := assert.New(t)

	db, err := sql.Open("gedb", "./TestOpenDatabase.gedb")
	assert.NoError(err)
	assert.NotNil(db)
	//defer os.Remove("./TestOpenDatabase.gedb")

	db.Query("SELECT * FROM users")
	//assert.NoError(err)
	//assert.NotNil(rows)
}
