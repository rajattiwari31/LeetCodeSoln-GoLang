package Medium

/*
Approach:
	- whenever problem requires combination we will use BACKTRACKING concept
	- CLICK will be, combinations to find
	- when coding in GOLANG
		- think about the pointers
		- copy the slice and then append other wise wont get correct result
*/
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	backtracking(&res, candidates, []int{}, 0, 0, target)
	return res
}

func backtracking(res *[][]int, cand, temp []int, currSum, pos, target int) {
	if currSum == target {
		temp2 := make([]int, len(temp))
		copy(temp2, temp)
		*res = append(*res, temp2)
		return
	} else {
		for i := pos; i < len(cand); i++ {
			if currSum > target {
				continue
			}
			temp = append(temp, cand[i])
			backtracking(res, cand, temp, currSum+cand[i], i, target)
			temp = temp[:len(temp)-1]
		}
	}
}
