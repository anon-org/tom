package querybuilder

import (
	"errors"
	"github.com/anon-org/tom/helper"
	"strings"
)

type joinQuery struct {
	rightTable  string
	alias       string
	leftField   string
	compareType CompareType
	rightField  string
	JoinType    JoinType
}

type joinQueries []joinQuery

type SelectQuery struct {
	fields     []string
	table      string
	joinStacks joinQueries
}

func Select(fields ...string) *SelectQuery {
	return &SelectQuery{
		fields,
		"",
		joinQueries{},
	}
}

func SelectAll() *SelectQuery {
	return Select("*")
}

func Count(field string) *SelectQuery {
	return Select("COUNT(" + field + ")")
}

func Distinct(field string) *SelectQuery {
	return Select("DISTINCT(" + field + ")")
}

func Min(field string) *SelectQuery {
	return Select("MIN(" + field + ")")
}

func Max(field string) *SelectQuery {
	return Select("MAX(" + field + ")")
}

func Sum(field string) *SelectQuery {
	return Select("SUM(" + field + ")")
}

func Avg(field string) *SelectQuery {
	return Select("AVG(" + field + ")")
}

func (sq *SelectQuery) Field(fields ...string) *SelectQuery {
	sq.fields = append(sq.fields, fields...)
	return sq
}

func (sq *SelectQuery) Count(field string) *SelectQuery {
	return sq.Field("COUNT(" + field + ")")
}

func (sq *SelectQuery) Distinct(field string) *SelectQuery {
	return sq.Field("DISTINCT(" + field + ")")
}

func (sq *SelectQuery) Min(field string) *SelectQuery {
	return sq.Field("MIN(" + field + ")")
}

func (sq *SelectQuery) Max(field string) *SelectQuery {
	return sq.Field("MAX(" + field + ")")
}

func (sq *SelectQuery) Sum(field string) *SelectQuery {
	return sq.Field("SUM(" + field + ")")
}

func (sq *SelectQuery) Avg(field string) *SelectQuery {
	return sq.Field("AVG(" + field + ")")
}

func (sq *SelectQuery) From(table string) *SelectQuery {
	sq.table = table
	return sq
}

func (sq *SelectQuery) join(rightTable, alias string, joinType JoinType) *SelectQuery {
	joinQuery := joinQuery{
		rightTable: rightTable,
		alias:      alias,
		leftField:  "",
		rightField: "",
		JoinType:   joinType,
	}

	sq.joinStacks = append(sq.joinStacks, joinQuery)
	return sq
}

func (sq *SelectQuery) Join(rightTable string) *SelectQuery {
	return sq.join(rightTable, "", InnerJoin)
}

func (sq *SelectQuery) JoinAs(rightTable, alias string) *SelectQuery {
	return sq.join(rightTable, alias, InnerJoin)
}

func (sq *SelectQuery) LeftJoinAs(rightTable, alias string) *SelectQuery {
	return sq.join(rightTable, alias, LeftJoin)
}

func (sq *SelectQuery) RightJoinAs(rightTable, alias string) *SelectQuery {
	return sq.join(rightTable, alias, RightJoin)
}

func (sq *SelectQuery) On(leftField string, compareType CompareType, rightField string) *SelectQuery {
	joinStackLength := len(sq.joinStacks)

	if joinStackLength == 0 {
		return sq
	}

	topStack := sq.joinStacks[joinStackLength-1]

	topStack.leftField = leftField
	topStack.compareType = compareType
	topStack.rightField = rightField

	return sq
}

func (sq *SelectQuery) Build() (string, error) {
	var b strings.Builder

	if sq.table == "" {
		return "", errors.New("missing table name")
	}

	b.WriteString(sq.getFieldQuery())
	b.WriteString(" FROM ")
	b.WriteString(sq.table)
	b.WriteString(sq.getJoinQuery())

	return b.String(), nil
}

func (sq *SelectQuery) getFieldQuery() string {
	if len(sq.fields) == 0 || helper.Has(sq.fields, "*") {
		sq.fields = []string{"*"}
	}

	return "SELECT " + strings.Join(sq.fields, ", ")
}

func (sq *SelectQuery) getJoinQuery() string {
	var b strings.Builder

	if len(sq.joinStacks) == 0 {
		return ""
	}

	for _, joinQuery := range sq.joinStacks {
		b.WriteString(joinQuery.JoinType.String())
		b.WriteString(joinQuery.rightTable)
		b.WriteString(" ON ")
		b.WriteString(joinQuery.leftField)
		b.WriteString(joinQuery.compareType.String())
		b.WriteString(joinQuery.rightField)
	}

	return b.String()
}
