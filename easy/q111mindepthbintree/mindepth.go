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
	node  *TreeNode
	depth int
}

// minDepthRec is slower using recursion
func minDepthRec(root *TreeNode) int {
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
