package minstack

// MinStack is a stack thst keeps track of the minimum
// TODO what is the current min after popping
type MinStack struct {
	stack []int
	mins  []int
}

// Constructor creates a new MinStack
func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

// Push adds an item to the top of the stack
func (m *MinStack) Push(x int) {
	m.stack = append(m.stack, x)

	l := len(m.mins)

	if l == 0 {
		m.mins = append(m.mins, x)
		return
	}

	min := m.mins[l-1]

	if x < min {
		min = x
	}
	m.mins = append(m.mins, min)

}

// Pop removes the top element
func (m *MinStack) Pop() {
	m.stack = m.stack[:len(m.stack)-1]
	m.mins = m.mins[:len(m.mins)-1]
}

// Top reports the top element withoug changing the stack
func (m *MinStack) Top() int {
	return last(m.stack)
}

func last(s []int) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	return s[l-1]
}

// GetMin returns the current minimum value
func (m *MinStack) GetMin() int {
	return last(m.mins)
}
