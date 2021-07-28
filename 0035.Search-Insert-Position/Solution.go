package leetcode

func searchInsert(nums []int, target int) int {
	low, hi := 0, len(nums)-1
	for low <= hi {
		mid := low + (hi-low)>>1
		if nums[mid] >= target {
			hi = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
