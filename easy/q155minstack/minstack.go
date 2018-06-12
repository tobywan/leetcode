package minstack

// MinStack is a stack thst keeps track of the minimum
type MinStack struct {
	stack []element
}

type element struct {
	value int
	min   int
}

// Constructor creates a new MinStack
func Constructor() MinStack {
	return MinStack{[]element{}}
}

// Push adds an item to the top of the stack
func (m *MinStack) Push(x int) {

	min := x
	l := len(m.stack)
	if l > 0 {
		min = m.stack[l-1].min
		if x < min {
			min = x
		}
	}

	m.stack = append(m.stack, element{x, min})

}

// Pop removes the top element
func (m *MinStack) Pop() {
	m.stack = m.stack[:len(m.stack)-1]
}

// Top reports the top element without changing the stack
func (m *MinStack) Top() int {
	return last(m.stack).value
}

func last(s []element) element {
	l := len(s)
	if l == 0 {
		return element{0, 0}
	}
	return s[l-1]
}

// GetMin returns the current minimum value
func (m *MinStack) GetMin() int {
	return last(m.stack).min
}
