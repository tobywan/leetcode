package tree

import "fmt"

// Node is the definition for a binary tree node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func isSymmetric(root *Node) bool {
	return isMirror(root.Left, root.Right)

}
func isMirror(a, b *Node) bool {
	switch {
	case a == nil && b == nil:
		return true
	case a == nil || b == nil:
		return false
	case a.Val != b.Val:
		return false
	default:
		return isMirror(a.Left, b.Right) && isMirror(a.Right, b.Left)
	}
}

// Equal compares two trees for equality. Nil nodes are equal
func Equal(a, b *Node) bool {

	switch {
	case a == nil && b == nil:
		return true
	case a == nil || b == nil:
		return false
	case a.Val != b.Val:
		return false
	case !Equal(a.Left, b.Left):
		return false
	default:
		return Equal(a.Right, b.Right)
	}
}

// NewBT creates a new tree
//         {1:1,2:2,3:3, 5:5,6:6}
//             1
//           /   \
//          2     3
//          \     /
//           5   6
//
//                1
//           /         \
//          2           3
//        /   \       /   \
//       4     5     6     7
//      / \   / \   / \   / \
// .    8  9 10 11 12 13 14 15
//
//
// node = 2^n -1 + m
// left child = 2^n+1 - 1 + 2(m - 1) +1
// = 2^n+1 +2m -2
// = 2(2^n+1 -1 +m )
//                1
//           /         \
//          10          11
//        /   \       /   \
//      100   101   110   111
//      / \   / \   / \   / \
//     8   9 10 11 12 13 14 15
// nth level, m from the left is (1<<(n-1)) + m
// parent of node n is n/2

// NewBT generates a new tree from a map of node number to value.
// where node number is the position in the tree :
// assumes well formed discription, or will panic
// e.g.
//                1
//           /         \
//          2           3
//        /   \       /   \
//       4     5     6     7
//      / \   / \   / \   / \
// .    8  9 10 11 12 13 14 15
func NewBT(vals map[int]int) *Node {

	todo := len(vals)

	done := make(map[int]*Node)
	rootval, ok := vals[1]

	if !ok {
		panic("no root node")
	}
	done[1] = &Node{Val: rootval}
	todo--

	for i := 2; todo > 0; i++ {
		val, ok := vals[i]

		if !ok {
			continue
		}
		iParent := i / 2
		even := i&1 == 0

		parent, ok := done[iParent]
		if !ok {
			panic(fmt.Errorf("[%d]%d has no parent", i, val))
		}
		done[i] = &Node{Val: val}
		if even {
			parent.Left = done[i]
		} else {
			parent.Right = done[i]
		}
		todo--

	}
	return done[1]

}
