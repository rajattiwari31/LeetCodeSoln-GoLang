package Medium

/*
Approach -
	- First, reverse the matrix vertically
	- second, reverse the matrix diagonally(Left-right diagonal)

TimeComplexity-
	- We have checked every single block once + rversing directly the array block
	- O(M) - m is the no. of block
SpaceComplexity-
	- O(1) - no use of DS

*/
func rotate(matrix [][]int) {
	if len(matrix) <= 1 {
		return
	}
	start, end := 0, len(matrix)-1

	// reverse vertically
	for start < end {
		tmp := matrix[start]
		matrix[start] = matrix[end]
		matrix[start] = tmp

		start++
		end--
	}

	//Now reverse diagonally
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[i]); j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
}
