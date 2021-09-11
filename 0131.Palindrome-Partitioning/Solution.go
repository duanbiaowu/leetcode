package leetcode

func partition(s string) [][]string {
	n := len(s)
	if n == 0 {
		return nil
	}

	var res [][]string
	var path []string
	backtrack(s, 0, path, &res)
	return res
}

func backtrack(s string, begin int, path []string, res *[][]string) {
	n := len(s)
	if begin == n {
		*res = append(*res, append([]string(nil), path...))
		return
	}
	
	for i := begin; i < n; i++ {
		if isPal(s[begin : i+1]) {
			backtrack(s, i+1, append(path, s[begin:i+1]), res)
		}
	}
}

func isPal(s string) bool {
	for left, right := 0, len(s)-1; left < right; {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
