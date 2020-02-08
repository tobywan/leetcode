package factorialzeros

import "testing"

// TestTrailingZeros does exactly that
func TestTrailingZeroes(t *testing.T) {
	tests := []struct {
		in   int
		want int
	}{
		{
			in:   5,
			want: 1,
		},
		{
			in:   10,
			want: 2,
		},
		{
			in:   6,
			want: 1,
		},
		{
			in:   19,
			want: 3,
		},
		{
			in:   25,
			want: 6,
		},
		{
			in:   26,
			want: 6,
		},
		{
			in:   30,
			want: 7,
		},
		{
			in:   3,
			want: 0,
		},
	}
	for _, test := range tests {
		got := trailingZeroes(test.in)
		if got != test.want {
			t.Errorf("trailingZeroes(%d)=%d, want %d)", test.in, got, test.want)
		}
	}
}
