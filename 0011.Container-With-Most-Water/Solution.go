package leetcode

// 如果直接看图，思路容易错误的地方在于：Y 轴坐标宽度占用宽度
// 其实题目要求只求出一个相对最大区域即可
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0

	for left < right {
		width := right - left
		short := 0
		if height[left] < height[right] {
			short = height[left]
			left++
		} else {
			short = height[right]
			right--
		}
		res = max(res, width*short)
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
