package leetcode

func longestOnes(nums []int, k int) int {
	res, cnt := 0, 0
	left, right := 0, 0

	for right < len(nums) {
		if nums[right] == 0 {
			cnt++
		}
		for cnt > k {
			if nums[left] == 0 {
				cnt--
			}
			left++
		}
		right++
		res = max(res, right-left)
	}

	return res
}

func longestOnes2(nums []int, k int) int {
	left, right := 0, 0

	for right < len(nums) {
		if nums[right] == 0 {
			k--
		}
		if k < 0 {
			if nums[left] == 0 {
				k++
			}
			left++
		}
		right++
	}

	return right - left
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
