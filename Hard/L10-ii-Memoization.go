package Hard

/*

* Memoization soln [DP]
* Time Complexity - O(MN)
* Approach -
	- Similar to what we did in Recursion
	- Here we will introduced a cache where we will store the result
	- This will avoid the unnecessary Recusrive calls
	- Catch is when we have '*' there we need to check,
		- if we need to take the previous char
		- or We dont need to.
*/

func isMatch(s string, p string) bool {
	//Creating a 2D slice
	mem := make([][]int, len(s)+1)
	for i := range mem {
		mem[i] = make([]int, len(p)+1)
	}

	//Intiliazing the slice with -1
	for i := range mem {
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	return isMatchUtil(s, p, 0, 0, &mem)
}

func isMatchUtil(s, p string, i, j int, mem *[][]int) bool {
	//If the elemnent is present then just return
	if (*mem)[i][j] != -1 {
		if (*mem)[i][j] == 0 {
			return false
		}
		return true
	}

	if i >= len(s) && j >= len(p) {
		(*mem)[i][j] = 1
		return true
	}

	if j >= len(p) {
		(*mem)[i][j] = 0
		return false
	}

	matchFisrtChar := (i < len(s)) && (s[i] == s[j] || p[j] == '.')

	if j+1 < len(p) && p[j+1] == '*' {
		checkBool := isMatchUtil(s, p, i, j+2, mem) ||
			(matchFisrtChar && isMatchUtil(s, p, i+1, j, mem))
		if checkBool {
			(*mem)[i][j] = 1
			return true
		}
	} else if matchFisrtChar {
		checkBool := isMatchUtil(s, p, i+1, j+1, mem)
		if checkBool {
			(*mem)[i][j] = 1
			return true
		}
	}

	(*mem)[i][j] = 0
	return false
}
