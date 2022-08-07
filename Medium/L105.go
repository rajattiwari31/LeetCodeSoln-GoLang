package Medium

/*
TimeComplexity -
	- o(n)
SpaceComplexity -
	- o(n)
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	mp := make(map[int]int)

	for i := range inorder {
		mp[inorder[i]] = i
	}

	return helper(preorder, inorder, mp, 0, len(preorder)-1, 0, len(inorder)-1)
}

func helper(preorder, inorder []int, mp map[int]int, startPre, endPre, startIn, endIn int) *TreeNode {
	if startPre > endPre || startIn > endIn {
		return nil
	}

	rootVal := preorder[startPre]

	root := &TreeNode{
		Val: rootVal,
	}

	rootIndex := mp[rootVal]
	numbersLeft := rootIndex - startIn

	root.Left = helper(preorder, inorder, mp, startPre+1, startPre+numbersLeft, startIn, rootIndex-1)
	root.Right = helper(preorder, inorder, mp, startPre+1+numbersLeft, endPre, rootIndex+1, endIn)
	return root

}
