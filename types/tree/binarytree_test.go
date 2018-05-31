package tree

import (
	"reflect"
	"testing"
)

func TestNewBT(t *testing.T) {

	i := map[int]int{1: 10, 2: 20, 3: 30}

	node := NewBT(i)

	want := []int{10, 20, 30}

	got := make([]int, 3, 3)
	got[0] = node.Val
	got[1] = node.Left.Val
	got[2] = node.Right.Val

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Wrong")
	}

}

func TestEqual(t *testing.T) {
	tests := []struct {
		in1  map[int]int
		in2  map[int]int
		want bool
	}{
		{
			in1:  map[int]int{1: 1, 2: 2},
			in2:  map[int]int{1: 1, 3: 2},
			want: false,
		},
		{
			in1:  map[int]int{1: 1, 2: 2, 3: 3},
			in2:  map[int]int{1: 1, 2: 2, 3: 3},
			want: true,
		},
		{
			in1:  map[int]int{1: 1, 2: 2, 3: 1},
			in2:  map[int]int{1: 1, 2: 1, 3: 2},
			want: false,
		},
	}
	for _, test := range tests {
		a := NewBT(test.in1)
		b := NewBT(test.in2)
		got := Equal(a, b)
		if got != test.want {
			t.Errorf("Equal(%v,%v)=%t", test.in1, test.in2, got)
		}
	}
}
