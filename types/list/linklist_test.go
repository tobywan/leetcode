package list

import "testing"

func TestString(t *testing.T) {
	root := &Node{Val: 1}
	s := root.String()

	want := "1"
	if s != want {
		t.Errorf("Expected %q got %q", want, s)
	}

	root = NewNode(1, 2, 3)
	want = "1->2->3"
	s = root.String()
	if s != want {
		t.Errorf("Expected %q got %q", want, s)
	}
}

func TestAppend(t *testing.T) {
	root := NewNode(1, 2, 3, 4, 5, 6, 7, 8, 9, 0, -1, -4, -8)
	s := root.String()

	want := "1->2->3->4->5->6->7->8->9->0->-1->-4->-8"
	if s != want {
		t.Errorf("Expected %q got %q", want, s)
	}

}
func TestTail(t *testing.T) {
	root := NewNode(1, 2, 3, 4, 5, 6, 7, 8, 9, 0, -1, -4, -99)
	tail := root.Tail()
	s := tail.String()

	want := "-99"
	if s != want {
		t.Errorf("Expected %q got %q", want, s)
	}
}

func TestNilList(t *testing.T) {
	got := NewNode()
	if got != nil {
		t.Errorf("Expected nil list, got %v", got)
	}
}

func TestNoCycle(t *testing.T) {
	head := NewNode(1, 2, 3, 4, 5, 6, 7, 8, 9, 0, -1, -4, -99)
	want := false
	got := hasCycle(head)
	if want != got {
		t.Errorf("TestNoCycle failed. Expected %t, got %t", want, got)
	}
}

func TestHasCycle(t *testing.T) {
	stem := NewNode(1, 2, 3, 4, 5, 6)
	cycle := NewNode(11, 22, 33, 44, 55, 66, 17, 28, 39, 40, -1, -4, -99)

	tail := cycle.Tail()
	tail.Next = cycle
	tail = stem.Tail()
	tail.Next = stem
	want := true
	got := hasCycle(stem)
	if want != got {
		t.Errorf("TestNoCycle failed. Expected %t, got %t", want, got)
	}
}
