package querybuilder

import (
	"testing"
)

func TestSelectMissingTableName(t *testing.T) {
	funcString := "Select().Build()"
	gotQuery, gotErr := Select().Build()
	wantQuery := ""
	wantError := "missing table name"

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr == nil || gotErr.Error() != wantError {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestSelectSingleValue(t *testing.T) {
	funcString := "Select(\"field_name\").From(\"table_name\").Build()"
	gotQuery, gotErr := Select("field_name").From("table_name").Build()
	wantQuery := "SELECT field_name FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestSelectWithoutValue(t *testing.T) {
	funcString := "Select().From(\"table_name\").Build()"
	gotQuery, gotErr := Select().From("table_name").Build()
	wantQuery := "SELECT * FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestSelectMultiWithoutValue(t *testing.T) {
	funcString := "Select(\"field1\", \"field2\").From(\"table_name\").Build()"
	gotQuery, gotErr := Select("field1", "field2").From("table_name").Build()
	wantQuery := "SELECT field1, field2 FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestSelectAll(t *testing.T) {
	funcString := "SelectAll().From(\"table_name\").Build()"
	gotQuery, gotErr := Select().From("table_name").Build()
	wantQuery := "SELECT * FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestSelectEmptyValueWithFieldEmptyValue(t *testing.T) {
	funcString := "Select().Field().From(\"table_name\").Build()"
	gotQuery, gotErr := Select().Field().From("table_name").Build()
	wantQuery := "SELECT * FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestSelectSingleValueWithFieldEmptyValue(t *testing.T) {
	funcString := "Select(\"field1\").Field().From(\"table_name\").Build()"
	gotQuery, gotErr := Select("field1").Field().From("table_name").Build()
	wantQuery := "SELECT field1 FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestSelectSingleValueWithFieldSingleValue(t *testing.T) {
	funcString := "Select(\"field1\").Field(\"field2\").From(\"table_name\").Build()"
	gotQuery, gotErr := Select("field1").Field("field2").From("table_name").Build()
	wantQuery := "SELECT field1, field2 FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestSelectEmptyValueWithFieldSingleValue(t *testing.T) {
	funcString := "Select().Field(\"field1\").From(\"table_name\").Build()"
	gotQuery, gotErr := Select().Field("field1").From("table_name").Build()
	wantQuery := "SELECT field1 FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestSelectMultiValueWithFieldEmptyValue(t *testing.T) {
	funcString := "Select(\"field1\", \"field2\").Field().From(\"table_name\").Build()"
	gotQuery, gotErr := Select("field1", "field2").Field().From("table_name").Build()
	wantQuery := "SELECT field1, field2 FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestSelectMultiValueWithFieldMultiValue(t *testing.T) {
	funcString := "Select(\"field1\", \"field2\").Field(\"field3\", \"field4\").From(\"table_name\").Build()"
	gotQuery, gotErr := Select("field1", "field2").Field("field3", "field4").From("table_name").Build()
	wantQuery := "SELECT field1, field2, field3, field4 FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestSelectAllWithAnotherField(t *testing.T) {
	funcString := "SelectAll().Field(\"field1\", \"field2\").From(\"table_name\").Build()"
	gotQuery, gotErr := SelectAll().Field("field1", "field2").From("table_name").Build()
	wantQuery := "SELECT * FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestCount(t *testing.T) {
	funcString := "Count(\"*\").From(\"table_name\").Build()"
	gotQuery, gotErr := Count("*").From("table_name").Build()
	wantQuery := "SELECT COUNT(*) FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestDistinct(t *testing.T) {
	funcString := "Distinct(\"field1\").From(\"table_name\").Build()"
	gotQuery, gotErr := Distinct("field1").From("table_name").Build()
	wantQuery := "SELECT DISTINCT(field1) FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestMin(t *testing.T) {
	funcString := "Min(\"field1\").From(\"table_name\").Build()"
	gotQuery, gotErr := Min("field1").From("table_name").Build()
	wantQuery := "SELECT MIN(field1) FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestMax(t *testing.T) {
	funcString := "Max(\"field1\").From(\"table_name\").Build()"
	gotQuery, gotErr := Max("field1").From("table_name").Build()
	wantQuery := "SELECT MAX(field1) FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestSum(t *testing.T) {
	funcString := "Sum(\"field1\").From(\"table_name\").Build()"
	gotQuery, gotErr := Sum("field1").From("table_name").Build()
	wantQuery := "SELECT SUM(field1) FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestAvg(t *testing.T) {
	funcString := "Avg(\"field1\").From(\"table_name\").Build()"
	gotQuery, gotErr := Avg("field1").From("table_name").Build()
	wantQuery := "SELECT AVG(field1) FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestCountWithAnotherField(t *testing.T) {
	funcString := "Select(\"field1\").Count(\"field2\").Field(\"field3\").From(\"table_name\").Build()"
	gotQuery, gotErr := Select("field1").Count("field2").Field("field3").From("table_name").Build()
	wantQuery := "SELECT field1, COUNT(field2), field3 FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestDistinctWithAnotherField(t *testing.T) {
	funcString := "Select(\"field1\").Distinct(\"field2\").Field(\"field3\").From(\"table_name\").Build()"
	gotQuery, gotErr := Select("field1").Distinct("field2").Field("field3").From("table_name").Build()
	wantQuery := "SELECT field1, DISTINCT(field2), field3 FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestMinWithAnotherField(t *testing.T) {
	funcString := "Select(\"field1\").Min(\"field2\").Field(\"field3\").From(\"table_name\").Build()"
	gotQuery, gotErr := Select("field1").Min("field2").Field("field3").From("table_name").Build()
	wantQuery := "SELECT field1, MIN(field2), field3 FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestMaxWithAnotherField(t *testing.T) {
	funcString := "Select(\"field1\").Max(\"field2\").Field(\"field3\").From(\"table_name\").Build()"
	gotQuery, gotErr := Select("field1").Max("field2").Field("field3").From("table_name").Build()
	wantQuery := "SELECT field1, MAX(field2), field3 FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestSumWithAnotherField(t *testing.T) {
	funcString := "Select(\"field1\").Sum(\"field2\").Field(\"field3\").From(\"table_name\").Build()"
	gotQuery, gotErr := Select("field1").Sum("field2").Field("field3").From("table_name").Build()
	wantQuery := "SELECT field1, SUM(field2), field3 FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}

func TestAvgWithAnotherField(t *testing.T) {
	funcString := "Select(\"field1\").Avg(\"field2\").Field(\"field3\").From(\"table_name\").Build()"
	gotQuery, gotErr := Select("field1").Avg("field2").Field("field3").From("table_name").Build()
	wantQuery := "SELECT field1, AVG(field2), field3 FROM table_name"
	wantError := ""

	if gotQuery != wantQuery {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantQuery, gotQuery)
	}

	if gotErr != nil {
		t.Errorf("%s was incorrect, want %s, got %s", funcString, wantError, gotErr)
	}
}
