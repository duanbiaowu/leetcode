package leetcode

func rotate(nums []int, k int) {
	if k > 0 && len(nums) > 0 {
		k %= len(nums)
		reverse(nums)
		reverse(nums[:k])
		reverse(nums[k:])
	}
}

func reverse(nums []int) {
	for left, right := 0, len(nums); left < right>>1; left++ {
		nums[left], nums[right-1-left] = nums[right-1-left], nums[left]
	}
}
