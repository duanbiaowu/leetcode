package leetcode

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	slow := 2
	for fast := 2; fast < n; fast++ {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}
