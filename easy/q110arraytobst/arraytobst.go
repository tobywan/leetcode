package arraytobst

// TreeNode is node of a tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {

	root := toNode(nums)

	return root
}

func toNode(s []int) *TreeNode {
	l := len(s)

	if l == 0 {
		return nil
	}
	mid := l / 2

	n := &TreeNode{Val: s[mid]}

	n.Left = toNode(s[:mid])
	n.Right = toNode(s[mid+1:])

	return n
}
