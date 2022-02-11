package leetcode

// 奇数位右移，偶数位左移，取或得结果
// 0xaaaaaaaa = 01010101010101010101010101010101  (偶数位为1，奇数位为0）
// 0x55555555 = 10101010101010101010101010101010  (偶数位为0，奇数位为1）
func exchangeBits(num int) int {
	return ((num & 0x55555555) << 1) | ((num & 0xaaaaaaaa) >> 1)
}
