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

func TestIsMirror(t *testing.T) {
	tests := []struct {
		in1  map[int]int
		want bool
	}{
		{
			in1:  map[int]int{1: 1},
			want: true,
		},
		{
			in1:  map[int]int{1: 1, 2: 2, 3: 2},
			want: true,
		},
		{
			in1:  map[int]int{1: 1, 2: 2, 3: 2, 4: 3, 5: 4, 6: 4, 7: 3},
			want: true,
		},
		{
			in1:  map[int]int{1: 1, 2: 2, 3: 2, 5: 3, 7: 3},
			want: false,
		},
	}
	for _, test := range tests {
		a := NewBT(test.in1)
		got := isSymmetric(a)
		if got != test.want {
			t.Errorf("isSymmetric(%v)=%t", test.in1, got)
		}
	}
}

func TestMaxDepth(t *testing.T) {

	tests := []struct {
		in1  map[int]int
		want int
	}{
		{
			in1:  map[int]int{1: 1},
			want: 1,
		},
		{
			in1:  map[int]int{1: 1, 2: 2, 3: 2},
			want: 2,
		},
		{
			in1:  map[int]int{1: 1, 2: 2, 3: 2, 4: 3, 5: 4, 6: 4, 7: 3},
			want: 3,
		},
		{
			in1:  map[int]int{1: 3, 2: 9, 3: 20, 6: 15, 7: 7},
			want: 3,
		},
		{
			in1:  map[int]int{1: 3, 2: 9, 4: 20, 8: 15, 16: 7},
			want: 5,
		},
	}
	for _, test := range tests {
		a := NewBT(test.in1)
		got := maxDepth(a)
		if got != test.want {
			t.Errorf("maxDepth(%v)=%d, want %d", test.in1, got, test.want)
		}
	}

}
