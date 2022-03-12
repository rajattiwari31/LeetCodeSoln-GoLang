package Hard

/*

* Brute Force soln
* Time Comlexity O(2^n)
* Approach
	- Match the first char
		- if the char is != '*' && and it is matching
			- Recusrively move to the next one
		- if the char is '*' then we have two option
				- Either we will take it
						- check first is the fisrt character is matching
						- increase the index for the s
						- keep the index same for the p
				- Dont take it
						- increase the index for the P by + 2

* Example - Str - "aaabb", pattern - "a*.*"
*/

func isMatch(s string, p string) bool {
	return isMatchUtil(s, p, 0, 0)
}

func isMatchUtil(s, p string, i, j int) bool {

	//If we reached at the end of the string & pattern i.e we are done
	if i >= len(s) && j >= len(p) {
		return true
	}

	//If we reached the end of the pattern but not the string i.e str is not matched with the pattern
	if j >= len(p) {
		return false
	}

	matchFisrtChar := i < len(s) && (s[i] == p[j] || p[j] == '.')

	if j+1 < len(p) && p[j+1] == '*' {
		return isMatchUtil(s, p, i, j+2) ||
			(matchFisrtChar && isMatchUtil(s, p, i+1, j))
	} else if matchFisrtChar {
		return isMatchUtil(s, p, i+1, j+1)
	}
	return false
}
