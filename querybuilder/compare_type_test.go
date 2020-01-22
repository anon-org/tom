package querybuilder

import "testing"

func TestCompareTypeEnum(t *testing.T) {
	testCase := []struct {
		val CompareType
		valString string
		want string
	} {
		{
			val: Equal,
			valString: "Equal",
			want: " = ",
		},
		{
			val: NotEqual,
			valString: "NotEqual",
			want: " <> ",
		},
		{
			val: GreaterThan,
			valString: "GreaterThan",
			want: " > ",
		},
		{
			val: GreaterThanEqual,
			valString: "GreaterThanEqual",
			want: " >= ",
		},
		{
			val: LessThan,
			valString: "LessThan",
			want: " < ",
		},
		{
			val: LessThanEqual,
			valString: "LessThanEqual",
			want: " <= ",
		},
		{
			val: Is,
			valString: "Is",
			want: " IS ",
		},
		{
			val: IsNot,
			valString: "IsNot",
			want: " IS NOT ",
		},
		{
			val: Like,
			valString: "Like",
			want: " LIKE ",
		},
		{
			val: NotLike,
			valString: "NotLike",
			want: " NOT LIKE ",
		},
		{
			val: In,
			valString: "In",
			want: " IN ",
		},
		{
			val: NotIn,
			valString: "NotIn",
			want: " NOT IN ",
		},
		{
			val: UndefinedCompare,
			valString: "UndefinedCompare",
			want: "UNDEFINED",
		},
	}

	for _, test := range testCase {
		got := test.val.String()

		if got != test.want {
			t.Errorf("%v.String() was incorrect, want: %s, got %s", test.valString, test.want, got)
		}
	}
}
