package rows

import (
	"database/sql/driver"
	"testing"
)

func TestImplementsRows(t *testing.T) {
	var _ driver.Rows = &tableScan{}
}
