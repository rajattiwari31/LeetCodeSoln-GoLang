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

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) != 0 {
		n := len(q)
		tempArr := make([]int, 0)

		for n > 0 {
			curr := q[0]
			q = q[1:]

			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
			tempArr = append(tempArr, curr.Val)
			n--
		}

		res = append(res, tempArr)
	}

	return res
}
