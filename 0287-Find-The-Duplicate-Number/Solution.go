package leetcode

func findDuplicate(nums []int) int {
	low, hi := 0, len(nums)-1
	for low < hi {
		mid := low + (hi-low)>>1
		cnt := 0
		for i := range nums {
			if nums[i] <= mid {
				cnt++
			}
		}
		// 有一半以上的元素，说明重复元素在左侧
		if cnt > mid {
			hi = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
