package Medium

import "sort"

/*
	Algorithm:
		- Similar to the TWO sum problem
		- Use Math's Associative Property - a + b + c = 0 => a + b = -c
		- The key point here is -
			- Skip the repetative element once found the triplets in order to avoid the duplicate triplets
	TimeComplexity:
		- Builtin SORT function takes - o(nlog(n))
		- Algo takes - O(n^2)
		- overall - O(n^2)
	SpaceComplexity:
		- O(1)
*/

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	if len(nums) <= 2 {
		return res
	}

	n := len(nums)
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		target := -nums[i]

		start := i + 1
		end := n - 1

		for start < end {
			sum := nums[start] + nums[end]

			if target > sum {
				start++
			} else if target < sum {
				end--
			} else {
				temp := []int{nums[i], nums[start], nums[end]}
				res = append(res, temp)

				//skip the duplicate
				for start < end && nums[start] == temp[1] {
					start++
				}

				//skip the duplicate last elem
				for start < end && nums[end] == temp[2] {
					end--
				}
			}
		}
		for i < n-2 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
