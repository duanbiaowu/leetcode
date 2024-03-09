package leetcode

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)

	// 从倒数第二天开始，因为最后一天必然是 0
	for i := n - 2; i >= 0; i-- {
		j := i + 1
		// 两个变量模拟单调栈
		for j < n {
			if temperatures[j] > temperatures[i] {
				// 找到了下一个更热的天
				// 更新结果，然后直接进行下一个循环
				res[i] = j - i
				break
			} else if res[j] == 0 {
				// 后面没有更热的天了，直接进行下一个循环
				break
			} else {
				// res[j] 表示下一个更热的天
				// 直接将 j 向前进 res[j] 步
				j += res[j]
			}
		}
	}

	return res
}

// 单调栈实现
func dailyTemperatures2(temperatures []int) []int {
	res := make([]int, len(temperatures))

	var stack []int
	for i := range temperatures {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			preIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[preIndex] = i - preIndex
		}

		stack = append(stack, i)
	}

	return res
}
