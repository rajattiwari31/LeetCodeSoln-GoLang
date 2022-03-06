package hard

/* We will use the stack approach, where we will track the valid Parnethesis.
Then calculate the gap between the valid parenthesis*/

func longestValidParentheses(s string) int {
	//"((()"
	//"))))))()"

	if len(s) == 0 {
		return 0
	}
	maxLength := 0

	//Create a stack for the tracking purpose
	st := make([]int, 0)

	//For first elem
	st = append(st, -1)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			st = append(st, i)
		} else {
			//pop the element
			st = st[:len(st)-1]

			//if stack is empty, i.e insert the index. Consider the example "))))("
			if len(st) == 0 {
				st = append(st, i)
			} else {
				maxLength = Max(maxLength, i-st[len(st)-1])
			}
		}
	}

	return maxLength
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
