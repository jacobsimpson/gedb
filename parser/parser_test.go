package parser

import (
	"testing"
)

func TestImplementsStmt(t *testing.T) {
	var _ AST = &SelectStatement{}
}
