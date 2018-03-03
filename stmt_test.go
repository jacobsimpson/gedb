package gedb

import (
	"database/sql/driver"
	"testing"
)

func TestImplementsStmt(t *testing.T) {
	var _ driver.Stmt = &gedbStmt{}
	// Optional implementation, replaces a deprecated method.
	//var _ driver.StmtExecContext = &gedbStmt{}
}
