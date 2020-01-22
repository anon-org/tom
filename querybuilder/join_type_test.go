package querybuilder

import "testing"

func TestJoinTypeEnum(t *testing.T) {
	testCase := []struct {
		val       JoinType
		valString string
		want      string
	}{
		{
			val:       InnerJoin,
			valString: "InnerJoin",
			want:      " JOIN ",
		},
		{
			val:       LeftJoin,
			valString: "LeftJoin",
			want:      " LEFT JOIN ",
		},
		{
			val:       RightJoin,
			valString: "RightJoin",
			want:      " RIGHT JOIN ",
		},
		{
			val:       UndefinedJoin,
			valString: "UndefinedJoin",
			want:      "UNDEFINED",
		},
	}

	for _, test := range testCase {
		got := test.val.String()

		if got != test.want {
			t.Errorf("%v.String() was incorrect, want: %s, got %s", test.valString, test.want, got)
		}
	}
}
