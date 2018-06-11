package list

import (
	"bytes"
	"strconv"
)

////////////////////////
// local code only for test and display

// Node is a Definition for singly-linked list.
type Node struct {
	Val  int
	Next *Node
}

// NewNode returns the root of a linked list
// represented by the slice values
func NewNode(values ...int) *Node {
	if len(values) == 0 {
		return nil
	}
	root := &Node{Val: values[0]}
	l1 := root

	for _, i := range values[1:] {
		l2 := l1.Append(i)
		l1 = l2
	}

	return root

}

func hasCycle(head *Node) bool {
	fast := head
	slow := head
	atEnd := false

	for {
		fast, atEnd = advance(fast, 2)
		if atEnd {
			return false
		}
		slow, _ = advance(slow, 1)
		if fast == slow {
			return true
		}
	}

}

func advance(node *Node, steps int) (child *Node, atEnd bool) {
	child = node
	for i := 0; i < steps; i++ {
		child = child.Next
		if child == nil {
			return nil, true
		}
	}
	return child, false
}

// Append adds a node
func (l *Node) Append(i int) *Node {
	n := &Node{i, nil}
	l.Next = n
	return n
}

// Tail gets te last node, O(n)
func (l *Node) Tail() *Node {
	n := l
	for {
		if n.Next == nil {
			break
		}
		n = n.Next
	}
	return n
}

func (l *Node) String() string {
	var b bytes.Buffer
	n := l
	b.WriteString(strconv.FormatInt(int64(l.Val), 10))
	for {
		if n.Next == nil {
			break
		}
		n = n.Next
		b.WriteString("->")
		b.WriteString(strconv.FormatInt(int64(n.Val), 10))
	}
	return b.String()
}
