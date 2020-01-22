package querybuilder

type JoinType int8

const (
	UndefinedJoin JoinType = iota
	InnerJoin
	LeftJoin
	RightJoin
)

func (jt JoinType) String() string {
	switch jt {
	case InnerJoin:
		return " JOIN "
	case LeftJoin:
		return " LEFT JOIN "
	case RightJoin:
		return " RIGHT JOIN "
	}

	return "UNDEFINED"
}
