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

func findDuplicate2(nums []int) int {
	// 数字都在 1 到 n 之间（包括 1 和 n）
	// 由于存在的重复的数字 target，因此 target 这个位置一定有起码两条指向它的边
	// 因此整张图一定存在环，且我们要找到的 target 就是这个环的入口
	slow, fast := nums[0], nums[nums[0]]
	// 直到快慢指针相遇
	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	// 两个指针每次同时移动一步，相遇的点就是 target
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
