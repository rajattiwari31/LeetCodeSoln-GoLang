package Easy

/*
Approach -
	- CHeck every element and add till that point
	- if it the sum till that point is smaller than the number at that point
	- then break the chain and start from that point

Time Complexity -
	- O(n)

Space Complexity -
	- O(1)
*/
func maxSubArray(nums []int) int {
	ans := nums[0]
	maxSoFar := nums[0]

	for i := 1; i < len(nums); i++ {
		maxSoFar = Max(maxSoFar+nums[i], nums[i])
		ans = Max(ans, maxSoFar)
	}
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
