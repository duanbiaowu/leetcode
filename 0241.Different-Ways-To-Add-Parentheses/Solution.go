package leetcode

import (
	"strconv"
)

var (
	memo = map[string][]int{} // 避免重复计算
)

// 分治过程
// 1. 按运算符分成两部分，分别求解
// 2. 实现计算表达式求解
// 3. 按运算符合并两部分的求解值
func diffWaysToCompute(expression string) []int {
	if val, ok := memo[expression]; ok {
		return val
	}

	// 数字直接返回
	if isDigit(expression) {
		val, _ := strconv.Atoi(expression)
		return []int{val}
	}

	var res []int
	for i, c := range expression {
		// 如果是运算符，则计算左右两边的计算表达式
		if c == '+' || c == '-' || c == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])

			memo[expression[:i]] = left
			memo[expression[i+1:]] = right

			for _, leftNum := range left {
				for _, rightNum := range right {
					var val int
					if c == '+' {
						val = leftNum + rightNum
					} else if c == '-' {
						val = leftNum - rightNum
					} else {
						val = leftNum * rightNum
					}
					res = append(res, val)
				}
			}
		}
	}

	return res
}

func isDigit(input string) bool {
	if len(input) == 0 {
		return false
	}

	for _, c := range input {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
