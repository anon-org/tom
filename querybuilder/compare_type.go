package querybuilder

type CompareType int8

const (
	UndefinedCompare CompareType = iota
	Equal
	NotEqual
	GreaterThan
	GreaterThanEqual
	LessThan
	LessThanEqual
	Is
	IsNot
	Like
	NotLike
	In
	NotIn
)

func (ct CompareType) String() string {
	switch ct {
	case Equal:
		return " = "
	case NotEqual:
		return " <> "
	case GreaterThan:
		return " > "
	case GreaterThanEqual:
		return " >= "
	case LessThan:
		return " < "
	case LessThanEqual:
		return " <= "
	case Is:
		return " IS "
	case IsNot:
		return " IS NOT "
	case Like:
		return " LIKE "
	case NotLike:
		return " NOT LIKE "
	case In:
		return " IN "
	case NotIn:
		return " NOT IN "
	}

	return "UNDEFINED"
}
