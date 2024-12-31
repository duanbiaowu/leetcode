package leetcode

// simple solution
func generateParenthesisSimple(n int) []string {
	if n <= 0 {
		return []string{}
	}

	path := make([]byte, n<<1)
	var res []string

	generateAll(0, path, &res)

	return res
}

func generateAll(start int, path []byte, res *[]string) {
	if start == len(path) {
		if isValid(path) {
			*res = append(*res, string(path))
		}
		return
	}

	path[start] = '('
	generateAll(start+1, path, res)
	path[start] = ')'
	generateAll(start+1, path, res)
}

// 检测字符 `(` 和 `)` 数量是否匹配
func isValid(path []byte) bool {
	balance := 0

	for _, c := range path {
		if c == '(' {
			balance++
		} else {
			balance--
		}
		if balance < 0 {
			return false
		}
	}

	return balance == 0
}

// backtrack solution
func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}
	var res []string
	backtrack(n, n, "", &res)
	return res
}

func backtrack(left, right int, str string, res *[]string) {
	// 左右两边同时结束，说明对称
	if left == 0 && right == 0 {
		*res = append(*res, str)
		return
	}

	// 下一个字符使用 (
	if left > 0 {
		backtrack(left-1, right, str+"(", res)
	}
	// 下一个字符使用 )
	if right > 0 && right > left {
		backtrack(left, right-1, str+")", res)
	}
}

// backtrack2(0, 0, n, path, &res)
func backtrack2(open, close, n int, path []byte, res *[]string) {
	if len(path) == n<<1 {
		*res = append(*res, string(path))
		return
	}

	// 下一个字符使用 (
	if open < n {
		path = append(path, '(')
		backtrack2(open+1, close, n, path, res)
		path = path[:len(path)-1]
	}

	// 下一个字符使用 )
	if close < open {
		path = append(path, ')')
		backtrack2(open, close+1, n, path, res)
		path = path[:len(path)-1]
	}
}
