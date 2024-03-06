package leetcode

func removeElement(nums []int, val int) int {
	slow := 0
	for fast := range nums {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
