package Medium

/*
Approach:
	- Traverse right->down->left->up
	- when traversing letf and up focus on the boundaru condtions
	- appproach 2 has the condition where we check the length of the result array
	- to the toatal elements
	- also increment and dercrement the index after evry iteration

TimeComplexity-
	- O(n*m)
SpaceComplexity-
	- o(n*m) - array
*/

//Approach 2
func spiralOrder(matrix [][]int) []int {
	ans := make([]int, 0)

	if matrix == nil || len(matrix) == 0 {
		return ans
	}
	n, m := len(matrix), len(matrix[0])
	rowBegin, rowEnd := 0, n-1
	colBegin, colEnd := 0, m-1

	for len(ans) < n*m {

		//Traversing right
		for j := colBegin; j <= colEnd && len(ans) < n*m; j++ {
			ans = append(ans, matrix[rowBegin][j])
		}
		rowBegin++

		//Travesing down
		for i := rowBegin; i <= rowEnd && len(ans) < n*m; i++ {
			ans = append(ans, matrix[i][colEnd])
		}
		colEnd--

		//Travesing left
		for j := colEnd; j >= colBegin && len(ans) < n*m; j-- {
			ans = append(ans, matrix[rowEnd][j])
		}
		rowEnd--

		//Traversing Up
		for i := rowEnd; i >= rowBegin && len(ans) < n*m; i-- {
			ans = append(ans, matrix[i][colBegin])
		}
		colBegin++
	}
	return ans
}

/* // Approach - 1
func spiralOrder(matrix [][]int) []int {

	if len(matrix) == 0 {
		return nil
	}

	ans := make([]int, 0)

	rowBegin, rowEnd := 0, len(matrix)-1
	colBegin, colEnd := 0, len(matrix[0])-1

	for colBegin <= colEnd && rowBegin <= rowEnd {

		//traversing right
		for j := colBegin; j <= colEnd; j++ {
			ans = append(ans, matrix[rowBegin][j])
		}
		rowBegin++

		//down
		for i := rowBegin; i <= rowEnd; i++ {
			ans = append(ans, matrix[i][colEnd])
		}
		colEnd--

		//left
		if rowBegin <= rowEnd {
			for j := colEnd; j >= colBegin; j-- {
				ans = append(ans, matrix[rowEnd][j])
			}
			rowEnd--
		}

		if colBegin <= colEnd {
			for j := rowEnd; j >= rowBegin; j-- {
				ans = append(ans, matrix[j][colBegin])
			}
			colBegin++
		}
	}

	return ans
}*/
