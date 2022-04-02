package Medium

/*
	Algorithm -
		- Try to think about Exapnding from the middle
		- Think about the logic where we just need to check the leftMost and RightMost of the string and
		- than exapnd the string, we can apply the same algo inorder to solve it by DP
		- Need to do it one for OddLength & also for the EvenLength
	TimeComplexity-
		- O(n^2) - iterating one by one and then expanding from the middle
	SpaceComplexity
		- O(1) 	 - No Extra Usage, all are constant variables
*/

func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	var result string

	for i := 0; i < n-1; i++ {
		var temp string
		oddLength := expandLongest(s, i, i)
		evenLength := expandLongest(s, i, i+1)
		if len(oddLength) > len(evenLength) {
			temp = oddLength
		} else {
			temp = evenLength
		}

		if len(temp) > len(result) {
			result = temp
		}
	}

	return result

}

func expandLongest(s string, low, high int) string {

	for low >= 0 && high < len(s) {
		if s[low] != s[high] {
			break
		}
		low--
		high++
	}
	return s[low+1 : high]
}
