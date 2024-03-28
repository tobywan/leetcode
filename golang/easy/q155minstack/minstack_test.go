package minstack

import "testing"

func TestStack(t *testing.T) {
	// MinStack minStack = new MinStack();
	s := Constructor()
	// minStack.push(-2);
	s.Push(-2)
	// minStack.push(0);
	s.Push(0)
	// minStack.push(-3);
	s.Push(-3)
	// minStack.getMin();   --> Returns -3.
	x := s.GetMin()
	expect(-3, x, t)
	// minStack.pop();
	s.Pop()
	// minStack.top();      --> Returns 0.
	x = s.Top()
	expect(0, x, t)
	// minStack.getMin();   --> Returns -2
	x = s.GetMin()
	expect(-2, x, t)

}

func expect(want, got int, t *testing.T) {
	if want != got {
		t.Errorf("Wanted %d got %d", want, got)
	}
}
