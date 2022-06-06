package Easy

//Approach - 2
func climbStairs(n int) int {
	mp := make(map[int]int)

	mp[1], mp[2], mp[3] = 1, 2, 3

	return dp(mp, n)
}

func dp(mp map[int]int, n int) int {
	if _, ok := mp[n]; ok {
		return mp[n]
	}

	mp[n] = dp(mp, n-2) + dp(mp, n-1)
	return mp[n]
}

// Approach 1
// func climbStairs(n int) int {
// 	if n <= 2 {
// 		return n
// 	}

// 	dp := make([]int, n)
// 	dp[n-2] = 1
// 	dp[n-1] = 1

// 	for i := n - 3; i >= 0; i-- {
// 		dp[i] = dp[i+2] + dp[i+1]
// 	}

// 	return dp[0]
// }
