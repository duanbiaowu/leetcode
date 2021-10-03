package leetcode

func removeDuplicates(nums []int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if fast < 2 || nums[fast] != nums[slow-2] {
			slow++
		}
	}

	return slow
}
