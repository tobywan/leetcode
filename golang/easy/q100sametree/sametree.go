package sametree

//
// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return Equal(p, q)
}

// Equal is defined and tested in
// github.com/tobywan/leetcode/types/tree/binarytree.go
func Equal(a, b *TreeNode) bool {

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
