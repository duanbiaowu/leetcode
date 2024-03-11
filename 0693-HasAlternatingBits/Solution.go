package leetcode

// 逐个二进制位进行比较
func hasAlternatingBits(n int) bool {
	for n != 0 {
		// 当前数字最后 1 位
		prev := n & 1
		n >>= 1
		// 当前数字倒数第 2 位
		cur := n & 1

		if prev == cur {
			return false
		}
	}

	return true
}

func hasAlternatingBitsOpt(n int) bool {
	// 例如参数 n = 21 =  01 0101
	// 右移 1 位  = 10 =  00 1010
	// 两者 ^ 计算之后  = 01 1111

	// 也就是说，只有参数 n 的二进制表示符合 0,1 交替出现形式时
	// n  ^ (n >> 1) 的计算才能得到结果 0111 ... 1 这种形式
	// 此时可以给异或结果 + 1 = m, 就变为 1000 ... 0 的形式
	// 0111 1111
	// 1000 0000
	// 0000 0000 // 结果为 0

	// 最后判断 m & (m+1) 是否等于 0 即可得出结果
	n = n ^ (n >> 1)
	return n&(n+1) == 0
}
