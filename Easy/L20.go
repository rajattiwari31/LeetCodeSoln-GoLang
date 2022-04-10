package Easy

/*
	Algorithm:
		- One by one check and put it in the stack
		- Consider these cases to handle
			- "(" only one elemnt in the stack
			- ")" nothing in the stack
		- Stack Algorithm

	TimeComplexity:
		- o(n) - iterating the string
	SpaceComplexity
		- o(n) - stack
*/

func isValid(s string) bool {

	stack := make([]rune, 0)

	for _, k := range s {

		if k == '(' || k == '{' || k == '[' {
			stack = append(stack, k)
		} else {
			if k == ')' && (len(stack) == 0 || stack[len(stack)-1] != '(') {
				return false
			} else if k == ']' && (len(stack) == 0 || stack[len(stack)-1] != '[') {
				return false
			} else if k == '}' && (len(stack) == 0 || stack[len(stack)-1] != '{') {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}
