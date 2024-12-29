package leetcode

func jump(nums []int) int {
	steps := 0

	pos, right := 0, 0
	for i, v := range nums[:len(nums)-1] {
		pos = max(pos, i+v)
		if i == right {
			right = pos
			steps++
		}
	}

	return steps
}

func jump2(nums []int) int {
	pos := len(nums) - 1
	steps := 0

	for pos > 0 {
		for i := 0; i < pos; i++ {
			if i+nums[i] >= pos {
				pos = i
				steps++
				break
			}
		}
	}

	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
