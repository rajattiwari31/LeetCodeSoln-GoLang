package Medium

/*

Time Complexity-
	-> finding the first character is straight forward - O(m*n)
	-> After first character we have 4 options to go over O(4^k)
	-> O(m*n*4^k)
	-> More info https://cs.stackexchange.com/questions/96626/whats-the-big-o-runtime-of-a-dfs-word-search-through-a-matrix

*/

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[0] {
			if word[0] == board[i][j] {
				if validate(board, word, i, j, 0) {
					return true
				}
			}
		}
	}

	return false
}

func validate(board [][]byte, word string, i, j, k int) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == '*' {
		return false
	}

	if word[k] != board[i][j] {
		return false
	}

	if k == len(word)-1 {
		return true
	}

	temp := board[i][j]
	board[i][j] = '*'

	dfsCheck := validate(board, word, i+1, j, k+1) ||
		validate(board, word, i-1, j, k+1) ||
		validate(board, word, i, j+1, k+1) ||
		validate(board, word, i, j-1, k+1)

	board[i][j] = temp
	return dfsCheck
}
