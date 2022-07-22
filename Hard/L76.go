package Hard

/*
* TimeComplexity
	- o(m+n)
* SpaceComplexity
	- O(m+n)
*/

type memo struct {
	length     int
	leftIndex  int
	rightIndex int
}

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	count := 0

	left, right := 0, 0

	mp := make(map[byte]int)
	for i := range t {
		mp[t[i]]++
	}
	uniqueChar := len(mp)
	tempMp := make(map[byte]int)

	memObj := &memo{
		length: -1,
	}

	for right < len(s) {
		tempMp[s[right]]++
		if val, ok := mp[s[right]]; ok {
			if val == tempMp[s[right]] {
				count++
			}
		}

		for left <= right && count == uniqueChar {
			if memObj.length == -1 || memObj.length > right-left+1 {
				memObj.length = right - left + 1
				memObj.leftIndex = left
				memObj.rightIndex = right
			}

			tempMp[s[left]]--

			if val, ok := mp[s[left]]; ok {
				if val > tempMp[s[left]] {
					count--
				}
			}
			//contract
			left++
		}
		//expand
		right++
	}
	if memObj.length == -1 {
		return ""
	} else {
		return s[memObj.leftIndex : memObj.rightIndex+1]
	}
}
