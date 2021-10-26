package leetcode

// 双指针, 结合图片理解
func trap(height []int) int {
	res := 0
	left, right := 0, len(height)-1
	maxLeft, maxRight := 0, 0

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
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				left := stack[len(stack)-1]
				w := i - left - 1
				h := min(height[i], height[left]) - height[top]
				res += w * h
			}
		}
		stack = append(stack, i)
	}

	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
