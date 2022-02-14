package leetcode

func calculate(s string) int {
	var stack []int
	var opt uint8
	// 初始运算符
	opt = '+'
	num := 0

	// c 的类型是 int32
	// for _, c := range s
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			num = num*10 + int(s[i]-'0')
		}
		// 结束前计算一次: i == len(s) - 1
		if (!isDigit(s[i]) && s[i] != ' ') || i == len(s)-1 {
			switch opt {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				num *= stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, num)
			case '/':
				num = stack[len(stack)-1] / num
				stack = stack[:len(stack)-1]
				stack = append(stack, num)
			}

			num = 0
			opt = s[i]
		}
	}

	res := 0
	for len(stack) > 0 {
		res += stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return res
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
