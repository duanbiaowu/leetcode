package leetcode

func search(nums []int, target int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return nums[0] == target
	}

	low, hi := 0, n-1
	for low <= hi {
		mid := low + (hi-low)>>1
		if nums[mid] == target {
			return true
		}

		if nums[low] == nums[mid] && nums[mid] == nums[hi] {
			low++
			hi--
		} else if nums[low] <= nums[mid] {
			if nums[low] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[n-1] {
				low = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}

	return false
}
