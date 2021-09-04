package leetcode

func findMin(nums []int) int {
	low, hi := 0, len(nums)-1
	for low <= hi {
		mid := low + (hi-low)>>1
		if nums[mid] < nums[hi] {
			hi = mid
		} else {
			low = mid + 1
		}
	}
	return nums[low]
}
