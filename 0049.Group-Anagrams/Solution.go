package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
	// prime 与 A～Z 对应，英文中出现概率越大的字母，选用的质数越小
	var prime = []int{
		5, 71, 37, 29, 2,
		53, 59, 19, 11, 83,
		79, 31, 43, 13, 7,
		67, 97, 23, 17, 3,
		41, 73, 47, 89, 61,
		101,
	}

	var encode = func(s string) int {
		res := 1
		for i := range s {
			res *= prime[s[i]-'a']
		}
		return res
	}

	m := map[int][]string{}
	for i := 0; i < len(strs); i++ {
		code := encode(strs[i])
		m[code] = append(m[code], strs[i])
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func groupAnagrams2(strs []string) [][]string {
	m := map[string][]string{}
	for i := 0; i < len(strs); i++ {
		s := []byte(strs[i])
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedS := string(s)
		m[sortedS] = append(m[sortedS], strs[i])
	}

	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
