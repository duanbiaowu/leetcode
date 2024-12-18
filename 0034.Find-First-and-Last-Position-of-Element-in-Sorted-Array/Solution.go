package leetcode

func searchRange(nums []int, target int) []int {
	// 查找目标元素 target 的插入位置
	left := search(nums, target)
	// 查找目标元素 target+1 的插入位置
	// 该位置正好就是 target 的最后一个出现位置
	right := search(nums, target+1)

	// 边界处理
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}

	return []int{left, right - 1}
}

// 寻找目标元素插入位置
// 针对存在重复元素的情况
// 如果没有重复元素，直接使用标准的二分查找即可
// 0035 原题
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

func searchRange2(nums []int, target int) []int {
	res := []int{-1, -1}
	n := len(nums)
	low, hi := 0, n-1
	if hi < 0 {
		return res
	}

	// 寻找左边界，由右侧逼近
	for low <= hi {
		mid := low + (hi-low)>>1
		if nums[mid] >= target {
			hi = mid - 1
		} else {
			low = mid + 1
		}
	}

	if low >= n || nums[low] != target {
		return res
	}

	res[0] = low
	// 寻找右边界，由左侧逼近
	hi = n - 1
	for low <= hi {
		mid := low + (hi-low)>>1
		if nums[mid] <= target {
			low = mid + 1
		} else {
			hi = mid - 1
		}
	}
	res[1] = low - 1

	return res
}
