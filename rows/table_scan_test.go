package rows

import (
	"database/sql/driver"
	"testing"
)

func TestTableScanImplementsRows(t *testing.T) {
	var _ driver.Rows = &tableScan{}
}
