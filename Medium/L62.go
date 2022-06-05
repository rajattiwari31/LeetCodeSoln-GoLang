package Medium

/*
Approach -
	- Start from the end and build the array
	- add index from down and right of that block
TimeComplexity -
	- o(n*n)
SpaceCompelixty -
	- o(1)
*/

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}

	for i := range res {
		for j := range res[i] {
			if i == m-1 || j == n-1 {
				res[i][j] = 1
			} else {
				res[i][j] = 0
			}
		}
	}
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			res[i][j] = res[i+1][j] + res[i][j+1]
		}
	}

	return res[0][0]
}
