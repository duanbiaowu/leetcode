package leetcode

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0

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

		tmp := moreHeight * width
		if tmp > max {
			max = tmp
		}

	}
	return max
}
