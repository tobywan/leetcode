package mindepthbintree

// TreeNode is A definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return minSubDepth([]*TreeNode{root}, 1)
}

func minSubDepth(level []*TreeNode, levelDepth int) int {

	next := make([]*TreeNode, 0, len(level)*2)
	for _, n := range level {
		if n == nil {
			continue
		}
		if n.Left == nil && n.Right == nil {
			return levelDepth
		}
		next = append(next, n.Left, n.Right)
	}
	return minSubDepth(next, levelDepth+1)
}
