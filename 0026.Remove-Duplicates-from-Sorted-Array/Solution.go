package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	index := 1
	for i, val := range nums {
		if val != nums[index-1] {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}
