package Hard

import "math"

/* Naive solution -
	- Compare one by one
	- Time Complexity - O(n+m)
   Efficient soln -
	- just bother only about the middle element
	- dont need to sort the whole array
	- Use BinarySearch **Thats the catch
	- Time Complexity - O(log(n+m))
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var ans float64
	if len(nums1) == 0 && len(nums1) == len(nums2) {
		return ans
	}

	//nums1 will always be a smaller array
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	n := len(nums1)
	m := len(nums2)

	low := 0
	high := n

	for low <= high {
		mid := (low + high) / 2
		cut1 := mid
		cut2 := (n+m+1)/2 - cut1

		l1, l2 := math.MinInt, math.MinInt
		r1, r2 := math.MaxInt, math.MaxInt

		if cut1 != 0 {
			l1 = nums1[cut1-1]
		}
		if cut2 != 0 {
			l2 = nums2[cut2-1]
		}

		if cut1 != n {
			r1 = nums1[cut1]
		}
		if cut2 != m {
			r2 = nums2[cut2]
		}

		if l1 <= r2 && l2 <= r1 {
			if (n+m)%2 == 0 {
				ans = float64(Max(l1, l2)+Min(r1, r2)) / 2
			} else {
				ans = float64(Max(l1, l2))
			}
			return ans
		} else if l1 > r2 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
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
