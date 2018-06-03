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

// NewBalancedBT creates a balanced BT where the left and right subtrees
// of each node differ in hight by a max of two
func NewBalancedBT(sorted []int) (root *Node) {

	root = toNode(sorted)

	return root
}

func toNode(s []int) *Node {
	l := len(s)

	if l == 0 {
		return nil
	}
	mid := l / 2

	n := &Node{Val: s[mid]}

	n.Left = toNode(s[:mid])
	n.Right = toNode(s[mid+1:])

	return n
}

func minDepth(root *Node) int {
	if root == nil {
		return 0
	}
	queue := []nodeDepth{{root, 1}}
	i := 0

	for {
		n := queue[i].node
		if n.Left == nil && n.Right == nil {
			return queue[i].depth
		}
		if n.Left != nil {
			queue = append(queue, nodeDepth{n.Left, queue[i].depth + 1})
		}
		if n.Right != nil {
			queue = append(queue, nodeDepth{n.Right, queue[i].depth + 1})
		}
		i++
	}
	// shouldn't be here

}

type nodeDepth struct {
	node  *Node
	depth int
}

func minDepthRec(root *Node) int {
	if root == nil {
		return 0
	}
	return minSubDepth([]*Node{root}, 1)
}

func minSubDepth(level []*Node, levelDepth int) int {

	next := make([]*Node, 0, len(level)*2)
	for _, n := range level {
		if n.Left == nil && n.Right == nil {
			return levelDepth
		}
		if n.Left != nil {
			next = append(next, n.Left)
		}
		if n.Right != nil {
			next = append(next, n.Right)
		}

	}
	return minSubDepth(next, levelDepth+1)
}

func hasPathSum(root *Node, sum int) bool {
	if root == nil {
		return false
	}
	return testSumPath(root, 0, sum)
}

func testSumPath(node *Node, subTotal int, sum int) bool {

	subTotal = subTotal + node.Val

	if node.Left == nil && node.Right == nil {
		return sum == subTotal
	}
	res := false
	if node.Left != nil {
		res = testSumPath(node.Left, subTotal, sum)
	}
	if res {
		return res
	}
	if node.Right != nil {
		res = testSumPath(node.Right, subTotal, sum)
	}

	return res

}

func isBalanced(root *Node) bool {
	_, ok := maxSubTreeHeight(root)
	return ok
}

// maxSubTreeHeight gets the height at a node, but if its left
// and right trees differ by more than 1, return ok -> false
func maxSubTreeHeight(n *Node) (height int, ok bool) {

	if n.Left == nil && n.Right == nil {
		return 1, true
	}

	hl, hr := 0, 0
	if n.Left != nil {
		hl, ok = maxSubTreeHeight(n.Left)

		if !ok {
			return -1, false
		}
	}
	if n.Right != nil {
		hr, ok = maxSubTreeHeight(n.Right)
		if !ok {
			return -1, false
		}
	}

	diff := hl - hr
	if diff > 1 || diff < -1 {
		return -1, false
	}

	return max(hl, hr) + 1, true

}
