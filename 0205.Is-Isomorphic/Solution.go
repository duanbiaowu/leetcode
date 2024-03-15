package leetcode

func isIsomorphic(s string, t string) bool {
	sMap := [128]int{}
	tMap := [128]int{}

	// 只需要记录连续字符的位置相对性
	for i := range s {
		if sMap[s[i]] != tMap[t[i]] {
			return false
		}

		if sMap[s[i]] == 0 {
			// 更新字符位置为最新索引 + 1
			// i + 1 是为了避免 map value 默认值为 0 引起的错误
			sMap[s[i]] = i + 1
			tMap[t[i]] = i + 1
		}
	}

	return true
}
