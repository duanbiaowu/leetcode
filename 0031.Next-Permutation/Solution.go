package leetcode

// https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-by-leetcode-solution/
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	if i < 0 {
		return
	}
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n>>1; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
