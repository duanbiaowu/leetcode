package leetcode

func numSubarrayProductLessThanK(nums []int, k int) int {
	res, cnt := 0, 1
	for low, hi := 0, 0; hi < len(nums); hi++ {
		cnt *= nums[hi]
		for low <= hi && cnt >= k {
			cnt /= nums[low]
			low++
		}
		if low <= hi {
			res += hi - low + 1
		}
	}
	return res
}
