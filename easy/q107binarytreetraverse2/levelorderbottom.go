package binarytreetraverse2

// TreeNode is a definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrderBottom is in github.com/tobywan/leetcode/types/tree/binarytree.go
func levelOrderBottom(root *TreeNode) [][]int {

	res := [][]int{}

	res = inOrderTraverse(root, 0, res)

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

// the level of a root node is 0
func inOrderTraverse(node *TreeNode, level int, seen [][]int) [][]int {
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
