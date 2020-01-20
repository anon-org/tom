package helper

import "testing"

func TestHas(t *testing.T) {
	params := []int{}

	got := Has(params, 0)
	want := false


	if got != want {
		t.Errorf("Has() failed\n" +
			"want: %v, got: %v", got, want)
	}
}
