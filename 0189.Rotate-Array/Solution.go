package leetcode

// Example:
// nums = [1, 2, 3, 4, 5, 6, 7]; k =3
// result = [5, 6, 7, 1, 2, 3, 4];

// reverse [1, 2, 3, 4, 5, 6, 7] we can get [7, 6, 5, 4, 3, 2, 1]
// reverse [7, 6, 5] we can get	[5, 6, 7]
// reverse [4, 3, 2, 1] we can get [1, 2, 3, 4]
func rotate(nums []int, k int) {
	if k > 0 && len(nums) > 0 {
		k %= len(nums)
		reverse(nums)
		reverse(nums[:k])
		reverse(nums[k:])
	}
}

func rotateSimple(nums []int, k int) {
	n := len(nums)
	if n == 0 || k == 0 {
		return
	}

	k %= n
	clone := make([]int, n)
	copy(clone, nums)

	for i := 0; i < n; i++ {
		nums[(i+k)%n] = clone[i]
	}
}

func reverse(nums []int) {
	for left, right := 0, len(nums); left < right>>1; left++ {
		nums[left], nums[right-1-left] = nums[right-1-left], nums[left]
	}
}
