package querybuilder

import "testing"

func TestSelect(t *testing.T) {
	got := Select("field1", "field2", "field3").From("table").Build()
	want := "SELECT field1, field2, field3 FROM table"

	if got != want {
		t.Errorf("Select(\"field1\", \"field2\", \"field3\").From(\"table\").Build() failed\n" +
			"want: %v, got: %v", got, want)
	}
}
