package leetcode

// 利用异或运算的性质:
// 1. 任何数和 0 做异或运算，结果仍然是原来的数
// 2. 任何数和其自身做异或运算，结果是 0
// 3. 异或运算满足交换律和结合律，a^b^a = b^a^a = b^(a^a) = b^0 = b
func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}
