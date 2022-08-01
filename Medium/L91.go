package Medium

/*
TimeComplexity -
	- Recurssion - 2^n
	- memoisation - o(n)

SpaceCompelxity - o(n)
*/

//Recursion
func numDecodingsRecurssion(s string) int {
	if s == "" {
		return 0
	}
	return recurHelper(0, s)
}

func recurHelper(pos int, s string) int {
	if pos == len(s) {
		return 1
	}

	if s[pos] == '0' {
		return 0
	}

	res := recurHelper(pos+1, s)

	if pos < len(s)-1 && (s[pos] == '1' || (s[pos] == '2' && s[pos+1] <= '6')) {
		res += recurHelper(pos+2, s)
	}

	return res
}

//Memoisation

func numDecodingsMemo(s string) int {
	if s == "" {
		return 0
	}
	arr := make([]int, len(s)+1)

	for i := range arr {
		arr[i] = -1
	}
	arr[len(s)] = 1
	return memoHelper(&arr, 0, s)
}

func memoHelper(arr *[]int, pos int, s string) int {
	if (*arr)[pos] != -1 {
		return (*arr)[pos]
	}

	if s[pos] == '0' {
		(*arr)[pos] = 0
		return (*arr)[pos]
	}

	res := memoHelper(arr, pos+1, s)

	if (pos < len(s)-1) && (s[pos] == '1' || (s[pos] == '2' && s[pos+1] <= '6')) {
		res += memoHelper(arr, pos+2, s)
	}

	(*arr)[pos] = res
	return (*arr)[pos]

}
