package helper

import "testing"

func TestHasEmptySlice(t *testing.T) {
	var haystack []int
	needle := 0

	got := Has(haystack, needle)
	want := false

	if got != want {
		t.Errorf("Has(%v, %v) was incorrect, want: %v, got: %v", haystack, needle, want, got)
	}
}

func TestHasNotEmptySlice(t *testing.T) {
	haystack := []int{
		0, 1, 2, 3,
	}

	testCase := []struct {
		needle int
		want   bool
	}{
		{
			needle: 2,
			want:   true,
		},
		{
			needle: 4,
			want:   false,
		},
	}

	for _, test := range testCase {
		got := Has(haystack, test.needle)

		if got != test.want {
			t.Errorf("Has(%v, %v) was incorrect, want: %v, got: %v", haystack, test.needle, test.want, got)
		}
	}
}
