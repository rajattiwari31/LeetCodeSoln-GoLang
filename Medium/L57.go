package Medium

import "sort"

/*
Approach -
	- It is similar to L56 only thing here is first we need to insert the newInterval
	- then do the same thing from L56
	- also thing where there the newInterval is not overlaping one
TimeCompleixty -
	- o(n) - when the newInterval is overlaping one
	- o(nlogn) - when we need to sort/shift to insert newInterval
SpaceComplexity -
	- o(n*2) ~ o(n) - for creating new array
*/

func insert(intervals [][]int, newInterval []int) [][]int {
	if (intervals == nil || len(intervals) == 0) && newInterval == nil {
		return nil
	}
	res := [][]int{}
	overlaping := false

	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] >= newInterval[0] && intervals[i][0] <= newInterval[1] {
			intervals[i][0] = min(intervals[i][0], newInterval[0])
			intervals[i][1] = max(intervals[i][1], newInterval[1])
			overlaping = true
			break
		}
	}
	if !overlaping {
		intervals = append(intervals, newInterval)
		sort.SliceStable(intervals, func(i, j int) bool {
			return intervals[i][0] < intervals[j][0]
		})
		return intervals
	}

	idx := 0
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		if res[idx][1] >= intervals[i][0] && res[idx][0] <= intervals[i][1] {
			res[idx][0] = min(intervals[i][0], res[idx][0])
			res[idx][1] = max(intervals[i][1], res[idx][1])
		} else {
			res = append(res, intervals[i])
			idx++
		}
	}
	return res
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
