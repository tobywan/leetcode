package timetosell

import (
	"testing"
)

func BenchmarkMaxProfit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxProfit([]int{1, 4, 65, 8, 90, 45, 3, 234, 5, 7, 8, 9, 8, 6, 4, 3, 3, 455, 7, 8, 8, 5, 456, 74, 78, 34, 6, 7, 8, 8, 65, 5, 34, 3, 3, 4, 5, 677, 8, 99, 7, 65, 4})
	}
}

func BenchmarkMaxProfit1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxProfit1([]int{1, 4, 65, 8, 90, 45, 3, 234, 5, 7, 8, 9, 8, 6, 4, 3, 3, 455, 7, 8, 8, 5, 456, 74, 78, 34, 6, 7, 8, 8, 65, 5, 34, 3, 3, 4, 5, 677, 8, 99, 7, 65, 4})
	}
}

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
