package Hard

/*

* BottomUp soln - DP
* TimeComplexity - O(MN)
* Approach -
	- First fill the first row and first column
	- Then check accordingly

	- example -     _ a a b b b
				 _  T F F F F F
				 a	F T F F F F
				 *  T T T F F F
				 b	F F F T F F
				 *  T T T T T T

*/

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(p)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s)+1)
	}

	for i := range dp {
		for j := range dp[0] {

			//Frst fill the first row and column
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = false
			} else if j == 0 {
				if p[i-1] == '*' {
					dp[i][j] = dp[i-2][j]
				} else {
					dp[i][j] = false
				}
			} else {
				if s[j-1] == p[i-1] || p[i-1] == '.' {
					dp[i][j] = dp[i-1][j-1]
				} else if p[i-1] == '*' {
					dp[i][j] = dp[i-2][j]

					if p[i-2] == '.' || s[j-1] == p[i-2] {
						dp[i][j] = dp[i][j] || dp[i][j-1]
					}
				} else {
					if s[j-1] != p[i-1] {
						dp[i][j] = false
					}
				}
			}
		}
	}

	return dp[len(p)][len(s)]
}
