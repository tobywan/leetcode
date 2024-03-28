package removeduplicates

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		in   []int
		want []int
	}{
		{
			in:   []int{},
			want: []int{},
		},
		{
			in:   []int{1},
			want: []int{1},
		},
		{
			in:   []int{1, 1},
			want: []int{1},
		},
		{
			in:   []int{1, 1, 2},
			want: []int{1, 2},
		},
		{
			in:   []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: []int{0, 1, 2, 3, 4},
		},
	}
	for _, test := range tests {

		got := removeDuplicates(test.in)
		if !reflect.DeepEqual(test.in[:got], test.want) {
			t.Errorf("removeDuplicates got %v(%d), want %v(%d)", test.in, got, test.want, len(test.want))
		}
	}
}
