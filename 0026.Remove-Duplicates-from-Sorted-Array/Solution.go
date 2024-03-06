package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
