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

	// 1. 初始化单调栈
	// 栈内存储数组索引
	var stack []int

	// 2. 遍历数组
	for i := range height {
		// 当前栈内索引对应的柱子依次递减，然后寻找右侧 “更高的柱子” 用来蓄水
		// 符合单调递减的特性
		// 也就是说，如果一个索引出现在栈内
		// 说明还未找到该索引对应的柱子 的右侧 “更高的柱子”

		// 3. 维护单调递减性
		// 如果当前柱子比栈定
		// 如果当前柱子高度 大于 栈顶元素高度
		// 说明当前单调递减性已经被破坏
		// 那么当前元素就是 右侧 “更高的柱子”
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			// 4. 处理具体逻辑

			// 获取栈顶元素索引
			top := stack[len(stack)-1]
			// 栈顶元素出栈
			stack = stack[:len(stack)-1]

			// 如果栈内只有一个元素
			// 那么此时 stack 的长度等于 0
			// 也就说明此时才刚刚找到 “左边的柱子”
			// 还没有找到 “右边的柱子”
			// 比如这种情况 [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]
			// 当 i 为 1 时，左边的柱子 (栈顶元素) 是 0
			// 这时接不到任何雨水，所以无需计算
			if len(stack) > 0 {
				left := stack[len(stack)-1]

				// 左右两个柱子之间的距离
				w := i - left - 1
				// 左右两个柱子之间的高度，以较低的为准，“木桶原理”
				h := min(height[i], height[left]) - height[top]

				// 将左右两个柱子之间的面积 (也就是蓄水容量) 累积到总计值
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
