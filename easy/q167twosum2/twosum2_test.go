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
	}
	for _, test := range tests {
		got := twoSum(test.in1, test.in2)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("twoSum(%v,%d)=%v, want %v", test.in1, test.in2, got, test.want)
		}
	}
}
