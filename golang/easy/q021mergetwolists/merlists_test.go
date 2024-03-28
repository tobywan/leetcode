package mergetwolists

import "testing"

func TestString(t *testing.T) {
	root := &ListNode{Val: 1}
	s := root.String()

	want := "1"
	if s != want {
		t.Errorf("Expected %q got %q", want, s)
	}

	root = NewListNode(1, 2, 3)
	want = "1->2->3"
	s = root.String()
	if s != want {
		t.Errorf("Expected %q got %q", want, s)
	}
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			l1:   nil,
			l2:   NewListNode(1, 3, 4),
			want: NewListNode(1, 3, 4),
		},
		{
			l1:   NewListNode(1, 2, 4),
			l2:   nil,
			want: NewListNode(1, 2, 4),
		},
		{
			l1:   nil,
			l2:   nil,
			want: nil,
		},
		{
			l1:   NewListNode(1, 2, 4),
			l2:   NewListNode(1, 3, 4),
			want: NewListNode(1, 1, 2, 3, 4, 4),
		},
	}
	for _, test := range tests {
		got := mergeTwoLists(test.l1, test.l2)
		if !equal(test.want, got) {
			t.Errorf("mergeTwoLists(%q, %q) = %q, want %q", test.l1, test.l2, got, test.want)
		}
	}
}

func equal(l1, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	if l1 == nil || l2 == nil {
		return false
	}
	return l1.String() == l2.String()
}
