package leetcode

func oneEditAway(first string, second string) bool {
	m, n := len(first), len(second)
	if m < n {
		first, second = second, first
		m, n = n, m
	}
	if m-n > 1 {
		return false
	}
	for i := 0; i < n; i++ {
		if first[i] != second[i] {
			// 只能编辑 1 次
			// 如果当前字符不一致，更改一个字符 OR 删除一个字符
			return first[i+1:] == second[i:] || first[i+1:] == second[i+1:]
		}
	}
	return true
}
