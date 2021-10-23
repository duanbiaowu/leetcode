package leetcode

// https://leetcode-cn.com/problems/longest-valid-parentheses/solution/zui-chang-you-xiao-gua-hao-by-leetcode-solution/
func longestValidParentheses(s string) int {
	res, left, right := 0, 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			res = max(res, right*2)
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			res = max(res, left*2)
		} else if left > right {
			left, right = 0, 0
		}
	}

	return res
}

func longestValidParenthesesStack(s string) int {
	res := 0
	// 边界情况: 始终保持栈底元素为当前已经遍历过的元素中「最后一个没有被匹配的右括号的下标」
	stack := []int{-1}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}

	return res
}

func longestValidParenthesesDp(s string) int {
	res := 0
	dp := make([]int, len(s))

	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			res = max(res, dp[i])
		}
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
