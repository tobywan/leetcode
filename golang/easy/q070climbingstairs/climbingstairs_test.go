package climbingstairs

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		in   int
		want int
	}{
		{
			in:   1,
			want: 1,
		},
		{
			in:   2,
			want: 2,
		},
		{
			in:   3,
			want: 3,
		},
		{
			in:   4,
			want: 5,
		},
		{
			in:   5,
			want: 8,
		},
		{
			in:   6,
			want: 13,
		},
		{
			in:   7,
			want: 21,
		},
	}

	for _, test := range tests {
		got := climbStairs(test.in)
		if got != test.want {
			t.Errorf("climbStairs(%d) = %d, want %d", test.in, got, test.want)
		}
	}

}
