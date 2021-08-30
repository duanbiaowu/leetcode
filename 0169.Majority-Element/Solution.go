package leetcode

func majorityElement(nums []int) int {
	var res, count int
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			res = nums[i]
		}
		if nums[i] == res {
			count++
		} else {
			count--
		}
	}
	return res
}
