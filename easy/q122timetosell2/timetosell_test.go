package timetosell2

import "testing"

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
			in:   []int{7, 1, 5, 3, 6, 4},
			want: 7,
		},
		{
			in:   []int{1, 2, 3, 4, 5},
			want: 4,
		},
		{
			in:   []int{7, 13, 1, 5},
			want: 10,
		},
		{
			in:   []int{7, 6, 4, 3, 1},
			want: 0,
		},
		{
			in:   []int{1, 1, 1, 1, 1, 1},
			want: 0,
		},
		{
			in:   []int{1, 1, 1, 1, 5, 1, 1, 1, 1, 1},
			want: 4,
		},
	}
	for _, test := range tests {
		got := maxProfit(test.in)
		if got != test.want {
			t.Errorf("maxProfit(%#v)=%d, want %d", test.in, got, test.want)
		}
	}
}
