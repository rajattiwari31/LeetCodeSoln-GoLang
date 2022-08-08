package Easy

/*
TimeComplexity-
	- o(n)
SpaceCOmplexity
	- o(1)
*/

func maxProfit(prices []int) int {
	result := 0

	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		minPrice = min(minPrice, prices[1])

		result = max(result, prices[i]-minPrice)
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
