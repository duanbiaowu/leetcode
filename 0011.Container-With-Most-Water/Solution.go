package leetcode

// 如果直接看图，思路容易错误的地方在于：
// Y 轴坐标宽度占用的宽度，会很容易联想到 “接雨水题目”，但是这两个题目的区别在于：
//  1. 接雨水题目是求出所有的区域，而本题目是求出一个相对最大区域
//  2. 接雨水题目是求出所有的区域，所以需要求出每个区域的高度
//     而本题目是要求只求出一个相对最大区域即可，所以只需要求出 (从左到右) 最小的高度差即可
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0

	for left < right {
		// 表示当前矩形的长度
		w := right - left
		// 表示当前矩形的高度 (宽度, 类似低水位)
		h := 0

		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}

		// 更新已知的最大矩形
		res = max(res, w*h)
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
