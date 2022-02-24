package leetcode

func exchange(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		for i < j && nums[i]&1 == 1 {
			i++
		}
		for j > i && nums[j]&1 == 0 {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
