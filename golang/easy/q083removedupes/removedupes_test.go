package removedupes

import (
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		in   []int
		want string
	}{
		{
			in:   []int{1, 2, 3},
			want: "1->2->3",
		}, {
			in:   []int{1, 1, 2},
			want: "1->2",
		},
		{
			in:   []int{1},
			want: "1",
		},
		{
			in:   []int{1, 1, 2, 2, 3, 3},
			want: "1->2->3",
		},
		{
			in:   []int{1, 1, 1, 1, 1, 1},
			want: "1",
		},
	}

	for _, test := range tests {
		head := NewListNode(test.in...)
		got := deleteDuplicates(head)

		if got.String() != test.want {
			t.Errorf("deleteDuplicates(%v)=%s, want %s", test.in, got, test.want)
		}
	}
}
