package Medium

import (
	"sort"
	"strings"
)

/*
Approach -
	- First sort all the strings
	- then store the sorted key in the hashmap
	- now compare all the string with stored sorted string

Time Complexity -
	- O(nlogn) - for soting the key

Space Complexity -
	- o(n) - for the hashmap
*/

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	res := make([][]string, 0)

	mp := make(map[string][]string)

	for _, s := range strs {
		sortString := sortAnagram(s)
		if _, ok := mp[sortString]; !ok {
			mp[sortString] = []string{}
		}
		mp[sortString] = append(mp[sortString], s)
	}

	for _, value := range mp {
		res = append(res, value)
	}
	return res
}

func sortAnagram(str string) string {
	sortString := strings.Split(str, "")
	sort.Strings(sortString)
	return strings.Join(sortString, "")
}
