package Medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}

	stack := make([]*TreeNode, 0)
	prev := &TreeNode{}
	curr := root
	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && prev.Val >= curr.Val {
			return false
		}

		prev = curr
		curr = curr.Right
	}
	return true
}
