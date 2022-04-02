package Medium

/*
	Algorithm -
		- Think as a basic Area of rectangle/square problem
		- Start from the widest container(length)
		- then check the breadth(height) one by one using two Pointer concept
	TimeComplexity
		- O(n)
	SpaceComplexity
		- O(1)
*/

func maxArea(height []int) int {

	start := 0
	end := len(height) - 1

	resultMax := 0

	for start <= end {
		temp := Min(height[end], height[start]) * (end - start)
		resultMax = Max(resultMax, temp)

		if height[end] > height[start] {
			start++
		} else {
			end--
		}
	}

	return resultMax
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
