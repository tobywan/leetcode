package mergetwolists

import (
	"bytes"
	"strconv"
)

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var a, b, ret *ListNode
	ret, l1, l2 = choose(l1, l2)
	if ret == nil {
		// both were empty
		return nil
	}
	a, l1, l2 = choose(l1, l2)
	ret.Next = a
	for a != nil {
		b, l1, l2 = choose(l1, l2)
		a.Next = b
		a = b
	}
	return ret
}

// choose picks the appropriate head value and
// the remainder of the lists as n1 and n2
func choose(l1, l2 *ListNode) (choice, n1, n2 *ListNode) {

	if l1 == nil && l2 == nil {
		return nil, nil, nil
	}

	if l1 == nil {
		choice = l2
		n1 = nil
		n2 = l2.Next
		return choice, n1, n2
	}
	if l2 == nil {
		choice = l1
		n1 = l1.Next
		n2 = nil
		return choice, n1, n2
	}
	if l2.Val <= l1.Val {
		choice = l2
		n1 = l1
		n2 = l2.Next
		return choice, n1, n2
	}
	choice = l1
	n1 = l1.Next
	n2 = l2
	return choice, n1, n2

}

/////////////////////////
// local code only for test and display

// ListNode is a Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode returns the root of a linked list
// represented by the slice values
func NewListNode(values ...int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	root := &ListNode{Val: values[0]}
	l1 := root

	for _, i := range values[1:] {
		l2 := l1.Append(i)
		l1 = l2
	}

	return root

}

// Append adds a node
func (l *ListNode) Append(i int) *ListNode {
	n := &ListNode{i, nil}
	l.Next = n
	return n
}

// Tail gets te last node, O(n)
func (l *ListNode) Tail() *ListNode {
	n := l
	for {
		if n.Next == nil {
			break
		}
		n = n.Next
	}
	return n
}

func (l *ListNode) String() string {
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
