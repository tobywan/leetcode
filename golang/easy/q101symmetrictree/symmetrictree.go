package symmetrictree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)

}

// tested at github.com/tobywan/leetcode/types/tree/binarytree.go
func isMirror(a, b *TreeNode) bool {
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

// TreeNode is node of a tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
