package removeelement

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		in1  []int
		in2  int
		want []int
	}{
		{
			in1:  []int{0, 1, 2, 2, 3, 0, 4, 2},
			in2:  2,
			want: []int{0, 1, 3, 0, 4},
		},
		{
			in1:  []int{3, 2, 2, 3},
			in2:  2,
			want: []int{3, 3},
		},
		{
			in1:  []int{1},
			in2:  1,
			want: []int{},
		},
		{
			in1:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			in2:  -1,
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		},
		{
			in1:  []int{1, 1, 1, 1, 3, 1, 1, 1, 1},
			in2:  1,
			want: []int{3},
		},
		{
			in1:  []int{},
			in2:  3,
			want: []int{},
		},
	}

	for _, test := range tests {
		got1 := removeElement(test.in1, test.in2)
		got2 := test.in1[:got1]
		if !reflect.DeepEqual(got2, test.want) {
			t.Errorf("removeElement=%v(%d), want %v(%d)", got1, got2, test.want, len(test.want))
		}
	}
}
