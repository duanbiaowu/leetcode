package leetcode

func searchRange(nums []int, target int) []int {
	left := search(nums, target)
	right := search(nums, target+1)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	return []int{left, right - 1}
}

func search(nums []int, target int) int {
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
