package tree

import "fmt"

// Node is the definition for a binary tree node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// toMap returns a map representation of the tree
// i is the index of the current node, e.g 0 for root, 1 for root.Left etc
func toMap(n *Node, index int) map[int]int {
	if n == nil {
		return nil
	}
	ret := make(map[int]int)
	ret[index] = n.Val
	ml := toMap(n.Left, index*2)
	mr := toMap(n.Right, index*2+1)
	for k, v := range ml {
		ret[k] = v
	}
	for k, v := range mr {
		ret[k] = v
	}

	return ret
}

func levelOrderBottom(root *Node) [][]int {

	res := [][]int{}

	res = inOrderTraverse(root, 0, res)

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

// the level of a root node is 0
func inOrderTraverse(node *Node, level int, seen [][]int) [][]int {
	if node == nil {
		return seen
	}

	if level == len(seen) {
		seen = append(seen, []int{})
	}
	seen = inOrderTraverse(node.Left, level+1, seen)
	seen[level] = append(seen[level], node.Val)
	seen = inOrderTraverse(node.Right, level+1, seen)

	return seen
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
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
