package leetcode

// 因为data中都是整数，所以可以稍微变一下，不是搜索k的两个位置，而是搜索k-0.5和k+0.5
// 这两个数应该插入的位置，然后相减即可
func searchCount(nums []int, target int) int {
	return searchLocation(nums, float64(target)+0.5) - searchLocation(nums, float64(target)-0.5)
}

func searchLocation(nums []int, target float64) int {
	low, hi := 0, len(nums)-1
	for low <= hi {
		mid := low + (hi-low)>>1
		cur := float64(nums[mid])
		if cur < target {
			low = mid + 1
		} else if cur > target {
			hi = mid - 1
		}
	}
	return low
}

// 标准二分
func searchCount2(nums []int, target int) int {
	left := searchLocationWithFlag(nums, target, true)
	// 极端情况：target is not in nums
	right := searchLocationWithFlag(nums, target, false) - 1
	if left <= right && right < len(nums) &&
		nums[left] == target && nums[right] == target {
		return right - left + 1
	}
	return 0
}

func searchLocationWithFlag(nums []int, target int, lower bool) int {
	low, hi := 0, len(nums)-1
	res := hi
	for low <= hi {
		mid := low + (hi-low)>>1
		if nums[mid] > target || (lower && nums[mid] >= target) {
			hi = mid - 1
			res = mid
		} else {
			low = mid + 1
		}
	}
	return res
}
