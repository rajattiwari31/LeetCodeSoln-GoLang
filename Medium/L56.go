package Medium

import "sort"

/*
Approach -
	- the array can be in any order so first we need to sort the array on the basis of the first element
TimeComplexity-
	- o(nlogn) because of sorting
SpaceComplexity
	- o(n*2) ~ o(n)
*/

func merge(intervals [][]int) [][]int {
	if intervals == nil || len(intervals) == 0 {
		return [][]int{}
	}

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	res = append(res, intervals[0])
	idx := 0
	for i := 1; i < len(intervals); i++ {
		if res[idx][1] > intervals[i][0] && res[idx][0] < intervals[i][1] {
			res[idx][0] = min(res[idx][0], intervals[i][0])
			res[idx][1] = max(res[idx][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
			idx++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
