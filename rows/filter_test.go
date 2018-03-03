package rows

import (
	"database/sql/driver"
	"testing"
)

func TestFilterImplementsRows(t *testing.T) {
	var _ driver.Rows = &filter{}
}
