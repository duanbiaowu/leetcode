package leetcode

func search(nums []int, target int) int {
	low, hi := 0, len(nums)-1
	for low <= hi {
		// 查找到最小索引，直接返回
		if nums[low] == target {
			return low
		}

		mid := low + (hi-low)>>1
		// 更新右边界
		if nums[mid] == target {
			hi = mid
		} else if nums[mid] > nums[low] { // 左边是升序的
			// 如果[low, mid - 1] 是有序数组
			// 且 target 的大小满足 ([nums[low],nums[mid]) 则将搜索范围缩小至 [low, mid - 1]
			// 否则在 [mid + 1, hi] 中寻找
			if nums[low] <= target && target <= nums[mid] {
				hi = mid
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[low] { // 右边是升序的
			// 如果[mid, hi - 1] 是有序数组
			// 且 target 的大小满足 ([nums[mid],nums[hi]) 则将搜索范围缩小至 [mid, hi - 1]
			// 否则在 [low, mid-1] 中寻找
			if nums[mid] <= target && target <= nums[hi] {
				low = mid
			} else {
				hi = mid - 1
			}
		} else {
			// 左边界前移，否则无限循环
			low++
		}
	}
	return -1
}
