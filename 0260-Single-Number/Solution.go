package leetcode

// reference: https://leetcode-cn.com/problems/single-number-iii/comments/23317
// 有两个数只出现了一次记为 num1、num2 初始化为 0, 其余的数出现了两次,
// 我们可以对所有的数进行一次异或操作, 易知这个结果就等于 num1 和 num2
// 的异或结果(相同的数异或结果都为 0, 0和任意数异或结果都为那个数).

// 那么我可以考虑异或结果的某个非 0 位如最后一个非 0 位, 因为我们知道只
// 有当 num1、num2 在该位不一样的时候才会出现异或结果为 1. 所以我们以该
// 位是否为 1 对数组进行划分, 只要该位为 1 就和 num1 异或, 只要该位为 0
// 就和 num2 异或, 这样最终得到就是只出现过一次的两个数
// 	(其他在该位为 1 或 0 的数必然出现 0/2 次对异或结果无影响)
func singleNumber(nums []int) []int {
	num1, num2 := 0, 0
	xor := 0
	for i := range nums {
		xor ^= nums[i]
	}

	// 保留最右边的 1
	bit := xor & -xor

	for i := range nums {
		if nums[i]&bit == 0 {
			num1 ^= nums[i]
		} else {
			num2 ^= nums[i]
		}
	}

	return []int{num1, num2}
}
