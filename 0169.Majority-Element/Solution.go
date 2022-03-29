package leetcode

func majorityElement(nums []int) int {
	res, cnt := 0, 0
	for i := range nums {
		if cnt == 0 {
			res = nums[i]
		}
		if res == nums[i] {
			cnt++
		} else {
			cnt--
		}
	}
	return res
}
