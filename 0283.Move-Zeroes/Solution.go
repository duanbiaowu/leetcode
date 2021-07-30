package leetcode

func moveZeroes(nums []int) {
	for slow, fast := 0, 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
