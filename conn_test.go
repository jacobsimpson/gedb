package gedb

import (
	"database/sql/driver"
	"testing"
)

func TestImplementsConn(t *testing.T) {
	var _ driver.Conn = &gedbConn{}
	// Optional implementation, replaces a deprecated method.
	//var _ driver.ConnBeginTx = &gedbConn{}
}
