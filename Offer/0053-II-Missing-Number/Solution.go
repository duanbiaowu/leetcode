package leetcode

// 数学公式
func missingNumber2(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) >> 1
	for i := 0; i < n; i++ {
		sum -= nums[i]
	}
	return sum
}

// x ^ x = 0
// x ^ 0 = x
// x ^ y = y ^ x
// 一个萝卜一个坑（下标和数组里的数）
// 然后找没有萝卜的那个坑
func missingNumber(nums []int) int {
	sum := 0
	for i := range nums {
		sum ^= i
		sum ^= nums[i]
	}
	return sum ^ len(nums)
}

func missingNumber3(nums []int) int {
	low, hi := 0, len(nums)-1
	for low <= hi {
		mid := low + (hi-low)>>1
		// 如果相等，说明缺失的数字不在前半部分
		if nums[mid] == mid {
			low = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return low
}
