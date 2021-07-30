package leetcode

// simple solution
//func generateParenthesis(n int) []string {
//	if n <= 0 {
//		return []string{}
//	}
//	var res []string
//	generateAll(make([]byte, n*2), 0, &res)
//	return res
//}

func generateAll(current []byte, index int, res *[]string) {
	if index == len(current) {
		if isValid(current) {
			*res = append(*res, string(current))
		}
	} else {
		current[index] = '('
		generateAll(current, index+1, res)
		current[index] = ')'
		generateAll(current, index+1, res)
	}
}

func isValid(current []byte) bool {
	balance := 0
	for _, c := range current {
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
	if left > 0 {
		backtrack(left-1, right, str+"(", res)
	}
	if right > 0 && right > left {
		backtrack(left, right-1, str+")", res)
	}
}
