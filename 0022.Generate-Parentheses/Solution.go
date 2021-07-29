package leetcode

// simple solution
func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}
	var res []string
	generateAll(make([]byte, n*2), 0, &res)
	return res
}

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
