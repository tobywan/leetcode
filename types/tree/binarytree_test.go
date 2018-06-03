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

func TestLevelOrderBottom(t *testing.T) {

	tests := []struct {
		in1  map[int]int
		want [][]int
	}{
		{
			in1:  map[int]int{1: 1, 2: 2, 3: 3},
			want: [][]int{{2, 3}, {1}},
		},
		{
			in1:  map[int]int{1: 1},
			want: [][]int{{1}},
		},
		{
			in1:  map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7},
			want: [][]int{{4, 5, 6, 7}, {2, 3}, {1}},
		},
		{
			in1:  map[int]int{1: 3, 2: 9, 3: 20, 6: 15, 7: 7},
			want: [][]int{{15, 7}, {9, 20}, {3}},
		},
		{
			in1:  map[int]int{1: 3, 2: 9, 4: 20, 8: 15, 16: 7},
			want: [][]int{{7}, {15}, {20}, {9}, {3}},
		},
	}
	for _, test := range tests {
		a := NewBT(test.in1)
		got := levelOrderBottom(a)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("levelOrderBottom(%v)=%v, want %v", test.in1, got, test.want)
		}
	}

}

func TestToMap(t *testing.T) {

	tests := []struct {
		in1  map[int]int
		want map[int]int
	}{
		{
			in1:  map[int]int{1: 1, 2: 2, 3: 3},
			want: map[int]int{1: 1, 2: 2, 3: 3},
		},
		{
			in1:  map[int]int{1: 1},
			want: map[int]int{1: 1},
		},
		{
			in1:  map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7},
			want: map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7},
		},
		{
			in1:  map[int]int{1: 3, 2: 9, 3: 20, 6: 15, 7: 7},
			want: map[int]int{1: 3, 2: 9, 3: 20, 6: 15, 7: 7},
		},
		{
			in1:  map[int]int{1: 3, 2: 9, 4: 20, 8: 15, 16: 7},
			want: map[int]int{1: 3, 2: 9, 4: 20, 8: 15, 16: 7},
		},
	}
	for _, test := range tests {
		a := NewBT(test.in1)
		got := toMap(a, 1)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("toMap(%v)=%v, want %v", test.in1, got, test.want)
		}
	}

}

func TestNewBalancedBT(t *testing.T) {
	tests := []struct {
		in   []int
		want map[int]int
	}{
		{
			in:   []int{1},
			want: map[int]int{1: 1},
		},
		{
			in:   []int{1, 2},
			want: map[int]int{1: 2, 2: 1},
		},
		{
			in:   []int{1, 2, 3},
			want: map[int]int{1: 2, 2: 1, 3: 3},
		},
		{
			in:   []int{1, 2, 3, 4, 5},
			want: map[int]int{1: 3, 2: 2, 4: 1, 3: 5, 6: 4},
		},
	}
	for _, test := range tests {
		w := NewBT(test.want)
		got := NewBalancedBT(test.in)
		if !Equal(got, w) {
			t.Errorf("NewBalancedBT(%v)=%v, want %v", test.in, toMap(got, 1), test.want)
		}
	}
}

func BenchmarkMinDepth(b *testing.B) {
	a := NewBT(map[int]int{1: 1, 2: 2, 3: 2, 4: 3, 5: 4, 6: 4, 7: 3})
	for i := 0; i < b.N; i++ {
		minDepth(a)
	}
}
func BenchmarkMinDepthRec(b *testing.B) {
	a := NewBT(map[int]int{1: 1, 2: 2, 3: 2, 4: 3, 5: 4, 6: 4, 7: 3})
	for i := 0; i < b.N; i++ {
		minDepthRec(a)
	}
}

func TestMinDepth(t *testing.T) {

	tests := []struct {
		in1  map[int]int
		want int
	}{
		// {
		// 	in1:  map[int]int{1: 1},
		// 	want: 1,
		// },
		{
			in1:  map[int]int{1: 1, 2: 2},
			want: 2,
		},
		{
			in1:  map[int]int{1: 1, 2: 2, 3: 2, 4: 3, 5: 4, 6: 4, 7: 3},
			want: 3,
		},
		{
			in1:  map[int]int{1: 3, 2: 9, 3: 20, 6: 15, 7: 7},
			want: 2,
		},
		{
			in1:  map[int]int{1: 3, 2: 9, 4: 20, 8: 15, 16: 7},
			want: 5,
		},
	}
	for _, test := range tests {
		a := NewBT(test.in1)
		got := minDepth(a)
		if got != test.want {
			t.Errorf("minDepth(%v)=%d, want %d", test.in1, got, test.want)
		}
	}

}

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		in1  map[int]int
		in2  int
		want bool
	}{
		{
			in1:  map[int]int{1: 1},
			in2:  1,
			want: true,
		},
		{
			in1:  map[int]int{1: 1, 2: 2},
			in2:  2,
			want: false,
		},
		{
			in1:  map[int]int{1: 1, 2: 2},
			in2:  3,
			want: true,
		},
		{
			in1:  map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7},
			in2:  7,
			want: true,
		},
		{
			in1:  map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7},
			in2:  8,
			want: true,
		},
		{
			in1:  map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7},
			in2:  10,
			want: true,
		},
		{
			in1:  map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7},
			in2:  11,
			want: true,
		},
		{
			in1:  map[int]int{1: -1, 2: -2, 3: -3, 6: -4, 7: -5},
			in2:  -3,
			want: true,
		},
		{
			in1:  map[int]int{1: 3, 2: 9, 4: 20, 8: 15, 16: 7},
			in2:  54,
			want: true,
		},
	}
	for _, test := range tests {
		a := NewBT(test.in1)
		got := hasPathSum(a, test.in2)
		if got != test.want {
			t.Errorf("hasPathSum(%v,%d)=%t", test.in1, test.in2, got)
		}
	}

}
