package leetcode

func isValid(s string) bool {
	// 如果 s 的长度为奇数，肯定不是有效括号
	if len(s)&1 == 1 {
		return false
	}

	// 将左右边界字符进行分组并形成映射
	charMap := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	// 初始化栈
	var stack []byte

	for i := range s {
		if _, ok := charMap[s[i]]; ok {
			// 左边界组的字符直接入栈
			stack = append(stack, s[i])
		} else {
			// 右边界组的字符与栈顶元素进行比较
			index := len(stack) - 1
			if index < 0 || s[i] != charMap[stack[index]] {
				return false
			}
			stack = stack[:index]
		}
	}

	// 此时栈内可能还存在一些左边界字符
	// 所以需要判断栈是否为空
	return len(stack) == 0
}
