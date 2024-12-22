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
	m := make(map[string][]string)

	for i := range strs {
		s := []byte(strs[i])

		// 排序之后，异位单次的 key 都是一样的
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})

		key := string(s)
		m[key] = append(m[key], strs[i])
	}

	res := make([][]string, 0, len(m))
	for i := range m {
		res = append(res, m[i])
	}

	return res
}

func groupAnagrams3(strs []string) [][]string {
	// 得益于 Go 语言的特性，数组可以作为 map 的 key
	// 所以这里直接用数组来统计每个字符串中字母出现的次数
	// 对于不支持数组作为 key 的语言，可以使用字符串 (自定义编码 key) 来代替
	// 例如 a 的次数为 2，b 的次数为 3，可以编码为 "2#3"
	// 其中 2 和 3 分别代表 a 和 b 出现的次数，按照 a-z 的顺序排列
	m := make(map[[26]int][]string)

	for i := range strs {
		cnt := [26]int{}

		for j := range strs[i] {
			cnt[strs[i][j]-'a']++
		}

		m[cnt] = append(m[cnt], strs[i])
	}

	res := make([][]string, 0, len(m))
	for i := range m {
		res = append(res, m[i])
	}

	return res
}
