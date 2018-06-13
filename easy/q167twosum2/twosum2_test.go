package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		in1  []int
		in2  int
		want []int
	}{
		{
			in1:  []int{2, 7, 11, 15},
			in2:  9,
			want: []int{1, 2},
		},
		{
			in1:  []int{2, 7, 11, 15},
			in2:  17,
			want: []int{1, 4},
		},
		{
			in1:  []int{2, 7, 11, 15},
			in2:  26,
			want: []int{3, 4},
		},
		{
			in1:  []int{2, 7},
			in2:  9,
			want: []int{1, 2},
		},
		{
			in1:  []int{1, 2, 3, 5, 8, 13, 21, 34, 55},
			in2:  39,
			want: []int{4, 8},
		},
		{
			in1:  []int{1, 2, 3, 5, 8, 13, 21, 34, 55},
			in2:  60,
			want: []int{4, 9},
		},
	}
	for _, test := range tests {
		got := twoSum(test.in1, test.in2)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("twoSum(%v,%d)=%v, want %v", test.in1, test.in2, got, test.want)
		}
	}
}
