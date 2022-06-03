package Medium

/*
Approach -
	- This is greddy soln
	- here we are just checking the max jumps we can take
	- every index we are finding the farthest distance
	- if the index is greater than the max jumps we can take than return false
TimeComplexity -
	- o(n)
SpaceComplexity -
	- o(1)
*/
func canJump(nums []int) bool {
	max := 0
	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		}

		if i+nums[i] > max {
			max = i + nums[i]
		}
	}
	return true
}
