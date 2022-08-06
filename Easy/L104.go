package Easy

/*
TimeComplexity -
	- o(number of nodes)
SpaceComplexity -
	- o(number of nodes)
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	lTree := maxDepth(root.Left) + 1
	rTree := maxDepth(root.Right) + 1

	if lTree > rTree {
		return lTree
	}

	return rTree
}
