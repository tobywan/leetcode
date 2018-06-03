package pathsum

// TreeNode is a definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return testSumPath(root, 0, sum)
}

func testSumPath(node *TreeNode, subTotal int, sum int) bool {

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
