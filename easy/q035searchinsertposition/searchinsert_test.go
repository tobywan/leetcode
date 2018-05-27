package searchinsertposition

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		in1  []int
		in2  int
		want int
	}{
		{
			in1:  []int{1, 3, 5, 6},
			in2:  5,
			want: 2,
		},
		{
			in1:  []int{1, 3, 5, 6},
			in2:  2,
			want: 1,
		},
		{
			in1:  []int{1, 3, 5, 6},
			in2:  7,
			want: 4,
		},
		{
			in1:  []int{1, 3, 5, 6},
			in2:  0,
			want: 0,
		},
		{
			in1:  []int{-27, -10, -5, -1},
			in2:  0,
			want: 4,
		},
		{
			in1:  []int{0},
			in2:  0,
			want: 0,
		},
		{
			in1:  []int{},
			in2:  20,
			want: 0,
		},
		{
			in1:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 20, 21},
			in2:  15,
			want: 10,
		},
		{
			in1:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 20, 21},
			in2:  1,
			want: 1,
		},
		{
			in1:  []int{-10},
			in2:  -20,
			want: 0,
		},
	}

	for _, test := range tests {
		//		t.Logf("Testing: %v", test.in1)
		got := searchInsert(test.in1, test.in2)
		if got != test.want {
			t.Errorf("searchInsert(%v, %d)=%d, want %d", test.in1, test.in2, got, test.want)
		}
	}
}
