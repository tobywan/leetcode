package pascalstriangle2

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	tests := []struct {
		in   int
		want []int
	}{
		{
			in:   0,
			want: []int{1},
		},
		{
			in:   1,
			want: []int{1, 1},
		},
		{
			in:   2,
			want: []int{1, 2, 1},
		},
		{
			in:   3,
			want: []int{1, 3, 3, 1},
		},
		{
			in:   4,
			want: []int{1, 4, 6, 4, 1},
		},
		{
			in:   5,
			want: []int{1, 5, 10, 10, 5, 1},
		},
	}

	for _, test := range tests {

		got := getRow(test.in)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("getRow(%d)=%v, want %v", test.in, got, test.want)
		}

	}
}
