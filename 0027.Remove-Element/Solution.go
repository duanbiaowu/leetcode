package leetcode

func removeElement(nums []int, val int) int {
	slow, n := 0, len(nums)
	for fast := 0; fast < n; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
