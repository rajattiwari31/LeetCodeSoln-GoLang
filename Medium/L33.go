package Medium

/*
Approach:
	- As it is saying sorted i.e means we need to use Binary Search somehow
	- The challenge is here we first need to find the PIVOT index then do Binaray Search from there
	- In the first half using Binary search we find PIOVOT
	- then use that pivot index to do the real binary search for the index
	- IMPORTANT: as it is rotated sorted array, remember to do mod % n

Time Complexity:
	- binary Search so O(logn)
Space Complexity:
	- O(1)

*/
func search(nums []int, target int) int {
	n := len(nums)
	left := 0
	right := n - 1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	pivot := left
	return binarySearch(nums, pivot, pivot-1+n, target)
}

func binarySearch(nums []int, left, right, target int) int {

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid%len(nums)] == target {
			return mid % len(nums)
		} else if nums[mid%len(nums)] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
