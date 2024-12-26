package leetcode

// 单调栈解决类似问题：对数组中每一个元素找到第一个比自己小的元素
func largestRectangleArea(heights []int) int {
	n := len(heights)
	// left[0...n] 	各元素左边界: 左边第一个比自己小的数
	// right[0...n] 各元素右边界: 右边第一个比自己小的数
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}

	var stack []int
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	res := 0
	for i := 0; i < n; i++ {
		res = max(res, (right[i]-left[i]-1)*heights[i])
	}
	return res
}

func largestRectangleArea2(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	res := 0

	// 单调递增栈
	var stack []int
	// i <= n [极端情况下，阶梯形状矩形]
	for i := 0; i <= n; i++ {
		var cur int
		if i == n {
			cur = 0
		} else {
			cur = heights[i]
		}

		// 当前高度小于栈内元素，则将栈内元素逐个弹出计算面积
		// 也就是递增栈被破坏了
		for len(stack) > 0 && cur < heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 栈顶元素为矩形高度
			h := heights[top]
			// 先以当前元素索引为宽度
			// Example: [2, 1, 5, 6, 2, 3]
			// 当 i = 1 时，h = 2, w = 1, area = 2
			w := i

			// 如果栈不为空，更新宽度为: 当前元素索引 - 比栈顶元素更小的元素索引 - 1
			if len(stack) > 0 {
				w = i - stack[len(stack)-1] - 1
			}

			res = max(res, h*w)
		}

		stack = append(stack, i)
	}

	return res
}

func largestRectangleArea3(heights []int) int {
	res := 0
	stack := []int{-1}

	for i := range heights {
		for len(stack) != 1 && heights[i] < heights[stack[len(stack)-1]] {
			// 矩形高度 = 栈顶元素高度
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			// 矩形宽度 = 当前元素索引 - 比栈顶元素次小的元素索引 - 1
			w := i - stack[len(stack)-1] - 1
			res = max(res, h*w)
		}

		stack = append(stack, i)
	}

	// 如果栈内还有元素，继续计算面积
	// 因为栈内元素是递增的，所以栈内元素的右边界都是数组长度
	// 主要针对极端场景: 阶梯形状矩形, 递增栈不会被破坏
	// Example: [1, 2, 3, 4, 5, 6, 7]
	for len(stack) != 1 {
		// 矩形高度 = 栈顶元素高度
		h := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]

		// 矩形宽度 = 数组长度 - 比栈顶元素次小的元素索引 - 1
		// 相当于一个已排序数组，从后向前遍历计算 “后缀和“
		w := len(heights) - stack[len(stack)-1] - 1
		res = max(res, h*w)
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func LargestRectangleArea(heights []int) int {
	return largestRectangleArea3(heights)
}
