package leetcode

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	var res []string
	var backtrack func(left, right int, path string)
	backtrack = func(left, right int, path string) {
		if left > n || right > left {
			return
		}
		if len(path) == n<<1 {
			res = append(res, path)
		} else {
			backtrack(left+1, right, path+"(")
			backtrack(left, right+1, path+")")
		}
	}

	backtrack(0, 0, "")
	return res
}
