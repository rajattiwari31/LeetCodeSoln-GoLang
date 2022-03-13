package Medium

/*

* Example -  s = "bbbbbadefs"
		   ans = "badefs" - 6
		   - s = "abcabc"
		    ans = "abc" - 3

* Approach -
		- Brute force approach -
				- Take two loops and iterate and check in the HashMap
				- It the elem is there then break out from the loop
				- Time Complexity - O(n^2)
		- We can improve this by using sliding window approach in one loop
		- We will have two index - start and end
		- Our sub string lies within the start and end.
		- We will iterate end, if we find the dupilcate then we increment start by 1
		- Delete all the previously elements from the start index
		- Increment start
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	mp := make(map[byte]int)
	start := 0
	result := 1

	for end := 0; end < len(s); end++ {
		if k, found := mp[s[end]]; found {
			result = Max(result, end-start)

			for j := start; j <= k; j++ {
				delete(mp, s[j])
			}

			start = k + 1
		}
		mp[s[end]] = end
	}
	result = Max(result, len(s)-start) // If the element is not repeated till the end
	return result
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
