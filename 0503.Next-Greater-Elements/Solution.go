package leetcode

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	for i := range nums {
		res[i] = -1

		j := i
		// 题目要求数组为循环数组，也就是说，检测每个元素的右侧最大值时，需要被检测的元素是除了当前元素外的其他所有元素
		// 长度也就是 N - 1
		for cnt := 1; cnt < n; cnt++ {
			j++
			if j == n {
				j = 0
			}
			if nums[j] > nums[i] {
				res[i] = nums[j]
				break
			}
		}
	}

	return res
}

func nextGreaterElementsOpt(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}

	// 单调栈，用于存储索引中间状态值
	stack := []int{}
	// 题目要求数组为循环数组，也就是说，检测每个元素的右侧最大值时，需要被检测的元素是除了当前元素外的其他所有元素
	// 长度也就是 N - 1

	// 如何实现循环数组呢？
	// 比较直观的方法是将数组容量翻倍，然后将当前所有元素追加一份到末尾，但是这会造成空间的浪费

	// 还有一种技巧: 使用取模运算，这样遍历时下标会控制在 0 ~ N - 1, 不会发生越界问题
	// 但是需要注意的是: 只遍历一次数组是不够的，例如数组 [2, 3, 1], 最后单调栈中的值为 [3, 1], 元素 1 的下一个最大值还是未知的
	// 此时可以将数组进行 “拉直翻倍”，将原有元素复制一份追加到数组后面
	// 当然并不会发生实际的 “扩容” 操作，我们使用取模运算来遍历数组两次

	for i := 0; i < n<<1-1; i++ {
		// 从前向后扫描，只要栈顶元素比当前元素小，更新栈顶元素 (数组索引) 对应的下一个最大值
		// 同时出栈顶部元素
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			res[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}

		// 当前索引入栈
		// 因为是处理循环数组，所以这里使用 i%n
		stack = append(stack, i%n)
	}

	return res
}
