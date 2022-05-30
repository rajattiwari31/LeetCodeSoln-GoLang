package Medium

import (
	"strings"
)

/*
Approach -
	Approach 1
		- First sort all the strings
		- then store the sorted key in the hashmap
		- now compare all the string with stored sorted string
	Approach 2
		- Instead of sorting the string, we should create a unique ID
		- Then we will use this UniqueID to check in the hashmap
		- By this we can reduce the O(nlogn) -> o(n)
Time Complexity -
	- O(nlogn) - for soting the key
	- o(n) 	- In second approach

Space Complexity -
	- o(n) - for the hashmap
*/

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	res := make([][]string, 0)

	mp := make(map[string][]string)
	/* Approach 1
	for _, s := range strs {
		sortString := sortAnagram(s)
		if _, ok := mp[sortString]; !ok {
			mp[sortString] = []string{}
		}
		mp[sortString] = append(mp[sortString], s)
	}*/

	for _, s := range strs {
		strId := getID(s)
		mp[strId] = append(mp[strId], s)
	}

	for _, value := range mp {
		res = append(res, value)
	}
	return res
}

/* Approach - 1
func sortAnagram(str string) string {
	sortString := strings.Split(str, "")
	sort.Strings(sortString)
	return strings.Join(sortString, "")
}*/

// Approach 2
func getID(s string) string {
	arr := make([]byte, 26)

	for _, k := range s {
		arr[k-'a']++
	}

	var sb strings.Builder

	for _, k := range arr {
		sb.WriteByte(k)
	}

	return sb.String()
}
