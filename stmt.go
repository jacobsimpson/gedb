package gedb

import (
	sqldriver "database/sql/driver"
	"fmt"

	"github.com/jacobsimpson/gedb/parser"
)

type gedbStmt struct {
	ast parser.AST
}

// Close closes the statement.
//
// As of Go 1.1, a Stmt will not be closed if it's in use
// by any queries.
func (stmt *gedbStmt) Close() error {
	return fmt.Errorf("unimplemented - could not close")
}

// NumInput returns the number of placeholder parameters.
//
// If NumInput returns >= 0, the sql package will sanity check
// argument counts from callers and return errors to the caller
// before the statement's Exec or Query methods are called.
//
// NumInput may also return -1, if the driver doesn't know
// its number of placeholders. In that case, the sql package
// will not sanity check Exec or Query argument counts.
func (stmt *gedbStmt) NumInput() int {
	return -1
}

// Exec executes a query that doesn't return rows, such
// as an INSERT or UPDATE.
//
// Deprecated: Drivers should implement StmtExecContext instead (or additionally).
func (stmt *gedbStmt) Exec(args []sqldriver.Value) (sqldriver.Result, error) {
	return nil, fmt.Errorf("unimplemented - could not 'Exec'")
}

// Query executes a query that may return rows, such as a
// SELECT.
//
// Deprecated: Drivers should implement StmtQueryContext instead (or additionally).
func (stmt *gedbStmt) Query(args []sqldriver.Value) (sqldriver.Rows, error) {
	fmt.Printf("the query to execute is : %+v\n", stmt.ast)
	return nil, fmt.Errorf("unimplemented - could not 'Query'")
}
