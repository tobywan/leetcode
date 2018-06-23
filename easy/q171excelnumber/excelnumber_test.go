package excelnumber

import "testing"

func TestTitleToNumber(t *testing.T) {

	tests := []struct {
		in   string
		want int
	}{
		{
			in:   "A",
			want: 1,
		},
		{
			in:   "AB",
			want: 28,
		},
		{
			in:   "AZ",
			want: 52,
		},
		{
			in:   "ZY",
			want: 701,
		},
		{
			in:   "AA",
			want: 27,
		},
		{
			in:   "AAA",
			want: 703,
		},
		{
			in:   "ZZ",
			want: 702,
		},
	}

	for _, test := range tests {
		got := titleToNumber(test.in)
		if got != test.want {
			t.Errorf("titleToNumber(%q)=%d, want %d", test.in, got, test.want)
		}
	}

}
