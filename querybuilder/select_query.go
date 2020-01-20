package querybuilder

import (
	"github.com/anon-org/tom/helper"
	"strings"
)

type SelectQuery struct {
	fields []string
	table string
}

func Select(fields ...string) *SelectQuery {
	return &SelectQuery{
		fields,
		"",
	}
}

func SelectAll() *SelectQuery {
	return &SelectQuery{
		[]string{"*"},
		"",
	}
}

func (sq *SelectQuery) Field(fields ...string) *SelectQuery {
	sq.fields = append(sq.fields, fields...)
	return sq
}

func (sq *SelectQuery) From(table string) *SelectQuery {
	sq.table = table
	return sq
}

func (sq *SelectQuery) Build() string {
	var b strings.Builder

	if helper.Has(sq.fields, "*") {
		sq.fields = []string{"*"}
	}

	fields := strings.Join(sq.fields, ", ")

	b.WriteString("SELECT ")
	b.WriteString(fields)
	b.WriteString(" FROM ")
	b.WriteString(sq.table)

	return b.String()
}