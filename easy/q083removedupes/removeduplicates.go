package removedupes

import (
	"bytes"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	return nil
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
