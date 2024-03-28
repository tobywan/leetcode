package mergesorted

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		in1  []int
		n1   int
		in2  []int
		n2   int
		want []int
	}{
		{
			in1:  []int{1, 2, 3, 0, 0, 0},
			n1:   3,
			in2:  []int{2, 5, 6},
			n2:   3,
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			in1:  []int{2, 2, 2, 2, 0, 0, 0, 0, 0, 0},
			n1:   4,
			in2:  []int{1, 1},
			n2:   2,
			want: []int{1, 1, 2, 2, 2, 2},
		},
		{
			in1:  []int{},
			n1:   0,
			in2:  []int{},
			n2:   0,
			want: []int{},
		},
		{
			in1:  []int{2, 4, 6, 8, 0, 0, 0, 0},
			n1:   4,
			in2:  []int{1, 3, 5, 7},
			n2:   4,
			want: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, test := range tests {
		merge(test.in1, test.n1, test.in2, test.n2)

		if !reflect.DeepEqual(test.in1[:test.n1+test.n2], test.want) {
			t.Errorf("merge(%v, %d, %v, %d), want %v", test.in1, test.n1, test.in2, test.n2, test.want)
		}
	}

}
