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

	// 1. 初始化单调栈
	// 因为题目要求寻找 “下一个” 更高的温度对应的索引 (间隔天数)
	// 所以栈内存储索引
	var stack []int

	// 2. 遍历数组
	for i := range temperatures {
		// 当前栈内索引对应的温度依次递减，符合单调递减的特性
		// 也就是说，如果一个索引出现在栈内
		// 说明还未找到该索引对应的温度 的  “下一个” 更高的温度

		// 3. 维护单调递减性
		// 如果当前元素温度 大于 栈顶元素温度
		// 说明当前单调递减性已经被破坏
		// 那么当前元素就是  “下一个” 更高的温度
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			// 4. 处理具体逻辑

			// 获取栈顶元素索引
			preIndex := stack[len(stack)-1]
			// 栈顶元素出栈
			stack = stack[:len(stack)-1]

			// 更新栈顶元素索引对应的温度 的  “下一个” 更高的温度
			// 因为题目要求寻找 “下一个”  更高的温度在几天后
			// 所以这里使用当前索引 减去 栈顶索引
			// 得出的差值就是中间的间隔天数
			res[preIndex] = i - preIndex
		}

		// 将当前元素的索引入栈
		stack = append(stack, i)
	}

	return res
}
