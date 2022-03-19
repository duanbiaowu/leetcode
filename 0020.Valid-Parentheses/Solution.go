package leetcode

func isValid(s string) bool {
	n := len(s)
	if n&1 == 1 {
		return false
	}

	charMap := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	var stack []byte

	for i := 0; i < n; i++ {
		if _, ok := charMap[s[i]]; ok {
			stack = append(stack, s[i])
		} else {
			index := len(stack) - 1
			if len(stack) == 0 || s[i] != charMap[stack[index]] {
				return false
			}
			stack = stack[:index]
		}
	}

	return len(stack) == 0
}
