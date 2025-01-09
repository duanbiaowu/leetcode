package leetcode

import (
	"strconv"
	"strings"
)

// Example:
// 3[a]2[bc] => aaabcbc
func decodeString(s string) string {
	var stack []string

	for i := range s {
		// 如果不是 ']' 直接入栈
		if s[i] != ']' {
			stack = append(stack, string(s[i]))
			continue
		}

		// 如果是 ']' 则计算 [] 内的字符串的重复拼接结果

		// 1. 计算 [] 内的字符串
		str := ""
		for top := len(stack) - 1; top >= 0 && stack[top] != "["; top-- {
			str = stack[top] + str
			stack = stack[:top]
		}

		// 2. pop '['
		if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}

		// 3. 计算 [] 前面的数字
		num := ""
		for top := len(stack) - 1; top >= 0 && isDigits(stack[top]); top-- {
			num = stack[top] + num
			stack = stack[:top]
		}

		// 4. 计算重复的字符串
		repeatStr := ""
		repeatCnt, _ := strconv.Atoi(num)
		for i := 0; i < repeatCnt; i++ {
			repeatStr += str
		}

		// 5. 将当前结果作为一个 “新的子字符串” 存入栈内
		// Example: 3[a2[c]] = acc acc acc
		// 当计算完 2[c] 后, `cc`直接存入栈内，下次就可以将 a2[c] 解析为 `acc`
		stack = append(stack, repeatStr)
	}

	return strings.Join(stack, "")
}

func isDigits(str string) bool {
	return str >= "0" && str <= "9"
}
