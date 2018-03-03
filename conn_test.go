package gedb

import (
	sqldriver "database/sql/driver"
	"testing"
)

func TestImplementsConn(t *testing.T) {
	var _ sqldriver.Conn = &gedbConn{}
	// Optional implementation, replaces a deprecated method.
	//var _ driver.ConnBeginTx = &gedbConn{}
}
