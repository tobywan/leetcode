package timetosell

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{},
			want: 0,
		},
		{
			in:   []int{7, 6, 4, 3, 1},
			want: 0,
		},
		{
			in:   []int{7, 1, 5, 3, 6, 4},
			want: 5,
		},
		{
			in:   []int{7, 13, 1, 5},
			want: 6,
		},
		{
			in:   []int{7, 13, 1, 8},
			want: 7,
		},
	}
	for _, test := range tests {
		got := maxProfit(test.in)
		if got != test.want {
			t.Errorf("maxProfit(%#v)=%d, want %d", test.in, got, test.want)
		}
	}
}
