package leetcode

func partition(s string) [][]string {
	if len(s) == 0 {
		return nil
	}

	var res [][]string
	var path []string

	backtrack(s, 0, path, &res)

	return res
}

// 回溯的优化: isPal是否回文子串预处理
// https://leetcode-cn.com/problems/palindrome-partitioning/solution/hui-su-you-hua-jia-liao-dong-tai-gui-hua-by-liweiw/
func backtrack(s string, start int, path []string, res *[][]string) {
	if start == len(s) {
		*res = append(*res, append([]string{}, path...))
		return
	}

	for i := start; i < len(s); i++ {
		if isPal(s[start : i+1]) {
			backtrack(s, i+1, append(path, s[start:i+1]), res)
		}
	}
}

func isPal(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
