package leetcode

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0

	for left < right {
		width := right - left
		moreHeight := 0
		if height[left] < height[right] {
			moreHeight = height[left]
			left++
		} else {
			moreHeight = height[right]
			right--
		}
		res = max(res, width*moreHeight)
	}
	
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}