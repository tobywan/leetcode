package pascalstriangle

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		in   int
		want [][]int
	}{
		{
			in:   0,
			want: [][]int{},
		},
		{
			in:   1,
			want: [][]int{{1}},
		},
		{
			in:   2,
			want: [][]int{{1}, {1, 1}},
		},
		{
			in:   3,
			want: [][]int{{1}, {1, 1}, {1, 2, 1}},
		},
		{
			in:   4,
			want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}},
		},
		{
			in:   5,
			want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			in:   6,
			want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}},
		},
	}

	for _, test := range tests {

		got := generate(test.in)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("generate(%d)=%v, want %v", test.in, got, test.want)
		}

	}
}
