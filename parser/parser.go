package parser

import "fmt"

type AST interface {
	String() string
}

type SelectStatement struct {
	Projection *Projection
	From       *Table
	Where      *Expression
}

func (s *SelectStatement) String() string {
	return fmt.Sprintf("SELECT %s FROM %s WHERE %s", s.Projection, s.From, s.Where)
}

type Projection struct {
	Columns *Column
}

func (p *Projection) String() string {
	if p.Columns == nil {
		return "nil"
	}
	return p.Columns.String()
}

type Column struct {
	Schema string
	Table  string
	Column string
}

func (c *Column) String() string {
	return fmt.Sprintf("%s.%s.%s", c.Schema, c.Table, c.Column)
}

type Table struct {
	Name string
}

func (t *Table) String() string {
	return t.Name
}

type Expression struct {
}

func Parse(query string) (AST, error) {
	if len(query) == 0 {
		return nil, fmt.Errorf("Unable to parse %q", query)
	}
	return &SelectStatement{
		Projection: &Projection{},
		From: &Table{
			Name: "users",
		},
		Where: &Expression{},
	}, nil
}
