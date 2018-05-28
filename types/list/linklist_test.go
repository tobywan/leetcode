package list

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

func TestAppend(t *testing.T) {
	root := NewListNode(1, 2, 3, 4, 5, 6, 7, 8, 9, 0, -1, -4, -8)
	s := root.String()

	want := "1->2->3->4->5->6->7->8->9->0->-1->-4->-8"
	if s != want {
		t.Errorf("Expected %q got %q", want, s)
	}

}
func TestTail(t *testing.T) {
	root := NewListNode(1, 2, 3, 4, 5, 6, 7, 8, 9, 0, -1, -4, -99)
	tail := root.Tail()
	s := tail.String()

	want := "-99"
	if s != want {
		t.Errorf("Expected %q got %q", want, s)
	}
}

func TestNilList(t *testing.T) {
	got := NewListNode()
	if got != nil {
		t.Errorf("Expected nil list, got %v", got)
	}
}
