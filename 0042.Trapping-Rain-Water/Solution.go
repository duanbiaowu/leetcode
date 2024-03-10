package leetcode

// 双指针, 结合图片理解
func trap(height []int) int {
	res := 0
	left, right := 0, len(height)-1
	maxLeft, maxRight := 0, 0

	// 如果右端有更高的高度，积水的高度依赖于当前方向的高度（从左到右）
	// 如果左侧有更高的高度，从相反的方向（从右到左）
	for left <= right {
		if height[left] <= height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				res += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] > maxRight {
				maxRight = height[right]
			} else {
				res += maxRight - height[right]
			}
			right--
		}
	}

	return res
}

// Stack
func trap2(height []int) int {
	res := 0

	var stack []int
	for i := range height {
		// height[i] > height[stack[len(stack)-1]]
		// 说明找到了右边 “较高的柱子” (相较当前元素而言)
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 如果栈内只有一个元素
			// 那么此时 stack 的长度等于 0
			// 也就说明此时才刚刚找到 “左边的柱子”
			// 还没有找到 “右边的柱子”
			// 比如这种情况 [0,1,0,2,1,0,1,3,2,1,2,1]
			// 当 i 为 1 时，左边的柱子 (栈顶元素) 是 0
			// 这时接不到任何雨水，所以无需计算
			if len(stack) > 0 {
				left := stack[len(stack)-1]

				// 左右两个柱子之间的距离
				w := i - left - 1
				// 左右两个柱子之间的高度，以较低的为准，“木桶原理”
				h := min(height[i], height[left]) - height[top]

				res += w * h
			}
		}

		stack = append(stack, i)
	}

	return res
}

// 总体积减去柱子体积就是水的容量
// reference: https://leetcode-cn.com/problems/volume-of-histogram-lcci/solution/shuang-zhi-zhen-an-xing-qiu-jie-xiang-xi-d162/
func trap3(height []int) int {
	sum := 0
	for i := range height {
		sum += height[i]
	}
	volume := 0
	level := 1
	left, right := 0, len(height)-1
	for left <= right {
		for left <= right && height[left] < level {
			left++
		}
		for left <= right && height[right] < level {
			right--
		}
		volume += right - left + 1
		level++
	}
	return volume - sum
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
