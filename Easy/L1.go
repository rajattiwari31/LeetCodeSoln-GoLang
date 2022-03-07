package Easy

// Time Complexity = o(n)
//Space COmplexity = o(n)
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	res := make([]int, 0)

	// a + b = target
	// a = target-b
	for i, k := range nums {
		if _, found := mp[target-k]; found {
			res = append(res, mp[target-k])
			res = append(res, i)
			return res
		}

		mp[k] = i
	}
	return res
}
