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

	var stack []int
	res := 0

	// i <= n [极端情况下，阶梯形状矩形]
	for i := 0; i <= n; i++ {
		var cur int
		if i == n {
			cur = 0
		} else {
			cur = heights[i]
		}

		// 当前高度小于栈内元素，则将栈内元素逐个弹出计算面积
		for len(stack) > 0 && cur < heights[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h := heights[pop]
			w := i
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				w = i - top - 1
			}
			res = max(res, h*w)
		}
		stack = append(stack, i)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
