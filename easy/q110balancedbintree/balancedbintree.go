package balancedbintree

// TreeNode is A definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// tested at github.com/tobywan/leetcode/types/tree/binarytree.go
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, ok := maxSubTreeHeight(root)
	return ok
}

// maxSubTreeHeight gets the height at a node, but if its left
// and right trees differ by more than 1, return ok -> false
func maxSubTreeHeight(n *TreeNode) (height int, ok bool) {

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

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
