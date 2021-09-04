package leetcode

func findMin(nums []int) int {
	low, hi := 0, len(nums)-1
	for low < hi {
		mid := low + (hi-low)>>1
		if nums[mid] < nums[hi] {
			hi = mid
		} else if nums[mid] > nums[hi] {
			low = mid + 1
		} else {
			hi--
		}
	}
	return nums[low]
}
