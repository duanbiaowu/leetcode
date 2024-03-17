package leetcode

func findMin(nums []int) int {
	low, hi := 0, len(nums)-1

	for low < hi {
		mid := low + (hi-low)>>1
		if nums[mid] < nums[hi] {
			// 为什么 high = mid, 而不是 high = mid-1 ?
			// 因为 nums[mid] 有可能就是最小值
			// 示例: {4, 5, 1, 2, 3}
			// 如果 high = mid - 1，则丢失了最小值 1
			hi = mid
		} else {
			// 为什么 low = mid + 1, 而不是 low = mid ?
			// 因为 nums[mid] 不可能是最小值
			// 因为如果 nums[mid] 是最小值
			// 那么其会满足上面的表达式: nums[mid] < nums[hi]
			// 代码就不会执行到这里了
			low = mid + 1
		}
	}

	return nums[low]
}
