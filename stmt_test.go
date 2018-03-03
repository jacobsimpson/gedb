package gedb

import (
	sqldriver "database/sql/driver"
	"testing"
)

func TestImplementsStmt(t *testing.T) {
	var _ sqldriver.Stmt = &gedbStmt{}
	// Optional implementation, replaces a deprecated method.
	//var _ driver.StmtExecContext = &gedbStmt{}
}
