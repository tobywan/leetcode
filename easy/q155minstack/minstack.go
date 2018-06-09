package minstack

// MinStack is a stack thst keeps track of the minimum
// TODO what is the current min after popping
type MinStack struct {
	stack []int
	min   int
}

// Constructor creates a new MinStack
func Constructor() MinStack {
	return MinStack{[]int{}, 0}
}

// Push adds an item to the top of the stack
func (m *MinStack) Push(x int) {
	m.stack = append(m.stack, x)

	if x < m.min {
		m.min = x
	}

}

// Pop removes the top element
func (m *MinStack) Pop() {
	m.stack = m.stack[:len(m.stack)-1]
}

// Top reports the top element withoug changing the stack
func (m *MinStack) Top() int {
	l := len(m.stack)
	if l == 0 {
		return 0
	}
	return m.stack[l-1]
}

// GetMin returns the current minimum value
func (m *MinStack) GetMin() int {
	return m.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
