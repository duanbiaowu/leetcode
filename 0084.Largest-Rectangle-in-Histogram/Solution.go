package leetcode

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}

	var monoStack []int
	for i := 0; i < n; i++ {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= heights[i] {
			right[monoStack[len(monoStack)-1]] = i
			monoStack = monoStack[:len(monoStack)-1]
		}

		if len(monoStack) == 0 {
			left[i] = -1
		} else {
			left[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, i)
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
