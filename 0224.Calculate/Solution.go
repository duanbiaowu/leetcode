package leetcode

func calculate(s string) int {
	// 使用一个栈来记录操作类型
	ops := []int{1}
	// 标记区分操作类型 (1: 加) (-1: 减)
	sign := 1

	var res int
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			// 空格忽略，继续向前扫描
			i++
		case '+':
			// 更新操作类型
			sign = ops[len(ops)-1]
			i++
		case '-':
			// 更新操作类型
			sign = -ops[len(ops)-1]
			i++
		case '(':
			// 当前操作类型入栈
			ops = append(ops, sign)
			i++
		case ')':
			// 最后的操作类型出栈
			// 因为最后一个操作类型在上次遇到 +/- 字符之后就更新了，所以后面就用不到了，这里出栈丢掉即可
			ops = ops[:len(ops)-1]
			i++
		default:
			// 扫描到数字之后，解析出具体的数字
			num := 0
			for ; i < n && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			res += num * sign
		}
	}

	return res
}
