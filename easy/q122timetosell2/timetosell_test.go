package timetosell2

import "testing"

func BenchmarkMaxProfit(b *testing.B) {
	in := []int{1, 2, 3, 5, 7, 54, 3, 2, 1, 3, 5, 6, 9, 0, 9, 7, 5, 34, 23, 3, 4, 5, 67, 8, 99, 75, 4, 3, 23}
	for i := 0; i < b.N; i++ {
		maxProfit(in)
	}

}

func BenchmarkMaxProfit1(b *testing.B) {
	in := []int{1, 2, 3, 5, 7, 54, 3, 2, 1, 3, 5, 6, 9, 0, 9, 7, 5, 34, 23, 3, 4, 5, 67, 8, 99, 75, 4, 3, 23}
	for i := 0; i < b.N; i++ {
		maxProfit1(in)
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
