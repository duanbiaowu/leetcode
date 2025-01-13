package leetcode

// https://leetcode-cn.com/problems/longest-valid-parentheses/solution/zui-chang-you-xiao-gua-hao-by-leetcode-solution/
func longestValidParentheses(s string) int {
	res := 0
	left, right := 0, 0

	// 贪心算法:
	// 从前向后遍历
	// 逐个计算以每个字符结尾的有效括号长度
	for _, c := range s {
		if c == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			res = max(res, left<<1)
		} else if right > left {
			// 当右括号数量大于左括号数量时
			// 说明当前的左右边界组成的子括号不符合规则
			// 这时就置为 0, 开始重新计算
			left, right = 0, 0
		}
	}

	// 为什么还需要再次从头向前遍历？
	// 兼容极端的边界情况 (遍历时左括号数量始终大于右括号数量)
	// 例如: (()
	// 此时正确的有效括号长度应该是 2
	// 如果只有从前向后遍历，那么计算出的结果是 0
	// 所以需要从头向前再遍历一次
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			res = max(res, left<<1)
		} else if left > right {
			left, right = 0, 0
		}
	}

	return res
}

func longestValidParenthesesStack(s string) int {
	res := 0
	// 边界情况: 始终保持栈底元素为当前已经遍历过的元素中「最后一个没有被匹配右括号的 左括号下标」
	stack := []int{-1}

	for i := 0; i < len(s); i++ {
		// 遇到左括号时直接入栈
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			// 遇到右括号时先将栈顶元素出栈，表示匹配当前的左括号
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				// 如果栈顶元素出栈后，栈变为空，说明之前栈内的元素只有一个 -1
				// 也就是当前的右括号，无法正常匹配
				// 既然无法匹配，那就直接将当前的右括号加入到栈内
				// 更新「最后一个没有被匹配右括号的 左括号下标」
				stack = append(stack, i)
			} else {
				// 如果栈顶元素出栈后，栈不为空
				// 当前有括号的索引减去栈顶元素就等于「以该右括号为结尾的最长有效括号的长度」
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}

	return res
}

func longestValidParenthesesDp(s string) int {
	res := 0

	// 状态转移方程:
	// 1. 当 s[i-1] = '(' && s[i] = ')'
	//    例如 ..()() 这种形式
	//    dp[i] = dp[i−2] + 2
	// 2. 当 s[i-1] = ')' && s[i] = ')"
	//    例如 ..)) 这种形式
	//    如果 s[ i - dp[i-1] - 1 ] = '('
	//    那么 dp[i] = dp[i-1] + dp[ i - dp[i-1] - 2 ] + 2
	dp := make([]int, len(s))

	for i := 1; i < len(s); i++ {
		if s[i] != ')' {
			continue
		}

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

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
