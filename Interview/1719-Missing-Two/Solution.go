package leetcode

// 抽屉排序
func missingTwo(nums []int) []int {
	var res []int
	nums = append(nums, -1)
	nums = append(nums, -1)
	for i := range nums {
		for nums[i] != -1 && nums[i] != i+1 {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := range nums {
		if nums[i] == -1 {
			res = append(res, i+1)
		}
	}
	return res
}
