package leetcode

func missingNumber(nums []int) int {
	low, hi := 0, len(nums)-1
	for low <= hi {
		mid := low + (hi-low)>>1
		// 如果相等，说明缺失的数字不在前半部分
		if nums[mid] == mid {
			low = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return low
}
