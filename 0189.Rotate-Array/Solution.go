package leetcode

func rotate(nums []int, k int) {
	if k > 0 && len(nums) > 0 {
		k %= len(nums)
		reverse(nums)
		reverse(nums[:k])
		reverse(nums[k:])
	}
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n>>1; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
