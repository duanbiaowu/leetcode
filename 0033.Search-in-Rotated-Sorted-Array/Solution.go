package leetcode

func search(nums []int, target int) int {
	low, hi := 0, len(nums)-1

	for low <= hi {
		mid := low + (hi-low)>>1
		if nums[mid] == target {
			return mid
		}

		// 如果[low, mid - 1] 是有序数组
		// 且 target 的大小满足 ([nums[low],nums[mid]) 则将搜索范围缩小至 [low, mid - 1]
		// 否则在 [mid + 1, hi] 中寻找

		// 通过区间左右两端的数字来确认该区间是否有序
		if nums[low] <= nums[mid] {
			if nums[low] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[hi] {
				low = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}

	return -1
}
