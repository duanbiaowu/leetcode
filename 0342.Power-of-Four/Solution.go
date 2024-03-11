package leetcode

func isPowerOfFour(n int) bool {
	// 数论
	// 如果 n 是 4 的幂，其除以 3 的余数一定为 1
	// 如果 n 是 2 的幂但不是 4 的幂，其除以 3 的余数一定为 2
	//return n > 0 && n&(n-1) == 0 && n%3 == 1

	return n > 0 && n&(n-1) == 0 && n&0xaaaaaaaa == 0
}
