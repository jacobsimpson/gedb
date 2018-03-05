package gedb

import (
	"database/sql/driver"
	"fmt"

	"github.com/jacobsimpson/gedb/parser"
	"github.com/jacobsimpson/gedb/storage"
)

type gedbConn struct {
	store storage.Store
}

// Prepare returns a prepared statement, bound to this connection.
func (conn *gedbConn) Prepare(query string) (driver.Stmt, error) {
	ast, err := parser.Parse(query)
	if err != nil {
		return nil, err
	}
	return &gedbStmt{store: conn.store, ast: ast}, nil
}

// Close invalidates and potentially stops any current
// prepared statements and transactions, marking this
// connection as no longer in use.
//
// Because the sql package maintains a free pool of
// connections and only calls Close when there's a surplus of
// idle connections, it shouldn't be necessary for drivers to
// do their own connection caching.
func (conn *gedbConn) Close() error {
	return fmt.Errorf("conn.Close() unimplemented")
}

// Begin starts and returns a new transaction.
//
// Deprecated: Drivers should implement ConnBeginTx instead (or additionally).
func (conn *gedbConn) Begin() (driver.Tx, error) {
	return nil, fmt.Errorf("conn.Begin() unimplemented")
}
