package maximumsubarray

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{1, 1, 1, 1, 1},
			want: 5,
		},
		{
			in:   []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
	}

	for _, test := range tests {
		got := maxSubArray(test.in)

		if got != test.want {
			t.Errorf("maxSubArray(%v) = %d, want %d", test.in, got, test.want)
		}
	}
}
