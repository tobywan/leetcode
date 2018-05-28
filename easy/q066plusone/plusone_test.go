package plusone

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		in   []int
		want []int
	}{
		{
			in:   []int{1, 2, 3},
			want: []int{1, 2, 4},
		},
		{
			in:   []int{4, 3, 2, 1},
			want: []int{4, 3, 2, 2},
		},
		{
			in:   []int{0},
			want: []int{1},
		},
		{
			in:   []int{1, 0, 9},
			want: []int{1, 1, 0},
		},
		{
			in:   []int{9, 9, 9, 9, 9, 9, 9, 9, 9},
			want: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, test := range tests {
		got := plusOne(test.in)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("plusOne(%v)=%v, want %v", test.in, got, test.want)
		}
	}
}
