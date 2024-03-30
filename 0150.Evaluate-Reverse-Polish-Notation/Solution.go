package leetcode

import "strconv"

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	var stack []int

	for _, token := range tokens {
		// 将当前 token 解析为数字
		val, err := strconv.Atoi(token)

		// 如果没有发生错误
		// 说明当前 token 为数字，直接入栈
		if err == nil {
			stack = append(stack, val)
		} else {
			// 说明当前 token 为运算符
			// 取出栈顶的两个数字
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			// 根据不同的运算符执行不同的计算
			// 同时将计算结果加入到栈中s
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}

	// 返回栈底元素
	return stack[0]
}
