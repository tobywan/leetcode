package twosum

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		in1  []int
		in2  int
		want []int
	}{
		{
			in1:  []int{1, 2, 4, 8, 16, 32, 64, 128},
			in2:  192,
			want: []int{6, 7},
		},
		{
			in1:  []int{1, 2, 4, 8, 16, 32, 64, 128},
			in2:  10,
			want: []int{1, 3},
		},
		{
			in1:  []int{1, 2, 4, 8, 16, 32, 64, 128},
			in2:  129,
			want: []int{0, 7},
		},
		{
			in1:  []int{2, 7, 11, 1},
			in2:  9,
			want: []int{0, 1},
		},
		{
			in1:  []int{3, 2, 4},
			in2:  6,
			want: []int{1, 2},
		},
		{
			in1:  []int{3, 3},
			in2:  6,
			want: []int{0, 1},
		},
	}
	for _, test := range tests {
		got := twoSum(test.in1, test.in2)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("twoSum(%v,%d)=%v, want%v", test.in1, test.in2, got, test.want)
		}
	}

}
